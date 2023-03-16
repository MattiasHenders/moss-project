package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/internal/services/auth"
	"github.com/MattiasHenders/moss-communication-server/pkg/constants"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	h "github.com/MattiasHenders/moss-communication-server/pkg/handler"
	"github.com/MattiasHenders/moss-communication-server/pkg/models/users"
)

func LoginHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		email := h.GetFormParam(r, "email")
		if email == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing email")
		}
		password := h.GetFormParam(r, "password")
		if password == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing password")
		}

		authToken, authErr := auth.Login(w, *email, *password)
		if authErr != nil {
			return authErr
		}

		_ = json.NewEncoder(w).Encode(authToken)
		return nil
	}
}

func SignUpNormalUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user, userErr := getUserParams(r, constants.UserTypeNormal)
		if userErr != nil {
			return userErr
		}

		authToken, authErr := auth.SignUp(w, user)
		if authErr != nil {
			return authErr
		}

		_ = json.NewEncoder(w).Encode(authToken)
		return nil
	}
}

func SignUpAdminUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user, userErr := getUserParams(r, constants.UserTypeAdmin)
		if userErr != nil {
			return userErr
		}

		authToken, authErr := auth.SignUp(w, user)
		if authErr != nil {
			return authErr
		}

		_ = json.NewEncoder(w).Encode(authToken)
		return nil
	}
}

func LogoutHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		auth.Logout(w)
		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}

func getUserParams(r *http.Request, userType string) (*users.User, *errors.HTTPError) {

	firstName := h.GetFormParam(r, "firstName")
	lastName := h.GetFormParam(r, "lastName")

	email := h.GetFormParam(r, "email")
	if email == nil {
		return nil, errors.NewHTTPError(nil, http.StatusBadRequest, "Missing email")
	}

	country := h.GetFormParam(r, "country")
	sex := h.GetFormParam(r, "sex")

	password := h.GetFormParam(r, "password")
	if password == nil {
		return nil, errors.NewHTTPError(nil, http.StatusBadRequest, "Missing password")
	}

	user := users.User{
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Country:        country,
		Sex:            sex,
		HashedPassword: password,
		UserType:       &userType,
	}

	return &user, nil
}
