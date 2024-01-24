package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/bar-raisers/intention/common/services"

	accounts "github.com/bar-raisers/intention/contracts/finance/accounts/v1"
)

func main() {
	serviceFlags := services.NewServiceFlags()
	serviceFlags.Init(os.Args[1:])

	address := fmt.Sprintf("%s:%d", serviceFlags.GetAddress(), serviceFlags.GetPort())
	log.Printf("starting accounts service at %s...\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to start accounts service: %v\n", err)
	}

	server := grpc.NewServer()
	service := NewAccountsService()

	accounts.RegisterAccountsServiceServer(server, service)
	reflection.Register(server)
	server.Serve(listener)
}
