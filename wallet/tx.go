package wallet

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nenadmitt/go3/transaction"
)

func (wallet *wallet) SignTransaction(tx transaction.Tx) (string, error) {

	txHash, _ := tx.Hash()
	signature, _ := crypto.Sign(txHash, &wallet.pk)

	return hex.EncodeToString(signature), nil
}
