/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// V100Cmd represents the V100 command
var V100Cmd = &cobra.Command{
	Use:   "V100",
	Short: "Version 100 of the challenge",
	Long: `Based on the vorto challenge, given a testfile 
of a list of Loads, with Pickup and dropOff lat longs,
this app will output a list of drivers and their assigned load.
The result will include the minimum drivers count so that they do not
Violate 12 hour rule.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("V100 called")
		slog := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

		ProblemPath, err := cmd.Flags().GetString("ProblemPath")
		if err != nil {
			slog.Println("Error getting ProblemPath flag", err)
			return
		}

		SaveOutput, err := cmd.Flags().GetString("SaveOutput")
		if err != nil {
			slog.Println("Error getting SaveOutput flag", err)
			return
		}

		if SaveOutput {
			util.MakeOutputDirs()
		}

	},
}

func init() {
	rootCmd.AddCommand(V100Cmd)
	V100Cmd.Flags().StringP("ProblemPath", "T", "../training/Problems/problem20.txt", "select a Single Problem to work on")
	V100Cmd.Flags().BoolP("SaveOutput", "s", true, "Save the separate stages to a JSON file for inspection")

}
