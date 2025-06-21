package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

const driver = "pgx"

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	Dbname          string
	SSLMode         bool
	ConnPoolEnabled bool
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

func NewConnection(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Dbname,
		cfg.Password,
	)

	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return nil, errors.New("Failed database connect: " + err.Error())
	}

	if cfg.ConnPoolEnabled {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
		db.SetMaxIdleConns(cfg.MaxIdleConns)
		db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, errors.New("Failed database ping: " + err.Error())
	}

	return db, nil
}
