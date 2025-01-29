package main

import (
	"log"
	"net/http"

	"Receipt-Processor-Challenge/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/receipts/process", handlers.ProcessReceipt)
	mux.HandleFunc("/receipts/", handlers.GetPoints)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Receipt Processor API"))
	})
	
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
