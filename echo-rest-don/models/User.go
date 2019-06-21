package models

import "github.com/jinzhu/gorm"

//User model for user

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"size:64"` //
	Password string `json:"password" gorm:"size:65"` //
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Role     int8   `json:"role"`
}
