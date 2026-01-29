package main

import "fmt"

func printReports(intro, body, outro string) {
    // Intro: 2x the message length
    printCostReport(func(msg string) int {
        return len(msg) * 2
    }, intro)
    
    // Body: 3x the message length
    printCostReport(func(msg string) int {
        return len(msg) * 3
    }, body)
    
    // Outro: 4x the message length
    printCostReport(func(msg string) int {
        return len(msg) * 4
    }, outro)
}

// don't touch below this line

func main() {
    printReports(
        "Welcome to the Hotel California",
        "Such a lovely place",
        "Plenty of room at the Hotel California",
    )
}

func printCostReport(costCalculator func(string) int, message string) {
    cost := costCalculator(message)
    fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
    fmt.Println()
}
