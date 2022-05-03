package dao

import (
	"go_blog/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
	CheckUser(username string) model.User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	db, err := gorm.Open(sqlite.Open("database/blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to INIT db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	log.Print("Init DB Successful")
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) CheckUser(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}
