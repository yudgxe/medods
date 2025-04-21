package services

import (
	"medods/database/dao"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dao.SetupAuthDao(authDaoMock{})

	os.Exit(m.Run())
}
