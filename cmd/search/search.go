/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"runtime"

	"github.com/imelon2/orbit-cli/prompt"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("bad path")
		}

		cmdName, err := prompt.SelectNextCmd(filename)
		if err != nil {
			log.Fatal(err)
		}

		nextCmd, _, err := cmd.Find([]string{cmdName})
		if err != nil {
			log.Fatal(err)
		}
		nextCmd.Run(nextCmd, args)
	},
}

func init() {
	SearchCmd.AddCommand(NodeCreatedCmd)
	SearchCmd.AddCommand(NodeConfirmedCmd)
	SearchCmd.AddCommand(BatchDeliveredCmd)
}
