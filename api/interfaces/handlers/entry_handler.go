package handlers

import (
	"net/http"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/presenter"
	"github.com/Khigashiguchi/khigashiguchi.com/api/usecase"
	"github.com/gin-gonic/gin/json"
)

// Handler is interface of handling request.
type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type getEntriesHandler struct {
	UseCase usecase.GetEntriesUseCase
}

// GetEntriesHandler handle request GET /entries.
func (h *getEntriesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	entries, err := h.UseCase.Run()
	if err != nil {
		presenter.RespondErrorJson(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res := presenter.GetEntriesResponse{
		Entities: entries,
	}
	presenter.RespondJson(w, res, http.StatusOK)
}

// NewGetEntriesHandler create new handler of getting entries.
func NewGetEntriesHandler(db repository.DB) Handler {
	return &getEntriesHandler{UseCase: usecase.NewGetEntriesUseCase(db)}
}

type postEntriesHandler struct {
	UseCase usecase.PostEntriesUseCase
}

func (h *postEntriesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	entry := entity.Entry{}
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		presenter.RespondErrorJson(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	if err := h.UseCase.Run(entry); err != nil {
		presenter.RespondJson(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewPostEntriesHandler(db repository.Beginner) Handler {
	return &postEntriesHandler{UseCase: usecase.NewPostEntriesUseCase(db)}
}
