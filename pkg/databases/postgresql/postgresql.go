package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close()
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
	Stat() *pgxpool.Stat
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PgConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewClient creates new postgres client.
func NewClient(ctx context.Context, maxAttempts int, maxDelay time.Duration, config PgConfig, binary bool) (pool *pgxpool.Pool, err error) {

	dsn := config.ConnStringFromCfg()

	pgxCfg, parseConfigErr := pgxpool.ParseConfig(dsn)
	if parseConfigErr != nil {
		log.Printf("Unable to parse config: %v\n", parseConfigErr)
		return nil, parseConfigErr
	}

	if binary {
		pgxCfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
	}

	pool, parseConfigErr = pgxpool.NewWithConfig(ctx, pgxCfg)
	if parseConfigErr != nil {
		log.Printf("Failed to parse PostgreSQL configuration due to error: %v\n", parseConfigErr)
		return nil, parseConfigErr
	}

	err = DoWithAttempts(func() error {
		pingErr := pool.Ping(ctx)
		if pingErr != nil {
			log.Printf("Failed to connect to postgres due to error %v... Going to do the next attempt\n", pingErr)
			return pingErr
		}

		return nil
	}, maxAttempts, maxDelay)
	if err != nil {
		log.Printf("All attempts are exceeded. unable to connect to PostgreSQL: %v", err)
		return nil, err
	}

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--
			continue
		}

		return nil
	}

	return err
}

func (c *PgConfig) ConnStringFromCfg() string {
	url := strings.Builder{}
	url.WriteString("postgresql://")
	url.WriteString(c.Username)
	url.WriteString(":")
	url.WriteString(c.Password)
	url.WriteString("@")
	url.WriteString(c.Host)
	url.WriteString(":")
	url.WriteString(c.Port)
	url.WriteString("/")
	url.WriteString(c.DBName)

	return url.String()
}
