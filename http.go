package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Glossary struct {
	GlossaryData `json:"glossary"`
}

type GlossaryData struct {
	Title        string `json:"title"`
	GlossDivData `json:"GlossDiv"`
}

type GlossDivData struct {
	Title         string `json:"title"`
	GlossListData `json:"GlossList"`
}

type GlossListData struct {
	GlossEntryData `json:"GlossEntry"`
}

type GlossEntryData struct {
	ID        string `json:"ID"`
	SortAs    string `json:"SortAs"`
	GlossTerm string `json:"GlossTerm"`
	Acronym   string `json:"Acronym"`
	GlossSee  string `json:"GlossSee"`
}

func getJSON() {
	r, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println("Error: %v", err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	var GlossaryYS []Glossary
	err = json.Unmarshal(body, &GlossaryYS)
	if err != nil {
		fmt.Println("Unmarshal Error:", err)
	}

	fmt.Println(body)

	fmt.Println(GlossaryYS)
}

func main() {
	getJSON()
}
