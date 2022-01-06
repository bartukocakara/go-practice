package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var err error

type People struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
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
	// Json dosyasını açalım
	jsonDosya, err := os.Open("people.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonDosya.Close()

	dec := json.NewDecoder(jsonDosya)

	// Başlangıç verisini okuyalım

	_, err = dec.Token()

	var person People

	for dec.More(){
		err = dec.Decode(&person)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(person.Name)
	}

	// kapanış verisini okuyalım
	_, err = dec.Token()
	if err != nil {
		println(err)
	}
}