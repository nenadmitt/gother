package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nenadmitt/go3/transaction"
	"golang.org/x/crypto/sha3"
	"strings"
)

type Wallet interface {
	Address() string
	SignTransaction(tx transaction.Tx) (string, error)
}

type wallet struct {
	pk      ecdsa.PrivateKey
	address string
}

func (w *wallet) Address() string {
	return w.address
}

func FromPrivateKey(privateKeyHex string) (Wallet, error) {

	/// Strip 0x prefix if present
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// Decode the private key from hex to bytes
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %v", err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create ECDSA private key: %v", err)
	}

	// Get the public key bytes in uncompressed format
	pubKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)

	// Hash the public key using Keccak-256 (SHA3-256)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes[1:]) // Remove the leading 0x04 byte
	pubKeyHash := hash.Sum(nil)

	// The Ethereum address is the last 20 bytes of the hash
	address := pubKeyHash[12:]
	addressString := hex.EncodeToString(address)

	return &wallet{
		pk:      *privateKey,
		address: toChecksumAddress(addressString),
	}, nil
}

// Function to convert an Ethereum address to EIP-55 checksum format
func toChecksumAddress(address string) string {
	address = strings.ToLower(address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(address))
	hashBytes := hash.Sum(nil)

	var checksumAddress strings.Builder
	for i := 0; i < len(address); i++ {
		if address[i] >= '0' && address[i] <= '9' {
			checksumAddress.WriteByte(address[i])
		} else {
			if hashBytes[i/2]&(0xf<<(4*(1-uint(i)%2))) >= 8 {
				checksumAddress.WriteByte(address[i] - 32) // convert to uppercase
			} else {
				checksumAddress.WriteByte(address[i])
			}
		}
	}
	return "0x" + checksumAddress.String()
}
