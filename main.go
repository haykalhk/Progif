package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// News struct (Model)
type News struct {
	ID     string  `json:"id"`
	Source string  `json:"src"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init MNews var as a slice News struct
var MNews []News

// Get all MNews
func getMNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MNews)
}

// Get single News
func getNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through MNews and find one with the id from the params
	for _, item := range MNews {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&News{})
}

// Add new News
func createNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var News News
	_ = json.NewDecoder(r.Body).Decode(&News)
	News.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	MNews = append(MNews, News)
	json.NewEncoder(w).Encode(News)
}

// Update News
func updateNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range MNews {
		if item.ID == params["id"] {
			MNews = append(MNews[:index], MNews[index+1:]...)
			var News News
			_ = json.NewDecoder(r.Body).Decode(&News)
			News.ID = params["id"]
			MNews = append(MNews, News)
			json.NewEncoder(w).Encode(News)
			return
		}
	}
}

// Delete News
func deleteNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range MNews {
		if item.ID == params["id"] {
			MNews = append(MNews[:index], MNews[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(MNews)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	MNews = append(MNews, News{ID: "1", Source: "Tribun", Title: "News One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	MNews = append(MNews, News{ID: "2", Source: "New York Times", Title: "News Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/MNews", getMNews).Methods("GET")
	r.HandleFunc("/MNews/{id}", getNews).Methods("GET")
	r.HandleFunc("/MNews", createNews).Methods("POST")
	r.HandleFunc("/MNews/{id}", updateNews).Methods("PUT")
	r.HandleFunc("/MNews/{id}", deleteNews).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
