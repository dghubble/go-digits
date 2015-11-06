package digits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAPIError = &APIError{
	Errors: []ErrorDetail{
		ErrorDetail{Message: "Could not authenticate you.", Code: 32},
	},
}
var errTestError = fmt.Errorf("unknown host")

func TestAPIError_ErrorString(t *testing.T) {
	err := &APIError{}
	assert.Equal(t, "", err.Error())
	assert.Equal(t, "digits: 32 Could not authenticate you.", testAPIError.Error())
}

func TestAPIError_Empty(t *testing.T) {
	err := APIError{}
	assert.True(t, err.Empty())
	assert.False(t, testAPIError.Empty())
}

func TestFirstError(t *testing.T) {
	cases := []struct {
		httpError error
		apiError  *APIError
		expected  error
	}{
		{nil, nil, nil},
		{nil, &APIError{}, nil},
		{nil, testAPIError, testAPIError},
		{errTestError, &APIError{}, errTestError},
		{errTestError, testAPIError, errTestError},
	}
	for _, c := range cases {
		err := firstError(c.httpError, c.apiError)
		assert.Equal(t, c.expected, err)
	}
}
