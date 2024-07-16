package client

import (
	"math/big"
)

func (client *web3Client) GasPrice() (*big.Int, error) {
	params := []interface{}{}
	request := client.rpcClient.NewRequest(MethodGetGasPrice, params, 1)

	response, err := client.rpcClient.Execute(request)

	gas := new(big.Int)

	if err != nil {
		return gas, err
	}
	gas.SetString(response.Result[2:], 16)
	return gas, nil
}
