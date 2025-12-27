package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// MODELS
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Event struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	EventDate       time.Time `json:"event_date"`
	Location        string    `json:"location"`
	TotalTicket     int       `json:"total_ticket"`
	AvailableTicket int       `json:"available_ticket"`
	OrganizerID     int       `json:"organizer_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Transaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	EventID   int       `json:"event_id"`
	Quantity  int       `json:"quantity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var users []User
var events []Event
var transactions []Transaction

// USER HANDLER
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = len(users) + 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// EVENT HANDLER
func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event Event
	json.NewDecoder(r.Body).Decode(&event)

	event.ID = len(events) + 1
	event.AvailableTicket = event.TotalTicket
	event.Status = "PUBLISHED"
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	events = append(events, event)
	json.NewEncoder(w).Encode(event)
}

// TRANSACTION HANDLER

func createTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var trs Transaction
	json.NewDecoder(r.Body).Decode(&trs)

	if trs.Quantity <= 0 {
		http.Error(w, "Quantity must be greater than 0", http.StatusBadRequest)
		return
	}

	for i := range events {
		if events[i].ID == trs.EventID {

			// Cek stok tiket
			if events[i].AvailableTicket < trs.Quantity {
				http.Error(w, "Ticket not available", http.StatusConflict)
				return
			}

			// Kurangi stok
			events[i].AvailableTicket -= trs.Quantity
			events[i].UpdatedAt = time.Now()

			// Simpan transaksi
			trs.ID = len(transactions) + 1
			trs.Status = "PAID"
			trs.CreatedAt = time.Now()
			trs.UpdatedAt = time.Now()

			transactions = append(transactions, trs)

			json.NewEncoder(w).Encode(trs)
			return
		}
	}

	http.Error(w, "Event not found", http.StatusNotFound)
}

// MAIN
func main() {
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/events", createEvent)
	http.HandleFunc("/transactions", createTransaction)

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
