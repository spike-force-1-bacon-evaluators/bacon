package neo4bacon

import (
	"fmt"
	"log"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/spf13/viper"
)

// Auth ...
type Auth struct {
	user     string
	password string
	baseURL  string
	port     string
	URL      string
}

const (
	connectErrMsg    = "Error connectiong to Neo4J"
	loadConfigErrMsg = "Error loading config file"
)

// LoadConfig from Viper
func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s: %s", loadConfigErrMsg, err)
	}
}

// newAuth returns an auth instance
func newAuth() *Auth {
	return &Auth{
		user:     viper.GetString("neo4j.user"),
		password: viper.GetString("neo4j.password"),
		baseURL:  viper.GetString("neo4j.baseURL"),
		port:     viper.GetString("neo4j.port"),
	}
}

// getURL creates a connection URL
func (a *Auth) getURL() {
	a.URL = fmt.Sprintf("bolt://%s:%s@%s:%s/",
		a.user, a.password, a.baseURL, a.port)
}

func (a *Auth) getConnection() (bolt.Conn, error) {

	// Ask for a new Neo4J Bolt Driver
	driver := bolt.NewDriver()

	// Open new connection with Neo4j
	conn, err := driver.OpenNeo(a.URL)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", connectErrMsg, err)
	}

	return conn, nil
}
