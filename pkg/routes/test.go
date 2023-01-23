package routes

import (
	"github.com/lerryjay/email-service/pkg/pb"
	"golang.org/x/net/context"
)

func SayingHello(ctx context.Context, req *pb.SayingHelloRequest) (*pb.SayingHelloResponse, error) {

	return &pb.SayingHelloResponse{Response: "Hello " + req.Name}, nil
}
