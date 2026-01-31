package main

import "fmt"

// Implement cost method for expense interface
func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

// Implement format method for formatter interface
func (e email) format() string {
	status := "Not Subscribed"
	if e.isSubscribed {
		status = "Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, status)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

