package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	doubles := transformNumbers(&numbers, double)
	fmt.Println(doubles)
	triples := transformNumbers(&numbers, triple)
	fmt.Println(triples)

	quadruple := createTransformer(4)
	quadruples := transformNumbers(&numbers, quadruple)
	fmt.Println(quadruples)

	fmt.Println(factorial(5))

	sum := sumup(1, 10, 15, -5)
	fmt.Println(sum)
	anotherSum := sumup(numbers...)
	fmt.Println(anotherSum)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// Closures, similar to factory functions, but uses the variables to lock a value inside the functions
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

}

// variadic functions
func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
