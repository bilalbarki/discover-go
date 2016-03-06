package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type user struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Response   string `json:"Response"`
}

func main() {
	imdbcode := "tt0372784"
	plot := "short"
	r := "json"
	resp, err := http.Get("http://www.omdbapi.com/?i=" + imdbcode + "&plot=" + plot + "&r=" + r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
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
