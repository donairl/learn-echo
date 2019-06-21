package models

import (
	"github.com/jinzhu/gorm"
)

//Article ini adalah model article
type Article struct {
	gorm.Model
	Title     string `gorm:"type:varchar(255)"`
	Body      string `gorm:"type:varchar(255)"`
	CreatedBy string
	User      User `gorm:"foreignkey:CreatedBy"`
}
