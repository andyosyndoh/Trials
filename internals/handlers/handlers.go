package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"groupie/internals/renders"
	"groupie/utils"
)

type ArtistDetails struct {
	Art       utils.Artists        `json:"artist"`
	Locations utils.Location      `json:"locations"`
	Dates     utils.Date          `json:"dates"`
	Relations utils.ArtistDetails `json:"relations"`
}

// type ArtistDetails struct {
//     Art       utils.Artists  // This stores a single artist
//     Locations utils.Locations      // Assuming it's a slice of strings, adjust as necessary
//     Dates     utils.Dates      // Assuming it's a slice of strings, adjust as necessary
//     Relations interface{}    // Adjust the type as necessary
// }

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		artists, err := utils.GetArtists()
		if err != nil {
			http.Error(w, "Failed to retrieve artists data", http.StatusInternalServerError)
			log.Printf("Error retrieving artists: %v", err)
			return
		}
		renders.RenderTemplate(w, "home.page.html", artists)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NotFoundHandler handles unknown routes; 404 status
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, "notfound.page.html", nil)
}

// BadRequestHandler handles bad requests routes
func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, "badrequest.page.html", nil)
}

// ServerErrorHandler handles server failures that result in status 500
func ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, "serverError.page.html", nil)
}

// AboutHandler handles the about page route '/about'
func Location(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		location, _ := utils.GetLocations(ID)
		artists, _ := utils.Getsingleartist(ID)
		dates, _ := utils.GetDates(ID)
		relations, err := utils.GetRelation(ID)

		artistDetails := ArtistDetails{
			Art:       artists, // Wrap single artist in a slice
			Locations: location,
			Dates:     dates,
			Relations: relations,
		}

		fmt.Println(artistDetails.Art)

		if err != nil {
			http.Error(w, "Failed to retrieve location data", http.StatusInternalServerError)
			log.Printf("Error retrieving location data: %v", err)
			return
		}

		renders.RenderTemplate(w, "location.page.html", artistDetails)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		dates, err := utils.GetDates(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve date data", http.StatusInternalServerError)
			log.Printf("Error retrieving date data: %v", err)
			return
		}

		renders.RenderTemplate(w, "date.page.html", dates)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		relation, err := utils.GetRelation(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve relation data", http.StatusInternalServerError)
			log.Printf("Error retrieving relation data: %v", err)
			return
		}
		renders.RenderTemplate(w, "relation.page.html", relation)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
