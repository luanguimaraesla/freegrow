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
	"fmt"

	"github.com/spf13/cobra"
        "github.com/sirupsen/logrus"

        "github.com/luanguimaraesla/freegrow/controller"
        "github.com/luanguimaraesla/freegrow/system/relay"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start freegrow server",
	Long: `start freegrow server`,
	Run: start,
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
        startCmd.Flags().StringP("board", "b", "raspberry", "controller board to use")

}

func start(cmd *cobra.Command, args []string) {
        board, err := cmd.Flags().GetString("board")
        if err != nil {
                fmt.Printf("error getting board name")
        }

        logger := logrus.New()
        log := logger.WithFields(logrus.Fields{
                "command": "start",
        })

        log.Info("starting modules")

        log.Info("configuring controller")
        controller.SetLogger(log)
        controller.StartController(board)

        r, err := relay.NewRelay("hello world", 14) // Test
        if err != nil {
                fmt.Printf("error creating relay device: %v", err)
        }
        r.Activate()
        r.Deactivate()
        r.Activate()
        r.Deactivate()

        log.Info("Finished")
}
