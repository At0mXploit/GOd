package main

func reformat(message string, formatter func(string) string) string {
    // Apply the formatter three times to the message
    result := message
    for i := 0; i < 3; i++ {
        result = formatter(result)
    }
    
    // Add the prefix and return
    return "TEXTIO: " + result
}
