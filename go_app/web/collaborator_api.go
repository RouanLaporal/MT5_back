package web

import (
	"back_project/structure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (handler *Handler) GetCollaboratorByShop() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id_shop")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		collaborators, err := handler.Store.GetCollaboratorByShop(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(writer).Encode(collaborators)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) AddCollaborator() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		collaborator := structure.Collaborator{}
		json.NewDecoder(request.Body).Decode(&collaborator)

		id, err := handler.Store.CollaboratorStoreInterface.AddCollaborator(collaborator)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status          string `json:"status"`
			Message         string `json:"message"`
			NewCollaborator int    `json:"newCollaborator"`
		}{
			Status:          "success",
			Message:         "Nouveau collaborateur inséré avec succès",
			NewCollaborator: id,
		})
	}
}

func (handler *Handler) DeleteCollaborator() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.CollaboratorStoreInterface.DeleteCollaborator(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Collaborateur supprimé avec succès",
		})
	}
}

func (handler *Handler) UpdateCollaborator() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)
		collaborator := structure.Collaborator{}
		json.NewDecoder(request.Body).Decode(&collaborator)
		writer.Header().Set("Content-Type", "application/json")
		err := handler.Store.CollaboratorStoreInterface.UpdateCollaborator(id, collaborator)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Collaborateur modifié avec succès",
		})

	}
}
