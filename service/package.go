package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RunnablePackage struct {
	router *gin.Engine

	Nodes      []RunnableNode     `json:"nodes"`
	Links      []RunnableLink     `json:"links"`
	Events     []PackageEvent     `json:"events"`
	Dispatches []RunnableDispatch `json:"responders"`
	Settings   RunnableSettings   `json:"settings"`
	Assets     AssetHandler       `json:"-"`
}

func (p *RunnablePackage) GetRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/nodes/execute", p.HandleNodeExecute)
	r.POST("/nodes/execute-mock", p.HandleNodeExecuteMock)
	r.GET("/nodes/:nid/ui", p.HandleNodeUI)

	r.POST("/links/execute", p.HandleLinkExecute)
	r.POST("/links/execute-mock", p.HandleLinkExecuteMock)
	r.GET("/links/:lid/ui", p.HandleLinkUI)

	r.POST("/dispatch/execute", p.HandleDispatchExecute)
	r.POST("/dispatch/execute-mock", p.HandleDispatchExecuteMock)

	r.GET("/settings/ui", p.HandleSettingsUI)

	r.GET("/assets/:filename", p.HandleAssetRequest)

	return r
}

func (p *RunnablePackage) GetNode(id uuid.UUID) *RunnableNode {
	for _, n := range p.Nodes {
		if n.ID == id {
			return &n
		}
	}

	return nil
}

func (p *RunnablePackage) GetLink(id uuid.UUID) *RunnableLink {
	for _, l := range p.Links {
		if l.ID == id {
			return &l
		}
	}

	return nil
}

func (p *RunnablePackage) GetDispatch(id uuid.UUID) *RunnableDispatch {
	for _, r := range p.Dispatches {
		if r.ID == id {
			return &r
		}
	}

	return nil
}
