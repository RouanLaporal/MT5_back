package web

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetKind() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")
		kinds, err := h.Store.GetAllKind()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(kinds)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
