package contracts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	ErrFailedContractCall = errors.New("calling smart contract failed")
	ErrInvalidEventData   = errors.New("event data could not be parsed")
	ErrNoCode             = errors.New("no contract code at given address")
)

// Event represents a raw smart contract event.
type Event struct {
	Address   common.Address
	TxHash    common.Hash
	BlockHash common.Hash
	Data      [][]byte
}

type SignTransaction func(*types.Transaction) (*types.Transaction, error)

type ConstructOpts struct {
	From     common.Address  // Ethereum account to send the transaction from
	Nonce    uint64          // Nonce to use for the transaction execution
	Value    *big.Int        // Funds to transfer along along the transaction (nil = 0 = no funds)
	GasPrice *big.Int        // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64          // Gas limit to set for the transaction execution (0 = estimate)
	Sign     SignTransaction // Method to use for signing the transaction (nil = return unsigned transaction)
	Context  context.Context // Network context to support cancellation and timeouts
}

// construct creates a signed transaction object from the provided parameters.
func construct(opts *ConstructOpts, eth *ethclient.Client, address *common.Address, data []byte) (*types.Transaction, error) {
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var err error
	gasPrice := opts.GasPrice
	if opts.GasPrice == nil {
		gasPrice, err = eth.SuggestGasPrice(opts.Context)
		if err != nil {
			return nil, err
		}
	}
	gasLimit := opts.GasLimit
	if opts.GasLimit == 0 {
		if address != nil {
			// Check if contract has code.
			code, err := eth.PendingCodeAt(opts.Context, *address)
			if err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, ErrNoCode
			}
		}
		gasLimit, err = eth.EstimateGas(opts.Context, ethereum.CallMsg{
			From:  opts.From,
			To:    address,
			Value: value,
			Data:  data,
		})
		if err != nil {
			return nil, err
		}
	}
	if address == nil && opts.Sign == nil {
		return types.NewContractCreation(opts.Nonce, new(big.Int), gasLimit+gasLimit/10, gasPrice, data), nil
	} else if address != nil && opts.Sign == nil {
		return types.NewTransaction(opts.Nonce, *address, new(big.Int), gasLimit+gasLimit/10, gasPrice, data), nil
	} else if address == nil && opts.Sign != nil {
		return opts.Sign(types.NewContractCreation(opts.Nonce, new(big.Int), gasLimit+gasLimit/10, gasPrice, data))
	} else {
		return opts.Sign(types.NewTransaction(opts.Nonce, *address, new(big.Int), gasLimit+gasLimit/10, gasPrice, data))
	}
}
