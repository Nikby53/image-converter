package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (s *Server) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AuthorizationHeader)
		if authHeader == "" {
			http.Error(w, "empty authorization handler", http.StatusUnauthorized)
			return
		}

		HeaderParts := strings.Split(authHeader, " ")
		if len(HeaderParts) != 2 || HeaderParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		if len(HeaderParts[1]) == 0 {
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
