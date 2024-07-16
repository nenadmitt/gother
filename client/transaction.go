package client

import (
	"errors"
	"github.com/nenadmitt/go3/transaction"
)

func (client *web3Client) sendRawTransaction(tx transaction.Tx) error {

	if client.signer == nil {
		return errors.New("signer not set")
	}

	signature, err := client.signer.SignTransaction(tx)

	if err != nil {
		return err
	}

	return nil
}
