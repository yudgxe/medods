package services

import (
	"fmt"
	"medods/database/model"
	"testing"
)

type authDaoMock struct{}

func (this authDaoMock) CreateOrUpdate(auth *model.Auth) error {
	return nil
}

func (this authDaoMock) IsExistsUUID(uuid string) (bool, error) {
	return true, nil

}

func (this authDaoMock) IsExistsToken(token string) (bool, error) {
	if token == "exist" {
		return true, nil
	}

	return false, nil
}

func Test_CreateTokens(t *testing.T) {
	if _, _, err := Auth().CreateTokens("UUID"); err != nil {
		t.Fatal(err)
	}
}

func Test_TryRefreshToken(t *testing.T) {
	const uuid string = "UUID"

	for _, test := range []struct {
		name        string
		token       string
		expectedErr error
	}{
		{
			name:        "not_exist",
			token:       "not_exist",
			expectedErr: ErrTokenNotExist,
		},
		{
			name:        "exist",
			token:       "exist",
			expectedErr: nil,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			_, _, err := Auth().TryRefreshToken(test.name, test.token)

			if got, want := err, test.expectedErr; got != want {
				t.Fatal(fmt.Sprintf("got: %v, want: %v", got, want))
			}
		})
	}
}
