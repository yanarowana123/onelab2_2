package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2_1/internal/models"
	"net/http"
)

// CreateTransaction
// @Summary Transaction
// @Description Transaction
// @Security ServerKeyAuth
// @Param transaction body models.CreateTransactionRequest true "body"
// @Success 201
// @Failure 400 "validation error"
// @Failure 500 "internal server error"
// @Router /transaction [post]
func (h *Manager) CreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var createTransactionRequest models.CreateTransactionRequest
		err := json.NewDecoder(r.Body).Decode(&createTransactionRequest)
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		err = h.validate.Struct(createTransactionRequest)
		if err != nil {
			h.respondWithErrorList(w, http.StatusBadRequest, err)
			return
		}

		err = h.service.Transaction.Create(r.Context(), createTransactionRequest)

		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// GetMoneySumByBookID
// @Summary Transaction
// @Description Transaction
// @Security ServerKeyAuth
// @Param bookID path string true "bookID uuid"
// @Success 200 {number} number "book sum"
// @Failure 400 "validation error"
// @Failure 500 "internal server error"
// @Router /sum/book/{bookID} [get]
func (h *Manager) GetMoneySumByBookID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		bookID, err := uuid.Parse(params["bookID"])
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		sum, err := h.service.Transaction.GetSumByBookID(r.Context(), bookID)
		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(w).Encode(sum)
	}
}
