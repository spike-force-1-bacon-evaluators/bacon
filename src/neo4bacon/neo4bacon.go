// Package neo4bacon implement API for communication with Neo4j
package neo4bacon

import (
	"log"
	"strconv"

	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Restaurant struct wraps information related
// to a single restaurant entity
type Restaurant struct {
	ID      string
	Name    string
	Lastpos int
	Newpos  int
	Change  string
}

// neo4j struct wraps information related
// to Neo4j communication access and results
// from queries
type neo4j struct {
	conn          bolt.Conn
	url           string
	oldlistquery  string
	newlistquery  string
	oldlistresult [][]interface{}
	newlistresult [][]interface{}
	result        []*Restaurant
}

const (
	queryOld = "MATCH (r:Restaurant)-->(b:Bacon) where b.last_points is not null return r.id as Rid, r.name as Restaurant ORDER BY b.last_points DESC"
	queryNew = "MATCH (r:Restaurant)-->(b:Bacon) where b.points is not null return r.id as Rid, r.name as Restaurant ORDER BY b.points DESC"
)

// Run Neo4j integration to get a new list of restaurants
func Run() []*Restaurant {

	// Load credentials for Neo4j connection
	loadConfig()

	// Initialize Auth struct and get URL for Neo4j connection
	a := newAuth()

	// Get URL to connect to Neo4j
	a.getURL()

	n := &neo4j{
		url:          a.URL,
		oldlistquery: queryOld,
		newlistquery: queryNew,
	}

	// Ask for connection
	conn, err := a.getConnection()
	if err != nil {
		log.Fatal(err)
	}
	n.conn = conn

	// Get lists
	if err := n.getRestaurants(a); err != nil {
		log.Fatal(err)
	}
	n.mapResult()

	n.conn.Close()
	return n.result
}

// Execute Neo4j query to retrive data and generate new ranking
func (n *neo4j) getRestaurants(a *Auth) error {

	// Query the old list of restaurants
	oldlist, err := n.queryNeo4j(n.oldlistquery)
	if err != nil {
		return fmt.Errorf("failed to receive the old list %s", err)
	}
	n.oldlistresult = oldlist

	// Query the new list of restaurants
	newlist, err := n.queryNeo4j(n.newlistquery)
	if err != nil {
		return fmt.Errorf("failed to receive the new list %s", err)
	}
	n.newlistresult = newlist
	return nil
}

// Query Neo4j
func (n *neo4j) queryNeo4j(str string) ([][]interface{}, error) {
	data, _, _, err := n.conn.QueryNeoAll(str, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// mapResults iterates over the lists of restaurants
// and creates maps containing the restaurant name
// and the position occupied by the restaurant in the
// original list
func (n *neo4j) mapResult() {

	var restaurants []*Restaurant

	for i, value := range n.newlistresult {
		r := &Restaurant{
			ID:     value[0].(string),
			Name:   value[1].(string),
			Newpos: i + 1,
		}
		restaurants = append(restaurants, r)
	}

	for _, r := range restaurants {
		r.Change = "N/A"
		for i, value := range n.oldlistresult {
			if r.ID == value[0].(string) {
				r.Lastpos = i + 1
				strChange := strconv.Itoa(r.Lastpos - r.Newpos)
				r.Change = strChange
			}
		}
	}
	n.result = restaurants
}
