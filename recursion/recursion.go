package recursion

import "fmt"

func Recu() {
	var double int = returnFactorial(16)
	fmt.Println(double)
}
func returnFactorial(num int) int {
	if num == 1 {
		return num
	}
	return num * returnFactorial(num-1)
}
