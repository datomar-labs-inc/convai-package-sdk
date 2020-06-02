package convai_package_sdk

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NodeExecHandler func(call *NodeCall) (NodeCallResult, error)

// NodeCall is Convai requesting that a package perform a node execution and return the result
type NodeCall struct {
	RequestID       uuid.UUID         `json:"request_id"` // The ID of the current request
	ID              string            `json:"id"`         // The ID of the node type, used by the plugin to determine which node
	Version         string            `json:"version"`    // Which version of this node was this config created on
	Config          MemoryContainer   `json:"config"`     // How this specific node was configured by the bot builder
	PackageSettings MemoryContainer   `json:"package_settings"`
	Memory          []MemoryContainer `json:"memory"`   // Any other memory containers that this package is allowed to see
	Sequence        int               `json:"sequence"` // The number of nodes that have been executed during this execution
}

// NodeCallResult is what a package returns after executing a node
type NodeCallResult struct {
	RequestID       uuid.UUID        `json:"request_id"` // The package is required to return the request id for security
	Transformations []Transformation `json:"transformations"`
	Logs            []LogEntry       `json:"logs"`
	Errors          []Error          `json:"errors"`
}

// Node stuff
// POST /nodes/:nid/execute
// POST /nodes/:nid/execute-mock
type NodeExecutionRequest struct {
	Calls []NodeCall `json:"calls"` // All nodes in a pack must complete execution before returning, so don't pack nodes
}

type NodeExecutionResponse struct {
	Results []NodeCallResult `json:"results"` // Results should be returned in the same order that the calls were provided
}

type RunnableNode struct {
	Handler     NodeExecHandler `json:"-"`
	MockHandler NodeExecHandler `json:"-"`
	UI          UIHandler       `json:"-"`

	Name          string    `json:"name"`
	ID            string    `json:"id"`
	Version       string    `json:"version"` // Valid semantic version
	Style         NodeStyle `json:"style"`
	Documentation string    `json:"documentation"` // Markdown format
}

func (p *RunnablePackage) HandleNodeExecute(c *gin.Context) {
	var input NodeExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := NodeExecutionResponse{[]NodeCallResult{}}

	for _, call := range input.Calls {
		n := p.GetNode(call.ID)

		if n == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "node missing"}) // TODO: better error messages and logging
			return
		}

		result, err := n.Handler(&call)
		if err != nil {
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}

func (p *RunnablePackage) HandleNodeExecuteMock(c *gin.Context) {
	var input NodeExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := NodeExecutionResponse{[]NodeCallResult{}}

	for _, call := range input.Calls {
		n := p.GetNode(call.ID)

		if n == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "node missing"}) // TODO: better error messages and logging
			return
		}

		result, err := n.MockHandler(&call)
		if err != nil {
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}

func (p *RunnablePackage) HandleNodeUI(c *gin.Context) {
	nid := c.Param("nid")

	node := p.GetNode(nid)

	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"erorr": "node not found"}) // TODO add better error response and logging
		return
	}

	html, err := node.UI()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erorr": "error rendering " + err.Error()}) // TODO add better error response and logging
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write([]byte(html))
	if err != nil {
		fmt.Println("Failed to send node ui " + err.Error())
		return
	}
}
