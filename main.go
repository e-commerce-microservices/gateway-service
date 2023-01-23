package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/e-commerce-microservices/gateway-service/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := runtime.NewServeMux()
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "u:8081", opts)
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "user-service:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth-service:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, "product-service:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = pb.RegisterShopServiceHandlerFromEndpoint(ctx, mux, "shop-service:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = pb.RegisterImageServiceHandlerFromEndpoint(ctx, mux, "image-service:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = pb.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, "search-service:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen on port: 8080")

	if err := http.ListenAndServe(":8080", allowCORS(mux)); err != nil {
		log.Fatal(err)
	}
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Printf("preflight request for %s", r.URL.Path)
	return
}
