package routers

import (
	"ginjwt/pkg/setting"
	"ginjwt/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuthHandler) // 跨文件，首字母大写

	return r
}
