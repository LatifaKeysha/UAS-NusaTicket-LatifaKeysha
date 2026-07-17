package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"nusaticket-backend/config"
	"nusaticket-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	_ = json.NewDecoder(r.Body).Decode(&ticket)

	collection := config.DB.Collection("tickets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("tickets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var tickets []models.Ticket
	for cursor.Next(ctx) {
		var ticket models.Ticket
		cursor.Decode(&ticket)
		tickets = append(tickets, ticket)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}
