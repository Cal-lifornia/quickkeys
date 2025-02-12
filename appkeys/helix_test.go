package appkeys

import (
	"os"
	"testing"

	"github.com/Cal-lifornia/quickkeys/config"
	"github.com/Cal-lifornia/quickkeys/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type HelixTestSuite struct {
	suite.Suite
	testAppConf types.AppConfig
	config.Config
	logger *zap.Logger
}

func (suite *HelixTestSuite) SetupTest() {
	os.Setenv("ENVIRONMENT", "testing")

	suite.Config = config.Config{
		LogLevel: "debug",
	}

	suite.logger = zaptest.NewLogger(suite.T(), zaptest.WrapOptions(zap.AddCaller()))

	config.SetLogger(suite.logger)

	suite.testAppConf = types.AppConfig{
		Name:       "Helix",
		Alias:      []string{"hx"},
		ConfigPath: "../test_data/helix/config.toml",
	}
}

func (suite *HelixTestSuite) TestGetHelixKeysEntries() {
	_, err := getHelixKeysEntries(suite.testAppConf.ConfigPath)
	if assert.NoError(suite.T(), err, "should be no error: ") {
		suite.logger.Error("successful test")
	}
}

func TestHelixTestSuite(t *testing.T) {
	suite.Run(t, new(HelixTestSuite))
}
