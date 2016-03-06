package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string
	DOB  string
	City string
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Phiadelphia"}

	now := time.Now()

	age := now.Year() - 1917

	fmt.Println("Hello", u.Name)

	fmt.Println(u.Name, "who was born in", u.City, "would be", age, "years old today")
}
