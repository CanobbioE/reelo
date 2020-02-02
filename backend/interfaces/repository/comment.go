package repository

import (
	"context"

	"github.com/CanobbioE/reelo/backend/domain"
)

// COMMENTREPO is the handler name
const COMMENTREPO = "commentRepo"

// DbCommentRepo id the repository for comments
type DbCommentRepo DbRepo

// NewDbCommentRepo istanciates and returns a comment repository
func NewDbCommentRepo(dbHandlers map[string]DbHandler) *DbCommentRepo {
	return &DbCommentRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers[COMMENTREPO],
	}
}

// Store inserts a new comment in the repository
func (db *DbCommentRepo) Store(ctx context.Context, c domain.Comment) (int, error) {
	return 0, nil
}

// FindByID retrieves a comment from the repository, given the id
func (db *DbCommentRepo) FindByID(ctx context.Context) (domain.Comment, error) {
}

// UpdateTextByPlayerID updates a comment defined by the player's id
func (db *DbCommentRepo) UpdateTextByPlayerID(ctx context.Context, id int, t string) error {}

// CheckExistenceByPlayerID returns true if the repo contains a comment with for the given player id
func (db *DbCommentRepo) CheckExistenceByPlayerID(ctx context.Context, id int) bool {}
