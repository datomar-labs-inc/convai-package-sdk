package convai_package_sdk

import (
	ctypes "github.com/datomar-labs-inc/convai-types"
)

type LinkExecHandler func(call *ctypes.LinkCall) (ctypes.LinkCallResult, error)

type RunnableLink struct {
	ctypes.PackageLink

	Handler     LinkExecHandler `json:"-"`
	MockHandler LinkExecHandler `json:"-"`
	UIHandler   LinkUIHandler   `json:"-"`
}
