package models

import (
	"chat/config"

	"chat/utils/errors"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Message    string   `json:"Message"`
	UserId     uint     `json:"UserId" gorm:"index"`
	User       User     `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	ChatRoomId uint     `json:"ChatRoomId" gorm:"index"`
	ChatRoom   ChatRoom `json:"ChatRoom" gorm:"foreignKey:ChatRoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Upvotes    uint     `json:"Upvotes" gorm:"index"`
	Downvotes  uint     `json:"Downvotes" gorm:"index"`
}

func (cr *Chat) List(roomId uint, cht *[]Chat) *gorm.DB {
	db := config.GetDB()
	db = db.Where(Chat{ChatRoomId: roomId}).Preload("ChatRoom").Preload("User").Find(&cht).Limit(50)
	return db
}

func (c *Chat) Add() *gorm.DB {
	db := config.GetDB()
	db = db.Where(c).Create(&c)
	return db
}

func (u *Chat) UpdateMessageUpvotes(messageID uint) *gorm.DB {
	db := config.GetDB()
	err := db.Exec("UPDATE chats SET upvotes = upvotes + 1 WHERE id = $1", messageID)
	errors.DBErrorCheck(err)
	db = db.Where("id=?", messageID).Find(&u)
	return db
}

func (u *Chat) UpdateMessageDownvotes(messageID uint) *gorm.DB {
	db := config.GetDB()
	err := db.Exec("UPDATE chats SET downvotes = downvotes + 1 WHERE id = $1", messageID)
	errors.DBErrorCheck(err)
	db = db.Where("id=?", messageID).Find(&u)
	return db
}
