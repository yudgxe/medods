package handlers

import (
	"errors"
	"fmt"
	httper "medods/http"
	"medods/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

var _ services.AuthService = (*authServiceMock)(nil)

type authServiceMock struct{}

func (this authServiceMock) CreateTokens(uuid string) (string, string, error) {
	if uuid == "valid" {
		return "access", "refresh", nil
	}

	return "", "", errors.New("invalid")
}

func (this authServiceMock) TryRefreshToken(token, uuid string) (string, string, error) {
	return "access", "refresh", nil
}

func Test_login(t *testing.T) {
	for _, test := range []struct {
		name         string
		uuid         string
		expectedCode int
	}{
		{
			name:         "empty_uuid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid_uuid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "valid_uuid",
			uuid:         "valid",
			expectedCode: http.StatusOK,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			wr := httptest.NewRecorder()

			httper.CreateHandler(login, authServiceMock{})(wr, httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?uuid=%v", test.uuid), nil))

			if got, want := wr.Result().StatusCode, test.expectedCode; got != want {
				t.Fatal(fmt.Sprintf("got: %v, want: %v", got, want))
			}
		})
	}

}
