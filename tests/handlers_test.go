package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"Receipt-Processor-Challenge/internal/handlers"
)

func loadTestFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func TestProcessReceipt(t *testing.T) {
	testFiles := []string{"examples/morning-receipt.json", "examples/simple-receipt.json"}

	for _, file := range testFiles {
		data, err := loadTestFile(file)
		if err != nil {
			t.Fatalf("Failed to load test file %s: %v", file, err)
		}

		req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handlers.ProcessReceipt(w, req)
		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d for file %s", resp.StatusCode, file)
		}
	}
}

func TestGetPoints_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/receipts/nonexistent-id/points", nil)
	w := httptest.NewRecorder()

	handlers.GetPoints(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}
}
