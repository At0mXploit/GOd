package main

import (
	"fmt"
	"time"
)

// Define the message interface
type message interface {
	getMessage() string
}

// sendMessage returns the message content and its cost
func sendMessage(msg message) (string, int) {
	content := msg.getMessage()
	cost := len(content) * 3
	return content, cost
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// Example usage
func main() {
	bm := birthdayMessage{
		birthdayTime:  time.Date(2026, time.January, 30, 0, 0, 0, 0, time.UTC),
		recipientName: "Alice",
	}

	sr := sendingReport{
		reportName:    "Weekly",
		numberOfSends: 42,
	}

	content1, cost1 := sendMessage(bm)
	fmt.Println("Message:", content1)
	fmt.Println("Cost:", cost1)

	content2, cost2 := sendMessage(sr)
	fmt.Println("Message:", content2)
	fmt.Println("Cost:", cost2)
}

