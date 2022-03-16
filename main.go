package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type Customer struct {
	Name   string
	Age    int
	Height float32
}

func encoding(w http.ResponseWriter, r *http.Request) {
	customer1 := []Customer{{Name: "Abenezer", Age: 26, Height: 1.70},
		{Name: "Abenezer", Age: 26, Height: 1.70}}

	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer1)

	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer1)
	}
}

func main() {

	// JSON encoded

	http.HandleFunc("/", encoding)

	// XML encoded
	//http.HandleFunc("/xml", xml_encoding)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
