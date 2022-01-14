package partner

import "github.com/gin-gonic/gin"

func GetDataKlaster(c *gin.Context) {
	c.XML(200, gin.H{"status": true, "message": "ok"})
}
