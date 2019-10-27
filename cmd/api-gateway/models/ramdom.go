package models

import (
	"context"

	//_ "github.com/micro/go-plugins/broker/rabbitmq"
	//_ "github.com/micro/go-plugins/client/grpc"
	//_ "github.com/micro/go-plugins/selector/static"
	//_ "github.com/micro/go-plugins/server/grpc"

	pb "github.com/emiyalee/micro-random/internal/pkg/proto-spec/random"
	micro "github.com/micro/go-micro"
)

var service micro.Service

type RandomClient struct {
}

func init() {
	service = micro.NewService(
		micro.Name("api"),
	)
	service.Init()
}

func NewRandomClient() (*RandomClient, error) {
	r := &RandomClient{}
	return r, nil
}

func (c *RandomClient) Random() (int32, error) {
	client := pb.NewRandomService("random", service.Client())
	rsp, err := client.Random(context.TODO(), &pb.RandomRequest{})

	if err != nil {
		return 0, err
	}

	return rsp.RandomNum, nil
}
