package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"nusaticket-backend/config"
	"nusaticket-backend/models"
	"time"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	_ = json.NewDecoder(r.Body).Decode(&order)

	collection := config.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
