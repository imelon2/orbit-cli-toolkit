/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/imelon2/orbit-cli/prompt"
	"github.com/imelon2/orbit-cli/utils"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send ETH from select wallet",
	Run: func(cmd *cobra.Command, args []string) {

		value, err := prompt.EnterValue("")
		if err != nil {
			log.Fatal(err)
		}

		to, err := prompt.EnterRecipient()
		if err != nil {
			log.Fatal(err)
		}

		toAddress := common.HexToAddress(to)

		wallet, _, account, err := prompt.SelectWalletForSign()
		if err != nil {
			log.Fatal(err)
		}

		provider, err := prompt.SelectProvider()
		if err != nil {
			log.Fatal(err)
		}

		client, err := ethclient.Dial(provider)
		if err != nil {
			log.Fatal(err)
		}

		nonce, err := client.PendingNonceAt(context.Background(), account.Address)
		if err != nil {
			log.Fatal(err)
		}

		gasLimit := uint64(23000 * 5)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		gasPrice.Mul(gasPrice, big.NewInt(2))
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil /* calldata */)
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := wallet.SignTx(account, tx, chainID)
		if err != nil {
			log.Fatal(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
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

		fmt.Println("Transaction receipt: ")
		utils.PrintPrettyJson(receipt)
	},
}

func init() {
}
