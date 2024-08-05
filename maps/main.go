package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "www.google.com",
		"amazon web services": "www.aws.com",
	}
	fmt.Println(websites["google"])
	delete(websites, "google")
	fmt.Println(websites)
}
