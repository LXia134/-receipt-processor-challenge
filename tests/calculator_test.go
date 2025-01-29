package tests

import (
	"encoding/json"
	"os"
	"testing"
	"Receipt-Processor-Challenge/internal/calculator"
	"Receipt-Processor-Challenge/internal/models"
)

func loadReceiptFromFile(filePath string) (models.Receipt, error) {
	var receipt models.Receipt
	data, err := os.ReadFile(filePath)
	if err != nil {
		return receipt, err
	}
	if err := json.Unmarshal(data, &receipt); err != nil {
		return receipt, err
	}
	return receipt, nil
}

func TestCalculatePoints(t *testing.T) {
	testFiles := []string{"examples/morning-receipt.json", "examples/simple-receipt.json"}

	for _, file := range testFiles {
		receipt, err := loadReceiptFromFile(file)
		if err != nil {
			t.Fatalf("Failed to load test file %s: %v", file, err)
		}

		expectedPoints := calculator.CalculatePoints(receipt) 
		actualPoints := calculator.CalculatePoints(receipt)
		if actualPoints != expectedPoints {
			t.Errorf("Expected %d points, got %d for file %s", expectedPoints, actualPoints, file)
		}
	}
}
