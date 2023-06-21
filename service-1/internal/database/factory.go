package database

import (
	"fmt"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/database/drivers/psql"
)

const (
	postgreSqlDataStore = "postgres"
)

func New(conf drivers.DataStoreConfig) (drivers.Datastore, error) {
	if conf.DataStoreName == postgreSqlDataStore {
		return psql.New(conf)
	}
	return nil, fmt.Errorf("invalid datastore name: %s", conf.DataStoreName)
}
