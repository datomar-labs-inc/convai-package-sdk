package convai_package_sdk

import (
	"io"

	ctypes "github.com/datomar-labs-inc/convai-types"
)

type IPackageProvider interface {
	GetManifest() *ctypes.Package
	ExecuteNode(input *ctypes.NodeCall) (*ctypes.NodeCallResult, error)
	ExecuteNodeMock(input *ctypes.NodeCall) (*ctypes.NodeCallResult, error)
	ExecuteLink(request *ctypes.LinkExecutionRequest) (*ctypes.LinkExecutionResponse, error)
	ExecuteLinkMock(request *ctypes.LinkExecutionRequest) (*ctypes.LinkExecutionResponse, error)
	Dispatch(request *ctypes.DispatchRequest) (*ctypes.DispatchResponse, error)
	DispatchMock(request *ctypes.DispatchRequest) (*ctypes.DispatchResponse, error)
	GetAsset(filename string) (io.Reader, error)
}
