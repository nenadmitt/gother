package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type httpRpcClient struct {
	rpcURL string
}

func (rpc *httpRpcClient) NewRequest(method string, parameters []interface{}, id int) EthereumRPCRequest {
	return EthereumRPCRequest{
		Method:  method,
		JsonRpc: RpcVersion,
		Params:  parameters,
		ID:      id,
	}
}

func (rpc *httpRpcClient) Execute(request EthereumRPCRequest) (*EthereumRPCResponse, error) {
	payload, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	response, err := http.Post(rpc.rpcURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	var ethResponse EthereumRPCResponse
	err = json.Unmarshal(bodyBytes, &ethResponse)

	if err != nil {
		return nil, err
	}

	return &ethResponse, nil
}

func NewHttpRpcClient(rpcURL string) Client {
	return &httpRpcClient{rpcURL: rpcURL}
}
