package convai_package_sdk

import (
	ctypes "github.com/datomar-labs-inc/convai-types"
)

type NodeExecHandler func(call *ctypes.NodeCall) (ctypes.NodeCallResult, error)

// Node stuff
// POST /nodes/:nid/execute
// POST /nodes/:nid/execute-mock
type RunnableNode struct {
	ctypes.PackageNode

	Handler     NodeExecHandler `json:"-"`
	MockHandler NodeExecHandler `json:"-"`
	UIHandler   NodeUIHandler   `json:"-"`
}
