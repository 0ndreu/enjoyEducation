package postgres

import "time"

type Storage interface {
}

type storage struct {
	client  Client
	timeout time.Duration
}

// NewStorage creates a new instance of storage.
func NewStorage(client Client, timeout time.Duration) Storage {
	return &storage{
		client:  client,
		timeout: timeout,
	}
}
