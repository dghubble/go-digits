package digits

import (
	"fmt"
	"reflect"
	"testing"
)

var testAPIError = &APIError{
	Errors: []ErrorDetail{
		ErrorDetail{Message: "Could not authenticate you.", Code: 32},
	},
}
var errTestError = fmt.Errorf("unknown host")

func TestAPIError_ErrorString(t *testing.T) {
	err := APIError{}
	if err.Error() != "" {
		t.Errorf("expected \"\", got %v", err)
	}
	expected := "digits: 32 Could not authenticate you."
	if err := testAPIError.Error(); err != expected {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestAPIError_Empty(t *testing.T) {
	err := APIError{}
	if !err.Empty() {
		t.Errorf("expected Empty() to return true for %v", err)
	}
	if testAPIError.Empty() {
		t.Errorf("expected Empty() to return false for %v", testAPIError)
	}
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
		if !reflect.DeepEqual(c.expected, err) {
			t.Errorf("not DeepEqual: expected %v, got %v", c.expected, err)
		}
	}
}
