package main

import (
	"log"
	"net/http"
	"nusaticket-backend/config"
	"nusaticket-backend/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Initialize Router
	r := mux.NewRouter()

	// Routes for Users
	r.HandleFunc("/api/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/login", controllers.LoginUser).Methods("POST")

	// Routes for Tickets
	r.HandleFunc("/api/tickets", controllers.GetTickets).Methods("GET")
	r.HandleFunc("/api/tickets", controllers.CreateTicket).Methods("POST")

	// Routes for Orders
	r.HandleFunc("/api/orders", controllers.CreateOrder).Methods("POST")

	// Start Server
	log.Println("Server NusaTicket berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
