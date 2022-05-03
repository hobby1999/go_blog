package plugin

import (
	"crypto/md5"
	"encoding/hex"
	"go_blog/dao"
	"go_blog/model"
	"log"
	"strconv"
	"time"
)

type PluginHelp interface {
	HashCode(hashneed string)
}

func HashCode(hashneed string) (string, string) {
	salt := time.Now().Unix()
	md5 := md5.New()
	md5.Write([]byte(hashneed))
	md5.Write([]byte(strconv.Itoa(int(salt))))
	hashPassword := md5.Sum(nil)
	return strconv.Itoa(int(salt)), hex.EncodeToString(hashPassword)
}

func CheckHashCode(hashneed, salt string) string {
	md5 := md5.New()
	md5.Write([]byte(hashneed))
	md5.Write([]byte(salt))
	hashResult := md5.Sum(nil)
	return hex.EncodeToString(hashResult)
}

func AddUser() {
	username := "admin"
	password := "password"
	salt, hashPassrod := HashCode(password)
	user := model.User{
		Username: username,
		Password: hashPassrod,
		Salt:     salt,
	}

	dao.Mgr.AddUser(&user)
}

func InitAdmin() {
	user := dao.Mgr.CheckUser("admin")
	if user.Username == "" {
		AddUser()
		log.Printf("User Init Successful")
	} else {
		log.Printf("User Has Been Init")
	}
	defaultUser := dao.Mgr.CheckUser("admin")
	log.Printf("默认管理员账户:%s", defaultUser.Username)
	log.Printf("默认管理员密码:%s", "password")
	log.Printf("请尽快修改默认管理密码！！！")
}
