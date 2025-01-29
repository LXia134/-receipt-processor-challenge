package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"Receipt-Processor-Challenge/internal/models"
	"Receipt-Processor-Challenge/internal/storage"
	"Receipt-Processor-Challenge/internal/calculator"
	"github.com/google/uuid"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		log.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	calculatedPoints := calculator.CalculatePoints(receipt)

	// Store both receipt and points
	storage.SaveReceipt(id, receipt, calculatedPoints)

	log.Printf("Receipt processed: ID=%s, Points=%d\n", id, calculatedPoints)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/receipts/"), "/")
	if len(pathParts) < 2 || pathParts[1] != "points" {
		log.Println("Invalid URL format for fetching points")
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	id := pathParts[0]
	log.Printf("Fetching points for receipt ID: %s\n", id)

	points, found := storage.GetPoints(id)
	if !found {
		log.Printf("No receipt found for ID: %s\n", id)
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	log.Printf("Points found for receipt ID: %s -> Points: %d\n", id, points)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
