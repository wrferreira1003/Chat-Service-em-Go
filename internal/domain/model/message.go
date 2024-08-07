package model

import "time"

// Message represents a message exchanged between users.
// swagger:model Message
type Message struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	SenderID       string    `json:"sender_id" bson:"sender_id"`
	ReceiverID     string    `json:"receiver_id" bson:"receiver_id"`
	Content        string    `json:"content" bson:"content"`
	ConversationID string    `json:"conversation_id" bson:"conversation_id"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
}
