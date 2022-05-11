package main

import (
	"context"
	"fmt"
	"log"
	// 新增 grpc library
	"google.golang.org/grpc"
	// pb, short for proto buffer
	pb "grpc/test"
)

func main() {

	// 建立與 server 的 channel
	conn,err := grpc.Dial("localhost:8789",grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	// 在 main() 結束後執行 conn.Close()
	defer conn.Close()
	
	client := pb.NewMessageServiceClient(conn)
	message, err := client.GetFeature(
		context.Background(),
		&pb.Point{ 
			Latitude: 357147651,
			Longitude: 139794466,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(message.GetName()[0])
	fmt.Printf("The type of message is: %T\n", message)
	fmt.Printf("The type of message.name is: %T\n", message.GetName())
}