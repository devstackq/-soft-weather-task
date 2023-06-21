package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"soft-weater/internal/config"
	"soft-weater/internal/database/drivers"
)

type PostgreSql struct {
	conf    drivers.DataStoreConfig
	connURL string
	dbName  string

	client *sql.DB

	userRepo    drivers.UserRepository
	accountRepo drivers.AccountRepository
}

func (m *PostgreSql) Connect(cfg config.DataStoreConfiguration) error {
	if cfg.Postgres.Cli != nil {
		m.client = cfg.Postgres.Cli
		return nil
	}
	if cfg.Postgres.Cli == nil {
		psqlInfo := fmt.Sprint(
			" user= ", cfg.Postgres.UserName,
			" password= ", cfg.Postgres.Password,
			" dbname= ", cfg.DSName,
			" host= ", cfg.Postgres.Host,
			" port= ", cfg.Postgres.Port,
			" sslmode= ", cfg.Postgres.SSLMode,
		)
		fmt.Println(psqlInfo)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return err
		}

		if err = db.Ping(); err != nil {
			return err
		}

		m.client = db
	}

	return nil
}

const (
	tableAccountName = "account"
	tableUserName    = "users"
)

func (m *PostgreSql) AccountRepository() drivers.AccountRepository {
	if m.accountRepo == nil {
		m.accountRepo = NewAccountRepository(m.client, tableAccountName)
	}
	return m.accountRepo
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

func (m *PostgreSql) UserRepository() drivers.UserRepository {
	if m.userRepo == nil {
		m.userRepo = NewUserRepository(m.client, tableUserName)
	}
	return m.userRepo
}
