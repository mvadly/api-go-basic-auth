package partner

import (
	"api-klasterku-partner/util"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(util.BasicAuth)
	v1 := r.Group("/v1/partner/")
	{
		v1.GET("get_data_klaster", GetDataKlaster)
	}

}
