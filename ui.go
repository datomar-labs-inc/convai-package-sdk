package convai_package_sdk

import (
	"io"
)

// NodeUIHandler should return a valid HTML page
type NodeUIHandler func(n *RunnableNode) (io.ReadCloser, error)

// LinkUIHandler should return a valid HTML page
type LinkUIHandler func(l *RunnableLink) (io.ReadCloser, error)

// SettingsUIHandler should return a valid HTML page
type SettingsUIHandler func(p *RunnablePackage) (io.ReadCloser, error)