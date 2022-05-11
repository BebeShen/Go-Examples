package main

import (
	"log"
	"context"

	"google.golang.org/grpc"
	pb "client/pb"
)

func main()  {
	conn, err := grpc.Dial("localhost:8789", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}
	client := pb.NewPNserviceClient(conn)
	ReplyMessage, err := client.UpdateToken(context.Background(), &pb.TokenRequest{UserId: 8839, Token: "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]"})
	if err != nil {
		log.Println(err)
	}
	log.Println(ReplyMessage)
}