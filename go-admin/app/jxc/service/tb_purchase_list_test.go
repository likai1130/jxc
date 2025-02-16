package service

import (
	"fmt"
	"testing"
	"time"
)

func TestTruncate(t *testing.T) {
	startOfDay := time.Now().Truncate(24 * time.Hour).UTC()
	formattedStartTime := startOfDay.Format("2006-01-02 15:04:05")
	endOfDay := startOfDay.Add(24 * time.Hour).UTC()
	formattedEndTime := endOfDay.Format("2006-01-02 15:04:05")
	fmt.Println(formattedStartTime)
	fmt.Println(formattedEndTime)
	fmt.Println(startOfDay)
	fmt.Println(endOfDay)
}
