package main

import (
	"grpc1/graphql/generated"
	"grpc1/graphql/resolver"
	"grpc1/grpc/client"
	pb "grpc1/grpc/proto"
	"grpc1/grpc/service" // Path ke item_service.go
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors" // Import library CORS
	"google.golang.org/grpc"
)

// Fungsi untuk memulai server gRPC
func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Gunakan implementasi CRUD dari grpc/service/item_service.go
	itemService := &service.ItemServiceServer{}
	pb.RegisterItemServiceServer(grpcServer, itemService)

	log.Println("Server gRPC berjalan di :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// Jalankan server gRPC secara paralel
	go startGRPCServer()

	// Buat koneksi gRPC client
	grpcClient, err := client.NewItemServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	// Resolver untuk GraphQL
	resolver := &resolver.Resolver{
		ItemServiceClient: grpcClient,
	}

	// Konfigurasi server GraphQL
	cfg := generated.Config{Resolvers: resolver}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	// Middleware CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500"}, // Sesuaikan origin frontend
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Rute untuk GraphQL dan Playground
	http.Handle("/graphql", corsMiddleware.Handler(srv)) // Terapkan middleware CORS di sini
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	log.Println("Server GraphQL berjalan di :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
