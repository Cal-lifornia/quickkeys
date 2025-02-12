package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type ConfigTestSuite struct {
	suite.Suite
	Config
	logger *zap.Logger
}

func TestDarwinConfTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (suite *ConfigTestSuite) SetupTest() {
	os.Setenv("ENVIRONMENT", "testing")

	suite.Config = Config{
		LogLevel: "debug",
	}

	suite.logger = zaptest.NewLogger(suite.T(), zaptest.WrapOptions(zap.AddCaller()))

	SetLogger(suite.logger)
}
