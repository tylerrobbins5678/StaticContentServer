package file

import (
	"os"

	"github.com/gin-gonic/gin"
)

// if it is a file, send the file
// if it is a dir, send file data and dirs
// if not either, return 404
func HandleItem(c *gin.Context) {
	baseDir := "V:"
	path := baseDir + c.Request.URL.Path
	info, err := os.Stat(path)
	if err != nil {
		// path/to/whatever exists
		c.File(path)
		return
	}
	if !info.IsDir() {
		getFile(c, path)
		return
	}
	c.Status(404)
}

func getFile(c *gin.Context, dir string) {
	c.File(dir)
}
