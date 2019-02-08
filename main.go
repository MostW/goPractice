package main

import (
	"app-version-manager/api_server"
	"app-version-manager/utils/initUtils"
)
import "app-version-manager/utils/model"

func main() {
	StartTable()
	defer initUtils.MYDB.Close()
	api_server.New().Start()
}

func StartTable() {
	initUtils.MYDB.AutoMigrate(
		&model.Admin{},
	)
	//initUtils.MYDB.Create(&model.Admin{Name: "test", EncryptedPassword: "123456", AuthToken: "testToken"})
}
