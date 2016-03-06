package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (r *user) helloName() {
	fmt.Println("Hello", r.Name)
}

func (r *user) details() {
	DOB, err := time.Parse("January 2, 2006", r.DOB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s who was born in %s would be %d years old today.\n", r.Name, r.City, time.Now().Year()-DOB.Year())
	}
}

func main() {

	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
	u.helloName()
	u.details()
	//value := u.DOB
	//layout := "January 2, 2006"
}
