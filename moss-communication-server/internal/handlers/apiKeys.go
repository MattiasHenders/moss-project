package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/moss-communication-server/internal/services/apiKeys"
	"github.com/MattiasHenders/moss-communication-server/moss-communication-server/pkg/errors"
	h "github.com/MattiasHenders/moss-communication-server/moss-communication-server/pkg/handler"
	"github.com/MattiasHenders/moss-communication-server/moss-communication-server/pkg/middleware"
)

func CreateApiKeyAttachedToUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user := middleware.GetUserFromAuthContext(r.Context())
		if user == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "User not found in request")
		}

		rawApiKey, apiKeyErr := apiKeys.CreateApiKeyAttachedToUser(user)
		if apiKeyErr != nil {
			return apiKeyErr
		} else if rawApiKey == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "Could not get API key for authenticated user")
		}

		_ = json.NewEncoder(w).Encode(rawApiKey)
		return nil
	}
}

func GetApiKeysFromUserHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		user := middleware.GetUserFromAuthContext(r.Context())
		if user == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "User not found in request")
		}

		apiKeys, apiKeysErr := apiKeys.GetApiKeysFromUserID(*user.ID)
		if apiKeysErr != nil {
			return apiKeysErr
		} else if apiKeys == nil {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "Could not get users API keys")
		}

		_ = json.NewEncoder(w).Encode(*apiKeys)
		return nil
	}
}

func UpdateApiKeyHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		permissions := h.GetFormParam(r, "permissions")
		if permissions == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing permissions")
		}
		name := h.GetFormParam(r, "name")
		if name == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing name")
		}
		apiKeyID := h.GetFormParam(r, "apiKeyID")
		if apiKeyID == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing apiKeyID")
		}

		apiKeysErr := apiKeys.UpdateApiKey(*permissions, *name, *apiKeyID)
		if apiKeysErr != nil {
			return apiKeysErr
		}

		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}

func DeleteApiKeyHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		apiKeyID := h.GetFormParam(r, "apiKeyID")
		if apiKeyID == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing apiKeyID")
		}

		apiKeysErr := apiKeys.DeleteApiKey(*apiKeyID)
		if apiKeysErr != nil {
			return apiKeysErr
		}

		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}
