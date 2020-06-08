package convai_package_sdk

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RunnablePackage struct {
	router *gin.Engine `json:"-"`

	Nodes      []RunnableNode     `json:"nodes"`
	Links      []RunnableLink     `json:"links"`
	Events     []RunnableEvent    `json:"events"`
	Dispatches []RunnableDispatch `json:"responders"`
	Settings   RunnableSettings   `json:"settings"`
	Assets     AssetHandler       `json:"-"`
}

func NewPackage() *RunnablePackage {
	return &RunnablePackage{
		Nodes:      []RunnableNode{},
		Links:      []RunnableLink{},
		Events:     []RunnableEvent{},
		Dispatches: []RunnableDispatch{},
		Settings:   RunnableSettings{},
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

func (p *RunnablePackage) SetSettings(settings RunnableSettings) {
	p.Settings = settings
}

func (p *RunnablePackage) SetAssets(handler AssetHandler) {
	p.Assets = handler
}

func (p *RunnablePackage) GetRouter(signingKey string) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/manifest", p.HManifest)
	r.GET("/assets/:filename", p.HandleAssetRequest)
	r.GET("/settings/ui", p.HandleSettingsUI)
	r.GET("/links/:lid/ui", p.HandleLinkUI)
	r.GET("/nodes/:nid/ui", p.HandleNodeUI)

	r.Use(signatureVerificationMiddleware(signingKey))

	r.POST("/nodes/execute", p.HandleNodeExecute)
	r.POST("/nodes/execute-mock", p.HandleNodeExecuteMock)

	r.POST("/links/execute", p.HandleLinkExecute)
	r.POST("/links/execute-mock", p.HandleLinkExecuteMock)

	r.POST("/dispatch/execute", p.HandleDispatchExecute)
	r.POST("/dispatch/execute-mock", p.HandleDispatchExecuteMock)

	return r
}

func (p *RunnablePackage) GetNode(id string) *RunnableNode {
	for _, n := range p.Nodes {
		if n.ID == id {
			return &n
		}
	}

	return nil
}

func (p *RunnablePackage) GetLink(id string) *RunnableLink {
	for _, l := range p.Links {
		if l.ID == id {
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
