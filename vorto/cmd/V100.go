/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/util"
	v100 "github.com/vinceyoumans/chal-vorto04/vorto/pkg/v100"
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

		ProblemFile, err := cmd.Flags().GetString("ProblemFile")
		if err != nil {
			slog.Println("Error getting ProblemFile flag", err)
			return
		}

		SaveOutput, err := cmd.Flags().GetBool("SaveOutput")
		if err != nil {
			slog.Println("Error getting SaveOutput flag", err)
			return
		}

		if SaveOutput == true {
			util.MakeOutputDirs()
		}

		ProblemPathFile := filepath.Join(ProblemPath, ProblemFile)
		slog.Println("--- ProblemPathFile : ", ProblemPathFile)

		v100.V100Start(ProblemPathFile, SaveOutput)

		// result := v100.V100Start(ProblemPathFile, SaveOutput)
		// for _, valResult := range result {
		// 	strNumbers := make([]string, len(valResult))
		// 	for i, num := range valResult {
		// 		strNumbers[i] = strconv.Itoa(num)
		// 	}
		// 	fmt.Printf("[%s]\n", strings.Join(strNumbers, ", "))
		// }

	},
}

func init() {
	rootCmd.AddCommand(V100Cmd)
	V100Cmd.Flags().StringP("ProblemPath", "P", "../training/Problems/", "select Problem Directory")
	V100Cmd.Flags().StringP("ProblemFile", "F", "problem20.txt", "select a Single Problem File to Calculate")
	V100Cmd.Flags().BoolP("SaveOutput", "S", true, "Save the separate stages to a JSON file for inspection")
}
