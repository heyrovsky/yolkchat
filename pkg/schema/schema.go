package schema

import (
	"time"
)

type UserList struct {
	Name     string
	LastSeen time.Time
	Online   bool
}
type ChatSchema struct {
	Received  bool
	Type      string // "text", "video", "image", "file"
	Content   string // if text, then the string; else path (file://)
	Timestamp time.Time
}

type ChatMessage struct {
	Ingress       bool
	Text          string
	Attatchements []string
	Timestamp     time.Time
}
