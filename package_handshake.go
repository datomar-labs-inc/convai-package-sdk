package convai_package_sdk

import (
	"github.com/google/uuid"
)

type NodeStyle struct {
	Color string   `json:"color"` // Valid hex code color
	Icons []string `json:"icons"` // File name (files will be served in a special format by the plugin)
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

type DispatchStyle struct {
	Color string `json:"color"` // Valid hex code color
	Icon  string `json:"icon"`  // File name (files will be served in a special format by the plugin)
}

type PackageModule struct {
}

type PackageTemplates struct {
}
