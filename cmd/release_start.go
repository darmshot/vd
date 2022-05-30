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

// releaseStartCmd represents the releaseStart command
var releaseStartCmd = &cobra.Command{
	Use:   "rs",
	Short: "Release start",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var releaseType = "minor"
		isMajor, _ := cmd.Flags().GetBool("major")
		isFirst, _ := cmd.Flags().GetBool("first")

		if isMajor {
			releaseType = "major"
		} else if isFirst {
			releaseType = "first"
		}

		err := git.ReleaseStart(releaseType)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(releaseStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// releaseStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// releaseStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	releaseStartCmd.PersistentFlags().BoolP("major", "m", false, "Release Start and upgrade major version")
	releaseStartCmd.PersistentFlags().BoolP("first", "f", false, "Release Start first")

}
