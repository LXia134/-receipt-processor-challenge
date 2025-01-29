package calculator

import (
	"Receipt-Processor-Challenge/internal/models"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleanRetailer := re.ReplaceAllString(receipt.Retailer, "")
	points := len(cleanRetailer)

	total, _ := strconv.ParseFloat(receipt.Total, 64)

	if total == math.Floor(total) {
		points += 50
	}
	
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil && date.Day()%2 != 0 {
		points += 6
	}

	if timeOfDay, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		if timeOfDay.Hour() >= 14 && timeOfDay.Hour() < 16 {
			points += 10
		}
	}

	return points
}
