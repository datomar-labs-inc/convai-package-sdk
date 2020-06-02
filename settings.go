package convai_package_sdk

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RunnableSettings struct {
	UI UIHandler `json:"-"`
}

func (p *RunnablePackage) HandleSettingsUI(c *gin.Context) {
	html, err := p.Settings.UI()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erorr": "error rendering " + err.Error()}) // TODO add better error response and logging
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write([]byte(html))
	if err != nil {
		fmt.Println("Failed to send link ui " + err.Error())
		return
	}
}
