package repository

import (
	"github.com/basheer-shahrour/gin-server/entity"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	SaveConversation(conversation entity.Conversation)
	UpdateConversation(conversation entity.Conversation)
	DeleteConversation(conversation entity.Conversation)
	FindAllConversations() []entity.Conversation
}

type conversationRepository struct {
	dbConnection *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepository {
	return &conversationRepository{dbConnection: db}
}

func (db *conversationRepository) SaveConversation(Conversation entity.Conversation) {
	db.dbConnection.Create(&Conversation)
}

func (db *conversationRepository) UpdateConversation(Conversation entity.Conversation) {
	db.dbConnection.Save(&Conversation)
}

func (db *conversationRepository) DeleteConversation(Conversation entity.Conversation) {
	db.dbConnection.Delete(&Conversation)
}

func (db *conversationRepository) FindAllConversations() []entity.Conversation {
	var Conversations []entity.Conversation
	db.dbConnection.Set("gorm:auto_preload", true).Find(&Conversations)
	return Conversations
}
