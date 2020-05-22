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
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/luanguimaraesla/freegrow/internal/controller"
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

	startCmd.Flags().StringP("board", "b", "raspberry", "controller board to use")

}

func preStart(cmd *cobra.Command, args []string) {
	logger = logger.With(
		zap.String("command", "start"),
	)
}

func start(cmd *cobra.Command, args []string) {
	board, err := cmd.Flags().GetString("board")
	if err != nil {
		logger.Fatal("failed to start", zap.Error(err))
	}

	logger.Info("starting system")
	controller.DefineController(board)

	logger.Info("finished")
}
