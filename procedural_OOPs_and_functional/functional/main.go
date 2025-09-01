package main

import "fmt"

// filter applies a function to each element of a slice and
// returns a new slice containing only the elements for which the function returns true.
// This is a higher-order function.
func filter(nums []int, f func(int) bool) []int {
    var result []int
    for _, num := range nums {
        if f(num) {
            result = append(result, num)
        }
    }
    return result
}

// mapFunc applies a function to each element of a slice and
// returns a new slice with the results.
// This is also a higher-order function.
func mapFunc(nums []int, f func(int) int) []int {
    var result []int
    for _, num := range nums {
        result = append(result, f(num))
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Use a pure function to filter out even numbers.
    // The original 'numbers' slice is not modified.
    isEven := func(n int) bool {
        return n%2 == 0
    }
    evenNumbers := filter(numbers, isEven)
    fmt.Println("Even numbers:", evenNumbers) // Output: Even numbers: [2 4 6 8 10]

    // Use a pure function to double each number.
    // The 'evenNumbers' slice is not modified.
    double := func(n int) int {
        return n * 2
    }
    doubledEvenNumbers := mapFunc(evenNumbers, double)
    fmt.Println("Doubled even numbers:", doubledEvenNumbers) // Output: Doubled even numbers: [4 8 12 16 20]
}



// Pure Functions:
// The anonymous functions isEven and double are pure. They have no side effects (e.g., they don't modify
// global variables or the original slices), and for the same input, they will always produce the same output.

// Immutability: 
// The program does not modify the original numbers or evenNumbers slices. Instead, it creates 
// new slices (evenNumbers, doubledEvenNumbers) to hold the results of the transformations. This avoids a 
// common source of bugs in large programs.

// First-Class and Higher-Order Functions:
// In Go, functions are first-class citizens, meaning they can be 
// passed as arguments to other functions, and returned as values. The filter and mapFunc functions are 
// higher-order functions because they accept another function as an argument. This is a powerful feature
// that allows for code reuse and composition.