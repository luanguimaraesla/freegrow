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
	"github.com/luanguimaraesla/freegrow/internal/cache"
	"github.com/luanguimaraesla/freegrow/internal/database"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"github.com/luanguimaraesla/freegrow/pkg/brain"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// brainCmd represents the start command
var brainCmd = &cobra.Command{
	Use:   "brain",
	Short: "freegrow API brain",
	Long:  `freegrow API brain`,
	Run:   startBrain,
}

func init() {
	rootCmd.AddCommand(brainCmd)

	brainCmd.Flags().String("bind", "", "server bind address")
}

func startBrain(cmd *cobra.Command, args []string) {
	log.L.Info("initializing system")

	initDB(cmd, args)
	initCache(cmd, args)
	initServer(cmd, args)

	log.L.Info("finished")
}

func initDB(cmd *cobra.Command, args []string) {
	database.Init(
		getEnvOrDefault("POSTGRES_HOST"),
		getEnvOrDefault("POSTGRES_PORT"),
		getEnvOrDefault("POSTGRES_DATABASE"),
		getEnvOrDefault("POSTGRES_USERNAME"),
		getEnvOrDefault("POSTGRES_PASSWORD"),
	)

	if err := database.Ping(); err != nil {
		log.L.Fatal("failed connecting to postgres", zap.Error(err))
	}

	log.L.Info("successfully connected to postgres!")
}

func initCache(cmd *cobra.Command, args []string) {
	cache.Init(
		getEnvOrDefault("REDIS_URL"),
	)

	if err := cache.Ping(); err != nil {
		log.L.Fatal("failed connecting to cache", zap.Error(err))
	}

	log.L.Info("successfully connected to cache!")
}

func initServer(cmd *cobra.Command, args []string) {
	bind, err := cmd.Flags().GetString("bind")
	if err != nil {
		log.L.Fatal("option --bind is missing", zap.Error(err))
	}

	if bind == "" {
		bind = getEnvOrDefault("BIND_ADDRESS")
	}

	b := brain.New()

	if err := b.Listen(bind); err != nil {
		log.L.Fatal("unable to start freegrow brain server", zap.Error(err))
	}
}
