package storage

import (
	"Receipt-Processor-Challenge/internal/models"
	"sync"
)

var (
	receipts = make(map[string]models.Receipt) 
	points   = make(map[string]int)            
	mutex    sync.Mutex
)

func SaveReceipt(id string, receipt models.Receipt, calculatedPoints int) {
	mutex.Lock()
	defer mutex.Unlock()
	receipts[id] = receipt
	points[id] = calculatedPoints
}

func GetPoints(id string) (int, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	pointsValue, found := points[id]
	return pointsValue, found
}
