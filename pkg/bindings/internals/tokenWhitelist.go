// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internals

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenWhitelistABI is the input ABI used to generate the binding from.
const TokenWhitelistABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"},{\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"name\":\"_burnable\",\"type\":\"bool[]\"},{\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_oracleNameHash\",\"type\":\"bytes32\"},{\"name\":\"_controllerNameHash\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"},{\"name\":\"_stabelcoinAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_burnable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"}]"

// TokenWhitelistBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistBin = `608060405234801561001057600080fd5b5060405160c08061189d83398101604090815281516020830151918301516060840151608085015160a09095015160018054600160a060020a03808716600160a060020a031992831617928390556000805483169382169390931790925560028590556003805489151574010000000000000000000000000000000000000000810260a060020a60ff021995881692909416919091179390931691909117905593959293919291839083906100fc5760408051600160a060020a038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b6040805160008152600160a060020a038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a15050600794909455505060068054600160a060020a031916600160a060020a03909316929092179091555061172190508061017c6000396000f3006080604052600436106100cf5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631f69565f81146100d45780632121dc751461019f57806334c73edc146101c85780633efec5e91461021d5780636c3824ef14610232578063715018a6146102525780637d73b2311461026757806380cc0dda146102985780638da5cb5b146102bf578063b242e534146102d4578063d545782e146102fa578063e2b4ce9714610321578063e81239ac14610336578063e9cbd8221461039b575b600080fd5b3480156100e057600080fd5b506100f5600160a060020a03600435166103b0565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b8381101561015e578181015183820152602001610146565b50505050905090810190601f16801561018b5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156101ab57600080fd5b506101b46104ac565b604080519115158252519081900360200190f35b3480156101d457600080fd5b5061021b602460048035828101929082013591813580830192908201359160443580830192908201359160643580830192908201359160843591820191013560a4356104ce565b005b34801561022957600080fd5b506100f5610965565b34801561023e57600080fd5b5061021b6004803560248101910135610a61565b34801561025e57600080fd5b5061021b610d00565b34801561027357600080fd5b5061027c610e38565b60408051600160a060020a039092168252519081900360200190f35b3480156102a457600080fd5b506102ad610e47565b60408051918252519081900360200190f35b3480156102cb57600080fd5b5061027c610e4d565b3480156102e057600080fd5b5061021b600160a060020a03600435166024351515610e5c565b34801561030657600080fd5b5061021b600160a060020a03600435166024356044356110ac565b34801561032d57600080fd5b506102ad61120e565b34801561034257600080fd5b5061034b611214565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561038757818101518382015260200161036f565b505050509050019250505060405180910390f35b3480156103a757600080fd5b5061027c611276565b600160a060020a0381166000908152600460208181526040808420600180820154600280840154600385015497850154855487516101009682161587026000190190911693909304601f81018990048902840189019097528683526060999889988998899889988998919788979296929560ff808216969382048116956201000090920416939289919083018282801561048b5780601f106104605761010080835404028352916020019161048b565b820191906000526020600020905b81548152906001019060200180831161046e57829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b60035474010000000000000000000000000000000000000000900460ff165b90565b600060606104db33611285565b1515610531576040805160e560020a62461bcd02815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b8b8a14801561053f57508b88145b801561054a57508b86145b801561055557508b86145b15156105ab576040805160e560020a62461bcd02815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b600091505b8b82101561095657600460008e8e858181106105c857fe5b60209081029290920135600160a060020a03168352508101919091526040016000206003015460ff1615610646576040805160e560020a62461bcd02815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b61066c6106678c8c8581811061065857fe5b60200291909101359050611337565b611363565b905060e0604051908101604052808281526020018a8a85818110151561068e57fe5b9050602002013581526020016000815260200160011515815260200188888581811015156106b857fe5b602090810292909201351515835250018686858181106106d457fe5b60209081029290920135151583525001849052600460008f8f868181106106f757fe5b90506020020135600160a060020a0316600160a060020a0316600160a060020a0316815260200190815260200160002060008201518160000190805190602001906107439291906115d6565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560058d8d848181106107ba57fe5b8354600181018555600094855260209485902001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0395909202939093013593909316929092179055507f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e8581811061083157fe5b90506020020135600160a060020a0316838c8c87818110151561085057fe5b905060200201358b8b88818110151561086557fe5b9050602002013515158a8a89818110151561087c57fe5b9050602002013515156040518087600160a060020a0316600160a060020a0316815260200186600160a060020a0316600160a060020a03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b8381101561090c5781810151838201526020016108f4565b50505050905090810190601f1680156109395780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a16001909101906105b0565b50505050505050505050505050565b600654600160a060020a03166000908152600460208181526040808420600180820154600280840154600385015497850154855487516101009682161587026000190190911693909304601f81018990048902840189019097528683526060999889988998899889988998919788979296929560ff8082169693820481169562010000909204169392899190830182828015610a425780601f10610a1757610100808354040283529160200191610a42565b820191906000526020600020905b815481529060010190602001808311610a2557829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b6000806000610a6f33611285565b1515610ac5576040805160e560020a62461bcd02815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b600092505b83831015610cf95760046000868686818110610ae257fe5b60209081029290920135600160a060020a03168352508101919091526040016000206003015460ff161515610b61576040805160e560020a62461bcd02815260206004820152601660248201527f746f6b656e206973206e6f7420617661696c61626c6500000000000000000000604482015290519081900360640190fd5b848484818110610b6d57fe5b60209081029290920135600160a060020a03166000818152600490935260408320909450919050610b9e8282611654565b506000600182018190556002820181905560038201805462ffffff19169055600490910181905590505b600554610bdc90600163ffffffff6113bd16565b811015610c975781600160a060020a0316600582815481101515610bfc57fe5b600091825260209091200154600160a060020a03161415610c8f5760058054610c2c90600163ffffffff6113bd16565b81548110610c3657fe5b60009182526020909120015460058054600160a060020a039092169183908110610c5c57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610c97565b600101610bc8565b6005805490610caa90600019830161169b565b5060408051338152600160a060020a038416602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a1600190920191610aca565b5050505050565b610d086113d4565b1515610d5e576040805160e560020a62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b60035474010000000000000000000000000000000000000000900460ff161515610dd2576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b60035460408051600160a060020a0390921682526000602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a16003805473ffffffffffffffffffffffffffffffffffffffff19169055565b600154600160a060020a031690565b60075490565b600354600160a060020a031690565b610e646113d4565b1515610eba576040805160e560020a62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b60035474010000000000000000000000000000000000000000900460ff161515610f2e576040805160e560020a62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600160a060020a0382161515610fb4576040805160e560020a62461bcd02815260206004820152602360248201527f6f776e65722063616e6e6f742062652073657420746f207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6003805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000831515908102919091179091556110345760408051600160a060020a038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60035460408051600160a060020a039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a1506003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60006110b96007546113e5565b90506110c433611285565b806110d7575033600160a060020a038216145b151561112d576040805160e560020a62461bcd02815260206004820152601b60248201527f656974686572206f7261636c65206f7220636f6e74726f6c6c65720000000000604482015290519081900360640190fd5b600160a060020a03841660009081526004602052604090206003015460ff1615156111a2576040805160e560020a62461bcd02815260206004820152601660248201527f746f6b656e206973206e6f7420617661696c61626c6500000000000000000000604482015290519081900360640190fd5b600160a060020a03841660008181526004602081815260409283902060028101889055909101859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60025490565b6060600580548060200260200160405190810160405280929190818152602001828054801561126c57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161124e575b5050505050905090565b600654600160a060020a031690565b60006112926002546113e5565b600160a060020a031663b429afeb836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561130557600080fd5b505af1158015611319573d6000803e3d6000fd5b505050506040513d602081101561132f57600080fd5b505192915050565b61133f6116c4565b604051602081016040528281528060208301525061135c826114dc565b8152919050565b606080600083600001516040519080825280601f01601f19166020018201604052801561139a578160200160208202803883390190505b5091506020820190506113b68185602001518660000151611592565b5092915050565b600080838311156113cd57600080fd5b5050900390565b600354600160a060020a0316331490565b60008054604080517f0178b8bf000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a0390921691630178b8bf9160248082019260209290919082900301818787803b15801561144c57600080fd5b505af1158015611460573d6000803e3d6000fd5b505050506040513d602081101561147657600080fd5b5051604080517f3b3b57de000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b15801561130557600080fd5b6000808215156114ef576000915061158c565b6fffffffffffffffffffffffffffffffff8316151561152257601001700100000000000000000000000000000000830492505b67ffffffffffffffff831615156115455760080168010000000000000000830492505b63ffffffff8316151561156057600401640100000000830492505b61ffff831615156115775760020162010000830492505b60ff83161515611585576001015b8060200391505b50919050565b60005b602082106115b7578251845260209384019390920191601f1990910190611595565b50905182516020929092036101000a6000190180199091169116179052565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061161757805160ff1916838001178555611644565b82800160010185558215611644579182015b82811115611644578251825591602001919060010190611629565b506116509291506116db565b5090565b50805460018160011615610100020316600290046000825580601f1061167a5750611698565b601f01602090049060005260206000209081019061169891906116db565b50565b8154818355818111156116bf576000838152602090206116bf9181019083016116db565b505050565b604080518082019091526000808252602082015290565b6104cb91905b8082111561165057600081556001016116e15600a165627a7a723058205085904703412b17faa0e10bf19cfe3a9f60fe59105757a1e8f92ab5102704150029`

// DeployTokenWhitelist deploys a new Ethereum contract, binding an instance of TokenWhitelist to it.
func DeployTokenWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _ens common.Address, _oracleNameHash [32]byte, _controllerNameHash [32]byte, _owner common.Address, _transferable bool, _stabelcoinAddress common.Address) (common.Address, *types.Transaction, *TokenWhitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistBin), backend, _ens, _oracleNameHash, _controllerNameHash, _owner, _transferable, _stabelcoinAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenWhitelist{TokenWhitelistCaller: TokenWhitelistCaller{contract: contract}, TokenWhitelistTransactor: TokenWhitelistTransactor{contract: contract}, TokenWhitelistFilterer: TokenWhitelistFilterer{contract: contract}}, nil
}

// TokenWhitelist is an auto generated Go binding around an Ethereum contract.
type TokenWhitelist struct {
	TokenWhitelistCaller     // Read-only binding to the contract
	TokenWhitelistTransactor // Write-only binding to the contract
	TokenWhitelistFilterer   // Log filterer for contract events
}

// TokenWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWhitelistSession struct {
	Contract     *TokenWhitelist   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWhitelistCallerSession struct {
	Contract *TokenWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWhitelistTransactorSession struct {
	Contract     *TokenWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWhitelistRaw struct {
	Contract *TokenWhitelist // Generic contract binding to access the raw methods on
}

// TokenWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWhitelistCallerRaw struct {
	Contract *TokenWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWhitelistTransactorRaw struct {
	Contract *TokenWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWhitelist creates a new instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelist(address common.Address, backend bind.ContractBackend) (*TokenWhitelist, error) {
	contract, err := bindTokenWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelist{TokenWhitelistCaller: TokenWhitelistCaller{contract: contract}, TokenWhitelistTransactor: TokenWhitelistTransactor{contract: contract}, TokenWhitelistFilterer: TokenWhitelistFilterer{contract: contract}}, nil
}

// NewTokenWhitelistCaller creates a new read-only instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistCaller(address common.Address, caller bind.ContractCaller) (*TokenWhitelistCaller, error) {
	contract, err := bindTokenWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistCaller{contract: contract}, nil
}

// NewTokenWhitelistTransactor creates a new write-only instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWhitelistTransactor, error) {
	contract, err := bindTokenWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistTransactor{contract: contract}, nil
}

// NewTokenWhitelistFilterer creates a new log filterer instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWhitelistFilterer, error) {
	contract, err := bindTokenWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistFilterer{contract: contract}, nil
}

// bindTokenWhitelist binds a generic wrapper to an already deployed contract.
func bindTokenWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelist *TokenWhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelist.Contract.TokenWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelist *TokenWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelist *TokenWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelist *TokenWhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelist *TokenWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelist *TokenWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.contract.Transact(opts, method, params...)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) ControllerNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCallerSession) ControllerNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetStablecoinInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getStablecoinInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.CallOpts)
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) GetTokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "getTokenAddressArray")
	return *ret0, err
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) GetTokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.GetTokenAddressArray(&_TokenWhitelist.CallOpts)
}

// GetTokenAddressArray is a free data retrieval call binding the contract method 0xe81239ac.
//
// Solidity: function getTokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) GetTokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.GetTokenAddressArray(&_TokenWhitelist.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetTokenInfo(opts *bind.CallOpts, _a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getTokenInfo", _a)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.CallOpts, _a)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.CallOpts, _a)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsTransferable() (bool, error) {
	return _TokenWhitelist.Contract.IsTransferable(&_TokenWhitelist.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCallerSession) IsTransferable() (bool, error) {
	return _TokenWhitelist.Contract.IsTransferable(&_TokenWhitelist.CallOpts)
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCaller) OracleNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "oracleNode")
	return *ret0, err
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) OracleNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.CallOpts)
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCallerSession) OracleNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) Owner() (common.Address, error) {
	return _TokenWhitelist.Contract.Owner(&_TokenWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) Owner() (common.Address, error) {
	return _TokenWhitelist.Contract.Owner(&_TokenWhitelist.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) Stablecoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "stablecoin")
	return *ret0, err
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) Stablecoin() (common.Address, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) Stablecoin() (common.Address, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _burnable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _burnable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "addTokens", _tokens, _symbols, _magnitude, _loadable, _burnable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _burnable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _burnable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _burnable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _burnable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _burnable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _burnable, _lastUpdate)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) RemoveTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "removeTokens", _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RemoveTokens(&_TokenWhitelist.TransactOpts, _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RemoveTokens(&_TokenWhitelist.TransactOpts, _tokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWhitelist *TokenWhitelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWhitelist *TokenWhitelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RenounceOwnership(&_TokenWhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RenounceOwnership(&_TokenWhitelist.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_TokenWhitelist *TokenWhitelistSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TransferOwnership(&_TokenWhitelist.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TransferOwnership(&_TokenWhitelist.TransactOpts, _account, _transferable)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) UpdateTokenRate(opts *bind.TransactOpts, _token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "updateTokenRate", _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.UpdateTokenRate(&_TokenWhitelist.TransactOpts, _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.UpdateTokenRate(&_TokenWhitelist.TransactOpts, _token, _rate, _updateDate)
}

// TokenWhitelistAddedTokenIterator is returned from FilterAddedToken and is used to iterate over the raw logs and unpacked data for AddedToken events raised by the TokenWhitelist contract.
type TokenWhitelistAddedTokenIterator struct {
	Event *TokenWhitelistAddedToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWhitelistAddedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWhitelistAddedToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWhitelistAddedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedToken represents a AddedToken event raised by the TokenWhitelist contract.
type TokenWhitelistAddedToken struct {
	Sender    common.Address
	Token     common.Address
	Symbol    string
	Magnitude *big.Int
	Loadable  bool
	Burnable  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _burnable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedToken(opts *bind.FilterOpts) (*TokenWhitelistAddedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedTokenIterator{contract: _TokenWhitelist.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _burnable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedToken(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedToken) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedToken)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenWhitelistLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the TokenWhitelist contract.
type TokenWhitelistLockedOwnershipIterator struct {
	Event *TokenWhitelistLockedOwnership // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWhitelistLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistLockedOwnership)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWhitelistLockedOwnership)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWhitelistLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistLockedOwnership represents a LockedOwnership event raised by the TokenWhitelist contract.
type TokenWhitelistLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*TokenWhitelistLockedOwnershipIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistLockedOwnershipIterator{contract: _TokenWhitelist.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *TokenWhitelistLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistLockedOwnership)
				if err := _TokenWhitelist.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenWhitelistRemovedTokenIterator is returned from FilterRemovedToken and is used to iterate over the raw logs and unpacked data for RemovedToken events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedTokenIterator struct {
	Event *TokenWhitelistRemovedToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWhitelistRemovedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWhitelistRemovedToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWhitelistRemovedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedToken represents a RemovedToken event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedToken struct {
	Sender common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedToken is a free log retrieval operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedToken(opts *bind.FilterOpts) (*TokenWhitelistRemovedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedTokenIterator{contract: _TokenWhitelist.contract, event: "RemovedToken", logs: logs, sub: sub}, nil
}

// WatchRemovedToken is a free log subscription operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedToken(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedToken) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedToken)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenWhitelistTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the TokenWhitelist contract.
type TokenWhitelistTransferredOwnershipIterator struct {
	Event *TokenWhitelistTransferredOwnership // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWhitelistTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistTransferredOwnership)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWhitelistTransferredOwnership)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWhitelistTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistTransferredOwnership represents a TransferredOwnership event raised by the TokenWhitelist contract.
type TokenWhitelistTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*TokenWhitelistTransferredOwnershipIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistTransferredOwnershipIterator{contract: _TokenWhitelist.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *TokenWhitelistTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistTransferredOwnership)
				if err := _TokenWhitelist.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenWhitelistUpdatedTokenRateIterator is returned from FilterUpdatedTokenRate and is used to iterate over the raw logs and unpacked data for UpdatedTokenRate events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRateIterator struct {
	Event *TokenWhitelistUpdatedTokenRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWhitelistUpdatedTokenRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWhitelistUpdatedTokenRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWhitelistUpdatedTokenRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenRate represents a UpdatedTokenRate event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRate struct {
	Sender common.Address
	Token  common.Address
	Rate   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenRate is a free log retrieval operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenRate(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenRateIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenRate")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenRateIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenRate is a free log subscription operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenRate(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenRate) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenRate)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
