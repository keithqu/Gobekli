package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HPIList ...
type HPIList struct {
	HPI []HPI
}

// HPI ...
type HPI struct {
	Year     uint    `json:"year"`
	Month    uint    `json:"month"`
	GeoType  string  `json:"geo_type"`
	GeoName  string  `json:"geo_name"`
	GeoCode  string  `json:"geo_code"`
	IndexNsa float64 `json:"index_nsa"`
	IndexSa  float64 `json:"index_sa"`
}

// Item ... items for sale
type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

type people struct {
	Number int `json:"number"`
}

// Astros ...
type Astros struct {
	Message string   `json:"message"`
	Number  uint     `json:"number"`
	People  []People `json:"people"`
}

// People ...
type People struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Function Called: homePage()")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory()")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON

	inventory = append(inventory, item) // Add item to inventory

	json.NewEncoder(w).Encode(item) // Show item in response JSON for verification
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUID(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUID(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			// Delete item from Slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON

	params := mux.Vars(r)

	_deleteItemAtUID(params["uid"])     // Delete item
	inventory = append(inventory, item) // Create it again with data from request

	json.NewEncoder(w).Encode(inventory)
}

func getStates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "https://api.namara.io/v0/data_sets/fa88065f-0f87-429b-8979-34130c45b317/data/en-0?geometry_format=wkt&api_key=1bcddff30001f86aa86bcfc16b56b6d89d23b0659479cb7a20adaa588acfe0d6&organization_id=5bfd52abb25822140d6e23fc&order=year%20DESC&select=year%2Cmonth%2Cgeo_type%2Cgeo_name%2Cgeo_code%2Cindex_nsa%2Cindex_sa"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error making request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// var hpi = new(HPIList)
	var hpi interface{}
	json.Unmarshal(body, &hpi)

	json.NewEncoder(w).Encode(hpi)
}

func getAstros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := "http://api.open-notify.org/astros.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err.Error())
	}

	var astros = new(Astros)
	err3 := json.Unmarshal(body, &astros)
	if err3 != nil {
		fmt.Println("whoops: ", err3)
	}

	json.NewEncoder(w).Encode(astros)
}

func handleRequests() {
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)
	log.Println("hello")

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory", createItem).Methods("POST")

	router.HandleFunc("/states", getStates).Methods("GET")
	router.HandleFunc("/astros", getAstros).Methods("GET")

	log.Fatal(http.ListenAndServe(":9999", router))
}

func main() {
	// Data store
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "A fine block of cheese.",
		Price: 4.99,
	})

	handleRequests()
}
