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

func (s UserService) Search(context.Context, *user.UserRequest) (*user.UserResponse, error) {
	times++
	log.Printf("UserService run [%d] times \n", times)
	name := "小明"
	return &user.UserResponse{
		Id:       200,
		Username: &name,
		Area:     []string{"1004", "3009"},
	}, nil
}
