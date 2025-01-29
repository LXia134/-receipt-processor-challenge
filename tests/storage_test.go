package tests

import (
	"Receipt-Processor-Challenge/internal/calculator"
	"Receipt-Processor-Challenge/internal/storage"
	"path/filepath"
	"testing"
)
func TestSaveAndGetReceipt(t *testing.T) {
	testFiles := []string{"examples/simple-receipt.json", "examples/morning-receipt.json"}

	for _, file := range testFiles {
		receipt, err := loadReceiptFromFile(file)
		if err != nil {
			t.Fatalf("Failed to load test file %s: %v", file, err)
		}

		receiptID := filepath.Base(file)
		expectedPoints := calculator.CalculatePoints(receipt)

		storage.SaveReceipt(receiptID, receipt, expectedPoints)
		actualPoints, found := storage.GetPoints(receiptID)
		if !found {
			t.Fatalf("Expected receipt ID %s to be found", receiptID)
		}
		if actualPoints != expectedPoints {
			t.Errorf("Expected %d points, got %d for file %s", expectedPoints, actualPoints, file)
		}
	}

	// Test non-existent receipt
	_, found := storage.GetPoints("non-existent-id")
	if found {
		t.Errorf("Expected non-existent receipt to not be found")
	}
}
