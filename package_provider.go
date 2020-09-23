package convai_package_sdk

import (
	"io"

	ctypes "github.com/datomar-labs-inc/convai-types"
)

type IPackageProvider interface {
	GetManifest() *ctypes.Package
	ExecuteNode(input *ctypes.NodeCall) (*ctypes.NodeCallResult, error)
	ExecuteNodeMock(input *ctypes.NodeCall) (*ctypes.NodeCallResult, error)
	GetNodeUI(typeID, version string) (io.ReadCloser, error)
	ExecuteLink(request *ctypes.LinkExecutionRequest) (*ctypes.LinkExecutionResponse, error)
	ExecuteLinkMock(request *ctypes.LinkExecutionRequest) (*ctypes.LinkExecutionResponse, error)
	GetLinkUI(typeID, version string) (io.ReadCloser, error)
	Dispatch(request *ctypes.DispatchRequest) (*ctypes.DispatchResponse, error)
	DispatchMock(request *ctypes.DispatchRequest) (*ctypes.DispatchResponse, error)
	GetSettingsUI() (io.ReadCloser, error)
	GetAsset(filename string) (io.ReadCloser, string, error)
}
