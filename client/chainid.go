package client

import (
	"strconv"
)

func (client *web3Client) ChainId() (int, error) {

	params := []interface{}{}
	request := client.rpcClient.NewRequest(MethodGetChainId, params, 1)

	response, err := client.rpcClient.Execute(request)

	if err != nil {
		return 0, err
	}

	chainId, _ := strconv.ParseInt(response.Result, 0, 64)

	return int(chainId), nil
}
