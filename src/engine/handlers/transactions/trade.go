package transactions

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type TradeRequest struct {
	Spent         string `json:"spent"`
	SpentAmount   int64  `json:"spent_amount"`
	ReceiveKey    string `json:"receive"`
	ReceiveAmount int64  `json:"receive_amount"`
}

type TradeResponse struct {
	Message string      `json:"message"`
	Trade   TradeResult `json:"trade"`
}

func TradeResources(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		var req TradeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if !models.IsValidResourceKey(req.Spent) || !models.IsValidResourceKey(req.ReceiveKey) {
			http.Error(w, "invalid resource key", http.StatusBadRequest)
			return
		}

		if req.SpentAmount <= 0 || req.ReceiveAmount <= 0 {
			http.Error(w, "amounts must be greater than zero", http.StatusBadRequest)
			return
		}

		if req.Spent == req.ReceiveKey {
			http.Error(w, "pay and receive resource must differ", http.StatusBadRequest)
			return
		}

		result, err := Trade(db, userID, req.Spent, req.SpentAmount, req.ReceiveKey, req.ReceiveAmount)
		if err != nil {
			if errors.Is(err, ErrInsufficientResources) {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			http.Error(w, "failed to apply transaction", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(TradeResponse{
			Message: "trade successful",
			Trade:   result,
		})
	})
}
