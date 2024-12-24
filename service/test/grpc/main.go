package main

import (
	"app/generated/grpc/enumgrpc"
	"app/generated/grpc/servicegrpc"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := servicegrpc.NewQuizzServiceClient(conn)

	res, err := client.CreateQuizz(context.Background(), &servicegrpc.CreateQuizzRequest{
		Quizz: &servicegrpc.Quizz{
			Ask:        "",
			Time:       10,
			Result:     []string{"1", "2"},
			Option:     []string{"1", "2", "3"},
			EntityType: enumgrpc.EntityType_QUIZZ_VIDEO_LESSON,
			EntityId:   1,
		},
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}
