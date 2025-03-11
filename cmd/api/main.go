package main

import (
	"books/app/config/bootstrap"
	"books/app/domain/usecase"
	handler_v1 "books/app/gateway/http/handler/v1"
	"books/app/gateway/postgres"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main()  {
	db, err := bootstrap.InitDB()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer db.Close()

	repo := postgres.NewBookRepository(db)
	uc := usecase.NewBookUseCase(repo)
	h := handler_v1.NewBookHandler(uc)


	router := chi.NewRouter()
	h.RegisterRoutes(router)

	http.ListenAndServe(":5000", router)
}