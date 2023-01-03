package main

import (
	"context"
	"log"
	"net/http"

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

	log.Println("listen on port: 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
