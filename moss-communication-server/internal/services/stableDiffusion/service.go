package stableDiffusion

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	"github.com/MattiasHenders/moss-communication-server/pkg/secrets"

	stableDiffusionModels "github.com/MattiasHenders/moss-communication-server/pkg/models/stableDiffusion"
)

const (
	baseURL                        = "https://api.runpod.ai"
	runpodID                       = "j0czf1bvlbn9r6"
	statusComplete                 = "COMPLETED"
	maxStableDiffusionRequestCount = 20
)

func CreateTextToImageRequest(prompt string, seed *int64, numOutputs *int, width *int, height *int, numInferenceSteps *int, guidanceScale *float64, initImage *int, strength *float64) ([]image.Image, *errors.HTTPError) {

	// Build the request
	requestInput := BuildRequestInput(prompt, seed, numOutputs, width, height, numInferenceSteps, guidanceScale, initImage, strength)

	// Run the request
	stableDiffusionRes, err := StableDiffusionRunRequest(stableDiffusionModels.StableDiffusionRequest{Input: requestInput})
	if err != nil {
		return []image.Image{}, errors.NewHTTPError(err, http.StatusInternalServerError, "Error sending request to stable diffusion server")
	}

	// Check the status to see when it is complete
	imageBase64, err := LoopUntilRequestFinishedAndImageIsGenerated(stableDiffusionRes.ID)
	if err != nil {
		return []image.Image{}, errors.NewHTTPError(err, http.StatusInternalServerError, "Error checking status from stable diffusion server")
	}

	// If all is good then parse the images from base64
	images := ConvertBase64StringsIntoImages(imageBase64)

	return images, nil
}

func CreateImageToImageRequest() *errors.HTTPError {

	return nil
}

func StableDiffusionRunRequest(request stableDiffusionModels.StableDiffusionRequest) (*stableDiffusionModels.StableDiffusionResponse, error) {

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	res, err := doRequest(stableDiffusionModels.Request{
		URL:          fmt.Sprintf("%s/v1/%s/run", baseURL, runpodID),
		Method:       http.MethodPost,
		OkStatusCode: http.StatusOK,
		Body:         body,
	})
	if err != nil {
		return nil, err
	}

	var stableDiffusionRes stableDiffusionModels.StableDiffusionResponse

	err = json.Unmarshal(res.Body, &stableDiffusionRes)
	if err != nil {
		return nil, err
	}

	return &stableDiffusionRes, nil
}

func StableDiffusionStatusRequest(runpodJobID string) (*stableDiffusionModels.StableDiffusionResponse, error) {

	res, err := doRequest(stableDiffusionModels.Request{
		URL:          fmt.Sprintf("%s/v1/%s/status/%s", baseURL, runpodID, runpodJobID),
		Method:       http.MethodPost,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}

	var stableDiffusionRes stableDiffusionModels.StableDiffusionResponse

	err = json.Unmarshal(res.Body, &stableDiffusionRes)
	if err != nil {
		return nil, err
	}

	return &stableDiffusionRes, nil
}

func doRequest(r stableDiffusionModels.Request) (stableDiffusionModels.Response, error) {

	secrets := secrets.LoadEnvAndGetSecrets()
	client := &http.Client{}

	var (
		req *http.Request
		err error
	)

	req, err = http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+secrets.RunpodAPIKey)

	if err != nil {
		return stableDiffusionModels.Response{}, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return stableDiffusionModels.Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return stableDiffusionModels.Response{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return stableDiffusionModels.Response{}, fmt.Errorf("error: %s details: %s", resp.Status, body)
	}

	return stableDiffusionModels.Response{Body: body, StatusCode: resp.StatusCode}, nil
}

func BuildRequestInput(prompt string, seed *int64, numOutputs *int, width *int, height *int, numInferenceSteps *int, guidanceScale *float64, initImage *int, strength *float64) stableDiffusionModels.StableDiffusionInput {

	one := 1

	return stableDiffusionModels.StableDiffusionInput{
		Prompt:           fmt.Sprintf("%s painted in makeshipBipedalV2 style", prompt),
		Seed:             seed,
		NumOutputs:       &one, // TODO remove the hardcoded
		Width:            width,
		Height:           height,
		NumInfrenceSteps: numInferenceSteps,
		GuidanceScale:    guidanceScale,
		InitImage:        initImage,
		Strength:         strength,
	}
}

func LoopUntilRequestFinishedAndImageIsGenerated(runpodID string) ([]string, error) {

	requestCount := 0

	for {

		res, err := StableDiffusionStatusRequest(runpodID)
		if err != nil {
			return []string{}, err
		}

		if res.Status == statusComplete {
			return res.Output.Images, nil
		}

		if requestCount > maxStableDiffusionRequestCount {
			return []string{}, fmt.Errorf("request limit reached")
		}

		requestCount++
		time.Sleep(time.Second)
	}
}

func ConvertBase64StringsIntoImages(imageBase64 []string) []image.Image {

	var imgArray []image.Image

	for _, imageBase64 := range imageBase64 {

		decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageBase64))
		png, errPng := png.Decode(decoder)
		if errPng != nil {
			continue
		}

		imgArray = append(imgArray, png)
	}
	return imgArray
}
