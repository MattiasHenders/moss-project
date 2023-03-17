package tests

import (
	"testing"

	"github.com/MattiasHenders/moss-communication-server/moss-communication-server/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErr(t *testing.T, expectedErr *errors.HTTPError, actualErr *errors.HTTPError) {
	// Should have error
	assert.NotNil(t, actualErr, "httpErr should not be nil.")

	// Status of error should match
	assert.Equal(t, expectedErr.Status, actualErr.Status, "HTTPError Incorrect Status Code")

	// Message should match
	assert.Equal(t, expectedErr.Message, actualErr.Message, "HTTPError Incorrect Error Message")
}
