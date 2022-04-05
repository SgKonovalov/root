package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

// urlExample := "postgres://username:password@localhost:5432/database_name" для подключения к postgres.

//Общий интерфейс для базы данных.
type GeneralDB interface {
	GetConnection() (*pgxpool.Pool, error)
}

// Структура для работы с PostgreSQL
type PostgreDBbyPGX struct {
	Config  *MainConfig
	Context context.Context
}

func NewPostgreDBbyPGX(Config *MainConfig, Context context.Context) *PostgreDBbyPGX {
	return &PostgreDBbyPGX{
		Config:  Config,
		Context: Context,
	}
}

func (pspgx *PostgreDBbyPGX) GetConnection() (*pgxpool.Pool, error) {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s",
		pspgx.Config.Database.Username,
		pspgx.Config.Database.Password,
		pspgx.Config.Database.Host,
		pspgx.Config.Database.DatabaseName)

	dbPool, err := pgxpool.Connect(pspgx.Context, dbURL)

	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
