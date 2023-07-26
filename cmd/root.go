/*
Copyright © 2023 Venkat Nagappan
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "numboz",
	Short: "maths challange",
	Long:  `command line program for practicing basic maths arithmetic skills`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
