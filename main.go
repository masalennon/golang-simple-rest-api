package main

import (
	"net/http"
	"fmt"
	"encoding/json"

)

func main() {

	http.HandleFunc("/trust", handleTrustRequest)
	http.HandleFunc("/belief", handleBeliefRequest)
	http.ListenAndServe(":8080", nil)
}

func handleTrustRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleTrustGet(w, r)
	case "POST":
		
	}
}

func handleBeliefRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleBeliefGet(w, r)
	case "POST":
		
	}
}

type trust struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type belief struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type allTrusts []trust

type allBeliefs []belief

var trusts = allTrusts{
	{
		ID:          "11",
		Title:       "trust",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

var beliefs = allBeliefs{
	{
		ID:          "12",
		Title:       "belief",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func handleTrustGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	json.NewEncoder(w).Encode(trusts)
}

func handleBeliefGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	json.NewEncoder(w).Encode(beliefs)
}

func handlePost(w http.ResponseWriter, r *http.Request) {

}