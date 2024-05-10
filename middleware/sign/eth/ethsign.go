package eth

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignVerify verify the sign.
func SignVerify(message, sign string, address string) (bool, error) {

	signature, err := hexutil.Decode(sign)
	if err != nil {
		return false, err
	}

	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	messageHash := accounts.TextHash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return false, err
	}

	return common.HexToAddress(address) == crypto.PubkeyToAddress(*pubKey), nil
}
