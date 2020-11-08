package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	File       string
	CurrentPage int
	Lines       int
}