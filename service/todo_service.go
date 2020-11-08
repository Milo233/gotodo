package service

import (
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-gonic/gin"
	"gotodo/model"
	"gotodo/serializer"
)

// CreateVideoService 视频投稿的服务
type CreateTodoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info   string `form:"info" json:"info" binding:"max=3000"`
	Url    string `form:"url" json:"url"`
	Notify    bool `form:"notify" json:"notify"`
}

// Create 创建todo
func (service *CreateTodoService) Create() serializer.Response {
	todo := model.Todo {
		Title:  service.Title,
		Info:   service.Info,
		Url:   service.Url,
		Notify:   service.Notify,
	}

	// 设置默认值，userid。。
	//session := sessions.Default(c)
	//uid := session.Get("user_id")
	//todo.UserId = uid

	err := model.DB.Create(&todo).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildTodo(todo),
	}
}

// ShowTodoService 投稿详情的服务
type ShowTodoService struct {
}

// Show 视频
func (service *ShowTodoService) Show(id string) serializer.Response {
	var todo model.Todo
	err := model.DB.First(&todo, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系问题
	//video.AddView() 增加点击etc
	return serializer.Response{
		Data: serializer.BuildTodo(todo),
	}
}


// ListTodoService 视频列表服务
type ListTodoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListTodoService) List() serializer.Response {
	todos := []model.Todo{}
	total := 0
	// default pagesize
	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := model.DB.Model(model.Todo{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&todos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildTodos(todos), uint(total))
}