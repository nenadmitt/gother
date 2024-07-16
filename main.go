package main

import (
	"fmt"
	"github.com/nenadmitt/go3/client"
	"github.com/nenadmitt/go3/wallet"
)

func main() {
	fmt.Println("Hello There")

	currnetWallet, err := wallet.FromPrivateKey()

	if err != nil {
		panic(err)
	}

	opts := client.Web3ClientOptions{
		Url: "",
	}

	web3Client := client.NewWeb3Client(opts)

	nonce, _ := web3Client.Nonce(currnetWallet.Address())
	balance, _ := web3Client.GetBalance(currnetWallet.Address())
	chainId, _ := web3Client.ChainId()

	fmt.Println(nonce)
	fmt.Println(balance.String())
	fmt.Println(chainId)
}
