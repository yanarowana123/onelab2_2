package handler

import (
	"net/http"
)

func (h *Manager) SecretKeyValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Secret-Key")
		if !h.service.Auth.IsGranted(key) {
			h.respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)

	}
}
