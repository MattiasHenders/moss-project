package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/internal/services/stableDiffusion"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
)

func CreateTextToImageRequestHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		httpErr := stableDiffusion.CreateTextToImageRequest()
		if httpErr != nil {
			return httpErr
		}

		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}

func CreateImageToImageRequestHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		httpErr := stableDiffusion.CreateImageToImageRequest()
		if httpErr != nil {
			return httpErr
		}

		_ = json.NewEncoder(w).Encode("success")
		return nil
	}
}
