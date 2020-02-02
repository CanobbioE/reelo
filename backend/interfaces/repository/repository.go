package repository

import "context"

// DbHandler is a high level interface for repository interrogation
type DbHandler interface {
	Execute(ctx context.Context, stmt string) error
	Query(ctx context.Context, stmt string) (Row, error)
}

// Row is a high level interface that allows repository's data manipulation
type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

// DbRepo represents a general repository.
// The handlers map lets every repository use any other repository
// respecting dependency injection
type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}
