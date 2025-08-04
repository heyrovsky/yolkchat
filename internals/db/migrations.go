package db

import (
	"github.com/heyrovsky/yolkchat/pkg/schema"
	"github.com/heyrovsky/yolkchat/pkg/users"
)

var (
	UserService users.Interface
)

func Init(path string) error {
	db, err := GetDbInstance(path)
	if err != nil {
		return err
	}
	db.AutoMigrate(&schema.UserList{}, &schema.ChatMessage{})

	userRepo := users.NewRepository(db)
	UserService = users.NewService(userRepo)

	return nil
}
