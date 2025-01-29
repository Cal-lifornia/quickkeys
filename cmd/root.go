/*
Copyright © 2025 William Hobson willhobson@live.com.au

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	conf "github.com/Cal-lifornia/quickkeys/config"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfgFile string

var config conf.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quickkeys",
	Short: "CLI app to quickly search keybinds for apps",
	Long: `CLI app to search and list keybinds for configured apps. A particular key or set of keys
		can be searched for to see what apps they are used in and what for, or all keybinds
		for a particular app can be listed out and filtered.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initAll)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/quickkeys/quickkeys.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".quickkeys" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("toml")
// 		viper.SetConfigName("quickkeys")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }

func initAll() {
	initConfig()
	initLogger()
}

func initConfig() {
	if environment == "dev" {
		cfgFile = "./config.toml"
	}
	configFile, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("failed to open config file: %s", err.Error())
	}
	_, err = toml.Decode(string(configFile), &config)
	if err != nil {
		log.Fatalf("failed to decode config file: %s\n", err.Error())
	}
}

var environment string

func InitEnv(env string) {
	environment = env
}

func initLogger() {
	var chosenLogLevel zapcore.Level = zapcore.InfoLevel

	switch logLevel := config.LogLevel; logLevel {
	case "debug":
		chosenLogLevel = zapcore.DebugLevel
	case "info":
		chosenLogLevel = zapcore.InfoLevel
	case "warn":
		chosenLogLevel = zapcore.WarnLevel
	}

	atom := zap.NewAtomicLevelAt(chosenLogLevel)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return (lvl >= zapcore.ErrorLevel && lvl.Enabled(atom.Level()))
	})

	// TODO: Set log level option
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return (lvl < zapcore.ErrorLevel && lvl.Enabled(atom.Level()))
	})

	var consoleEncoderConfig zapcore.EncoderConfig
	var logFile string

	if environment == "prod" {
		consoleEncoderConfig = zap.NewProductionEncoderConfig()

		logFile = "/var/tmp/quickkeys.log"

	} else {
		consoleEncoderConfig = zap.NewDevelopmentEncoderConfig()
		logFile = "./debug.log"
	}

	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	consoleErrors := zapcore.Lock(os.Stderr)

	fileLog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     7,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, fileLog, lowPriority),
	)

	var logger *zap.Logger

	if environment == "prod" {
		logger = zap.New(core)
	} else {
		logger = zap.New(core, zap.AddCaller())
	}

	zap.ReplaceGlobals(logger)

	zap.L().Debug("logger initialised")

}
