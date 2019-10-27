package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/micro/go-micro"

	"github.com/emiyalee/micro-random/internal/pkg/proto-spec/random"
)

type RandomServiceImpl struct {
	randomNum int32
}

func NewRandomService() *RandomServiceImpl {
	randomNum := rand.Int31n(int32(time.Now().Nanosecond()))
	log.Printf("random num:%d\n", randomNum)

	return &RandomServiceImpl{randomNum: randomNum}
}

func (impl *RandomServiceImpl) Random(ctx context.Context, req *random.RandomRequest, rsp *random.RandomResponse) error {
	rsp.RandomNum = impl.randomNum

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("random"),
	)

	service.Init()

	randomService := NewRandomService()

	random.RegisterRandomServiceHandler(service.Server(), randomService)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
