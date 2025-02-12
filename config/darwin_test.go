//go:build darwin

package config

func (suite *ConfigTestSuite) TestDarwinInitKeys() {
	suite.Suite.Run("Test With No Symbols", func() {
		suite.Config.Symbols = false
		SetConfig(&suite.Config)
		initKeys()

		suite.Assert().Equal("Command", meta)
		suite.Assert().Equal("Ctrl", ctrl)
		suite.Assert().Equal("Shift", shift)
		suite.Assert().Equal("Option", altKey)
	})

	suite.Suite.Run("Test With Symbols", func() {
		suite.Config.Symbols = true
		SetConfig(&suite.Config)
		initKeys()

		suite.Assert().Equal("󰘳", meta)
		suite.Assert().Equal("󰘴", ctrl)
		suite.Assert().Equal("", shift)
		suite.Assert().Equal("󰘵", altKey)
	})
}
