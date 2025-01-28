package gateway

import (
	"context"
	"net/http"
	"strings"

	authv1 "github.com/avalance-rl/otiva/proto/gen/avalance.auth.v1"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (g *Gateway) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		// Проверяем токен через auth service
		resp, err := g.authClient.ValidateToken(r.Context(), &authv1.TokenRequest{
			Token: token,
		})
		if err != nil {
			http.Error(w, "Auth service error", http.StatusInternalServerError)
			return
		}

		if !resp.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", resp.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (g *Gateway) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := g.authClient.Register(r.Context(), &authv1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (g *Gateway) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := g.authClient.Login(r.Context(), &authv1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
