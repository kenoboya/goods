package psql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PSQlConfig struct {
	Host     string
	Port     int
	Username string
	Name     string `mapstructure:"dbname"`
	SSLmode  string
	Password string
}

func NewPostgresConnection(inf PSQlConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		inf.Host, inf.Port, inf.Username, inf.Name, inf.SSLmode, inf.Password))
	if err != nil {
		return nil, err
	}
	return db, nil
}