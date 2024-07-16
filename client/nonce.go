package client

import (
	"strconv"
)

func (client *web3Client) Nonce(address string) (int, error) {

	params := []interface{}{
		address,
		BlockLatest,
	}
	request := client.rpcClient.NewRequest(MethodGetTransactionCount, params, 1)
	response, err := client.rpcClient.Execute(request)
	if err != nil {
		panic(err)
	}
	nonceHex := response.Result
	nonceInt, _ := strconv.ParseInt(nonceHex, 0, 64)

	return int(nonceInt), nil
}
