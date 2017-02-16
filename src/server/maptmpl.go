package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bacon/bacon/src/neo4bacon"
)

type GeoRest struct {
	Name  string
	Score float64
	Lng   float64
	Lat   float64
}

type JsSnippets struct {
	Name1      template.JS
	Name2      template.JS
	Name3      template.JS
	Name4      template.JS
	Name5      template.JS
	Name6      template.JS
	Name7      template.JS
	Name8      template.JS
	Name9      template.JS
	Name10     template.JS
	Location1  template.JS
	Location2  template.JS
	Location3  template.JS
	Location4  template.JS
	Location5  template.JS
	Location6  template.JS
	Location7  template.JS
	Location8  template.JS
	Location9  template.JS
	Location10 template.JS
	Score1     template.JS
	Score2     template.JS
	Score3     template.JS
	Score4     template.JS
	Score5     template.JS
	Score6     template.JS
	Score7     template.JS
	Score8     template.JS
	Score9     template.JS
	Score10    template.JS
}

func maptmpl(w http.ResponseWriter) {

	// Call neo4bacon.GetRestGeo
	rawGeoRestaurantList, err := neo4bacon.GetGeoLocation()
	if err != nil {
		log.Fatal(err)
	}

	geoList := parseGeoRests(rawGeoRestaurantList)

	js := &JsSnippets{
		Name1:      template.JS(fmt.Sprint(geoList[0].Name)),
		Location1:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[0].Lat, geoList[0].Lng)),
		Score1:     template.JS(fmt.Sprint(geoList[0].Score / 10)),
		Name2:      template.JS(fmt.Sprint(geoList[1].Name)),
		Location2:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[1].Lat, geoList[1].Lng)),
		Score2:     template.JS(fmt.Sprint(geoList[1].Score / 10)),
		Name3:      template.JS(fmt.Sprint(geoList[2].Name)),
		Location3:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[2].Lat, geoList[2].Lng)),
		Score3:     template.JS(fmt.Sprint(geoList[2].Score / 10)),
		Name4:      template.JS(fmt.Sprint(geoList[3].Name)),
		Location4:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[3].Lat, geoList[3].Lng)),
		Score4:     template.JS(fmt.Sprint(geoList[3].Score / 10)),
		Name5:      template.JS(fmt.Sprint(geoList[4].Name)),
		Location5:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[4].Lat, geoList[4].Lng)),
		Score5:     template.JS(fmt.Sprint(geoList[4].Score / 10)),
		Name6:      template.JS(fmt.Sprint(geoList[5].Name)),
		Location6:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[5].Lat, geoList[5].Lng)),
		Score6:     template.JS(fmt.Sprint(geoList[5].Score / 10)),
		Name7:      template.JS(fmt.Sprint(geoList[6].Name)),
		Location7:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[6].Lat, geoList[6].Lng)),
		Score7:     template.JS(fmt.Sprint(geoList[6].Score / 10)),
		Name8:      template.JS(fmt.Sprint(geoList[7].Name)),
		Location8:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[7].Lat, geoList[7].Lng)),
		Score8:     template.JS(fmt.Sprint(geoList[7].Score / 10)),
		Name9:      template.JS(fmt.Sprint(geoList[8].Name)),
		Location9:  template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[8].Lat, geoList[8].Lng)),
		Score9:     template.JS(fmt.Sprint(geoList[8].Score / 10)),
		Name10:     template.JS(fmt.Sprint(geoList[9].Name)),
		Location10: template.JS(fmt.Sprintf("{lat: %f, lng: %f}", geoList[9].Lat, geoList[9].Lng)),
		Score10:    template.JS(fmt.Sprint(geoList[9].Score / 10)),
	}

	// Create template
	tmpl, err := template.New("map").ParseFiles("tmpl/map.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "map", js)
}

func parseGeoRests(restaurants [][]interface{}) []*GeoRest {
	var geoRests []*GeoRest
	for _, r := range restaurants {
		g := &GeoRest{
			Name:  r[0].(string),
			Score: r[1].(float64),
			Lng:   r[2].(float64),
			Lat:   r[3].(float64),
		}
		geoRests = append(geoRests, g)
	}
	return geoRests
}
