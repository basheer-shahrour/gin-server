package entity

import "time"

type Conversation struct {
	Id          uint64    `json:"id" gorm:"primary_key; auto_increment"`
	ConnectorId uint64    `json:"-"`
	ReceiverId  uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default=CURRENT_TIMESTAMP"`
	Connector   User      `json:"connector" binding:"required" gorm:"foreignkey:ConnectorId"`
	Receiver    User      `json:"receiver" binding:"required" gorm:"foreignkey:ReceiverId"`
}
