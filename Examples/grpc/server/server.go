package main

import (
	"log" // 新增
	"net" // 新增
	"fmt"
	"context"
	"google.golang.org/protobuf/proto"
	// 新增 grpc library
	"google.golang.org/grpc"  
	// pb, short for proto buffer
	pb "grpc/test"
)

func dbServer() *MessageServiceServer {
	return &MessageServiceServer{
		savedFeatures: []*pb.Feature{
			{
				Name: []string{"東京鐵塔", "淺草寺"},
				Location: &pb.Point {
					Latitude: 353931000,
					Longitude: 139444400,
				},
			},
			{
				Name: []string{"東京鐵塔", "淺草寺"},
				Location: &pb.Point {
					Latitude: 357147651,
					Longitude: 139794466,
				},
			},
			{
				Name: []string{"東京鐵塔", "淺草寺", "晴空塔"},
				Location: &pb.Point {
					Latitude: 357100670,
					Longitude: 139808511,
				},
			},
		},
	}
}

// 定義 MessageServiceServer 中的 structure
type MessageServiceServer struct {
    pb.UnimplementedMessageServiceServer
    savedFeatures []*pb.Feature
}

// 根據 proto 中的 service 實作 function
// 		在 proto 中有定義這個 service 會接收 point 最為參數，並且會回傳 Feature
func (s *MessageServiceServer) GetFeature(ctx context.Context, point *pb.Point) (
    *pb.Feature, error,
) {
    for _, feature := range s.savedFeatures {
        if proto.Equal(feature.Location, point) {
            return feature, nil
        }
    }

    // No feature was found, return an unnamed feature
    return &pb.Feature{Location: point}, nil
}

func main()  {
	// 定義要監聽的 port 號
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8789))
    if err != nil {
        log.Fatalf("failed to listed: %v", err)
    }

    // 使用 gRPC 的 NewServer 方法來建立 gRPC Server 的實例
    grpcServer := grpc.NewServer()

    // 在 gRPC Server 中註冊 service 
    // 使用 proto 提供的 RegisterMessageServiceServer 方法，並將 MessageServiceServer 作為參數傳入
    pb.RegisterMessageServiceServer(grpcServer, dbServer())

    // 啟動 grpcServer，並阻塞在這裡直到該程序被 kill 或 stop
    err = grpcServer.Serve(lis)
    if err != nil {
        log.Fatalf("failed to serve: %v", err)
	}
}
