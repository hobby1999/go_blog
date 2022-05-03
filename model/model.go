package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Salt     string `json:"salt"`
	Password string `json:"password"`
}
