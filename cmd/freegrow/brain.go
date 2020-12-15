// Copyright © 2020 Luan Guimarães Lacerda <luang@riseup.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "github.com/lib/pq"
	"github.com/luanguimaraesla/freegrow/internal/database"
	"github.com/luanguimaraesla/freegrow/pkg/brain"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// startBrainCmd represents the start command
var startBrainCmd = &cobra.Command{
	Use:   "brain",
	Short: "start freegrow API brain",
	Long:  `start freegrow API brain`,
	Run:   startBrain,
}

func init() {
	startCmd.AddCommand(startBrainCmd)

	startBrainCmd.Flags().String("bind", "", "server bind address")
}

func startBrain(cmd *cobra.Command, args []string) {
	logger.Info("initializing system")

	initDB(cmd, args)
	initServer(cmd, args)

	logger.Info("finished")
}

func initDB(cmd *cobra.Command, args []string) {
	var err error

	o := &database.ConnectionOptions{
		Host:     getEnvOrDefault("POSTGRES_HOST"),
		Port:     getEnvOrDefault("POSTGRES_PORT"),
		Database: getEnvOrDefault("POSTGRES_DATABASE"),
		Username: getEnvOrDefault("POSTGRES_USERNAME"),
		Password: getEnvOrDefault("POSTGRES_PASSWORD"),
	}

	err = database.Connect(o)
	if err != nil {
		logger.Fatal("failed connecting to postgres", zap.Error(err))
	}

	logger.Info("successfully connected!")
}

func initServer(cmd *cobra.Command, args []string) {
	bind, err := cmd.Flags().GetString("bind")
	if err != nil {
		logger.Fatal("option --bind is missing", zap.Error(err))
	}

	if bind == "" {
		bind = getEnvOrDefault("BIND_ADDRESS")
	}

	b := brain.New()

	if err := b.Init(); err != nil {
		logger.Fatal("unable to initialize system", zap.Error(err))
	}

	logger.Info("starting freegrow brain server")

	if err := b.Listen(bind); err != nil {
		logger.Fatal("unable to start freegrow brain server", zap.Error(err))
	}
}
