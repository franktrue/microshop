package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/franktrue/microshop/services/demo/api"
	pb "github.com/franktrue/microshop/services/demo/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	address  = "127.0.0.1:9999"
	grpcPort = ":9999"
	httpPort = "8888"
	appName  = "Demo Service"
)

type DemoService struct{}

func (d *DemoService) SayHello(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	return &pb.DemoResponse{Text: "你好，" + req.Name}, nil
}

func main() {
	mode := flag.String("mode", "grpc", "mode:grpc/http/client")
	flag.Parse()
	fmt.Println("run mode:", *mode)

	switch *mode {
	case "http":
		fmt.Println("启动HTTP服务", appName)
		// 启动HTTP服务
		api.StartWebServer(httpPort)
	case "client":
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("连接到gRPC服务失败：%v", err)
		}

		defer conn.Close()

		client := pb.NewDemoServiceClient(conn)
		req := &pb.DemoRequest{Name: "Rick"}
		rsp, err := client.SayHello(context.Background(), req)
		if err != nil {
			log.Fatalf("调用 gRPC 服务接口失败: %v", err)
		}
		log.Printf("%s", rsp.Text)
	case "grpc":
		fallthrough
	default:
		listener, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("监听指定端口失败：%v", err)
		}

		server := grpc.NewServer()
		pb.RegisterDemoServiceServer(server, &DemoService{})

		reflection.Register(server)

		if err := server.Serve(listener); err != nil {
			log.Fatalf("grpc服务启动失败: %v", err)
		}
	}
}
