package main

import (
	"net/http"

	"github.com/Ba1vo/Proektirovanie/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	FileDirectory := http.Dir("./assets/")
	FileHandler := http.StripPrefix("/assets/", http.FileServer(FileDirectory))
	r.PathPrefix("/api/assets/").Handler(FileHandler).Methods("GET")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/verify", handlers.Verify).Methods("POST")
	r.HandleFunc("/api/main", handlers.Main).Methods("GET")
	r.HandleFunc("/api/search", handlers.Search).Methods("POST")
	r.HandleFunc("/api/book", handlers.Book).Methods("GET")
	r.HandleFunc("/api/createOrder", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/api/confirmOrder", handlers.ConfirmOrder).Methods("POST")
	r.HandleFunc("/api/addBook", handlers.AddBook).Methods("POST")
	r.HandleFunc("/api/redactBook", handlers.RedactBook).Methods("POST")
	r.HandleFunc("/api/deleteBook", handlers.DeleteBook).Methods("DELETE")
	http.ListenAndServe(":2406", r)
}
