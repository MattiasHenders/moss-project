package users

import (
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/moss-communication-server/pkg/errors"
	"github.com/MattiasHenders/moss-communication-server/moss-communication-server/pkg/models/users"
)

func GetUserByID(id string) (*users.User, *errors.HTTPError) {

	user, userErr := getUserByIDDAO(id)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to get user from DB: GetUserByID")
	}

	return user, nil
}

func GetUserByEmail(email string) (*users.User, *errors.HTTPError) {

	user, userErr := getUserByEmailDAO(email)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to get user from DB: GetUserByEmail")
	}

	return user, nil
}

func GetUserByHashedKey(hashedKey string) (*users.User, *errors.HTTPError) {

	user, userErr := getUserByHashedKeyDAO(hashedKey)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to get user from DB: GetUserByHashedKey")
	}

	return user, nil
}

func GetUserByEmailAndHashedPassword(email string, password string) (*users.User, *errors.HTTPError) {

	user, userErr := getUserByEmailAndHashedPasswordDAO(email, password)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to get user from DB: GetUserByEmailAndHashedPassword")
	} else if user == nil {
		return nil, errors.NewHTTPError(nil, http.StatusNotFound, "User not found from DB: GetUserByEmailAndHashedPassword")
	}

	return user, nil
}

func CreateUser(first string, last string, email string, country string, sex string, hashedPassword string, userType string) *errors.HTTPError {

	userErr := createUserDAO(first, last, email, country, sex, hashedPassword, userType)
	if userErr != nil {
		return errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to create new user from DB: CreateUser")
	}

	return nil
}

func UpdateUser(first string, last string, country string, sex string, id string) (*users.User, *errors.HTTPError) {

	user, userErr := updateUserDAO(first, last, country, sex, id)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to update user from DB: UpdateUser")
	}

	return user, nil
}

func DeleteUser(id string) *errors.HTTPError {

	userErr := deleteUserDAO(id)
	if userErr != nil {
		return errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to delete user from DB: DeleteUser")
	}

	return nil
}

func DeleteUserAdmin(userEmail string) *errors.HTTPError {

	userErr := deleteUserAdminDAO(userEmail)
	if userErr != nil {
		return errors.NewHTTPError(userErr, http.StatusInternalServerError, "Failed to delete user via admin from DB: DeleteUserAdmin")
	}

	return nil
}
