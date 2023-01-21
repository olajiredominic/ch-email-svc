package main

import (
	"log"
	"net"

	"github.com/lerryjay/email-service/pkg/config"
	"github.com/lerryjay/email-service/pkg/pb"

	routes "github.com/lerryjay/email-service/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := routes.Init(config)

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, &h)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.EMAIL_SVC_PORT)
	if err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
