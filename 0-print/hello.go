package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, we are Holberton School")
	t := time.Now()
	fmt.Println("the date is", t)
	fmt.Println("the year is", t.Year())
	fmt.Println("the month is", t.Month())
	fmt.Println("the day is", t.Day())
}
