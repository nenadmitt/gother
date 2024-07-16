package transaction

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
)

type Tx struct {
	From     string
	To       string
	Gas      big.Int
	GasPrice big.Int
	Value    big.Int
	Nonce    int
	Data     string
}

type RawTx struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Nonce    string `json:"nonce"`
	Input    string `json:"input"`
}

func (tx *Tx) ToRaw() *RawTx {
	return &RawTx{
		From:     tx.From,
		To:       tx.To,
		Gas:      tx.Gas.Text(16),
		GasPrice: tx.GasPrice.Text(16),
		Value:    tx.Value.Text(16),
		Nonce:    fmt.Sprintf("%x", tx.Nonce),
		Input:    tx.Data,
	}
}

func (tx *Tx) Hash() ([]byte, error) {

	txBytes, err := json.Marshal(tx.ToRaw())
	if err != nil {
		return []byte{}, err
	}
	hash := sha256.Sum256(txBytes)
	return hash[:], nil
}
