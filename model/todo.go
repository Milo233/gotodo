package model

import "github.com/jinzhu/gorm"

// Video 视频模型
type Todo struct {
	gorm.Model
	Title  string
	Info   string
}


