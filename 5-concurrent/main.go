package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

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
	resp, err := http.Get("http://www.omdbapi.com/?i=" + imdbcode + "&plot=short&r=json")
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
	i = (i / 10) * 100
	if err3 != nil {
		fmt.Println(err)
	}
	y := int(i)
	if (int(i*10))%10 >= 5 {
		y++
	}
	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes\n", m.Title, m.Year, y, m.ImdbVotes)
}
func main() {
	moviePtr := flag.String("movie", "batman", "a string")
	flag.Parse()
	resp, err := http.Get("http://www.omdbapi.com/?s=" + *moviePtr)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var m movie
	err2 := json.Unmarshal(body, &m)
	if err2 != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(m.Search); i++ {
		go req(m.Search[i].ImdbID)
		time.Sleep(time.Second * 5)
	}
}
