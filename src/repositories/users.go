package repositories

import (
	"api-rede-social/src/models"
	"database/sql"
	"fmt"
)

// Users is a representation of user repositorie
type Users struct {
	db *sql.DB
}

// NewUserRepositorie create a user repositories
func NewUserRepositorie(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert a user in db
func (ur *Users) Create(user models.User) (uint64, error) {
	statement, err := ur.db.Prepare(
		`INSERT INTO users (name, nick_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`,
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var lastInsertID uint64
	err = statement.QueryRow(user.Name, user.NickName, user.Email, user.Password).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Find retuns all register equals to search conditions
func (ur *Users) Find(searchConditions string) ([]models.User, error) {
	searchConditions = fmt.Sprintf("%%%s%%", searchConditions)

	rows, err := ur.db.Query("SELECT id, name, nick_name, email, created_at FROM users WHERE name LIKE $1 OR nick_name LIKE $2", searchConditions, searchConditions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID returns an user search by id
func (ur *Users) FindByID(ID uint64) (models.User, error) {
	row, err := ur.db.Query("SELECT id, name, nick_name, email, created_at FROM users WHERE id = $1", ID)
	if err != nil {
		return models.User{}, err

	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
		return user, nil
	}
	return models.User{}, fmt.Errorf("User not found")
}

// UpdateUser update a user
func (ur *Users) UpdateUser(ID uint64, user models.User) error {
	statement, err := ur.db.Prepare("UPDATE users SET name=$1, nick_name=$2, email=$3 WHERE id = $4")
	if err != nil {
		return err

	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.NickName, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete delete a usar by id
func (ur *Users) Delete(ID uint64) error {
	statement, err := ur.db.Prepare("DELETE FROM users where id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// FindByEmail search an usar by email
func (ur *Users) FindByEmail(email string) (models.User, error) {

	row, err := ur.db.Query("SELECT id, password FROM users where email =$1", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {

		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err

		}

	}

	return user, err
}

// Follow follow a user
func (ur *Users) Follow(userID, followID uint64) error {

	statement, err := ur.db.Prepare("INSERT INTO followers (users_id, follower_id) VALUES ($1, $2) ON CONFLICT (users_id, follower_id) DO NOTHING;")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userID, followID)
	if err != nil {
		return err
	}

	return nil

}

// Unfollow unfollow a user
func (ur *Users) Unfollow(userID, followID uint64) error {

	statement, err := ur.db.Prepare("DELETE FROM followers where users_id = $1 AND follower_id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

// SearchFollowers list followers from a user
func (ur *Users) SearchFollowers(userID uint64) ([]models.User, error) {

	statement := `SELECT u.id, u.name, u.nick_name, u.email, u.created_at
				FROM users u
				INNER JOIN followers f ON u.id = f.follower_id
				WHERE f.users_id = $1;`

	rows, err := ur.db.Query(statement, userID)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

// SearchFollowing lists all users that a user follows
func (ur *Users) SearchFollowing(userID uint64) ([]models.User, error) {

	statement := `SELECT u.id, u.name, u.nick_name, u.email, u.created_at
				FROM users u
				INNER JOIN followers f ON u.id = f.users_id
				WHERE f.follower_id = $1;`

	rows, err := ur.db.Query(statement, userID)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

// SearchPasswordByID return a password in hash for a user
func (ur *Users) SearchPasswordByID(userID uint64) (string, error) {

	row, err := ur.db.Query("SELECT password FROM users WHERE id = $1", userID)
	if err != nil {
		return "", err

	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err := row.Scan(
			&user.Password,
		); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// UpdatePassword update a password
func (ur *Users) UpdatePassword(userID uint64, password string) error {

	statement, err := ur.db.Prepare("UPDATE users SET password = $1 WHERE id = $2")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(password, userID); err != nil {
		return err
	}
	return nil
}
