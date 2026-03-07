package router

import (
    "github.com/gorilla/mux"
    "github.com/Pavel90544/VSRPP/lab4/internal/handlers"
    "github.com/Pavel90544/VSRPP/lab4/internal/repository"
)

func SetupRoutes(repo *repository.ProductRepository) *mux.Router {
    router := mux.NewRouter()
    handler := handlers.NewProductHandler(repo)

    api := router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/products", handler.CreateProduct).Methods("POST")
    api.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
    api.HandleFunc("/products/{id}", handler.GetProduct).Methods("GET")
    api.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
    api.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")

    return router
}
