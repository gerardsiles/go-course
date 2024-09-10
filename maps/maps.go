package main

import "fmt"

type floatMap map[string]string

func main() {
	websites := map[string]string{
		"Google":              "www.google.com",
		"Amazon web services": "https://www.aws.com",
	}
	making := make(floatMap, 3)
	fmt.Println(making)
	// make(map[string]string,3) this would create a map with thee items spaces on it
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	for key, value := range websites {
		fmt.Println(key)
		fmt.Println(value)
	}
}
