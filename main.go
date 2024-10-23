package main

<<<<<<< HEAD
func main() {
	userNames := make([]string, 2, 5)

=======
import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// var productNames = [4]string{"A book"}
	courses := make(floatMap, 5)

	courses["go"] = 4.7
	courses["react"] = 7.4
	courses["angular"] = 20.1
	// courses.output()
	for key, value := range courses {
		fmt.Println(key, value)
	}
>>>>>>> origin/homeGo
}
