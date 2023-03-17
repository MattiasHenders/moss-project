package auth

import (
	"net/http"
	"time"

	"github.com/MattiasHenders/moss-communication-server/internal/services/users"
	"github.com/MattiasHenders/moss-communication-server/pkg/constants"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	"github.com/MattiasHenders/moss-communication-server/pkg/middleware"
	authModels "github.com/MattiasHenders/moss-communication-server/pkg/models/auth"
	userModels "github.com/MattiasHenders/moss-communication-server/pkg/models/users"
)

func Login(w http.ResponseWriter, email string, password string) (*authModels.AuthToken, *errors.HTTPError) {

	hashedPassword := middleware.HashPassword(password)

	user, userErr := users.GetUserByEmailAndHashedPassword(email, hashedPassword)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "User not found in request: Login")
	}

	if *user.HashedPassword != hashedPassword {
		return nil, errors.NewHTTPError(nil, http.StatusUnauthorized, "Password Not Correct: Login")
	}

	return middleware.Generate1DayAuthToken(w, email)
}

func SignUp(w http.ResponseWriter, user *userModels.User) (*authModels.AuthToken, *errors.HTTPError) {

	hashedPassword := middleware.HashPassword(*user.HashedPassword)

	userErr := users.CreateUser(*user.FirstName, *user.LastName, *user.Email, *user.Country, *user.Sex, hashedPassword, *user.UserType)
	if userErr != nil {
		return nil, errors.NewHTTPError(userErr, http.StatusInternalServerError, "User not found in request: SignUp")
	}

	return middleware.Generate1DayAuthToken(w, *user.Email)
}

func Logout(w http.ResponseWriter) {

	http.SetCookie(w, &http.Cookie{
		Name:    constants.AuthAccessTokenKey,
		Expires: time.Now(),
	})
}
