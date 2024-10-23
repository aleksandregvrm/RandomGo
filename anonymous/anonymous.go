package anonymous

import "fmt"

func Anony() {
	numbers := []int{1, 2, 3}

	// double := createTransformers(2)
	triple := createTransformers(3)

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 8
	})

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformers(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
