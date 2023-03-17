package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MattiasHenders/moss-communication-server/internal/services/apiKeys"
	"github.com/MattiasHenders/moss-communication-server/internal/services/users"
	"github.com/MattiasHenders/moss-communication-server/pkg/constants"
	"github.com/MattiasHenders/moss-communication-server/pkg/handler"
	apiKeyModels "github.com/MattiasHenders/moss-communication-server/pkg/models/apiKeys"
	authModels "github.com/MattiasHenders/moss-communication-server/pkg/models/auth"
	userModels "github.com/MattiasHenders/moss-communication-server/pkg/models/users"

	"github.com/MattiasHenders/moss-communication-server/internal/utils"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	"github.com/MattiasHenders/moss-communication-server/pkg/secrets"
	"github.com/golang-jwt/jwt/v4"
)

var (
	userCtxKey = &contextKey{"User"}
)

type contextKey struct {
	name string
}

func HashPassword(rawPassword string) string {
	secretData := secrets.LoadEnvAndGetSecrets()

	return utils.HashStringWithSHA256AndSalt(secretData.PasswordSecret, rawPassword)
}

func Generate1DayAuthToken(w http.ResponseWriter, email string) (*authModels.AuthToken, *errors.HTTPError) {

	secretData := secrets.LoadEnvAndGetSecrets()

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &authModels.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretData.PasswordSecret))
	if err != nil {
		return nil, errors.NewHTTPError(nil, http.StatusInternalServerError, "Could not get API key for authenticated user")
	}

	http.SetCookie(w, &http.Cookie{
		Name:    constants.AuthAccessTokenKey,
		Value:   tokenString,
		Expires: expirationTime,
	})

	return &authModels.AuthToken{
		AccessToken: tokenString,
		Expires:     expirationTime.UnixMilli(),
	}, nil
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(handler.Handler(func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		secretData := secrets.LoadEnvAndGetSecrets()

		tknStr := getJWTFromCookie(r)
		claims := &authModels.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretData.PasswordSecret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, "Unauthorized")
			}
			return errors.NewHTTPError(err, http.StatusBadRequest, "Bad Auth Request Parsing Claims")
		}
		if !tkn.Valid {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Unauthorized: Not Valid")
		}

		if time.Until(claims.ExpiresAt.Time) < 0 {
			return errors.NewHTTPError(err, http.StatusUnauthorized, "Unauthorized: Expired")
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(24 * time.Hour)
		claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secretData.PasswordSecret))
		if err != nil {
			return errors.NewHTTPError(err, http.StatusBadRequest, "Bad Auth Request")
		}

		http.SetCookie(w, &http.Cookie{
			Name:    constants.AuthAccessTokenKey,
			Value:   tokenString,
			Expires: expirationTime,
		})

		// Add the token to the context to be passed on if needed
		ctx := context.Background()
		ctx = context.WithValue(ctx, constants.AuthContextEmailKey, claims.Email)

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r.WithContext(ctx))
		return nil
	}))
}

func GetUserFromAuthContext(ctx context.Context) *userModels.User {
	userEmail := ctx.Value(constants.AuthContextEmailKey).(string)
	user, _ := users.GetUserByEmail(userEmail)
	return user
}

func VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(handler.Handler(func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
		authUser := r.Context().Value(userCtxKey).(*userModels.User)

		if *authUser.UserType != constants.UserTypeAdmin {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Auth user must be an admin")
		}

		next.ServeHTTP(w, r)
		return nil
	}))
}

func getJWTFromCookie(r *http.Request) string {
	cookie, cookieErr := r.Cookie(constants.AuthAccessTokenKey)
	if cookieErr != nil {
		return getJWTFromBearerAuth(r.Header.Get("Authorization"))
	} else {
		return cookie.Value
	}
}

func getJWTFromBearerAuth(fullAuth string) string {
	if strings.Contains(fullAuth, "Bearer") {
		splitAuth := strings.Split(fullAuth, " ")
		if len(splitAuth) > 1 {
			return splitAuth[1]
		}
	}
	return ""
}

func VerifyAPIKey(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

			rawApiKey := r.Header.Get("apiKey")
			if rawApiKey == "" {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, "No apiKey found in header")
			}

			if rawApiKey != apiKey {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, "apiKey is incorrect")
			}

			next.ServeHTTP(w, r)
			return nil
		}
		return http.HandlerFunc(handler.Handler(fn))
	}
}

func VerifyAPIKeyPermissions(permissionsNeeded []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

			rawApiKey := r.Header.Get("apiKey")
			if rawApiKey == "" {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, "No apiKey found in header: VerifyAPIKeyPermissions")
			}

			apiKey, apiKeyErr := apiKeys.GetApiKeyFromUnhashedAPIKey(rawApiKey)
			if apiKeyErr != nil {
				return errors.NewHTTPError(apiKeyErr, http.StatusInternalServerError, "Failed getting apiKey: VerifyAPIKeyPermissions")
			} else if apiKey == nil {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, "No apiKey found: VerifyAPIKeyPermissions")
			}

			ok, missingPermissions := apiKeyHasPermissionsOrGetMissingPermissions(apiKey, permissionsNeeded)
			if !ok {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, fmt.Sprintf("API Key does not have all needed permissions.\nMissing permissions: %s", *missingPermissions))
			}

			next.ServeHTTP(w, r)
			return nil
		}
		return http.HandlerFunc(handler.Handler(fn))
	}
}

func apiKeyHasPermissionsOrGetMissingPermissions(apiKey *apiKeyModels.ApiKey, rawPermissionsNeeded []string) (bool, *string) {

	if len(rawPermissionsNeeded) == 0 {
		return true, nil
	}

	rawPermissionsFound := apiKey.Permissions
	if rawPermissionsFound == nil {
		return false, formatMissingPermissions(rawPermissionsNeeded)
	}

	permissionsFound := strings.Split(*rawPermissionsFound, ",")
	permissionsNeeded := rawPermissionsNeeded

	missingPermissions := findMissingPermissions(permissionsNeeded, permissionsFound)
	if len(missingPermissions) > 0 {
		return false, formatMissingPermissions(missingPermissions)
	}
	return true, nil
}

func findMissingPermissions(a []string, b []string) []string {
	var missingPermissions []string
	for i := 0; i < len(a); i++ {
		m := len(b)
		var j int
		for j = 0; j < m; j++ {
			if a[i] == b[j] {
				break
			}
		}
		if j == m {
			missingPermissions = append(missingPermissions, a[i])
		}
	}
	return missingPermissions
}

func formatMissingPermissions(permissionsMissing []string) *string {
	formattedPermissions := strings.Join(permissionsMissing, ", ")
	formattedPermissions = strings.TrimSuffix(formattedPermissions, ", ")
	return &formattedPermissions
}
