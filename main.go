package main

import (
	"github.com/liyuanwu2020/goods/model"
	"github.com/liyuanwu2020/msgo"
	"net/http"
)

func main() {
	engine := msgo.Default()
	group := engine.Group("goods")
	group.Get("/find", func(ctx *msgo.Context) {
		goods := &model.Goods{Name: "一个花瓶", Id: 1008}
		ctx.JSON(http.StatusOK, &model.Result{Code: 1000, Msg: "success", Data: goods})
	})
	engine.Run(":9002")
}