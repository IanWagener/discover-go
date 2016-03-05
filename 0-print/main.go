package main

import (
"fmt"
"time"
)
func main() {
	now := time.Now()

	fmt.Println("Hello, we are Holberton School")

	fmt.Println("the date is", time.Now())

	fmt.Println("the year is", now.Year())

	fmt.Println("the month is", now.Month())

	fmt.Println("the day is", now.Day())
}