package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://developer.com:80/hello?name=Leonardo&age=20"
	url, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("url: %s\n", url.String())

	fmt.Printf("protocol: %s\n", url.Scheme)
	fmt.Printf("host: %s\n", url.Host)
	fmt.Printf("path: %s\n", url.Path)

	var name = url.Query()["name"][0]
	var age = url.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}
