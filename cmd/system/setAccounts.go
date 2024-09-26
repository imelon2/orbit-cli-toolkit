/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethLib "github.com/imelon2/orbit-cli/ethLib"
	"github.com/imelon2/orbit-cli/prompt"
	"github.com/imelon2/orbit-cli/solgen/go/precompilesgen"
	"github.com/imelon2/orbit-cli/utils"
	"github.com/spf13/cobra"
)

const (
	setL1PricingRewardRecipient = iota
	setInfraFeeAccount
	setNetworkFeeAccount
)

var setAccountsCommand = []string{"SetL1PricingRewardRecipient", "SetInfraFeeAccount", "SetNetworkFeeAccount"}

// setAccountsCmd represents the setAccounts command
var SetAccountsCmd = &cobra.Command{
	Use:   "setAccounts",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		var qs = &survey.Select{
			Message: "Select Command: ",
			Options: setAccountsCommand,
		}

		answerIndex := 0
		err := survey.AskOne(qs, &answerIndex)
		if err != nil {
			log.Fatal(err)
		}

		var client *ethclient.Client
		var signedTx *types.Transaction

		switch answerIndex {
		case setL1PricingRewardRecipient:
			client, signedTx = SetL1PricingRewardRecipient()
		case setInfraFeeAccount:
			client, signedTx = SetInfraFeeAccount()
		case setNetworkFeeAccount:
			client, signedTx = SetNetworkFeeAccount()
		}

		txResponse, _, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("\n\nTransaction Response: \n")
		utils.PrintPrettyJson(txResponse)
		fmt.Print("\n\nWait Mined Transaction ... \n\n")

		receipt, err := bind.WaitMined(context.Background(), client, signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Transaction receipt: \n")
		utils.PrintPrettyJson(receipt)
	},
}

func init() {

}

func SetL1PricingRewardRecipient() (*ethclient.Client, *types.Transaction) {
	newRewarderAccount, err := prompt.EnterAddress("new L1 Rewarder account")

	if err != nil {
		log.Fatal(err)
	}

	client, auth, err := ethLib.GenerateAuth()
	if err != nil {
		log.Fatal(err)
	}

	ArbOwner, err := precompilesgen.NewArbOwner(types.ArbOwnerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := ArbOwner.SetL1PricingRewardRecipient(auth, common.HexToAddress(newRewarderAccount))

	if err != nil {
		log.Fatal(err)
	}

	return client, signedTx
}

func SetInfraFeeAccount() (*ethclient.Client, *types.Transaction) {
	newInfraAccount, err := prompt.EnterAddress("new infra account")

	if err != nil {
		log.Fatal(err)
	}

	client, auth, err := ethLib.GenerateAuth()
	if err != nil {
		log.Fatal(err)
	}

	ArbOwner, err := precompilesgen.NewArbOwner(types.ArbOwnerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := ArbOwner.SetInfraFeeAccount(auth, common.HexToAddress(newInfraAccount))

	if err != nil {
		log.Fatal(err)
	}

	return client, signedTx
}

func SetNetworkFeeAccount() (*ethclient.Client, *types.Transaction) {
	newInfraAccount, err := prompt.EnterAddress("new Network Fee account")

	if err != nil {
		log.Fatal(err)
	}

	client, auth, err := ethLib.GenerateAuth()
	if err != nil {
		log.Fatal(err)
	}

	ArbOwner, err := precompilesgen.NewArbOwner(types.ArbOwnerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := ArbOwner.SetNetworkFeeAccount(auth, common.HexToAddress(newInfraAccount))

	if err != nil {
		log.Fatal(err)
	}

	return client, signedTx
}
