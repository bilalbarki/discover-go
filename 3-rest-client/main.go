package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	imdbcode := "tt0372784"
	plot := "short"
	r := "json"
	resp, err := http.Get("http://www.omdbapi.com/?i=" + imdbcode + "&plot=" + plot + "&r=" + r)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var m user
	err2 := json.Unmarshal(body, &m)
	if err2 != nil {
		fmt.Println(err)
	}
	i, err3 := strconv.ParseFloat(m.ImdbRating, 32)
	if err3 != nil {
		fmt.Println(err)
	}
	i = (i / 10) * 100
	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes\n", m.Title, m.Year, int(i), m.ImdbVotes)
}
