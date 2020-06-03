package convai_package_sdk

import (
	"fmt"
	"net/http"

	ctypes "github.com/datomar-labs-inc/convai-types"
	"github.com/gin-gonic/gin"
)

type LinkExecHandler func(call *ctypes.LinkCall) (ctypes.LinkCallResult, error)

type RunnableLink struct {
	ctypes.PackageLink

	Handler     LinkExecHandler `json:"-"`
	MockHandler LinkExecHandler `json:"-"`
	UI          UIHandler       `json:"-"`
}

func (p *RunnablePackage) HandleLinkExecute(c *gin.Context) {
	var input ctypes.LinkExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := ctypes.LinkExecutionResponse{[]ctypes.LinkCallResult{}}

	for _, call := range input.Calls {
		l := p.GetLink(call.ID)

		if l == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "link missing"}) // TODO: better error messages and logging
			return
		}

		result, err := l.Handler(&call)
		if err != nil {
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}

func (p *RunnablePackage) HandleLinkExecuteMock(c *gin.Context) {
	var input ctypes.LinkExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := ctypes.LinkExecutionResponse{[]ctypes.LinkCallResult{}}

	for _, call := range input.Calls {
		l := p.GetLink(call.ID)

		if l == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "link missing"}) // TODO: better error messages and logging
			return
		}

		result, err := l.MockHandler(&call)
		if err != nil {
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}

func (p *RunnablePackage) HandleLinkUI(c *gin.Context) {
	lid := c.Param("lid")

	link := p.GetLink(lid)

	if link != nil {
		c.JSON(http.StatusNotFound, gin.H{"erorr": "link not found"}) // TODO add better error response and logging
		return
	}

	html, err := link.UI()
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
