/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/justmeandopensource/numboz/internal/challenges"
	"github.com/justmeandopensource/numboz/internal/common"
	"github.com/spf13/cobra"
)

var challengeCmd = &cobra.Command{
	Use:   "challenge",
	Short: "a new numboz challenge",
	Run: func(cmd *cobra.Command, _ []string) {

		challengeType, _ := cmd.Flags().GetString("type")
		questions, _ := cmd.Flags().GetInt("questions")
		digits, _ := cmd.Flags().GetInt("digits")

		common.ValidateParams(challengeType, digits)

		common.ClearTerminal()

		switch strings.ToLower(challengeType) {
		case "mixed":
			for i := 1; i < questions+1; i++ {
				common.ClearTerminal()
				fmt.Printf("Question [%v/%v]\n\n", i, questions)
				challenges.Start(*common.PickRandomChallenge(), digits)
			}
		default:
			for i := 1; i < questions+1; i++ {
				common.ClearTerminal()
				fmt.Printf("Question [%v/%v]\n\n", i, questions)
				challenges.Start(challengeType, digits)
			}
		}
		challenges.PrintChallengeReport()
	},
}

func init() {

	challengeCmd.Flags().StringP("type", "t", "mixed", "type of challenge")
	challengeCmd.Flags().IntP("questions", "n", 10, "number of questions for the challenge")
	challengeCmd.Flags().IntP("digits", "d", 2, "number of digits to be used in the challenge [between 1 and 5]")

	rootCmd.AddCommand(challengeCmd)
}
