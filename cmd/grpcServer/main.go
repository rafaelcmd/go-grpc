package main

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rafaelcmd/gRPC/internal/database"
	"github.com/rafaelcmd/gRPC/internal/pb"
	"github.com/rafaelcmd/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	host, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
