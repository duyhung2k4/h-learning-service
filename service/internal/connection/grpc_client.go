package connection

import (
	"app/generated/grpc/servicegrpc"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func connectGrpcService() {
	connectGrpcServiceQuizz()
	connectGrpcServerStream()
	connectGrpcServerQuantity()
}

func connectGrpcServiceQuizz() {
	connGrpc, err := grpc.NewClient(fmt.Sprintf("localhost:%s", conn.QuizzService.Grpc), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpcClientQuizz = servicegrpc.NewQuizzServiceClient(connGrpc)

	log.Println("connected service grpc quizz")
}

func connectGrpcServerStream() {
	connGrpc, err := grpc.NewClient(fmt.Sprintf("localhost:%s", conn.StreamService.Grpc), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpcClientStream = servicegrpc.NewStreamServiceClient(connGrpc)

	log.Println("connected service grpc stream")
}

func connectGrpcServerQuantity() {
	connGrpc, err := grpc.NewClient(fmt.Sprintf("localhost:%s", conn.QuantityBlobService.Grpc), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpcClientQuantity = servicegrpc.NewQuantityServiceClient(connGrpc)

	log.Println("connected service grpc quantity")
}
