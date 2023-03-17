package users

import (
	"database/sql"

	"github.com/MattiasHenders/moss-communication-server/pkg/db"
	"github.com/MattiasHenders/moss-communication-server/pkg/models/users"
)

const (
	getUserByIDQuery = `
		SELECT * FROM users
		WHERE id = $1
		;
	`
	getUserByEmailQuery = `
		SELECT * FROM users
		WHERE email = $1
		;
	`
	getUserByHashedKeyQuery = `
		SELECT u.* FROM users u
		JOIN user_apiKeys k ON u.id = k.user_id	
		WHERE k.hashed_api_key = $1
		;
	`
	getUserByEmailAndHashedPasswordQuery = `
		SELECT * FROM users
		WHERE email = $1 AND hashed_password = $2
		;
	`
	createUserQuery = `
		INSERT INTO users (id, first_name, last_name, email, country, sex, hashed_password, user_type)
		VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7)
		;
	`
	updateUserQuery = `
		UPDATE users 
		SET first_name = $1, last_name = $2, country = $3, sex = $4
		WHERE id = $5
		RETURNING *
		;
	`
	deleteUserViaIDQuery = `
		DELETE FROM users 
		WHERE id = $1
		;
	`
	deleteUserViaEmailQuery = `
		DELETE FROM users 
		WHERE email = $1
		;
	`
)

func getUserByIDDAO(id string) (*users.User, error) {
	var user users.User

	err := db.DB.Get(&user, getUserByIDQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func getUserByEmailDAO(email string) (*users.User, error) {
	var user users.User

	err := db.DB.Get(&user, getUserByEmailQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func getUserByHashedKeyDAO(hashedKey string) (*users.User, error) {
	var user users.User

	err := db.DB.Get(&user, getUserByHashedKeyQuery, hashedKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func getUserByEmailAndHashedPasswordDAO(email string, password string) (*users.User, error) {
	var user users.User

	err := db.DB.Get(&user, getUserByEmailAndHashedPasswordQuery, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func createUserDAO(first string, last string, email string, country string, sex string, hashedPassword string, userType string) error {

	_, err := db.DB.Exec(createUserQuery, first, last, email, country, sex, hashedPassword, userType)
	if err != nil {
		return err
	}

	return nil
}

func updateUserDAO(first string, last string, country string, sex string, id string) (*users.User, error) {
	var user users.User

	err := db.DB.Get(&user, updateUserQuery, first, last, country, sex, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func deleteUserDAO(id string) error {

	_, err := db.DB.Exec(deleteUserViaIDQuery, id)
	if err != nil {
		return err
	}

	return nil
}

func deleteUserAdminDAO(userEmail string) error {

	_, err := db.DB.Exec(deleteUserViaEmailQuery, userEmail)
	if err != nil {
		return err
	}

	return nil
}
