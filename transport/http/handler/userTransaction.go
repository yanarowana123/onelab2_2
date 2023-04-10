package handler

import (
	"encoding/json"
	"github.com/yanarowana123/onelab2_1/internal/models"
	"net/http"
)

func (h *Manager) CreateUserTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var userTransaction models.CreateUserTransactionRequest
		json.NewDecoder(r.Body).Decode(&userTransaction)

		err := h.service.UserTransaction.CreateTransaction(r.Context(), userTransaction)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
