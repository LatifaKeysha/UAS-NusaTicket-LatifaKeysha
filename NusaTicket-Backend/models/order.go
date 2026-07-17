package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     string             `json:"user_id" bson:"user_id"`
	TicketID   string             `json:"ticket_id" bson:"ticket_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	TotalPrice float64            `json:"total_price" bson:"total_price"`
}
