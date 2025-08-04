package chats

import "github.com/heyrovsky/yolkchat/pkg/schema"

type Interface interface {
	AddChat(s schema.ChatMessage) error
}
