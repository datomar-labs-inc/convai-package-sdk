package service

// UIHandler should return a HTML string to render the UI
type UIHandler func() (string, error)