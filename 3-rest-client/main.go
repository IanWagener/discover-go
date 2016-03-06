package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, error := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	fmt.Println("status code is", response.Status)
	if error != nil {
		fmt.Println(error)
	}
}
