package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/bar-raisers/intention/common/services"

	tasks "github.com/bar-raisers/intention/scheduling/contracts/tasks/v1"
)

func main() {
	serviceFlags := services.NewServiceFlags()
	serviceFlags.Init(nil)

	address := fmt.Sprintf("%s:%d", serviceFlags.GetAddress(), serviceFlags.GetPort())
	log.Printf("starting accounts service at %s...\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to start accounts service: %v\n", err)
	}

	server := grpc.NewServer()
	service := NewTasksService()

	tasks.RegisterTasksServiceServer(server, service)
	reflection.Register(server)
	server.Serve(listener)
}
