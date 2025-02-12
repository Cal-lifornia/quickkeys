//go:build linux
// +build linux

package config

func (suite *ConfigTestSuite) TestLinuxInitKeys() {
	suite.Suite.Run("Test With No Symbols", func() {
		suite.Config.Symbols = false
		SetConfig(&suite.Config)
		initKeys()

		suite.Assert().Equal("Super", meta)
		suite.Assert().Equal("Ctrl", ctrl)
		suite.Assert().Equal("Shift", shift)
		suite.Assert().Equal("Alt", altKey)
	})

	suite.Suite.Run("Test With Symbols", func() {
		suite.Config.Symbols = true
		SetConfig(&suite.Config)
		initKeys()

		suite.Assert().Equal("", meta)
		suite.Assert().Equal("Ctrl", ctrl)
		suite.Assert().Equal("", shift)
		suite.Assert().Equal("Alt", altKey)
	})
}
