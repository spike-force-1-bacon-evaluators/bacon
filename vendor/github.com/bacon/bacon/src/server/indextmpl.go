package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/bacon/bacon/src/neo4bacon"
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

var restaurantList []*neo4bacon.Restaurant

// Create and execute index template
// TODO: change this function accordingly to values returned from Neo4j,
// remove mock data, and write test
func indextmpl(w http.ResponseWriter) {

	restaurantList = neo4bacon.Run()
	var restaurants []*Restaurant

	for _, r := range restaurantList {
		restaurant := &Restaurant{
			Name:      r.Name,
			Position:  r.Newpos,
			Variation: r.Change,
		}
		restaurants = append(restaurants, restaurant)
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

	if r.Variation != "N/A" {

		variation, err := strconv.Atoi(r.Variation)

		if err != nil {
			log.Fatal(err)
		}

		if variation > 0 {
			return &Indicator{
				Direction: "⇧",
				Colour:    "green",
			}
		}

		if variation < 0 {
			variationMod, err := strconv.Atoi(r.Variation)
			if err != nil {
				log.Fatal(err)
			}

			variationMod = -variationMod
			r.Variation = strconv.Itoa(variationMod)

			return &Indicator{
				Direction: "⇩",
				Colour:    "red",
			}
		}
		return &Indicator{
			Direction: "▭",
		}
	}

	ind := &Indicator{
		Direction: "",
		Colour:    "black",
	}
	r.Indicator = ind
	return r.Indicator
}
