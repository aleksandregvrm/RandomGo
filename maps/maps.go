package maps

import "fmt"

func main() {
	website := map[string]string{"Google": "https://google.com", "AWS": "https://aws.com"}
	println(website["Google"])
	delete(website, "Google")
	fmt.Println(website)
}
