package main

import (
	"context"
	"fmt"

	"github.com/georgejr3211/grpc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewCategoryServiceClient(conn)

	in := &pb.CategoryRequest{
		Name:        "test",
		Description: "test",
	}

	resp, err := client.CreateCategory(context.Background(), in)
	if err != nil {
		panic(err)
	}

	fmt.Println("New Category", resp)
}
