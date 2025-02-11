package appkeys

import (
	"os"
	"testing"

	"github.com/Cal-lifornia/quickkeys/config"
	"github.com/Cal-lifornia/quickkeys/types"
	"github.com/alecthomas/repr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HelixTestSuite struct {
	suite.Suite
	testAppConf types.AppConfig
	config.Config
}

func (suite *HelixTestSuite) SetupTest() {
	os.Setenv("ENVIRONMENT", "test")

	suite.Config = config.Config{
		LogLevel: "debug",
	}

	config.InitLogger(&suite.Config, "test")

	suite.testAppConf = types.AppConfig{
		Name:       "Helix",
		Alias:      []string{"hx"},
		ConfigPath: "../test_data/helix/config.toml",
	}
}

func (suite *HelixTestSuite) TestGetHelixKeysEntries() {

	results, err := getHelixKeysEntries(&suite.testAppConf)
	if assert.NoError(suite.T(), err, "should be no error: ") {
		repr.Println(results)
	}
}

func TestHelixTestSuite(t *testing.T) {
	suite.Run(t, new(HelixTestSuite))
}
