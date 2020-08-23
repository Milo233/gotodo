package server

import (
	"github.com/gin-gonic/gin"
	"gotodo/api"
	"gotodo/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors()) // 配置接收跨域请求

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		// 增
		v1.POST("todo",api.CreateTodo)
		// 删
		// 查
		v1.GET("todos",api.ListTodos)
		// 改
	}

	return r
}
