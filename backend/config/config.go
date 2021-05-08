package config

import (
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// the base of the struct that the data goes into
type Config struct {
	Mode string   `hcl:"Mode"`
	Port string   `hcl:"Port"`
	DB   DBConfig `hcl:"DB,block"`
}

type DBConfig struct {
	Host     string `hcl:"Host"`
	Port     string `hcl:"Port"`
	User     string `hcl:"User"`
	Password string `hcl:"Password"`
	DBName   string `hcl:"DBName"`
	SSLMode  string `hcl:"SSLMode"`
}

// a global variable to be able to access anywhere
var CONFIG Config

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

// Load in from the config
func Load() {
	_, isDeployed := os.LookupEnv("DEPLOY_ENV")
	if isDeployed {
		DBCONFIG := DBConfig{
			Host:     mustGetenv("DB_HOST"),
			Port:     "5432",
			User:     mustGetenv("DB_USER"),
			Password: mustGetenv("DB_PASSWORD"),
			DBName:   mustGetenv("DB_NAME"),
			SSLMode:  "disable",
		}
		CONFIG = Config{
			Mode: "prod",
			Port: "443",
			DB:   DBCONFIG,
		}
	} else {
		err := hclsimple.DecodeFile("../config/config.hcl", nil, &CONFIG)
		if err != nil {
			log.Fatalf("Failed to load config: %s", err)
		}
	}

	// if its in dev mode
	// print out entire config
	if CONFIG.Mode == "dev" {
		log.Printf("CONFIG: %#v", CONFIG)
	}
}
