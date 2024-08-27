package main

import (
	"log"
	"net"

	"github.com/shania-kawa/Billing-Project/api"
	"github.com/shania-kawa/Billing-Project/pkg/db"
	"github.com/shania-kawa/Billing-Project/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	//connecting to data base
	err := db.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database :%v", err)

	}
	db.RunTables()
	db.Close()

	//grpc
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)

	}

	grpcServer := grpc.NewServer()
	api.RegisterBillingServiceServer(grpcServer, &service.BillingService{})
	log.Println("grpc server is running on port 50051")

	if err := grpcServer.serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
