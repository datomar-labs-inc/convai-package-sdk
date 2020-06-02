package convai_package_sdk

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LinkExecHandler func(call *LinkCall) (LinkCallResult, error)

type RunnableLink struct {
	Handler     LinkExecHandler `json:"-"`
	MockHandler LinkExecHandler `json:"-"`
	UI          UIHandler       `json:"-"`

	Name          string    `json:"name"`
	ID            string    `json:"id"`
	Version       string    `json:"version"` // Valid semantic version
	Style         LinkStyle `json:"style"`
	Documentation string    `json:"documentation"` // Markdown format
}

// LinkCall is Convai requesting that a package perform a link execution and return the result
type LinkCall struct {
	RequestID       uuid.UUID         `json:"request_id"` // The id of the current request
	ID              string            `json:"id"`         // The ID of the link type, used by the plugin to determine which link should be executed
	Version         string            `json:"version"`    // Which version of this link was this config created on
	Config          MemoryContainer   `json:"config"`     // How this specific link was configured by the bot builder
	PackageSettings MemoryContainer   `json:"package_settings"`
	Memory          []MemoryContainer `json:"memory"`   // Any other memory containers that this package is allowed to see
	Sequence        int               `json:"sequence"` // The number of links that have been executed during this execution
}

// LinkCallResult is what a package returns after executing a link
type LinkCallResult struct {
	RequestID uuid.UUID  `json:"request_id"` // The package is required to return the request id for security
	Logs      []LogEntry `json:"logs"`
	Errors    []Error    `json:"errors"`
	Passable  bool       `json:"passable"` // Can the execution of this module proceed down this link
}

// Link stuff
// POST /links/:lid/execute
// POST /links/:lid/execute-mock
type LinkExecutionRequest struct {
	Calls []LinkCall `json:"calls"` // All links in a call should be processed concurrently
}

type LinkExecutionResponse struct {
	Results []LinkCallResult `json:"results"` // Results should be returned in the same order that the calls were provided
}

func (p *RunnablePackage) HandleLinkExecute(c *gin.Context) {
	var input LinkExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := LinkExecutionResponse{[]LinkCallResult{}}

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
	var input LinkExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := LinkExecutionResponse{[]LinkCallResult{}}

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
