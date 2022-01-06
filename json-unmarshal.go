package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var err error

type People struct {
	Name      string    `json:"name"`
	Height    int       `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}


func main() {
	jsonDosya, err := os.Open("people.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonDosya.Close()

	icerik, err := ioutil.ReadAll(jsonDosya)
	if err != nil {
		fmt.Println(err)
	}

	var people []People
	json.Unmarshal(icerik, &people)

	for _, person := range people {
		fmt.Println(person)
	}


}