package client

const (
	BlockLatest  = "latest"
	BlockPending = "pending"
)

const (
	MethodGetTransactionCount = "eth_getTransactionCount"
	MethodGetBalance          = "eth_getBalance"
	MethodGetChainId          = "eth_chainId"
	MethodGetGasPrice         = "eth_gasPrice"
)

const (
	RegularTransactionGas = 21_000
)
