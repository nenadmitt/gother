package client

import (
	"github.com/nenadmitt/go3/rpc"
	"github.com/nenadmitt/go3/wallet"
	"math/big"
)

type Web3Client interface {
	Nonce(address string) (int, error)
	GetBalance(address string) (*big.Int, error)
	GasPrice() (*big.Int, error)
	ChainId() (int, error)
}

type web3Client struct {
	rpcClient rpc.Client
	signer    wallet.Wallet
}

type Web3ClientOptions struct {
	Url    string
	Signer wallet.Wallet
}

func NewWeb3Client(opts Web3ClientOptions) Web3Client {
	rpcClient := rpc.NewHttpRpcClient(opts.Url)
	return &web3Client{
		rpcClient: rpcClient,
		signer:    opts.Signer,
	}
}
