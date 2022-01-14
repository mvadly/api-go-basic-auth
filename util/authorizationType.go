package util

import (
	"api-klasterku-partner/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserBasicAuth struct {
	Username string
	Password string
}

func BasicAuth(c *gin.Context) {
	var result UserBasicAuth
	username, password, hasAuth := c.Request.BasicAuth()
	authorization := fmt.Sprintf("%v", c.Request.Header["Authorization"])
	if len(authorization) > 0 && authorization[1:6] == "Basic" {
		check := config.ConnectDB().Debug().
			Table("user_basic_auth").
			Select("username, password").
			Where("username = ? AND password = MD5(?) ", username, password).First(&result)

		if hasAuth && check.Error != nil {
			c.JSON(401, gin.H{"status": true, "message": "Unauthorized"})
			c.Abort()
		}

	} else {
		c.JSON(401, gin.H{"status": true, "message": "Unauthorized"})
		c.Abort()
	}

	c.Next()
}
