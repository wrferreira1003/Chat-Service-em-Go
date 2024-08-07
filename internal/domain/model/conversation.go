package model

import "time"

// Conversation represents a conversation between users.
// swagger:model Conversation
type Conversation struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Participants []string  `json:"participants" bson:"participants"`
	LastMessage  *Message  `json:"last_message" bson:"last_message"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
