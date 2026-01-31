package main

func maxMessages(thresh int) int {
	total := 0
	messages := 0

	for {
		cost := 100 + messages
		if total+cost > thresh {
			break
		}
		total += cost
		messages++
	}

	return messages
}

