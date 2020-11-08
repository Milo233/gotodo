package server

import (
	"github.com/gin-gonic/gin"
	"gotodo/api"
	"gotodo/middleware"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 初始化session、、不加的话 登录的时候会出现问题
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors()) // 配置接收跨域请求
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		//todo add 查询link的接口
		v1.GET("ping", api.Ping)
		// 用户注册 不让注册...
		// v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		// 用户退出登录
		v1.DELETE("user/logout", api.UserLogout)
		// 查询当前登录用户信息
		v1.GET("user/me", api.UserMe)

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// 增
			authed.POST("todo",api.CreateTodo)
			// 删
			// 查
			authed.GET("todos",api.ListTodos)
			// 查
			authed.GET("todo/:id", api.ShowTodo)
			// 改
			//authed.DELETE("user/logout", api.UserLogout)
			authed.GET("book/read", api.ReadBook)
		}

	}

	return r
}
