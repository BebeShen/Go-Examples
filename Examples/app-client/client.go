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
	// ReplyMessage, err := client.UpdateToken(context.Background(), &pb.TokenRequest{UserId: 8839, Token: "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]"})
	sendNotificationReplyMessage, err := client.SendNotification(
		context.Background(), 
		&pb.NotificationRequest{
			UserId: []uint32{8839}, 
			Content: "gRPC大成功(ಥ﹏ಥ)",
			NotificationType: "chat",
			UidFrom: 8840})
	log.Println(sendNotificationReplyMessage)
	updateNotificationReplyMessage, err := client.UpdateNotificationSettings(context.Background(), &pb.UpdateSettingRequest{UserId: 8839, Setting: "1,1,1,0,1,1,0"})
	if err != nil {
		log.Println(err)
	}
	log.Println(updateNotificationReplyMessage)
}