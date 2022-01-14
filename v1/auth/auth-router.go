package auth

import (
	"api-klasterku-partner/util"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(util.BasicAuth)
	v1 := r.Group("/v1")
	{
		v1.GET("auth/credential/login", Login)
	}

}
