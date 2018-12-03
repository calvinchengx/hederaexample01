package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	hedera "github.com/hashgraph/hedera-sdk-go"
)

type config struct {
	Networks map[string]network
}

type network struct {
	Address     string
	NodeAccount string `toml:"node_account"`
}

var conf config

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic(err)
	}
}

func main() {
	testnetAddress := conf.Networks["testnet"].Address
	client, err := hedera.Dial(testnetAddress)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	accountID := hedera.AccountID{Account: 1017}

	balance, err := client.GetAccountBalance(accountID).Answer()
	if err != nil {
		panic(err)
	}

	fmt.Printf("balance = %v tinybars\n", balance)
	fmt.Printf("balance = %.5f hbars\n", float64(balance)/100000000.0)
}
