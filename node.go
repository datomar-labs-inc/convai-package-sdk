package convai_package_sdk

import (
	"fmt"
	"net/http"

	ctypes "github.com/datomar-labs-inc/convai-types"
	"github.com/gin-gonic/gin"
)

type NodeExecHandler func(call *ctypes.NodeCall) (ctypes.NodeCallResult, error)

// Node stuff
// POST /nodes/:nid/execute
// POST /nodes/:nid/execute-mock
type RunnableNode struct {
	ctypes.PackageNode

	Handler     NodeExecHandler `json:"-"`
	MockHandler NodeExecHandler `json:"-"`
	UI          UIHandler       `json:"-"`
}

func (p *RunnablePackage) HandleNodeExecute(c *gin.Context) {
	var input ctypes.NodeExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := ctypes.NodeExecutionResponse{[]ctypes.NodeCallResult{}}

	for _, call := range input.Calls {
		n := p.GetNode(call.TypeID)

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
	var input ctypes.NodeExecutionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := ctypes.NodeExecutionResponse{[]ctypes.NodeCallResult{}}

	for _, call := range input.Calls {
		n := p.GetNode(call.TypeID)

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
