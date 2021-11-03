package api

import (
	"ginjwt/models"
	"ginjwt/pkg/logging"
	"ginjwt/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuthHandler(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//var a auth
	a := auth{Username: username, Password: password}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&a)
	var code int
	var msg string
	var token string
	var err error
	if ok {
		//检测账号密码
		isExit := models.ChuckAuth(username, password)
		if isExit {
			token, err = util.GenerateToken(username, password)
			if err != nil {
				code = -1
				msg = err.Error()
				logging.Error(err)
				return
			} else {
				code = 0
				msg = "ok"
				logging.Info("账号密码正确验证成功")
			}

		} else {
			logging.Fatal("账号密码不匹配")
			msg = "账号密码不匹配"
			code = -1
			return
		}

	} else {
		for _, err := range valid.Errors {
			logging.Error(err)
		}
		code = -1
		msg = "账号验证错误"
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"token": token,
	})
}
