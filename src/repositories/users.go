package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users representes a user repositorie
type Users struct {
	db *sql.DB
}

// NewUserRepositorie create an repositorie of users
func NewUserRepositorie(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert an user on database
func (u Users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		`INSERT INTO users (
			name, 
			nick, 
			email, 
			password
		) values (?, ?, ?, ?)`,
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		user.Password,
	)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
