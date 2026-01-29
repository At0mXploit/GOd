package main

func adder() func(int) int {
    sum := 0  // This variable is enclosed/captured by the inner function
    return func(x int) int {
        sum += x  // Add input to the sum
        return sum  // Return the new total
    }
}
