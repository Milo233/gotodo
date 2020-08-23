package serializer

import "gotodo/model"

// Video 视频序列化器
type Todo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildTodo 序列化视频
func BuildTodo(item model.Todo) Todo {
	return Todo{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildTodos 序列化视频列表
func BuildTodos(items []model.Todo) (videos []Todo) {
	for _, item := range items {
		video := BuildTodo(item)
		videos = append(videos, video)
	}
	return videos
}

