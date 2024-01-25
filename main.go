package main

import (
	"fmt"
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
		} else {
			fmt.Println(r + " : " + os.Getenv(r))
		}
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/portfolio/")
		//c.File(os.Getenv("STATICCONTENTSERVER_BASE_PATH") + "/index.html")
	})
	r.GET("/static/*filepath", file.HandleItem)
	r.GET("/portfolio/*filepath", file.HandleItem)
	//r.GET("/assets/*filepath", file.HandleItem)
	//r.Any("/*filepath", file.HandleItem)
	r.Run(":" + os.Getenv("STATICCONTENTSERVER_PORT"))
}
