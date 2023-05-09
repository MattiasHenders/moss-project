package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MattiasHenders/moss-communication-server/internal/services/stableDiffusion"
	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	h "github.com/MattiasHenders/moss-communication-server/pkg/handler"
)

func CreateTextToImageRequestHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		prompt := h.GetFormParam(r, "prompt")
		if prompt == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing prompt")
		}

		seed := h.GetFormParamInt64(r, "seed")
		numOutputs := h.GetFormParamInt(r, "num_outputs")
		width := h.GetFormParamInt(r, "width")
		height := h.GetFormParamInt(r, "height")
		numInferenceSteps := h.GetFormParamInt(r, "num_inference_steps")
		guidanceScale := h.GetFormParamFloat64(r, "guidance_scale")

		// initImage := h.GetFormParamInt(r, "init_image")  // TODO
		// strength := h.GetFormParamFloat64(r, "strength") // TODO

		images, httpErr := stableDiffusion.CreateTextToImageRequest(*prompt, seed, numOutputs, width, height, numInferenceSteps, guidanceScale, nil, nil)
		if httpErr != nil {
			return httpErr
		} else if len(images) == 0 {
			return errors.NewHTTPError(nil, http.StatusInternalServerError, "Didnt make any images")
		}

		_, err := w.Write([]byte(images[0]))
		if err != nil {
			return errors.NewHTTPError(err, http.StatusInternalServerError, "Failed to convert image")
		}

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
