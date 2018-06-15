package contracts

import (
	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrFailedContractCall = errors.New("calling smart contract failed")
	ErrInvalidEventData   = errors.New("event data could not be parsed")
)

type Event struct {
	Address   common.Address
	TxHash    common.Hash
	BlockHash common.Hash
	Data      [][]byte
}
