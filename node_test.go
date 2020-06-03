package convai_package_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	ctypes "github.com/datomar-labs-inc/convai-types"
	"github.com/google/uuid"
)

func TestNodeExecuteRoute(t *testing.T) {
	nodeID := uuid.Must(uuid.NewRandom()).String()

	node := RunnableNode{
		Handler: func(call *ctypes.NodeCall) (result ctypes.NodeCallResult, err error) {
			ncr := ctypes.NodeCallResult{
				RequestID: call.RequestID,
			}

			return ncr, nil
		},

		PackageNode: ctypes.PackageNode{
			Name:    "TestNode",
			ID:      nodeID,
			Version: "0.1.0",
		},
	}

	p := RunnablePackage{
		Nodes:    []RunnableNode{node},
		Settings: RunnableSettings{},
	}

	w := httptest.NewRecorder()

	reqID := uuid.Must(uuid.NewRandom())

	req, err := http.NewRequest("POST", "/nodes/execute", bytes.NewReader(mustJSONify(
		ctypes.NodeExecutionRequest{
			Calls: []ctypes.NodeCall{
				{
					RequestID: reqID,
					ID:        nodeID,
					Version:   "0.1.0",
				},
			},
		})))
	if err != nil {
		t.Error(err)
		return
	}

	r := p.GetRouter()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error(fmt.Sprintf("expected status code 200, got %d", w.Code))
	}

	expectedBody := mustJSONify(ctypes.NodeExecutionResponse{
		[]ctypes.NodeCallResult{
			{
				RequestID: reqID,
			},
		},
	})

	if w.Body.String() != string(expectedBody) {
		t.Errorf("invalid body return, expected\n%s, got \n%s", string(expectedBody), w.Body.String())
	}
}

func mustJSONify(in interface{}) []byte {
	res, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	return res
}
