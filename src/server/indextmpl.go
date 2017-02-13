package server

import (
	"html/template"
	"log"
	"net/http"
)

// Meta information
type Meta struct {
	Title    string
	Subtitle string
	TabTitle string
}

// Indicator wrap info related to restaurant change symbol (the arrows)
type Indicator struct {
	Direction string
	Colour    string
}

// Restaurant wrap restaurant represetation for each row in ranking table
type Restaurant struct {
	Position  int
	Name      string
	Variation string
	Indicator *Indicator
}

// Index wrap all the content to be displayed on the index page
type Index struct {
	Meta        *Meta
	Restaurants []*Restaurant
}


// Create and execute index template
// TODO: change this function accordingly to values returned from Neo4j,
// remove mock data, and write test
func indextmpl(w http.ResponseWriter) {

	restaurants := []*Restaurant{
		&Restaurant{
			Position:  1,
			Name:      "The Restaurant at the End of the Universe",
			Variation: 2},
		&Restaurant{
			Position:  2,
			Name:      "Barriga Estufada",
			Variation: 1},
		&Restaurant{
			Position:  3,
			Name:      "PratoFeito",
			Variation: -2},
		&Restaurant{
			Position:  4,
			Name:      "Só Maminha",
			Variation: 0},
	}

	for _, r := range restaurants {
		r.Indicator = r.newIndicator()
	}

	index := &Index{
		Meta: &Meta{
			Title:    "Bacon Evaluators",
			Subtitle: "A Bodacious Adviser for Cuisine Over N-joyment",
			TabTitle: "Bacon",
		},
		Restaurants: restaurants,
	}

	// Create template
	tmpl, err := template.New("index").ParseFiles("tmpl/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.ExecuteTemplate(w, "index", index)
}

// This method assigns a symbol (arrow or rectangle) to
// a new Indicator accordingly the Variation value.
func (r *Restaurant) newIndicator() *Indicator {
	if r.Variation > 0 {
		return &Indicator{
			Direction: "⇧",
			Colour:    "green",
		}
	}
	if r.Variation < 0 {
		r.Variation = -r.Variation
		return &Indicator{
			Direction: "⇩",
			Colour:    "red",
		}
	}
	return &Indicator{
		Direction: "▭",
	}
}
