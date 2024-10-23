package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinStr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {

	fmt.Println(joinStr())

	fmt.Println(joinStr("GEEK", "GFG"))
	fmt.Println(joinStr("Geeks", "for", "Geeks"))
	fmt.Println(joinStr("G", "E", "E", "k", "S"))

}
