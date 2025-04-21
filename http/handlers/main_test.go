package handlers

import (
	"medods/logger"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logger.InitLogger(logrus.DebugLevel)

	os.Exit(m.Run())
}
