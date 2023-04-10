package main

import (
	"github.com/liyuanwu2020/goods/model"
	"github.com/liyuanwu2020/goods/service"
	"github.com/liyuanwu2020/micro.service.pb/go/user"
	"github.com/liyuanwu2020/msgo"
	"github.com/liyuanwu2020/msgo/register"
	"github.com/liyuanwu2020/msgo/rpc"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	goods := &model.Goods{Name: "一个花瓶", Id: 1008}

	engine := msgo.Default()
	group := engine.Group("goods")
	group.Get("/find", func(ctx *msgo.Context) {

		_ = ctx.JSON(http.StatusOK, &model.Result{Code: 1000, Msg: "success", Data: goods})
	})
	//listen, _ := net.Listen("tcp", ":9111")
	//server := grpc.NewServer()
	//user.RegisterUserServiceServer(server, &service.UserService{})
	//err := server.Serve(listen)
	//log.Println(err)

	nacosClient, nacosErr := register.CreateNacosClient()
	if nacosErr != nil {
		panic(nacosErr)
	}
	conf := register.NacosServiceConfig{
		Ip:          "localhost",
		Port:        9111,
		ServiceName: "user",
	}
	registerService, nacosErr := register.NacosServiceRegister(nacosClient, conf)
	if nacosErr != nil {
		panic(nacosErr)
	}
	log.Println(registerService)

	grpcServer, grpcErr := rpc.NewGrpcServer(":9111")
	grpcServer.Register(func(g *grpc.Server) {
		user.RegisterUserServiceServer(g, &service.UserService{})
	})
	grpcErr = grpcServer.Run()
	panic(grpcErr)
	engine.Run(":9002")
}
