package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Important: Only exported fields (Capital letters) are marshaled and turned into JSON objects
type Movie struct {
	Title  string
	Year   int      `json:"released"`        // Changes to released in JSON output becaused of the field tag
	Color  bool     `json:"color,omitempty"` // "omitempty" - not to include this field in the JSON output if it contains zero value for its type, for bool : false, for int 0 and so on.
	Actors []string //Field tag is a matadata associted with the struct at compile time

}

func main() {

	movies := []Movie{
		{Title: "Die another day",
			Year:   2002,
			Color:  true,
			Actors: []string{"Pierce Brosnan", "Halle Berry"},
		},
		{Title: "Skyfall",
			Year:   2012,
			Color:  true,
			Actors: []string{"Daniel Craig", "Javier Bardem"},
		},
		{Title: "No Time to Die",
			Year:   2020,
			Color:  false, // This field is not going to be included in the JSON output because it has omitempty and it is zero (false) for its type
			Actors: []string{"Daniel Craig", "Rami Malek"},
		},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Error marshaling movies struct")
	}
	fmt.Println(string(data))

	fmt.Printf("\n\n\n")

	data, err = json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("Error marshaling movies with marshalident")
	}

	fmt.Printf("%s\n", data)

	fmt.Printf("\n\n\n")

	// Unmarshaling
	// This is anonimous struct
	var titles []struct {
		Title string
	}

	err = json.Unmarshal(data, &titles) // returns only error
	if err != nil {
		log.Fatalf("Error unmarshaling data into titles")
	}

	fmt.Printf("%v\n", titles)
}
