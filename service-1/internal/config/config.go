package config

import (
	"database/sql"
	"os"
)

const (
	defaultPort     = ":9090"
	defaultDbName   = "postgres"
	defaultDBEngine = "postgres"
	defaultDbURL    = "localhost:5432"
	defaultHost     = "postgresdb"
	defaultDBPort   = "5432"
	defaultUsername = "postgres"
	defaultPassword = "postgres"
	defaultSSL      = "disable"
)

func (cfg *Configuration) setEnvValues() {

	cfg.DataStoreConfiguration.DSURL = os.Getenv("DB_URL")
	if cfg.DataStoreConfiguration.DSURL == "" {
		cfg.DataStoreConfiguration.DSURL = defaultDbURL
	}
	cfg.DataStoreConfiguration.DSDB = os.Getenv("DS_DB")
	if cfg.DataStoreConfiguration.DSDB == "" {
		cfg.DataStoreConfiguration.DSDB = defaultDBEngine
	}
	cfg.DataStoreConfiguration.DSName = os.Getenv("DB_NAME")
	if cfg.DataStoreConfiguration.DSName == "" {
		cfg.DataStoreConfiguration.DSName = defaultDbName
	}
	cfg.ServerConfiguration.Port = os.Getenv("PORT")
	if cfg.ServerConfiguration.Port == "" {
		cfg.ServerConfiguration.Port = defaultPort
	}
	cfg.DataStoreConfiguration.Postgres.Host = os.Getenv("DB_HOST")
	if cfg.DataStoreConfiguration.Postgres.Host == "" {
		cfg.DataStoreConfiguration.Postgres.Host = defaultHost
	}
	cfg.DataStoreConfiguration.Postgres.Port = os.Getenv("DB_PORT")
	if cfg.DataStoreConfiguration.Postgres.Port == "" {
		cfg.DataStoreConfiguration.Postgres.Port = defaultDBPort
	}
	cfg.DataStoreConfiguration.Postgres.UserName = os.Getenv("DB_USERNAME")
	if cfg.DataStoreConfiguration.Postgres.UserName == "" {
		cfg.DataStoreConfiguration.Postgres.UserName = defaultUsername
	}
	cfg.DataStoreConfiguration.Postgres.Password = os.Getenv("DB_USERNAME")
	if cfg.DataStoreConfiguration.Postgres.Password == "" {
		cfg.DataStoreConfiguration.Postgres.Password = defaultPassword
	}
	cfg.DataStoreConfiguration.Postgres.SSLMode = os.Getenv("DB_SSL")
	if cfg.DataStoreConfiguration.Postgres.SSLMode == "" {
		cfg.DataStoreConfiguration.Postgres.SSLMode = defaultSSL
	}

}

func LoadConfig() Configuration {
	var opts Configuration
	opts.setEnvValues()
	return opts
}

type Configuration struct {
	DataStoreConfiguration
	ServerConfiguration
}

type ServerConfiguration struct {
	Port string `json:"port" env:"port" envDefault:":8080"`
}

type DataStoreConfiguration struct {
	DSName   string `json:"db_name" env:"db_name" envDefault:"testdb"`
	DSDB     string `json:"db_engine" env:"db_engine" envDefault:"postgres"`
	DSURL    string `json:"db_url" env:"db_url"  envDefault:"localhost:5432://postgres:postgres@testdb"`
	Postgres Postgres
	Mongo    Mongo
}

type Mongo struct{}

type Postgres struct {
	Cli                *sql.DB
	Port               string
	Host               string
	UserName, Password string
	SSLMode            string
}
