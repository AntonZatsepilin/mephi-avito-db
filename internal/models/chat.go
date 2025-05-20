package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Messages []Message
	Users    []User    `gorm:"many2many:user_chats;constraint:OnDelete:CASCADE"`
}
