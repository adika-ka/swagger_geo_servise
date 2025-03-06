package main

import (
	"fmt"
	"log"
	"net/http"

	_ "swagger_geo_service/docs"
	"swagger_geo_service/internal/handlers"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Geo Service API
// @version 1.0
// @description	API для работы с адресами и геокодингом
// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()

	r.Post("/api/address/search", handlers.SearchHandler)
	r.Post("/api/address/geocode", handlers.GeocodeHandler)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
