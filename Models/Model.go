package models

import (
	"github.com/jinzhu/gorm"
)

//Todo is ...
type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      uint   `json:"rating"`
}

//TableName is ... explicitly set tablestruct
func (b *Todo) TableName() string {
	return "todo"
}
