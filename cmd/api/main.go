package main

import (
	"books/app/config/bootstrap"
	"books/app/domain/usecase"
	handler "books/app/gateway/http/handler"
	handler_v1 "books/app/gateway/http/handler/v1"
	"books/app/gateway/postgres"
	"books/app/library/telemetry"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	context := context.Background()

	shutDownTracer := telemetry.InitTracer(context, "books-api")
	defer shutDownTracer()

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

	if strings.EqualFold(os.Getenv("DEBUG"), "true") {
		handler.RegisterDebugRoutes(router)
	}	

	wrappedHandler := otelhttp.NewHandler(router, "http-server")

	log.Println("🚀 Server running on port 5000")
	if err := http.ListenAndServe("0.0.0.0:5000", wrappedHandler); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
