package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	rows, err := h.DB.Query("SELECT * FROM listings order by created_at desc limit 10")
	if err != nil {
		http.Error(w, "Error fetching listings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	listings := []listing{}
	for rows.Next() {
		var l listing
		err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.CreatedAt)
		if err != nil {
			http.Error(w, "Error scanning listing", http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error occurred while iterating rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := h.DB.Exec("DELETE FROM listings WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Error deleting listing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
