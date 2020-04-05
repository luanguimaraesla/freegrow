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
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/luanguimaraesla/freegrow/internal/controller"
	"github.com/luanguimaraesla/freegrow/internal/device"
	"github.com/luanguimaraesla/freegrow/internal/system"
	"github.com/luanguimaraesla/freegrow/pkg/gadgets"
	"github.com/luanguimaraesla/freegrow/pkg/gadgets/irrigator"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start freegrow server",
	Long:  `start freegrow server`,
	Run:   start,
}

var log *logrus.Entry

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("board", "b", "raspberry", "controller board to use")

}

func start(cmd *cobra.Command, args []string) {
	board, err := cmd.Flags().GetString("board")
	if err != nil {
		log.WithError(err).Fatal("error getting board name")
	}

	log.Info("starting system")
	controller.StartController(board)

	i, err := irrigator.New("main_irrigator", 14, time.Second*10) // Test
	if err != nil {
		log.WithError(err).Error("failed to create irrigator gadget")
	}
	i.Start()

	log.Info("Finished")
}

func getLogger() *logrus.Entry {
	logger := logrus.New()
	return logger.WithFields(logrus.Fields{
		"command": "start",
	})
}

func init() {
	log = getLogger()
	device.SetLogger(log)
	controller.SetLogger(log)
	gadgets.SetLogger(log)
	system.SetLogger(log)
}
