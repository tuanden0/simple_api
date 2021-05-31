package main

import (
	"fmt"
	"net"

	"github.com/tuanden0/simple_api/internal/api"
	"github.com/tuanden0/simple_api/internal/models"
	"github.com/tuanden0/simple_api/internal/repository"
	"github.com/tuanden0/simple_api/internal/services"
	"google.golang.org/grpc"
)

const (
	netStr  = "tcp"
	addrStr = ":50001"
)

func main() {

	// Init server
	fmt.Println("Starting server...")
	lis, err := net.Listen(netStr, addrStr)
	if err != nil {
		panic(err)
	}

	// Connect DB
	db := models.ConnectDatabase()

	// Create Student repo
	studentRepo := repository.NewStudentRepo(db)
	studentSrv := services.NewStudentGRPCService(studentRepo)

	// Init GRPC
	s := grpc.NewServer()
	api.RegisterStudentServiceServer(s, studentSrv)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
