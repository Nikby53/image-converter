package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// UserIDCtx is custom type.
type UserIDCtx string

const (
	authorizationHeader = "Authorization"
	// UserCtxKey is user's id context key.
	UserCtxKey = "userID"
)

// UserIdentity checks if the user is authorized or not.
func (s *Server) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authorizationHeader)
		if authHeader == "" {
			http.Error(w, "empty authorization handler", http.StatusUnauthorized)
			return
		}

		HeaderParts := strings.Split(authHeader, " ")
		if len(HeaderParts) != 2 || HeaderParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		if HeaderParts[1] == "" {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}
		token := HeaderParts[1]
		userID, err := s.services.ParseToken(token)
		if err != nil {
			http.Error(w, "can't parse jwt token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDCtx(UserCtxKey), userID)
		r = r.Clone(ctx)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		log.WithFields(log.Fields{
			"request path":    r.URL.Path,
			"request method":  r.Method,
			"time for answer": begin.Format(time.RFC822),
		}).Infoln("get request")
		handler.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"request path":    r.URL.Path,
			"request method":  r.Method,
			"time for answer": time.Since(begin),
		}).Infoln("request handled")
	})
}

// GetIDFromContext get user's id from context.
func (s *Server) GetIDFromContext(ctx context.Context) (int, error) {
	userIDCtx := ctx.Value(UserIDCtx(UserCtxKey))
	if userIDCtx == nil {
		return 0, fmt.Errorf("can't get id from context")
	}
	userID, ok := userIDCtx.(int)
	if !ok {
		return 0, fmt.Errorf("can't convert to int")
	}
	return userID, nil
}
