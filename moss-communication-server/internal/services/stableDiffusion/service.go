package stableDiffusion

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/MattiasHenders/moss-communication-server/pkg/errors"

	stableDiffusionModels "github.com/MattiasHenders/moss-communication-server/pkg/models/stableDiffusion"
)

func CreateTextToImageRequest() *errors.HTTPError {

	// Run the request

	// Check the status to see when it is complete

	// Check the return value

	// If all is good then parse the images from base64

	return nil
}

func CreateImageToImageRequest() *errors.HTTPError {

	return nil
}

func RunStableDiffusionRequest() *errors.HTTPError {

	return nil
}

func doRequest(r stableDiffusionModels.StableDiffusionRequest, isJSONContentType bool) (stableDiffusionModels.StableDiffusionResponse, error) {
	client := &http.Client{}

	var (
		req *http.Request
		err error
	)

	if isJSONContentType {
		req, err = http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))

		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(r.Method, r.URL, strings.NewReader(r.Data))

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Set("Authorization", "Bearer "+r.Bearer)

	// Wait some amount of msecs to avoid triggering Airtable API limits
	time.Sleep(airtableAPIWaitTimeMsecs * time.Millisecond)

	if err != nil {
		return airtableModel.Response{}, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return airtableModel.Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return airtableModel.Response{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return airtableModel.Response{}, fmt.Errorf("error: %s details: %s", resp.Status, body)
	}

	return airtableModel.Response{Body: body, StatusCode: resp.StatusCode}, nil
}
