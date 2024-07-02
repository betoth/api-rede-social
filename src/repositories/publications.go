package repositories

import (
	"api-rede-social/src/models"
	"database/sql"
)

// Publications represents a repository of Publications
type Publications struct {
	db *sql.DB
}

// NewPublicationRepositorie creates a new instance of the Publication struct and returns
func NewPublicationRepositorie(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create Generates a new publication of a user in the database
func (p *Publications) Create(pub models.Publication, userID uint64) (uint64, error) {

	statement, err := p.db.Prepare(
		`INSERT INTO publications ( title, content, author_id) VALUES ($1, $2, $3) RETURNING id`,
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var lastInsertID uint64

	if err = statement.QueryRow(pub.Title, pub.Content, userID).Scan(&lastInsertID); err != nil {
		return 0, err
	}

	return lastInsertID, nil

}
