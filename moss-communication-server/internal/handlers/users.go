package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/internal/services/users"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	h "github.com/MattiasHenders/moss-communication-server/pkg/handler"
	"github.com/MattiasHenders/moss-communication-server/pkg/middleware"
)

func GetUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user := middleware.GetUserFromAuthContext(r.Context())
		if user == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "User not found in request")
		}

		_ = json.NewEncoder(w).Encode(user)
		return nil
	}
}

func UpdateUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user := middleware.GetUserFromAuthContext(r.Context())
		if user == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "User not found in request")
		}

		first := h.GetFormParam(r, "first")
		if first == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing first")
		}

		last := h.GetFormParam(r, "last")
		if last == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing last")
		}

		country := h.GetFormParam(r, "country")
		if country == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing country")
		}

		sex := h.GetFormParam(r, "sex")
		if sex == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing sex")
		}

		user, userErr := users.UpdateUser(*first, *last, *country, *sex, *user.ID)
		if userErr != nil {
			return userErr
		}

		_ = json.NewEncoder(w).Encode(user)
		return nil
	}
}

func DeleteUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user := middleware.GetUserFromAuthContext(r.Context())
		if user == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "User not found in request")
		}

		userErr := users.DeleteUser(*user.ID)
		if userErr != nil {
			return userErr
		}

		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}

func DeleteUserAdminHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		userEmail := h.GetFormParam(r, "userEmail")
		if userEmail == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing userEmail")
		}

		userErr := users.DeleteUserAdmin(*userEmail)
		if userErr != nil {
			return userErr
		}

		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}
