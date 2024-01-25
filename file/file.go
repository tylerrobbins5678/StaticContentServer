package file

import (
	"os"

	"github.com/gin-gonic/gin"
)

// if it is a file, send the file
// if it is a dir, send file data and dirs
// if not either, return 404
func HandleItem(c *gin.Context) {
	baseDir := os.Getenv("STATICCONTENTSERVER_BASE_PATH")
	path := c.Param("filepath")
	if path == "" || path == "/" {
		path = "/index.html"
	}
	path = baseDir + path
	info, err := os.Stat(path)
	if err != nil {
		c.File(baseDir + "/index.html")
		return
	}
	if !info.IsDir() {
		getFile(c, path)
		return
	}
	// c.Status(404)
}

func getFile(c *gin.Context, dir string) {
	c.File(dir)
}
