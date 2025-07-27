package main

import (
	"database/sql"
	"grpc/internal/database"
	"grpc/internal/pb"
	"grpc/internal/service"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./grpc.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	// course_category_grpc.pb.go line 78
	// Register the CategoryService with the gRPC server
	// This is where the service is registered to handle incoming gRPC requests
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	// Register reflection service on gRPC server
	// This allows clients to discover the service and its methods at runtime
	// Useful for debugging and development purposes
	reflection.Register(grpcServer)

	// Listen on port 50051 for incoming gRPC requests
	// This is the port where the gRPC server will be accessible
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	// Start the gRPC server
	// This is where the server starts accepting incoming requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
