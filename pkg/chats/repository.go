package chats

import (
	"github.com/heyrovsky/yolkchat/pkg/schema"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Interface {
	return &repo{
		DB: db,
	}
}

func (r *repo) AddChat(message schema.ChatMessage) error {
	panic(0)
}
