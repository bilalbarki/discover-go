package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//s := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(s)
	//ran := r.Intn(100)
	ran := rand.Intn(100)
	if ran > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", ran)
	} else {
		fmt.Printf("my random number is %d and is not greater than 50\n", ran)
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	} else {
		fmt.Println("I am not a student of the", school)
	}
	beautifulWeather := true
	if beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	} else {
		fmt.Println("It's not a beautiful weather!")
	}
	holbertonFounders := [3]string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for i := 0; i < len(holbertonFounders); i++ {
		fmt.Println(holbertonFounders[i], "is a founder")
	}
}
