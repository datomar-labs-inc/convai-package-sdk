package convaipkgsdk

import (
	"github.com/google/uuid"
)

type Package struct {
	Nodes     []PackageNode     `json:"nodes"`
	Links     []PackageLink     `json:"links"`
	Events    []PackageEvent    `json:"events"`
	Dispatchs []PackageDispatch `json:"responders"`
	Settings  RunnableSettings  `json:"settings"`
}

type PackageDescription struct {
	Name          string `json:"name"`
	BaseURL       string `json:"base_url"`
	Documentation string `json:"documentation"`
}

type PackageNode struct {
	Name          string    `json:"name"`
	ID            uuid.UUID `json:"id"`
	Version       string    `json:"version"` // Valid semantic version
	Style         NodeStyle `json:"style"`
	Documentation string    `json:"documentation"` // Markdown format
}

type NodeStyle struct {
	Color string   `json:"color"` // Valid hex code color
	Icons []string `json:"icons"` // File name (files will be served in a special format by the plugin)
}

type PackageLink struct {
	Name          string    `json:"name"`
	ID            uuid.UUID `json:"id"`
	Version       string    `json:"version"` // Valid semantic version
	Style         LinkStyle `json:"style"`
	Documentation string    `json:"documentation"` // Markdown format
}

type LinkStyle struct {
	Color string   `json:"color"` // Valid hex code color
	Icons []string `json:"icons"` // File name (files will be served in a special format by the plugin)
}

type PackageEvent struct {
	Name          string    `json:"name"`
	ID            uuid.UUID `json:"id"`
	Documentation string    `json:"documentation"` // Markdown format
	Style         NodeStyle `json:"style"`
}

type PackageDispatch struct {
	Name          string    `json:"name"`
	ID            uuid.UUID `json:"id"`
	Documentation string    `json:"documentation"` // Markdown format
}

type DispatchStyle struct {
	Color string `json:"color"` // Valid hex code color
	Icon  string `json:"icon"`  // File name (files will be served in a special format by the plugin)
}

type PackageModule struct {
}

type PackageTemplates struct {
}
