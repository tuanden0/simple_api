package main

import (
	"context"
	"fmt"

	"github.com/tuanden0/simple_api/internal/api"
	"google.golang.org/grpc"
)

const addrStr = ":50001"

func main() {
	// Call GRPC server
	con, err := grpc.Dial(addrStr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// Init Client
	client := api.NewStudentServiceClient(con)

	// Test create
	req := &api.CreateRequest{
		Name: "chep",
		Gpa:  0.1,
	}

	res, err := client.Create(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
