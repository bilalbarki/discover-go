package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	ran := r.Intn(100)
	if ran > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", ran)
	} else {
		fmt.Printf("my random number is %d and is smaller than 50\n", ran)
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	} else {
		fmt.Println("I am a not student of the", school)
	}
	beautifulWeather := true
	if beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	} else {
		fmt.Println("It's not a beautiful weather!")
	}
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for i := 0; i < len(holbertonFounders); i++ {
		fmt.Println(holbertonFounders[i])
	}
}
