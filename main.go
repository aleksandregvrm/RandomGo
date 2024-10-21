package main

import (
	"fmt"
)

func main() {
	// var productNames = [4]string{"A book"}
	prices := [4]float64{1, 2, 3, 4}

	const myNum int = 4
	featuredPrices := prices[1:len(prices)]
	highlitedPrices := featuredPrices[:1]
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))
}
