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

package cmd

import (
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
}

func startBrain(cmd *cobra.Command, args []string) {
	logger.Info("initializing system")

	b := brain.New()

	if err := b.Init(); err != nil {
		logger.Fatal("unable to initialize system", zap.Error(err))
	}

	logger.Info("starting freegrow brain server")

	if err := b.Listen("0.0.0.0:8000"); err != nil {
		logger.Fatal("unable to start freegrow brain server", zap.Error(err))
	}

	logger.Info("finished")
}
