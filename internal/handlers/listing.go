package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/fahreyad/go_api/internal/httpx"
	"github.com/fahreyad/go_api/internal/middleware"
)

type Handler struct {
	DB     *sql.DB
	Logger *slog.Logger
}

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewHandler(db *sql.DB, logger *slog.Logger) *Handler {
	return &Handler{DB: db, Logger: logger}

}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetRequestID(r.Context())
	rows, err := h.DB.QueryContext(r.Context(), "SELECT * FROM listings order by created_at desc limit 10")
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error fetching listings", httpx.INTERNAL_ERROR)
		return
	}
	defer rows.Close()
	listings := []listing{}
	for rows.Next() {
		var l listing
		err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.CreatedAt)
		if err != nil {
			h.Logger.Error("Error scanning listing", "error", err)
			http.Error(w, "Error scanning listing", http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}
	h.Logger.Info("Fetched listings", "count", len(listings), "request_id", requestID)
	if err := rows.Err(); err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error occurred while iterating rows", httpx.INTERNAL_ERROR)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error encoding JSON", httpx.INTERNAL_ERROR)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	requestID := middleware.GetRequestID(ctx)
	_, err := h.DB.ExecContext(ctx, "DELETE FROM listing WHERE id = $1", id)
	if err != nil {
		h.Logger.Error("Error deleting listing with id", "id", id, "request_id", requestID, "error", err)
		httpx.Error(w, http.StatusInternalServerError, "Error deleting listing", httpx.INTERNAL_ERROR)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
