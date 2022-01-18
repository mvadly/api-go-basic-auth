package main

import (
	"api-klasterku-partner/config"
	"api-klasterku-partner/v1/auth"
	"api-klasterku-partner/v1/partner"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	routes := gin.Default()
	auth.Router(routes)
	partner.Router(routes)

	routes.Run(":1000")
}
