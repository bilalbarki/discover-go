package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	fmt.Println(r.Intn(100))
}
