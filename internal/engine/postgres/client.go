package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	// errNilDriver is returned by call any method which use connection to database if connection is nil
	errNilDriver = errors.New("nil mssql driver")
)

// Client interface
type Client interface {
	Start() (err error)
	GetClient() (db *sqlx.DB, err error)
	Ping(ctx context.Context) (err error)
	Shutdown() (err error)
}

type client struct {
	db *sqlx.DB

	engine      string
	dsn         string
	maxOpenConn int
	maxIdleConn int
	maxLifeConn time.Duration
	timeout     time.Duration
}

// Start opens a database specified by its database driver name and verifies the connection
func (c *client) Start() (err error) {
	db, err := sqlx.Open(c.engine, c.dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(c.maxOpenConn)
	db.SetMaxIdleConns(c.maxIdleConn)
	db.SetConnMaxLifetime(c.maxLifeConn)
	c.db = db

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	err = c.Ping(ctx)
	return
}

// GetClient returns connection to database
func (c *client) GetClient() (db *sqlx.DB, err error) {
	if c.db == nil {
		err = errNilDriver
		return
	}
	db = c.db
	return
}

// Ping verifies the connection to the database is still alive
func (c *client) Ping(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	err = c.db.PingContext(ctx)
	return
}

// Close closes the database
func (c *client) Shutdown() (err error) {
	err = c.db.Close()
	return
}

// NewClient creates a new instance of client
func NewClient(
	engine string,
	dsn string,
	maxOpenConn int,
	maxIdleConn int,
	maxLifeConn time.Duration,
	timeout time.Duration,
) Client {
	return &client{
		engine:      engine,
		dsn:         dsn,
		maxOpenConn: maxOpenConn,
		maxIdleConn: maxIdleConn,
		maxLifeConn: maxLifeConn,
		timeout:     timeout,
	}
}
