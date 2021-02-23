package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// List get all users on database and get per filter too
func (u Users) List(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	db, err := (u.db.Query(
		`SELECT 
			id, 
			name, 
			nick, 
			email, 
			createat 
		from users 
		where name LIKE ? or nick LIKE ?`,
		nameOrNick,
		nameOrNick,
	))

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var users []models.User

	for db.Next() {
		var user models.User
		if err = db.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID returns an user based on ID
func (u Users) GetUserByID(userID uint64) (models.User, error) {
	db, err := (u.db.Query(
		`SELECT 
      id, 
      name, 
      nick, 
      email, 
      createat 
		from users 
		where id = ?`,
		userID,
	))

	if err != nil {
		return models.User{}, err
	}

	defer db.Close()

	var user models.User

	if db.Next() {
		if err = db.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// UpdateUserByID update an user by ID, but not update key
func (u Users) UpdateUserByID(userID uint64, user models.User) error {
	statement, err := u.db.Prepare(
		`UPDATE users set
			name = ?, 
			nick = ?, 
			email = ?, 
		where id = ?`,
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}
