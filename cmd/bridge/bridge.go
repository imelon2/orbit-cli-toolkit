/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"runtime"

	"github.com/imelon2/orbit-cli/prompt"
	"github.com/imelon2/orbit-cli/utils"
	"github.com/spf13/cobra"
)

const ()

// bridgeCmd represents the bridge command
var BridgeCmd = &cobra.Command{
	Use:   "bridge",
	Short: "Doposit and withdraw native, erc20 asset",
	Run: func(cmd *cobra.Command, args []string) {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("bad path")
		}

		root := utils.GetRootDir(filename)
		selected, err := prompt.SelectCommand(root)

		if err != nil {
			log.Fatal(err)
		}

		nextCmd, _, err := cmd.Find([]string{selected})
		if err != nil {
			log.Fatal(err)
		}

		nextCmd.Run(nextCmd, args)
	},
}

func init() {
	BridgeCmd.AddCommand(Erc20Cmd)
}
