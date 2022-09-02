/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/darmshot/vd/git"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "c",
	Short: "Make commit (in dev)",
	Long: `Command in dev

This help make speed commit with url of task.

need add to .env COMMIT_MESSAGE_PREFIX=

then you can make simple commits commands like vd c -t <number_task>
`,
	Run: func(cmd *cobra.Command, args []string) {

		task, _ := cmd.Flags().GetString("task")
		message, _ := cmd.Flags().GetString("message")
		err := git.Commit(task, message)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	commitCmd.PersistentFlags().StringP("task", "t", "", "number task or numbers like 100_101")
	commitCmd.PersistentFlags().StringP("message", "m", "", "commit message")
}
