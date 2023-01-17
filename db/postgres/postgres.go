package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type Config struct {
	Username, Password, Host, Port, Database string
}

func NewPostgres(ctx context.Context, conf Config) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	conn, err := pgx.Connect(context.Background(), os.Getenv(dsn))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}
