package main

import (
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	FileDirectory := http.Dir("./ui/build/")
	FileHandler := http.StripPrefix("/", http.FileServer(FileDirectory))
	r.PathPrefix("/static").Handler(FileHandler).Methods("GET")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/verify", handlers.Verify).Methods("POST")
	r.HandleFunc("/api/popular", handlers.PopBooks).Methods("POST")
	r.HandleFunc("/api/discount", handlers.DiscBooks).Methods("POST")
	r.HandleFunc("/api/search", handlers.Search).Methods("GET")
	r.HandleFunc("/api/book", handlers.Book).Methods("POST")
	r.HandleFunc("/api/createOrder", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/api/deleteOrder", handlers.DeleteOrder).Methods("POST")
	r.HandleFunc("/api/addFavourite", handlers.AddFavourite).Methods("POST")
	r.HandleFunc("/api/deleteFavourite", handlers.DeleteFavourite).Methods("POST")
	r.HandleFunc("/api/getFavourites", handlers.GetFavourites).Methods("POST")
	r.HandleFunc("/api/redactOrder", handlers.UpdateOrder).Methods("POST")
	r.HandleFunc("/api/refresh", handlers.Refresh).Methods("POST")
	r.HandleFunc("/api/cookie", handlers.CookieLogin).Methods("POST")
	r.HandleFunc("/api/filters", handlers.GetFilters).Methods("POST")

	r.HandleFunc("/api/emp/deleteBook", handlers.DeleteBook).Methods("POST")
	r.HandleFunc("/api/emp/addBook", handlers.AddBook).Methods("POST")
	r.HandleFunc("/api/emp/redactBook", handlers.RedactBook).Methods("POST")

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS", "DELETE"},
	})
	handler := cors.Handler(r)
	fmt.Println(http.ListenAndServe(":83", handler).Error())
}
