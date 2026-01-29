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
