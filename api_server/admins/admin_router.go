package admins

import (
	"app-version-manager/utils/crypt"
	"app-version-manager/utils/initUtils"
	"app-version-manager/utils/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrorAdminSignInParam = errors.New("admin sign in error")
	ErrorAdminSignName = errors.New("name is not allow")
	ErrorAdminSignUpName = errors.New("name is exists")
)

var (
	NameInfo = fmt.Sprintf("Please Check Name length")
)

type ListSignInParam struct {
	Data struct {
		Name     string `json:"name" binding:"required" example:"18717711819"`
		Password string `json:"password" binding:"required" example:"xxxxxx"`
	} `json:"admin" binding:"required"`
}

func SignIn(c *gin.Context) {
	var param ListSignInParam
	if err := c.ShouldBindJSON(&param); err != nil {
		//c.JSON(400, c.AbortWithError(400, ErrorAdminSignInParam))
		c.JSON(
			400,
			gin.H{
				"code": 3000,
				"msg": "参数错误",
				"data": c.AbortWithError(400, ErrorAdminSignInParam)},
		)
		return
	}
	ok := crypt.CheckSignInName(param.Data.Name)
	if ok == false {
		fmt.Println(NameInfo)
		return
	}
	var admin model.Admin
	if dbError := initUtils.MYDB.Where("name = ?", param.Data.Name).First(&admin).Error; dbError != nil {
		//c.JSON(400, c.AbortWithError(400, dbError))
		c.JSON(
			400,
			gin.H{
				"code": 4000,
				"msg": "数据库查询数据失败",
				"data": c.AbortWithError(400, dbError)},
		)
		return
	}
	//c.JSON(http.StatusOK, admin.Serializer())
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": 0,
			"msg": "登录成功",
			"data": admin.Serializer()},
	)
}

type ListSignUpParam struct {
	Data struct {
		Name     string `json:"name" binding:"required" example:"18717711819"`
		Password string `json:"password" binding:"required" example:"xxxxx"`
		//Phone    string `json:"phone" binding:"required" example:"18717711819"`
	} `json:"admin" binding:"required"`
}

func SignUp(c *gin.Context) {
	var param ListSignUpParam
	if err := c.ShouldBindJSON(&param); err != nil {
		//c.JSON(400, c.AbortWithError(400, ErrorAdminSignInParam))
		c.JSON(200, gin.H{
			"code": 3000,
			"msg": "参数错误",
			"data": c.AbortWithError(400, ErrorAdminSignInParam),
		})
		return
	}

	if ok := crypt.CheckSignInName(param.Data.Name); ok != true {
		//c.JSON(400, c.AbortWithError(400, ErrorAdminSignName))
		c.JSON(200, gin.H{
			"code": 3001,
			"msg": "name格式错误",
			"data": c.AbortWithError(400, ErrorAdminSignName),
		})
		return
	}
	var admin model.Admin
	if dbError := initUtils.MYDB.Where("name = ?", param.Data.Name).First(&admin).Error; dbError == nil {
		//c.JSON(400, c.AbortWithError(400, ErrorAdminSignUpName))
		c.JSON(200, gin.H{
			"code": 4000,
			"msg": "数据库查询错误",
			"data": c.AbortWithError(400, dbError),
		})
		return
	}
	var newAdmin model.Admin
	newAdmin = model.Admin{
		Name:              param.Data.Name,
		AuthToken:         crypt.GenerateToken(),
		EncryptedPassword: crypt.PassWordEncrypted(param.Data.Password),
	}
	if dbError := initUtils.MYDB.Create(&newAdmin).Error; dbError != nil {
		//c.JSON(400, c.AbortWithError(400, dbError))
		c.JSON(200, gin.H{
			"code": 4001,
			"msg": "数据库插入数据错误",
			"data": c.AbortWithError(400, dbError),
		})
		return
	}

	//c.JSON(http.StatusOK, newAdmin.Serializer())
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "注册成功",
		"data": "",
	})
}
