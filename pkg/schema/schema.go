package schema

import (
	"time"

	"gorm.io/datatypes"
)

type UserList struct {
	Name     string        `gorm:"primaryKey;column:name;size:100;not null"`
	LastSeen time.Time     `gorm:"column:last_seen;not null"`
	Online   bool          `gorm:"column:online;not null"`
	Messages []ChatMessage `gorm:"foreignKey:UserName;references:Name"`
}

type ChatMessage struct {
	ID          uint                        `gorm:"primaryKey;autoIncrement;column:id"`
	UserName    string                      `gorm:"column:user_name;size:100;not null;index"`
	Ingress     bool                        `gorm:"column:ingress;not null"`
	Text        string                      `gorm:"column:text;type:text;not null"`
	Attachments datatypes.JSONSlice[string] `gorm:"column:attachments;type:json"`
	Timestamp   time.Time                   `gorm:"column:timestamp;not null"`
}
