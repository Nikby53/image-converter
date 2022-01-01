package handler

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
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
			s.errorJSON(w, http.StatusUnauthorized, fmt.Errorf("empty authorization handler"))
			return
		}

		HeaderParts := strings.Split(authHeader, " ")
		if len(HeaderParts) != 2 || HeaderParts[0] != "Bearer" {
			s.errorJSON(w, http.StatusUnauthorized, fmt.Errorf("invalid auth header"))
			return
		}

		if HeaderParts[1] == "" {
			s.errorJSON(w, http.StatusUnauthorized, fmt.Errorf("token is empty"))
			return
		}
		token := HeaderParts[1]
		userID, err := s.services.ParseToken(token)
		if err != nil {
			s.errorJSON(w, http.StatusUnauthorized, fmt.Errorf("can't parse jwt token: %w", err))
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
		defer panicHandler(w)
		handler.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"request path":    r.URL.Path,
			"request method":  r.Method,
			"time for answer": time.Since(begin),
		}).Infoln("request handled")
	})
}

func panicHandler(w http.ResponseWriter) {
	if err := recover(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err.(error).Error(), string(debug.Stack()))
	}
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
