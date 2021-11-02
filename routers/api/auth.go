package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
