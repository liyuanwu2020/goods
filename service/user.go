package service

import (
	"context"
	"github.com/liyuanwu2020/micro.service.pb/go/user"
	"log"
)

type UserService struct {
	user.Service
}

var times int

func (u UserService) Search(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("grpc request [%d]", times)
	defer func() {
		times++
	}()
	username := "李四"
	return &user.UserResponse{
		Id:       100,
		Username: &username,
		Area:     []string{"1003", "2048"},
	}, nil
}
