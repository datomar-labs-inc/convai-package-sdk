package convai_package_sdk

import (
	ctypes "github.com/datomar-labs-inc/convai-types"
)

type DispatchHandler func(call *ctypes.DispatchCall) (ctypes.DispatchCallResult, error)

type RunnableDispatch struct {
	ctypes.PackageDispatch

	Handler     DispatchHandler `json:"-"`
	MockHandler DispatchHandler `json:"-"`
}
