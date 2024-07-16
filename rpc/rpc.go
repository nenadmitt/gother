package rpc

const (
	RpcVersion = "2.0"
)

type Client interface {
	NewRequest(method string, params []interface{}, id int) EthereumRPCRequest
	Execute(req EthereumRPCRequest) (*EthereumRPCResponse, error)
}

type EthereumRPCRequest struct {
	Method  string        `json:"method"`
	JsonRpc string        `json:"jsonrpc"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type EthereumRPCResponse struct {
	ID      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
