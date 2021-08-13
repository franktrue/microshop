package main

import (
	"context"
	pb "github.com/franktrue/microshop/services/demo-service/proto/demo"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("services.demo-cli"),
	)
	service.Init()

	client := pb.NewDemoServiceClient("services.demo.service", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "Franktrue"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}
