package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Artists struct {
	ID              int      `json:"id"`
	ImageURL        string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	RelationsURL    string   `json:"relations"`
}

type Locations struct {
	Index []Location `json:"index"`
}
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Index []ArtistDetails `json:"index"`
}

type ArtistDetails struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}


const APIURL = "https://groupietrackers.herokuapp.com/api"

// GetAndUnmarshalArtists returns a list of artists by fetching or using cached data
func GetArtists() ([]Artists, error) {
	artists := []Artists{}
	err := unmarshalData("/artists", &artists)
	return artists, err
}

func GetLocations(ID int) (Location, error) {
	locations := Locations{}
	err := unmarshalData("/locations", &locations)
	if err != nil {
		return Location{}, err
	}

	for _, v := range locations.Index {
		if v.ID == ID {
			return v, nil
		}
	}
	return Location{}, fmt.Errorf("location with ID %d not found", ID)
}

func GetDates(ID int) (Date, error) {
	dates := Dates{}
	err := unmarshalData("/dates", &dates)
	if err != nil {
		return Date{}, err
	}

	for _, v := range dates.Index {
		if v.ID == ID {
			return v, nil
		}
	}
	return Date{}, fmt.Errorf("date with ID %d not found", ID)
}

func GetRelation(ID int) (ArtistDetails, error) {
	relation := Relation{}
	err := unmarshalData( "/relation", &relation)
	if err != nil {
		return ArtistDetails{}, err
	}

	for _, v := range relation.Index {
		if v.ID == ID {
			return v, nil
		}
	}
	return ArtistDetails{}, fmt.Errorf("relation with ID %d not found", ID)
}

func unmarshalData(endpoint string, out interface{}) error {
	jsonData, err := getJSONData(endpoint)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, out)
}

func getJSONData(endpoint string) (json.RawMessage, error) {
	resp, err := http.Get(APIURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s json data: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received a non-200 response code for %s: %d", endpoint, resp.StatusCode)
	}

	var jsonString json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&jsonString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s json data: %w", endpoint, err)
	}

	return jsonString, nil
}

// errors is a map of error output value in ErrorHandler
var errors = map[string]string{
	"colors":     "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\" standard",
	"justify":    "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard",
	"color":      "🤯 Oops! We couldn't recognise your color\n\nKindly search supported colors here: https://htmlcolorcodes.com/",
	"output":     "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard",
	"txt":        "😣 Oops! We currently only support text files\n\nSee Documentation in: ../README.md",
	"web":        "😮 Oops! Something went wrong",
	"restricted": "😣 Oops! this is a restricted path.\nplease use another path.",
}

// ErrorHandler outputs errors and safely exits the program
func ErrorHandler(errType string) {
	if errType == "fatal" {
		fmt.Printf("For color:\n%s\n", strings.Split(errors["colors"], "\n")[2])
		fmt.Printf("For output:\n%s\n", strings.Split(errors["output"], "\n")[2])
		fmt.Printf("For justify:\n%s\n", strings.Split(errors["justify"], "\n")[2])
		fmt.Println("For web:\ngo run . -web")
		os.Exit(0)
	}
	fmt.Println(errors[errType])
	os.Exit(0)
}
