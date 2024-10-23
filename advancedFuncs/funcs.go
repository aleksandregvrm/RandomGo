package funcs

import "fmt"

type transformFn func(int) int

type anotherFunc func(int, []string, map[string][]int) ([]int, string)

func Funcs() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var myDoubledNumbers = returnDoubled(&numbers, triple)
	transformerGet := getTransformerFunction(&numbers)
	moreTransformedNumbers := returnDoubled(&numbers, transformerGet)
	fmt.Println(moreTransformedNumbers)
}

func returnDoubled(arr *[]int, transform transformFn) []int {
	doubledNumbers := []int{}
	for index, val := range *arr {
		doubledNumbers = append(doubledNumbers, (transform(val))+index)
	}
	return doubledNumbers
}
func getTransformerFunction(numbers *[]int) func(int) int {
	return triple
}

func triple(number int) int {
	return number * 3
}
