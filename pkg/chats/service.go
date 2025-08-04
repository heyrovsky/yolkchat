package chats

import "github.com/heyrovsky/yolkchat/pkg/schema"

type Service struct {
	repo Interface
}

func NewService(r Interface) Interface {
	return &Service{
		repo: r,
	}
}

func (s *Service) AddChat(message schema.ChatMessage) error {

}
