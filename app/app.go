package main

import (
	"os"

	"github.com/GibMe/backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	registerHandlers(router)
	router.Run()
}

func registerHandlers(router *gin.Engine) {

	router.GET("/myfiles/:userid/:fileid", func (c *gin.Context) {
		userId := c.Param("userid")
		fileId := c.Param("fileid")

		if err := config.DB.getFile(userId, fileId); err != nil {
			c.Error()
		}
	})
}
