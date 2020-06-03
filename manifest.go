package convai_package_sdk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *RunnablePackage) HManifest(c *gin.Context) {
	c.JSON(http.StatusOK, p)
}
