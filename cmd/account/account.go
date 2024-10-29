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

// accountCmd represents the account command
var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "A brief description of your command",
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
	AccountCmd.AddCommand(NewCmd)
	AccountCmd.AddCommand(ListCmd)
	AccountCmd.AddCommand(ImportCmd)
}
