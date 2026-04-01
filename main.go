package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Listing struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Harga  float64 `json:"harga"`
	Gender string  `json:"gender"`
	Size   string  `json:"size"`
}

var listings = []Listing{
	{ID: 1, Nama: "Nike Court Trackzip", Harga: 420000, Gender: "Hommes", Size: "L"},
	{ID: 2, Nama: "Puma x BMW Hooded Bomber", Harga: 550000, Gender: "Hommes", Size: "M"},
	{ID: 3, Nama: "Adidas Terrex", Harga: 605000, Gender: "Hommes", Size: "XL"},
	{ID: 4, Nama: "Baselayer Rebook", Harga: 100000, Gender: "Unisex", Size: "S"},
	{ID: 5, Nama: "Kappa Trackpants", Harga: 75000, Gender: "Femmes", Size: "S"},
	{ID: 6, Nama: "Ellese Remini Tracktop", Harga: 650000, Gender: "Unisex", Size: "M"},
}

func getListings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listings)
}

func main() {
	http.HandleFunc("/listing", getListings)
	fmt.Println("Udah jalan.. http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
