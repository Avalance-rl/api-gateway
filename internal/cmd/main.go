package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avalance-rl/otiva/services/api-gateway/internal/internal/config"
	"github.com/avalance-rl/otiva/services/api-gateway/internal/internal/gateway"
)

func main() {
	cfg, err := config.Load(os.Getenv("CONFIG_FILE"))
	gw, err := gateway.NewGateway(
		cfg.AuthService.Host + cfg.AuthService.Port,
	)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", gw.Register)
	mux.HandleFunc("POST /auth/login", gw.Login)

	// Защищенные эндпоинты продуктов
	// mux.HandleFunc("/products/{.*}", gateway.AuthMiddleware(gateway.ProxyToProduct)).Methods("GET", "POST", "PUT", "DELETE")

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
