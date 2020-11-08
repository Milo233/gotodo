package model

import (
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	Address       string
	Name string
}
