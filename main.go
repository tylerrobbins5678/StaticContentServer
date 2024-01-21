package main

import (
	"os"
	"staticcontentserver/file"

	"github.com/gin-gonic/gin"
)

func main() {

	required_envs := []string{
		"STATICCONTENTSERVER_PORT",
		"STATICCONTENTSERVER_BASE_PATH",
	}

	for _, r := range required_envs {
		if os.Getenv(r) == "" {
			panic(r + " not in envirenment")
		}
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Thank you for visiting. This page is still in development and coming soon")
	})
	r.GET("/static/*filepath", file.HandleItem)
	r.Run(":" + os.Getenv("STATICCONTENTSERVER_PORT"))
}
