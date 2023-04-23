package main

import (
	"fmt"
	"net"

	"github.com/georgejr3211/grpc/internal/pb"
	"github.com/georgejr3211/grpc/internal/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	categoryService := service.NewCategoryService()

	pb.RegisterCategoryServiceServer(s, categoryService)

	fmt.Println("Runing on tcp://0.0.0.0:" + port)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
