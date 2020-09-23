package convai_package_sdk

import (
	"errors"
	"fmt"
	"io"

	ctypes "github.com/datomar-labs-inc/convai-types"
)

type RunnablePackage struct {
	Nodes      []RunnableNode     `json:"nodes"`
	Links      []RunnableLink     `json:"links"`
	Events     []RunnableEvent    `json:"events"`
	Dispatches []RunnableDispatch `json:"responders"`

	settingsUI SettingsUIHandler
	assets     AssetHandler
}

func (p *RunnablePackage) GetNodeUI(typeID, version string) (io.ReadCloser, error) {
	node := p.GetNode(typeID)

	if node == nil {
		return nil, errors.New("node did not exist")
	}

	return node.UIHandler(node)
}

func (p *RunnablePackage) GetLinkUI(typeID, version string) (io.ReadCloser, error) {
	link := p.GetLink(typeID)

	if link == nil {
		return nil, errors.New("link did not exist")
	}

	return link.UIHandler(link)
}

func (p *RunnablePackage) GetSettingsUI() (io.ReadCloser, error) {
	return p.settingsUI(p)
}

// AssetHandler should take a filename and return a reader for the file, and the mime type
type AssetHandler func(filename string) (io.ReadCloser, string, error)

func (p *RunnablePackage) GetManifest() *ctypes.Package {
	panic("implement me")
}

func (p *RunnablePackage) ExecuteNode(input *ctypes.NodeCall) (*ctypes.NodeCallResult, error) {
	node := p.GetNode(input.TypeID)

	if node == nil {
		return nil, errors.New("node did not exist")
	}

	result, err := node.Handler(input)
	return &result, err
}

func (p *RunnablePackage) ExecuteNodeMock(input *ctypes.NodeCall) (*ctypes.NodeCallResult, error) {
	node := p.GetNode(input.TypeID)

	if node == nil {
		return nil, errors.New("node did not exist")
	}

	result, err := node.MockHandler(input)
	return &result, err
}

func (p *RunnablePackage) ExecuteLink(request *ctypes.LinkExecutionRequest) (*ctypes.LinkExecutionResponse, error) {
	results := ctypes.LinkExecutionResponse{[]ctypes.LinkCallResult{}}

	for _, call := range request.Calls {
		l := p.GetLink(call.TypeID)

		if l == nil {
			return nil, errors.New("link did not exist")
		}

		result, err := l.Handler(&call)
		if err != nil {
			fmt.Println("Error was errored")
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	return &results, nil
}

func (p *RunnablePackage) ExecuteLinkMock(request *ctypes.LinkExecutionRequest) (*ctypes.LinkExecutionResponse, error) {
	results := ctypes.LinkExecutionResponse{[]ctypes.LinkCallResult{}}

	for _, call := range request.Calls {
		l := p.GetLink(call.TypeID)

		if l == nil {
			return nil, errors.New("link did not exist")
		}

		result, err := l.MockHandler(&call)
		if err != nil {
			fmt.Println("Error was errored")
			// TODO handle this appropriately
		}

		results.Results = append(results.Results, result)
	}

	return &results, nil
}

func (p *RunnablePackage) Dispatch(request *ctypes.DispatchRequest) (*ctypes.DispatchResponse, error) {
	results := ctypes.DispatchResponse{[]ctypes.DispatchCallResult{}}

	for _, call := range request.Dispatches {
		d := p.GetDispatch(call.ID)

		// No dispatch, stop and let convai know
		if d == nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    ctypes.ErrDispatchNotFound,
					Message: "dispatch missing",
				},
			}

			results.Results = append(results.Results, errResp)
			continue
		}

		// TODO determine level of error exposure
		// Handler didn't function
		result, err := d.Handler(&call)
		if err != nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    ctypes.ErrHandlerFailure,
					Message: fmt.Sprintf("handler failed: %s", err.Error()),
				},
			}
			results.Results = append(results.Results, errResp)
			continue
		}

		results.Results = append(results.Results, result)
	}

	return &results, nil
}

func (p *RunnablePackage) DispatchMock(request *ctypes.DispatchRequest) (*ctypes.DispatchResponse, error) {
	results := ctypes.DispatchResponse{[]ctypes.DispatchCallResult{}}

	for _, call := range request.Dispatches {
		d := p.GetDispatch(call.ID)

		// No dispatch, stop and let convai know
		if d == nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    ctypes.ErrDispatchNotFound,
					Message: "dispatch missing",
				},
			}

			results.Results = append(results.Results, errResp)
			continue
		}

		// TODO determine level of error exposure
		// Handler didn't function
		result, err := d.MockHandler(&call)
		if err != nil {
			errResp := ctypes.DispatchCallResult{
				Successful: false,
				Error: &ctypes.Error{
					Code:    ctypes.ErrHandlerFailure,
					Message: fmt.Sprintf("handler failed: %s", err.Error()),
				},
			}
			results.Results = append(results.Results, errResp)
			continue
		}

		results.Results = append(results.Results, result)
	}

	return &results, nil
}

func (p *RunnablePackage) GetAsset(filename string) (io.ReadCloser, string, error) {
	return p.assets(filename)
}

func NewPackage() *RunnablePackage {
	return &RunnablePackage{
		Nodes:      []RunnableNode{},
		Links:      []RunnableLink{},
		Events:     []RunnableEvent{},
		Dispatches: []RunnableDispatch{},
	}
}

func (p *RunnablePackage) AddNode(node RunnableNode) {
	p.Nodes = append(p.Nodes, node)
}

func (p *RunnablePackage) AddLink(link RunnableLink) {
	p.Links = append(p.Links, link)
}

func (p *RunnablePackage) AddEvent(event RunnableEvent) {
	p.Events = append(p.Events, event)
}

func (p *RunnablePackage) AddDispatch(dispatch RunnableDispatch) {
	p.Dispatches = append(p.Dispatches, dispatch)
}

func (p *RunnablePackage) SetAssetHandler(handler AssetHandler) {
	p.assets = handler
}

func (p *RunnablePackage) SetSettingsUIHandler(handler SettingsUIHandler) {
	p.settingsUI = handler
}

func (p *RunnablePackage) GetNode(id string) *RunnableNode {
	for _, n := range p.Nodes {
		if n.TypeID == id {
			return &n
		}
	}

	return nil
}

func (p *RunnablePackage) GetLink(id string) *RunnableLink {
	for _, l := range p.Links {
		if l.TypeID == id {
			return &l
		}
	}

	return nil
}

func (p *RunnablePackage) GetDispatch(id string) *RunnableDispatch {
	for _, r := range p.Dispatches {
		if r.ID == id {
			return &r
		}
	}

	return nil
}
