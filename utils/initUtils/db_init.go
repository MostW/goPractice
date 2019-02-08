package initUtils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MYDB *gorm.DB

func init() {
	//db initUtils
	db, err := gorm.Open("mysql", "root@(127.0.0.1:3306)/app_version_manager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("connect mysql failed")
	}
	fmt.Println("Login mysql database success!")
	MYDB = db

}