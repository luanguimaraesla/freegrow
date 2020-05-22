// Copyright © 2019 Luan Guimarães Lacerda <luang@riseup.net>
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
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start freegrow server",
	Long:  `start freegrow server`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Panic(err)
		}
		os.Exit(0)
	},
	PersistentPreRun: preStart,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringP("file", "f", "", "resource manifest")

	if err := cobra.MarkFlagRequired(startCmd.PersistentFlags(), "file"); err != nil {
		logger.Fatal("please set --file flag", zap.Error(err))
	}
}

func preStart(cmd *cobra.Command, args []string) {
	logger = logger.With(
		zap.String("command", "start"),
	)
}
