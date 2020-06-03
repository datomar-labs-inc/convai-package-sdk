package convai_package_sdk

import (
	"net/http"

	ctypes "github.com/datomar-labs-inc/convai-types"
	"github.com/gin-gonic/gin"
)

type DispatchHandler func(call *ctypes.DispatchCall) (ctypes.DispatchCallResult, error)

type RunnableDispatch struct {
	ctypes.PackageDispatch

	Handler     DispatchHandler `json:"-"`
	MockHandler DispatchHandler `json:"-"`
}

func (p *RunnablePackage) HandleDispatchExecute(c *gin.Context) {
	var input ctypes.DispatchRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := ctypes.DispatchResponse{[]ctypes.DispatchCallResult{}}

	for _, call := range input.Dispatches {
		d := p.GetDispatch(call.ID)

		if d == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "dispatch missing"}) // TODO: better error messages and logging
			return
		}

		result, err := d.Handler(&call)
		if err != nil {
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}

func (p *RunnablePackage) HandleDispatchExecuteMock(c *gin.Context) {
	var input ctypes.DispatchRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := ctypes.DispatchResponse{[]ctypes.DispatchCallResult{}}

	for _, call := range input.Dispatches {
		d := p.GetDispatch(call.ID)

		if d == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "dispatch missing"}) // TODO: better error messages and logging
			return
		}

		result, err := d.MockHandler(&call)
		if err != nil {
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}
