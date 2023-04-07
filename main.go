package main

import (
	"github.com/liyuanwu2020/goods/model"
	"github.com/liyuanwu2020/goods/service"
	"github.com/liyuanwu2020/micro.service.pb/go/user"
	"github.com/liyuanwu2020/msgo"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	engine := msgo.Default()
	group := engine.Group("goods")
	group.Get("/find", func(ctx *msgo.Context) {
		goods := &model.Goods{Name: "一个花瓶", Id: 1008}
		_ = ctx.JSON(http.StatusOK, &model.Result{Code: 1000, Msg: "success", Data: goods})
	})
	listen, _ := net.Listen("tcp", ":9111")
	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &service.UserService{})
	err := server.Serve(listen)
	log.Println(err)
	engine.Run(":9002")
}
