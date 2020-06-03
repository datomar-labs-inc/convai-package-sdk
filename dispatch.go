package convai_package_sdk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DispatchRequest struct {
	Dispatches []DispatchCall `json:"dispatches"`
}

type DispatchResponse struct {
	Results []DispatchCallResult `json:"dispatch_result"`
}

type DispatchCall struct {
	RequestID       uuid.UUID       `json:"request_id"`
	ID              string          `json:"id"`           // The ID of the type of dispatch being called
	MessageBody     string          `json:"message_body"` // XML format message body that the package should parse, post templating
	PackageSettings MemoryContainer `json:"package_settings"`
	Sequence        int             `json:"sequence"` // The order of the message (the order is per request id)
}

type DispatchCallResult struct {
	RequestID  uuid.UUID  `json:"request_id"`
	Successful bool       `json:"successful"` // Did the dispatch operation succeed
	Logs       []LogEntry `json:"logs"`
	Error      *Error     `json:"error"`
}

type DispatchHandler func(call *DispatchCall) (DispatchCallResult, error)

type RunnableDispatch struct {
	Handler     DispatchHandler `json:"-"`
	MockHandler DispatchHandler `json:"-"`

	Name          string `json:"name"`
	ID            string `json:"id"`
	Documentation string `json:"documentation"` // Markdown format
}

func (p *RunnablePackage) HandleDispatchExecute(c *gin.Context) {
	var input DispatchRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := DispatchResponse{[]DispatchCallResult{}}

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
	var input DispatchRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: better error messages and logging
		return
	}

	results := DispatchResponse{[]DispatchCallResult{}}

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
