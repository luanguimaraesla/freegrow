// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"github.com/luanguimaraesla/freegrow/pkg/machine"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:              "start",
	Short:            "start freegrow server",
	Long:             `start freegrow server`,
	Run:              start,
	PersistentPreRun: preStart,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("file", "f", "", "machine manifest")

	if err := cobra.MarkFlagRequired(startCmd.Flags(), "file"); err != nil {
		logger.Fatal("please set --file flag", zap.Error(err))
	}
}

func preStart(cmd *cobra.Command, args []string) {
	logger = logger.With(
		zap.String("command", "start"),
	)
}

func start(cmd *cobra.Command, args []string) {
	logger.Info("starting system")

	filename, err := cmd.Flags().GetString("file")
	if err != nil {
		logger.Fatal("unable to get file flag", zap.Error(err))
	}

	m := machine.New()

	logger.With(
		zap.String("file", filename),
		zap.String("stage", "loading"),
	).Info("loading file")

	if err := m.Load(filename); err != nil {
		logger.With(
			zap.String("file", filename),
			zap.String("stage", "loading"),
		).Fatal("unable to load file", zap.Error(err))
	}

	logger.With(
		zap.String("file", filename),
		zap.String("stage", "initializing"),
	).Info("initializing machine")

	if err := m.Init(); err != nil {
		logger.With(
			zap.String("file", filename),
			zap.String("stage", "initializing"),
		).Fatal("unable to initialize machine", zap.Error(err))
	}

	logger.With(
		zap.String("file", filename),
		zap.String("stage", "running"),
	).Info("running machine")

	if err := m.Run(); err != nil {
		logger.With(
			zap.String("file", filename),
			zap.String("stage", "running"),
		).Fatal("unable to run machine", zap.Error(err))
	}

	logger.Info("finished")
}
