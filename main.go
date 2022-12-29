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
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
