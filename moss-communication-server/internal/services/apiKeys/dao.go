package apiKeys

import (
	"database/sql"

	"github.com/MattiasHenders/moss-communication-server/pkg/db"
	apiKeyModels "github.com/MattiasHenders/moss-communication-server/pkg/models/apiKeys"
	"github.com/google/uuid"
)

const (
	createApiKeyAttachedToUserQuery = `
		INSERT INTO api_keys
		(id, hashed_api_key) 
		VALUES ($1, $3)
		;
		INSERT INTO user_api_keys
		(id, user_id, api_key_id) 
		VALUES (gen_random_uuid(), $2, $1) 
		;
	`
	getApiKeyByHashedKeyQuery = `
		SELECT id, permissions, name, created_on
		FROM api_keys
		WHERE hashed_api_key = $1
		;
	`
	getApiKeysByUserIDQuery = `
		SELECT k.id, k.permissions, k.name, k.created_on
		FROM user_api_keys u
		JOIN api_keys k 
		ON u.api_key_id = k.id 
		WHERE u.user_id = $1
		;
	`
	updateApiKeyQuery = `
		UPDATE api_keys
		SET permissions = $1, name = $2
		WHERE id = $3
		RETURNING *
		;
	`
	deleteApiKeyQuery = `
		DELETE FROM api_keys
		WHERE id = $1
		;
	`
)

func createApiKeyAttachedToUserDAO(userID string, hashedKey string) error {

	apiKeyUUID := uuid.New()

	_, err := db.DB.Exec(createApiKeyAttachedToUserQuery, apiKeyUUID, userID, hashedKey)
	if err != nil {
		return err
	}

	return nil
}

func getApiKeyByHashedKeyDAO(hashedKey string) (*apiKeyModels.ApiKey, error) {
	var key apiKeyModels.ApiKey

	err := db.DB.Get(&key, getApiKeyByHashedKeyQuery, hashedKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &key, nil
}

func getApiKeysByUserIDDAO(userID string) (*[]apiKeyModels.ApiKey, error) {
	var keys []apiKeyModels.ApiKey

	err := db.DB.Get(&keys, getApiKeysByUserIDQuery, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &keys, nil
}

func updateApiKeyDAO(permissions string, name string, apiKeyID string) error {
	_, err := db.DB.Exec(updateApiKeyQuery, permissions, name, apiKeyID)
	if err != nil {
		return err
	}

	return nil
}

func deleteApiKeyDAO(apiKeyID string) error {
	_, err := db.DB.Exec(deleteApiKeyQuery, apiKeyID)
	if err != nil {
		return err
	}

	return nil
}
