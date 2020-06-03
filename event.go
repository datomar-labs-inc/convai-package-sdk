package convai_package_sdk

import (
	ctypes "github.com/datomar-labs-inc/convai-types"
)

type RunnableEvent struct {
	Name          string           `json:"name"`
	ID            string           `json:"id"`
	Documentation string           `json:"documentation"` // Markdown format
	Style         ctypes.NodeStyle `json:"style"`
}
