package apiKeys

import (
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/internal/utils"
	"github.com/MattiasHenders/moss-communication-server/pkg/constants"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	apiKeyModels "github.com/MattiasHenders/moss-communication-server/pkg/models/apiKeys"
	"github.com/MattiasHenders/moss-communication-server/pkg/models/users"
	"github.com/MattiasHenders/moss-communication-server/pkg/secrets"
)

func CreateApiKeyAttachedToUser(user *users.User) (*string, *errors.HTTPError) {

	secretData := secrets.LoadEnvAndGetSecrets()

	unhashedAPIKey := utils.GenerateUnhashedAPIKeyWithSHA1(constants.ApiKeysPrefix)
	hashedKey := utils.HashStringWithSHA256AndSalt(secretData.HashSalt, unhashedAPIKey)

	keyErr := createApiKeyAttachedToUserDAO(*user.ID, hashedKey)
	if keyErr != nil {
		return nil, errors.NewHTTPError(keyErr, http.StatusInternalServerError, "Failed to attach key to user: CreateApiKeyAttachedToUser")
	}

	return &unhashedAPIKey, nil
}

func GetApiKeyFromUnhashedAPIKey(unhashedAPIKey string) (*apiKeyModels.ApiKey, *errors.HTTPError) {

	secretData := secrets.LoadEnvAndGetSecrets()

	hashedKey := utils.HashStringWithSHA256AndSalt(secretData.HashSalt, unhashedAPIKey)
	key, keyErr := getApiKeyByHashedKeyDAO(hashedKey)
	if keyErr != nil {
		return nil, errors.NewHTTPError(keyErr, http.StatusInternalServerError, "Failed to get key: GetApiKeyFromUnhashedAPIKey")
	}

	return key, nil
}

func GetApiKeysFromUserID(userID string) (*[]apiKeyModels.ApiKey, *errors.HTTPError) {

	keys, keyErr := getApiKeysByUserIDDAO(userID)
	if keyErr != nil {
		return nil, errors.NewHTTPError(keyErr, http.StatusInternalServerError, "Failed to get key: GetApiKeysFromUserID")
	}

	return keys, nil
}

func UpdateApiKey(permissions string, name string, apiKeyID string) *errors.HTTPError {

	keyErr := updateApiKeyDAO(permissions, name, apiKeyID)
	if keyErr != nil {
		return errors.NewHTTPError(keyErr, http.StatusInternalServerError, "Failed to update key: UpdateApiKey")
	}

	return nil
}

func DeleteApiKey(apiKeyID string) *errors.HTTPError {

	keyErr := deleteApiKeyDAO(apiKeyID)
	if keyErr != nil {
		return errors.NewHTTPError(keyErr, http.StatusInternalServerError, "Failed to delete key: DeleteApiKey")
	}

	return nil
}
