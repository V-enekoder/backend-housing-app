package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);uniqueIndex"`
	Password string `gorm:"type:varchar(255)"`

	Projects []Project `gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(255)"`
	Models []Model `gorm:"many2many:category_model;"`
}

type Project struct {
	gorm.Model
	UserID  uint
	Name    string  `gorm:"type:varchar(255)"`
	Height  float64 `gorm:"type:float"`
	Width   float64 `gorm:"type:float"`
	Length  float64 `gorm:"type:float"`
	Address string  `gorm:"type:text"`

	User   User    `gorm:"foreignKey:UserID"`
	Models []Model `gorm:"many2many:model_project;"`
}

type Model struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(255)"`
	Address    string     `gorm:"type:text"`
	Mimetype   string     `gorm:"type:varchar(255)"`
	Categories []Category `gorm:"many2many:category_model;"`
	Projects   []Project  `gorm:"many2many:model_project;"`
}
