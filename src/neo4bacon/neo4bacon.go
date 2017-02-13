package neo4bacon

import (
	"log"
	"strconv"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Restaurant struct {
	Name    string
	Lastpos int
	Newpos  int
	Change  string
}

type neo4j struct {
	conn          bolt.Conn
	url           string
	oldlistquery  string
	newlistquery  string
	oldlistresult [][]interface{}
	newlistresult [][]interface{}
	result        []*Restaurant
}

// var queryTest = "MATCH (r:Restaurant) return r"
var queryOld = "MATCH (r:Restaurant)-->(b:Bacon) where b.last_points is not null return r.name as Restaurant ORDER BY b.last_points DESC LIMIT 1000"
var queryNew = "MATCH (r:Restaurant)-->(b:Bacon) where b.points is not null return r.name as Restaurant ORDER BY b.points DESC LIMIT 1000"

// Run Neo4j integration to get a new list of restaurants
func Run() []*Restaurant {

	loadConfig()

	// Initialize Auth struct and get URL for Neo4j connection
	a := newAuth()
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
	n.getRestaurants(a)
	n.mapResult()
	n.conn.Close()
	return n.result
}

// Execute Neo4j query to retrieve data to generate new ranking
func (n *neo4j) getRestaurants(a *Auth) {

	oldlist, err := n.queryNeo4j(n.oldlistquery)
	if err != nil {
		log.Fatalf("failed to receive the old list %s", err)
	}
	n.oldlistresult = oldlist

	newlist, err := n.queryNeo4j(n.newlistquery)
	if err != nil {
		log.Fatalf("failed to receive the new list %s", err)
	}
	n.newlistresult = newlist
}

func (n *neo4j) queryNeo4j(str string) ([][]interface{}, error) {
	data, _, _, err := n.conn.QueryNeoAll(str, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (n *neo4j) mapResult() {

	var restaurants []*Restaurant

	for i, value := range n.newlistresult {
		r := &Restaurant{
			Name:   value[0].(string),
			Newpos: i + 1,
		}
		restaurants = append(restaurants, r)
	}

	for _, r := range restaurants {
		r.Change = "N/A"
		for i, value := range n.oldlistresult {
			if r.Name == value[0].(string) {
				r.Lastpos = i + 1
				strChange := strconv.Itoa(r.Lastpos - r.Newpos)
				r.Change = strChange
			}
		}
	}
	n.result = restaurants
}
