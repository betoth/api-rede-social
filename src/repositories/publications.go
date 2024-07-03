package repositories

import (
	"api-rede-social/src/models"
	"database/sql"
	"errors"
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

// FindPublicationByID search a user from ID
func (p *Publications) FindPublicationByID(pubID uint64) (models.Publication, error) {

	row, err := p.db.Query(
		`select p.*, u.nick_name  from publications p 
		inner join users u on p.author_id = u.id 
		where p.id = $1`, pubID)
	if err != nil {
		return models.Publication{}, err
	}
	defer row.Close()

	var pub models.Publication

	if row.Next() {
		if err := row.Scan(
			&pub.ID,
			&pub.Title,
			&pub.Content,
			&pub.AuthorID,
			&pub.Likes,
			&pub.CreatedAt,
			&pub.AuthorNickName,
		); err != nil {
			return models.Publication{}, err
		}

		return pub, nil
	}
	return models.Publication{}, errors.New("Publication not found")
}

// FindPublications list publications from a user and users he follows
func (p *Publications) FindPublications(userID uint64) ([]models.Publication, error) {

	rows, err := p.db.Query(
		`select p.*, u.nick_name from publications p
		inner join users u on p.author_id = u.id 
		inner join followers f  on f.users_id = u.id 
		where p.author_id = 69 or f.follower_id = $1
		ORDER BY 1 DESC`, userID)
	if err != nil {
		return []models.Publication{}, err
	}
	defer rows.Close()

	var pubs []models.Publication

	for rows.Next() {
		var pub models.Publication
		if err := rows.Scan(

			&pub.ID,
			&pub.Title,
			&pub.Content,
			&pub.AuthorID,
			&pub.Likes,
			&pub.CreatedAt,
			&pub.AuthorNickName,
		); err != nil {
			return []models.Publication{}, err
		}
		pubs = append(pubs, pub)
	}

	return pubs, nil
}

// UpdatePublication update a publication
func (p *Publications) UpdatePublication(pubID uint64, pub models.Publication) error {
	statement, err := p.db.Prepare("UPDATE publications SET title = $1, content = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(pub.Title, pub.Content, pubID)
	if err != nil {
		return err
	}
	return nil
}

// DeletePublication delete a publication by id
func (p *Publications) DeletePublication(pubID uint64) error {

	stement, err := p.db.Prepare("DELETE FROM publications WHERE id = $1")
	if err != nil {
		return err
	}
	defer stement.Close()

	_, err = stement.Exec(pubID)
	if err != nil {
		return err
	}

	return nil
}

// FindPublicationByUser find all publications from an user
func (p *Publications) FindPublicationByUser(userID uint64) ([]models.Publication, error) {

	rows, err := p.db.Query(
		`select p.*, u.nick_name  from publications p 
		inner join users u on p.author_id = u.id 
		where U.id = $1`, userID)
	if err != nil {
		return []models.Publication{}, err
	}
	defer rows.Close()

	var pubs []models.Publication

	for rows.Next() {
		var pub models.Publication
		if err := rows.Scan(

			&pub.ID,
			&pub.Title,
			&pub.Content,
			&pub.AuthorID,
			&pub.Likes,
			&pub.CreatedAt,
			&pub.AuthorNickName,
		); err != nil {
			return []models.Publication{}, err
		}
		pubs = append(pubs, pub)
	}

	return pubs, nil

}

// LikePublication like a publication
func (p *Publications) LikePublication(pubID uint64) error {
	statement, err := p.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(pubID)
	if err != nil {
		return err
	}

	return nil
}

// UnlikePublication like a publication
func (p *Publications) UnlikePublication(pubID uint64) error {
	statement, err := p.db.Prepare(`UPDATE publications 
		SET likes = likes - 1 
		WHERE id = $1 AND likes > 0`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(pubID)
	if err != nil {
		return err
	}

	return nil
}
