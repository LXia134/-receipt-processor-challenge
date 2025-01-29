package storage

import (
	"Receipt-Processor-Challenge/internal/models"
	"log"
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

	log.Printf("SaveReceipt storing: ID=%s, Points=%d\n", id, calculatedPoints)

	receipts[id] = receipt
	points[id] = calculatedPoints
}

func GetPoints(id string) (int, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	pointsValue, found := points[id]
	if !found {
		log.Printf("Receipt ID %s not found in stored receipts\n", id)
	}
	return pointsValue, found
}
