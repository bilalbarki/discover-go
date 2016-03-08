package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

type movie struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

func req(imdbcode string) {
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
	defer wg.Done()
	resp, err := http.Get("http://www.omdbapi.com/?i=" + imdbcode + "&plot=short&r=json")
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

	i, err3 := strconv.ParseFloat(m.ImdbRating, 64)
	if err3 != nil {
		fmt.Println(err)
	}
	i = i * 10
	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes\n", m.Title, m.Year, int(i), m.ImdbVotes)
}
func main() {
	moviePtr := flag.String("movie", "batman", "a string")
	flag.Parse()
	resp, err := http.Get("http://www.omdbapi.com/?s=" + *moviePtr)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var m movie
	err2 := json.Unmarshal(body, &m)
	if err2 != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(m.Search); i++ {
		wg.Add(1)
		go req(m.Search[i].ImdbID)
	}
	wg.Wait()
}
