package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `json:"members" gorm:"many2many:project_members;"`
}

type Member struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}