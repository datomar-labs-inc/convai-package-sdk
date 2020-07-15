package convai_package_sdk

import (
	"fmt"
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
		respErr := ctypes.Error{
			Code:    400,
			Message: "bad json request format",
		}
		c.JSON(http.StatusBadRequest, respErr) // TODO: better error messages and logging
		return
	}

	results := ctypes.DispatchResponse{[]ctypes.DispatchCallResult{}}

	for _, call := range input.Dispatches {
		d := p.GetDispatch(call.ID)

		// No dispatch, stop and let convai know
		if d == nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    400,
					Message: "dispatch missing",
				},
			}
			c.JSON(http.StatusBadRequest, errResp) // TODO: better and logging
			return
		}

		// TODO determine level of error exposure
		// Handler didn't function
		result, err := d.Handler(&call)
		if err != nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    500,
					Message: fmt.Sprintf("handler failed: %s", err.Error()),
				},
			}
			// TODO handle this appropriately
			c.JSON(http.StatusInternalServerError, errResp) // TODO: better and logging
			return
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}

func (p *RunnablePackage) HandleDispatchExecuteMock(c *gin.Context) {
	var input ctypes.DispatchRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		respErr := ctypes.Error{
			Code:    400,
			Message: "bad json request format",
		}
		c.JSON(http.StatusBadRequest, respErr) // TODO: better error messages and logging
		return
	}

	results := ctypes.DispatchResponse{[]ctypes.DispatchCallResult{}}

	for _, call := range input.Dispatches {
		d := p.GetDispatch(call.ID)

		if d == nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    400,
					Message: "dispatch missing",
				},
			}
			c.JSON(http.StatusBadRequest, errResp) // TODO: better (any) logging
			return
		}

		result, err := d.MockHandler(&call)
		if err != nil {
			// TODO logging
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    500,
					Message: fmt.Sprintf("handler failed: %s", err.Error()),
				},
			}
			c.JSON(http.StatusInternalServerError, errResp)
			return
		}

		results.Results = append(results.Results, result)
	}

	c.JSON(http.StatusOK, results)
}
