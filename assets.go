package convaipkgsdk

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AssetHandler when given a filename, should return a reader and mime type
type AssetHandler func(filename string) (io.ReadCloser, string, error)

func (p *RunnablePackage) HandleAssetRequest(c *gin.Context) {
	filename := c.Param("filename")

	reader, mime, err := p.Assets(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", mime)
	c.Status(http.StatusOK)

	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		fmt.Println("Failed to serve asset", err.Error())
	}

	reader.Close()
}