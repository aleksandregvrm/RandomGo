package lists

type MyObject struct {
	name    string
	surname string
	balance float64
}

func MyFunc() {

}

// func main() {
// 	hobbiesArr := [3]string{"skiing", "programming", "hiking"}
// 	newArr := hobbiesArr[:2]
// 	newArr = newArr[0:cap(newArr)]
// 	dynamicArr := []string{"learnin`g go", "learning gin with kubernetes"}
// 	fmt.Println(dynamicArr)

// 	structArr := []MyObject{MyObject{"dachi", "imedadze", 10.99}, MyObject{"dachiko", "imedadze", 99.75}}
// 	structArr = append(structArr, MyObject{"krebsona", "dabrundashvili", 75.80})
// 	fmt.Println(structArr)
// }

// func main() {
// 	prices := []float64{5.99, 4.99}
// 	fmt.Println(prices[1])

// 	prices = append(prices, 10.99)
// 	prices = prices[1:]
// 	fmt.Println(len(prices), cap(prices))
// 	otherPrices := []float64{2, 3, 4, 5, 6}
// 	prices = append(prices, otherPrices...)
// }

// func main() {
// 	// var productNames = [4]string{"A book"}
// 	prices := [4]float64{1, 2, 3, 4}

// 	const myNum int = 4
// 	featuredPrices := prices[1:len(prices)]
// 	highlitedPrices := featuredPrices[:1]
// 	fmt.Println(len(highlitedPrices), cap(highlitedPrices))
// }
