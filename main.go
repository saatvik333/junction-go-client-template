package main

import (
	"fmt"
	"log"
	"time"

	"github.com/saatvik333/junction-go-client-template/src"
	"github.com/saatvik333/junction-go-client-template/utils"
)

func main() {
	// Define configuration
	config := utils.Config{
		AccountPath: "./accounts",
		JsonRPC:     "http://localhost:26657",
		Token:       "stake",
		ChainPrefix: "air",
		AccountNames: []string{
			"account1",
			"account2",
		},
		AccountAddresses: []string{
			"air1vk4490xlxtz20lcrhzvwvku3hxun6gwmz4pr0y",
			// "air1g0xze04mljvf64xwrkxmvmcqs9pu245y5fh7rv",

		},
	}

	//setupAccounts(config)
	//checkBalances(config)
	//accountsData := connectAllClients(config)
	if len(config.AccountAddresses) == 0 || len(config.AccountAddresses) == 1 {
		setupAccounts(config)
	} else {
		//
		return

	}
	checkBalances(config)
	_, err := src.InitRollup()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(20 * time.Second)

	_, err = src.InitProver()
	if err != nil {
		fmt.Println(err)
	}
}

// setupAccounts creates accounts using the provided configuration.
func setupAccounts(cfg utils.Config) {
	for _, name := range cfg.AccountNames {
		exists, _ := utils.CheckIfAccountExists(name, cfg.AccountPath, cfg.ChainPrefix)
		if exists {
			log.Printf("Account %s already exists. Skipping creation.", name)
			continue
		}

		utils.CreateAccount(name, cfg.AccountPath) // Simply call it
	}
}

// checkBalances retrieves and logs the balance of each account.
func checkBalances(cfg utils.Config) {
	for _, addr := range cfg.AccountAddresses {
		success, balance, err := utils.CheckBalance(cfg.JsonRPC, cfg.Token, addr)
		if err != nil {
			log.Printf("Error checking balance for %s: %v", addr, err)
			continue
		}
		if !success {
			log.Printf("Failed to retrieve balance for %s", addr)
			continue
		}
		log.Printf("%s balance: %d", addr, balance)
	}
}

// // connectAllClients connects to all accounts and returns their connection data.
// func connectAllClients(cfg utils.Config) []utils.AccountData {
// 	var accountsData []utils.AccountData
// 	for _, name := range cfg.AccountNames {
// 		account, addr, client := junction.ClientConnect(cfg.AccountPath, name, cfg.ChainPrefix, cfg.JsonRPC)
// 		accountsData = append(accountsData, utils.AccountData{
// 			Account: account,
// 			Addr:    addr,
// 			Client:  client,
// 		})
// 	}
// 	return accountsData
// }
