package handler_v1

import (
	"books/app/domain/usecase"
	"books/app/library/telemetry"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"books/app/domain/dto"
)

type BookHandler struct {
	u usecase.BookUseCase
}

// NewBookHandler returns new BookHandler.
func NewBookHandler(u usecase.BookUseCase) *BookHandler {
	return &BookHandler{u: u}
}

func (h *BookHandler) RegisterRoutes(r chi.Router) {
	r.Route("/books", func(r chi.Router) {
		r.Post("/", h.CreateBook)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	spanCtx, span := telemetry.Tracer.Start(r.Context(), "/handler/create-book")
	defer span.End()

	var bookDto dto.CreateBookRequest
	if err := render.DecodeJSON(r.Body, &bookDto); err != nil {
		http.Error(w, "Error to parse request", http.StatusUnprocessableEntity)
		return
	}

	id, err := h.u.CreateBook(spanCtx, bookDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, &dto.BookResponse{
		ID: id,
	})
}
