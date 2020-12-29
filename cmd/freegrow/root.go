// Copyright © 2020 Luan Guimarães Lacerda
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/luanguimaraesla/freegrow/internal/log"
)

var (
	logger                       *zap.Logger
	cfgFile, logLevel, logFormat string
)

var defaults = map[string]string{
	"POSTGRES_USERNAME":        "freegrow",
	"POSTGRES_PASSWORD":        "freegrow",
	"POSTGRES_HOST":            "localhost",
	"POSTGRES_PORT":            "5432",
	"POSTGRES_DATABASE":        "freegrow",
	"BIND_ADDRESS":             "127.0.0.1:8000",
	"REDIS_URL":                "redis://redis:6379",
	"SESSION_ACCESS_SECRET":    "freegrow",
	"SESSION_ACCESS_DURATION":  "15m",
	"SESSION_REFERSH_SECRET":   "freegrow",
	"SESSION_REFERSH_DURATION": "24h",
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "freegrow",
	Short: "freegrow is a opensource API server used to control greenhouses remote devices",
	Long:  "freegrow is a opensource API server used to control greenhouses remote devices",
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogger)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.freegrow.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log", "info", "log level [debug, info, warn, error, panic, fatal]")
	rootCmd.PersistentFlags().StringVar(&logFormat, "format", "console", "log output format [json, console]")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".freegrow" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".freegrow")
	}

	viper.SetEnvPrefix("FREEGROW")
	viper.AutomaticEnv() // read in environment variables that match

	for env := range defaults {
		viper.BindEnv(env)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger() {
	rawJSON := []byte(fmt.Sprintf(`{
          "level": "%s",
          "encoding": "%s",
          "outputPaths": ["stdout"],
          "errorOutputPaths": ["stderr"],
          "encoderConfig": {
            "messageKey": "message",
            "levelKey": "level",
            "levelEncoder": "lowercase"
          }
        }`,
		logLevel,
		logFormat,
	))

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	defaultLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer defaultLogger.Sync()

	log.L = defaultLogger
}

// getEnvOrDefault returns the environment variable
// value returned by viper or the hardcoded default
func getEnvOrDefault(key string) string {
	defaultValue, ok := defaults[key]
	if !ok {
		logger.Fatal(fmt.Sprintf("default environment variable %s doesn't exist", key))
	}

	v := viper.GetString(strings.ToLower(key))
	if v == "" {
		v = defaultValue
	}

	return v
}
