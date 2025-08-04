package db

import "github.com/heyrovsky/yolkchat/pkg/schema"

var ()

func Init(path string) error {
	db, err := GetDbInstance(path)
	if err != nil {
		return err
	}
	db.AutoMigrate(&schema.UserList{}, &schema.ChatMessage{})

	return nil
}
