package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Port               string
	Host               string
	Username, Password string
	HtmlFilePath       string
}

const (
	defaultPort  = "587"
	defaultHost  = "smtp-relay.sendinblue.com"
	connectURI   = "mongodb://localhost:27016"
	testUsername = "8akebaev@gmail.com"
	testPassword = "t8O0zvxD3XrbJpA7"

	MONGO_DB_NAME      = "mailer"
	MONGO_TEMPLATE_COL = "templates"
	MONGO_USERS_COL    = "users"
)

func New() *Config {
	return &Config{}
}

func (c *Config) Load() {
	if value, exists := os.LookupEnv("port"); exists {
		c.Port = value
	} else {
		c.Port = defaultPort
	}
	c.Host = defaultHost
	c.Password = testPassword
	c.Username = testUsername
}

func Connect() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(connectURI))
}
