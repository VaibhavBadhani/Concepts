package main

import "fmt"

func calculateRectangleArea(length float64, width float64) float64 {
	return length * width
}

func printResult(area float64) {
	fmt.Println("The area of the rectangle is:", area)
}

func main() {
	// 1. Declare variables to store data
	var length float64 = 10.5
	var width float64 = 5.2

	// 2. Call the function to perform a specific task (calculation)
	area := calculateRectangleArea(length, width)

	// 3. Call another function to perform a different task (printing)
	printResult(area)
}



// Global State (Implicitly): While this example doesn't use global variables,
// a classic characteristic of procedural programming is that multiple 
// functions can modify shared, global data. In Go, this is less common due to
// the emphasis on clear data flow, but the paradigm still holds. The main 
// function holds the data (length, width, area) and orchestrates the flow of that data.

