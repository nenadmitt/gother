package client

import (
	"math/big"
)

func (client *web3Client) GetBalance(address string) (*big.Int, error) {
	params := []interface{}{
		address,
		BlockLatest,
	}
	request := client.rpcClient.NewRequest(MethodGetBalance, params, 1)

	response, err := client.rpcClient.Execute(request)
	balance := new(big.Int)

	balance.SetString(response.Result[2:], 16)

	if err != nil {
		return balance, err
	}

	return balance, nil
}
