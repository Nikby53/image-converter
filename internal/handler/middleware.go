package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
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
		UserID, err := s.services.ParseToken(token)
		if err != nil {
			http.Error(w, "can't parse jwt token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userCtx, UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetIDFromToken gets the id from user token.
func (s *Server) GetIDFromToken(r *http.Request) (int, error) {
	authHeader := r.Header.Get(authorizationHeader)
	HeaderParts := strings.Split(authHeader, " ")
	token := HeaderParts[1]
	UserID, err := s.services.ParseToken(token)
	if err != nil {
		return 0, fmt.Errorf("can't parse jwt token %w", err)
	}
	return UserID, nil
}
