package controller

import (
	"app-version-manager/utils/initUtils"
	"app-version-manager/utils/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	ErrorAuth      = errors.New("please add token: 'Authorization: Bearer xxxx'")
	ErrorAuthWrong = errors.New("token is not right，example: Bearer xxxx")
)

func ConfigHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		var origin string
		origin = c.Request.Header.Get("Origin")
		fmt.Println("origin is ", origin)
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		//c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Headers", "Origin,Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		fmt.Println("process option", method)
		if method == "OPTIONS" {
			fmt.Println("header: ", c.Request.Header)
			//c.AbortWithStatus(http.StatusNoContent)
			//c.AbortWithStatus(200)
			c.JSON(200, gin.H{
				"code": 3000,
				"msg": "处理跨域",
				"data": "",
			})

		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.AbortWithError(200, ErrorAuth)
			//c.JSON(200, gin.H{
			//	"code": 4000,
			//	"msg": "请求header不能为空",
			//	"data": "",
			//})
			return
		}

		authHeader := strings.Split(header, " ")

		if len(authHeader) != 2 {
			c.AbortWithError(200, ErrorAuthWrong)
			//c.JSON(200, gin.H{
			//	"code": 4001,
			//	"msg": "请求header错误",
			//	"data": "",
			//})
			return
		}

		token := authHeader[1]

		var admin model.Admin
		fmt.Println(token)
		if dbError := initUtils.MYDB.Where("auth_token = ?", token).First(&admin).Error; dbError != nil {
			c.AbortWithError(200, dbError)
			//c.JSON(200, gin.H{
			//	"code": 4002,
			//	"msg": "数据库查询错误",
			//	"data": "",
			//})
		} else {
			//c.Set("current_admin", admin)
			c.Next()
		}
	}
}