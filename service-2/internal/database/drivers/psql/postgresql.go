package psql

import (
	"database/sql"
	"fmt"
	"soft-weater/internal/config"
	"soft-weater/internal/database/drivers"
)

type PostgreSql struct {
	conf    drivers.DataStoreConfig
	connURL string
	dbName  string

	client      *sql.DB
	taskRepo    drivers.TaskRepository
	accountRepo drivers.AccountRepository
	historyRepo drivers.HistoryRepository
}

const (
	taskTableName    = "task"
	historyTableName = "history"
)

func (m *PostgreSql) Connect(cfg config.DataStoreConfiguration) error {
	if m.client != nil {
		return nil
	}
	if m.client == nil {
		psqlInfo := fmt.Sprint("user= postgres password=postgres dbname=testdb host=localhost port=5432 sslmode=disable") //temp todo fix

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return err
		}

		if err = db.Ping(); err != nil {
			return err
		}
		m.client = db
		//cfg.Postgres.Cli = db
	}

	return nil
}

func New(conf drivers.DataStoreConfig) (drivers.Datastore, error) {
	if conf.URL == "" {
		return nil, drivers.ErrInvalidConfigStruct
	}

	if conf.DataBaseName == "" {
		return nil, drivers.ErrInvalidConfigStruct
	}

	return &PostgreSql{
		connURL: conf.URL,
		dbName:  conf.DataBaseName,
	}, nil
}

func (m *PostgreSql) TaskRepository() drivers.TaskRepository {
	if m.taskRepo == nil {
		m.taskRepo = NewTaskRepository(m.client, taskTableName)
	}
	return m.taskRepo
}

func (m *PostgreSql) HistoryRepository() drivers.HistoryRepository {
	if m.historyRepo == nil {
		m.historyRepo = NewHistoryRepository(m.client, historyTableName)
	}
	return m.historyRepo
}
