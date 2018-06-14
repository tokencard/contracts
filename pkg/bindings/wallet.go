// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setGasLimitConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingRemoval\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setGasLimitCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingDailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addToWhitelistConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setDailyLimitConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeDailyLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setDailyLimitCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addToWhitelistCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeFromWhitelistCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAddition\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeFromWhitelistConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"initializeWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_controllers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TopUpGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetDailyLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoval\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `0x608060405234801561001057600080fd5b50604051612612380380612612833981018060405281019080805190602001909291908051906020019092919080518201929190505050828282600042600681905550670de0b6b3a764000060078190555060075460088190555083600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600a60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600090505b8151811015610176576001600080848481518110151561010e57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080806001019150506100f2565b5050505067016345785d8a0000600b81905550600b54600c8190555050505061246e806101a46000396000f3006080604052600436106101a1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630a726ac5146102185780631dafafbd146102455780632ea72b521461025c5780633af32abf146102c85780633d3343ef14610323578063441397a81461034e57806347064644146103655780634c58ab1d14610390578063548db174146103a75780635c9302c91461040d57806367eeba0c146104385780637cb8c9cb146104635780637dc0d1d01461048e5780637f649783146104e55780638da5cb5b1461054b5780639074d0ca146105a2578063a7fc7a07146105b9578063aaddce72146105fc578063b20d30a914610629578063b429afeb14610656578063b43c120d146106b1578063beabacc8146106c8578063d65b583114610735578063d9e98d4514610760578063da97cb8214610777578063dd0e91701461078e578063e30baedf146107fa578063e3d670d714610811578063e61c51ca14610868578063ee7d72b414610895578063f4199bb8146108c2578063f68016b714610928578063f6a74ed714610953575b6000341115610216577fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b005b34801561022457600080fd5b5061024360048036038101908080359060200190929190505050610996565b005b34801561025157600080fd5b5061025a610a94565b005b34801561026857600080fd5b50610271610b90565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102b4578082015181840152602081019050610299565b505050509050019250505060405180910390f35b3480156102d457600080fd5b50610309600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c1e565b604051808215151515815260200191505060405180910390f35b34801561032f57600080fd5b50610338610c3e565b6040518082815260200191505060405180910390f35b34801561035a57600080fd5b50610363610c62565b005b34801561037157600080fd5b5061037a610cde565b6040518082815260200191505060405180910390f35b34801561039c57600080fd5b506103a5610ce4565b005b3480156103b357600080fd5b5061040b60048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610ee3565b005b34801561041957600080fd5b50610422611025565b6040518082815260200191505060405180910390f35b34801561044457600080fd5b5061044d61102b565b6040518082815260200191505060405180910390f35b34801561046f57600080fd5b50610478611031565b6040518082815260200191505060405180910390f35b34801561049a57600080fd5b506104a3611055565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f157600080fd5b506105496004803603810190808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061107b565b005b34801561055757600080fd5b506105606111bd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105ae57600080fd5b506105b76111e3565b005b3480156105c557600080fd5b506105fa600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112b4565b005b34801561060857600080fd5b5061062760048036038101908080359060200190929190505050611365565b005b34801561063557600080fd5b5061065460048036038101908080359060200190929190505050611439565b005b34801561066257600080fd5b50610697600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114d6565b604051808215151515815260200191505060405180910390f35b3480156106bd57600080fd5b506106c66114f6565b005b3480156106d457600080fd5b50610733600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611572565b005b34801561074157600080fd5b5061074a6119c9565b6040518082815260200191505060405180910390f35b34801561076c57600080fd5b506107756119cf565b005b34801561078357600080fd5b5061078c611a51565b005b34801561079a57600080fd5b506107a3611ad3565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156107e65780820151818401526020810190506107cb565b505050509050019250505060405180910390f35b34801561080657600080fd5b5061080f611b61565b005b34801561081d57600080fd5b50610852600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d60565b6040518082815260200191505060405180910390f35b34801561087457600080fd5b5061089360048036038101908080359060200190929190505050611e7e565b005b3480156108a157600080fd5b506108c0600480360381019080803590602001909291905050506120e3565b005b3480156108ce57600080fd5b50610926600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192905050506121aa565b005b34801561093457600080fd5b5061093d612345565b6040518082815260200191505060405180910390f35b34801561095f57600080fd5b50610994600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061234b565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109f257600080fd5b600e60019054906101000a900460ff16151515610a0e57600080fd5b8066038d7ea4c6800011158015610a2d575067016345785d8a00008111155b1515610a3857600080fd5b80600b819055506001600e60016101000a81548160ff0219169083151502179055507ff5bdeca176beddded5a1132996e4edf0d16be5100214b17d8bada54bf8676297816040518082815260200191505060405180910390a150565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610aeb57600080fd5b600e60009054906101000a900460ff161515610b0657600080fd5b60095466038d7ea4c6800011158015610b29575067016345785d8a000060095411155b1515610b3157fe5b600954600b819055506000600e60006101000a81548160ff0219169083151502179055507ff5bdeca176beddded5a1132996e4edf0d16be5100214b17d8bada54bf86762976009546040518082815260200191505060405180910390a1565b60606004805480602002602001604051908101604052809291908181526020018280548015610c1457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610bca575b5050505050905090565b60026020528060005260406000206000915054906101000a900460ff1681565b60006201518060065401421115610c5957600b549050610c5f565b600c5490505b90565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610cb957600080fd5b60006009819055506000600e60006101000a81548160ff021916908315150217905550565b60095481565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610d3d57600080fd5b6000600380549050118015610d5e5750600560009054906101000a900460ff165b1515610d6957600080fd5b600090505b600380549050811015610e1957600160026000600384815481101515610d9057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050610d6e565b6000600560006101000a81548160ff0219169083151502179055507fc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434600360405180806020018281038252838181548152602001915080548015610ed257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e88575b50509250505060405180910390a150565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f4157600080fd5b600560019054906101000a900460ff16151515610f5d57600080fd5b6014825111151515610f6e57600080fd5b600090505b81518110156110065760048282815181101515610f8c57fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508080600101915050610f73565b6001600560016101000a81548160ff0219169083151502179055505050565b60065481565b60075481565b6000620151806006540142111561104c576007549050611052565b60085490505b90565b600a60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110d957600080fd5b600560009054906101000a900460ff161515156110f557600080fd5b601482511115151561110657600080fd5b600090505b815181101561119e576003828281518110151561112457fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050808060010191505061110b565b6001600560006101000a81548160ff0219169083151502179055505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561123a57600080fd5b600a60009054906101000a900460ff16151561125557600080fd5b6009546007819055506000600a60006101000a81548160ff0219169083151502179055507f4c6fb4d469797cfdfb4ac382c2fe84a494711f5b9676334ff9fac5f7601aef706009546040518082815260200191505060405180910390a1565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561130b57600080fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113c157600080fd5b600a60019054906101000a900460ff161515156113dd57600080fd5b806007819055506001600a60016101000a81548160ff0219169083151502179055507f4c6fb4d469797cfdfb4ac382c2fe84a494711f5b9676334ff9fac5f7601aef70816040518082815260200191505060405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561149557600080fd5b600a60009054906101000a900460ff161515156114b157600080fd5b806009819055506001600a60006101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915054906101000a900460ff1681565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561154d57600080fd5b60006009819055506000600a60006101000a81548160ff021916908315150217905550565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115d157600080fd5b82600081141515156115e257600080fd5b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156117cc5760008573ffffffffffffffffffffffffffffffffffffffff161415156117635783600a60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ba9d8ca876040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561171357600080fd5b505af1158015611727573d6000803e3d6000fd5b505050506040513d602081101561173d57600080fd5b81019080805190602001909291905050500292506000831415151561175e57fe5b611767565b8392505b62015180600654014211156117aa5762015180600654420381151561178857fe5b0491506201518082026006600082825401925050819055506007546008819055505b60085483111515156117bb57600080fd5b826008600082825403925050819055505b60008573ffffffffffffffffffffffffffffffffffffffff161415156118da578473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561188f57600080fd5b505af11580156118a3573d6000803e3d6000fd5b505050506040513d60208110156118b957600080fd5b810190808051906020019092919050505015156118d557600080fd5b611922565b8573ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f19350505050158015611920573d6000803e3d6000fd5b505b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef868686604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050505050565b600d5481565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611a2657600080fd5b60036000611a3491906123fc565b6000600560006101000a81548160ff021916908315150217905550565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611aa857600080fd5b60046000611ab691906123fc565b6000600560016101000a81548160ff021916908315150217905550565b60606003805480602002602001604051908101604052809291908181526020018280548015611b5757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611b0d575b5050505050905090565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611bba57600080fd5b6000600480549050118015611bdb5750600560019054906101000a900460ff165b1515611be657600080fd5b600090505b600480549050811015611c9657600060026000600484815481101515611c0d57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050611beb565b6000600560016101000a81548160ff0219169083151502179055507f4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94600460405180806020018281038252838181548152602001915080548015611d4f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d05575b50509250505060405180910390a150565b6000808273ffffffffffffffffffffffffffffffffffffffff16141515611e5e578173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611e1c57600080fd5b505af1158015611e30573d6000803e3d6000fd5b505050506040513d6020811015611e4657600080fd5b81019080805190602001909291905050509050611e79565b3073ffffffffffffffffffffffffffffffffffffffff163190505b919050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680611f245750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611f2f57600080fd5b8160008114151515611f4057600080fd5b6000600c54111515611f5157600080fd5b6201518060065401421115611f9457620151806006544203811515611f7257fe5b049150620151808202600660008282540192505081905550600b54600c819055505b600c54831115611fa457600c5492505b82600c60008282540392505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f1935050505015801561201c573d6000803e3d6000fd5b507f2235de9f3363e464311d990c51aeef966703c87d1c77e80737831d6944d87c8632600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561213f57600080fd5b600e60009054906101000a900460ff1615151561215b57600080fd5b8066038d7ea4c680001115801561217a575067016345785d8a00008111155b151561218557600080fd5b806009819055506001600e60006101000a81548160ff02191690831515021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561220857600080fd5b600560029054906101000a900460ff1615151561222457600080fd5b600090505b81518110156122ae57600160026000848481518110151561224657fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050612229565b6001600560026101000a81548160ff0219169083151502179055507fc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434826040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561232e578082015181840152602081019050612313565b505050509050019250505060405180910390a15050565b600b5481565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156123a257600080fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b508054600082559060005260206000209081019061241a919061241d565b50565b61243f91905b8082111561243b576000816000905550600101612423565b5090565b905600a165627a7a72305820ebb6f30bd0fa445f09e8e3db7050b6eed97e9eae3edce9a5727f3fbedd22e5380029`

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _oracle common.Address, _controllers []common.Address) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, _owner, _oracle, _controllers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(_asset address) constant returns(uint256)
func (_Wallet *WalletCaller) Balance(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "balance", _asset)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(_asset address) constant returns(uint256)
func (_Wallet *WalletSession) Balance(_asset common.Address) (*big.Int, error) {
	return _Wallet.Contract.Balance(&_Wallet.CallOpts, _asset)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(_asset address) constant returns(uint256)
func (_Wallet *WalletCallerSession) Balance(_asset common.Address) (*big.Int, error) {
	return _Wallet.Contract.Balance(&_Wallet.CallOpts, _asset)
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() constant returns(uint256)
func (_Wallet *WalletCaller) CurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "currentDay")
	return *ret0, err
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() constant returns(uint256)
func (_Wallet *WalletSession) CurrentDay() (*big.Int, error) {
	return _Wallet.Contract.CurrentDay(&_Wallet.CallOpts)
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() constant returns(uint256)
func (_Wallet *WalletCallerSession) CurrentDay() (*big.Int, error) {
	return _Wallet.Contract.CurrentDay(&_Wallet.CallOpts)
}

// DailyAvailable is a free data retrieval call binding the contract method 0x7cb8c9cb.
//
// Solidity: function dailyAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) DailyAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "dailyAvailable")
	return *ret0, err
}

// DailyAvailable is a free data retrieval call binding the contract method 0x7cb8c9cb.
//
// Solidity: function dailyAvailable() constant returns(uint256)
func (_Wallet *WalletSession) DailyAvailable() (*big.Int, error) {
	return _Wallet.Contract.DailyAvailable(&_Wallet.CallOpts)
}

// DailyAvailable is a free data retrieval call binding the contract method 0x7cb8c9cb.
//
// Solidity: function dailyAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) DailyAvailable() (*big.Int, error) {
	return _Wallet.Contract.DailyAvailable(&_Wallet.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_Wallet *WalletCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "dailyLimit")
	return *ret0, err
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_Wallet *WalletSession) DailyLimit() (*big.Int, error) {
	return _Wallet.Contract.DailyLimit(&_Wallet.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) DailyLimit() (*big.Int, error) {
	return _Wallet.Contract.DailyLimit(&_Wallet.CallOpts)
}

// GasAvailable is a free data retrieval call binding the contract method 0x3d3343ef.
//
// Solidity: function gasAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) GasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasAvailable")
	return *ret0, err
}

// GasAvailable is a free data retrieval call binding the contract method 0x3d3343ef.
//
// Solidity: function gasAvailable() constant returns(uint256)
func (_Wallet *WalletSession) GasAvailable() (*big.Int, error) {
	return _Wallet.Contract.GasAvailable(&_Wallet.CallOpts)
}

// GasAvailable is a free data retrieval call binding the contract method 0x3d3343ef.
//
// Solidity: function gasAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) GasAvailable() (*big.Int, error) {
	return _Wallet.Contract.GasAvailable(&_Wallet.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() constant returns(uint256)
func (_Wallet *WalletCaller) GasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "gasLimit")
	return *ret0, err
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() constant returns(uint256)
func (_Wallet *WalletSession) GasLimit() (*big.Int, error) {
	return _Wallet.Contract.GasLimit(&_Wallet.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) GasLimit() (*big.Int, error) {
	return _Wallet.Contract.GasLimit(&_Wallet.CallOpts)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController( address) constant returns(bool)
func (_Wallet *WalletCaller) IsController(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isController", arg0)
	return *ret0, err
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController( address) constant returns(bool)
func (_Wallet *WalletSession) IsController(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsController(&_Wallet.CallOpts, arg0)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController( address) constant returns(bool)
func (_Wallet *WalletCallerSession) IsController(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsController(&_Wallet.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted( address) constant returns(bool)
func (_Wallet *WalletCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isWhitelisted", arg0)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted( address) constant returns(bool)
func (_Wallet *WalletSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsWhitelisted(&_Wallet.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted( address) constant returns(bool)
func (_Wallet *WalletCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsWhitelisted(&_Wallet.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Wallet *WalletCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Wallet *WalletSession) Oracle() (common.Address, error) {
	return _Wallet.Contract.Oracle(&_Wallet.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Wallet *WalletCallerSession) Oracle() (common.Address, error) {
	return _Wallet.Contract.Oracle(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// PendingAddition is a free data retrieval call binding the contract method 0xdd0e9170.
//
// Solidity: function pendingAddition() constant returns(address[])
func (_Wallet *WalletCaller) PendingAddition(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingAddition")
	return *ret0, err
}

// PendingAddition is a free data retrieval call binding the contract method 0xdd0e9170.
//
// Solidity: function pendingAddition() constant returns(address[])
func (_Wallet *WalletSession) PendingAddition() ([]common.Address, error) {
	return _Wallet.Contract.PendingAddition(&_Wallet.CallOpts)
}

// PendingAddition is a free data retrieval call binding the contract method 0xdd0e9170.
//
// Solidity: function pendingAddition() constant returns(address[])
func (_Wallet *WalletCallerSession) PendingAddition() ([]common.Address, error) {
	return _Wallet.Contract.PendingAddition(&_Wallet.CallOpts)
}

// PendingDailyLimit is a free data retrieval call binding the contract method 0x47064644.
//
// Solidity: function pendingDailyLimit() constant returns(uint256)
func (_Wallet *WalletCaller) PendingDailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingDailyLimit")
	return *ret0, err
}

// PendingDailyLimit is a free data retrieval call binding the contract method 0x47064644.
//
// Solidity: function pendingDailyLimit() constant returns(uint256)
func (_Wallet *WalletSession) PendingDailyLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingDailyLimit(&_Wallet.CallOpts)
}

// PendingDailyLimit is a free data retrieval call binding the contract method 0x47064644.
//
// Solidity: function pendingDailyLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) PendingDailyLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingDailyLimit(&_Wallet.CallOpts)
}

// PendingGasLimit is a free data retrieval call binding the contract method 0xd65b5831.
//
// Solidity: function pendingGasLimit() constant returns(uint256)
func (_Wallet *WalletCaller) PendingGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingGasLimit")
	return *ret0, err
}

// PendingGasLimit is a free data retrieval call binding the contract method 0xd65b5831.
//
// Solidity: function pendingGasLimit() constant returns(uint256)
func (_Wallet *WalletSession) PendingGasLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingGasLimit(&_Wallet.CallOpts)
}

// PendingGasLimit is a free data retrieval call binding the contract method 0xd65b5831.
//
// Solidity: function pendingGasLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) PendingGasLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingGasLimit(&_Wallet.CallOpts)
}

// PendingRemoval is a free data retrieval call binding the contract method 0x2ea72b52.
//
// Solidity: function pendingRemoval() constant returns(address[])
func (_Wallet *WalletCaller) PendingRemoval(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingRemoval")
	return *ret0, err
}

// PendingRemoval is a free data retrieval call binding the contract method 0x2ea72b52.
//
// Solidity: function pendingRemoval() constant returns(address[])
func (_Wallet *WalletSession) PendingRemoval() ([]common.Address, error) {
	return _Wallet.Contract.PendingRemoval(&_Wallet.CallOpts)
}

// PendingRemoval is a free data retrieval call binding the contract method 0x2ea72b52.
//
// Solidity: function pendingRemoval() constant returns(address[])
func (_Wallet *WalletCallerSession) PendingRemoval() ([]common.Address, error) {
	return _Wallet.Contract.PendingRemoval(&_Wallet.CallOpts)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(_account address) returns()
func (_Wallet *WalletTransactor) AddController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addController", _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(_account address) returns()
func (_Wallet *WalletSession) AddController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddController(&_Wallet.TransactOpts, _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(_account address) returns()
func (_Wallet *WalletTransactorSession) AddController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddController(&_Wallet.TransactOpts, _account)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactor) AddToWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addToWhitelist", _addresses)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(_addresses address[]) returns()
func (_Wallet *WalletSession) AddToWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddToWhitelist(&_Wallet.TransactOpts, _addresses)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactorSession) AddToWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddToWhitelist(&_Wallet.TransactOpts, _addresses)
}

// AddToWhitelistCancel is a paid mutator transaction binding the contract method 0xd9e98d45.
//
// Solidity: function addToWhitelistCancel() returns()
func (_Wallet *WalletTransactor) AddToWhitelistCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addToWhitelistCancel")
}

// AddToWhitelistCancel is a paid mutator transaction binding the contract method 0xd9e98d45.
//
// Solidity: function addToWhitelistCancel() returns()
func (_Wallet *WalletSession) AddToWhitelistCancel() (*types.Transaction, error) {
	return _Wallet.Contract.AddToWhitelistCancel(&_Wallet.TransactOpts)
}

// AddToWhitelistCancel is a paid mutator transaction binding the contract method 0xd9e98d45.
//
// Solidity: function addToWhitelistCancel() returns()
func (_Wallet *WalletTransactorSession) AddToWhitelistCancel() (*types.Transaction, error) {
	return _Wallet.Contract.AddToWhitelistCancel(&_Wallet.TransactOpts)
}

// AddToWhitelistConfirm is a paid mutator transaction binding the contract method 0x4c58ab1d.
//
// Solidity: function addToWhitelistConfirm() returns()
func (_Wallet *WalletTransactor) AddToWhitelistConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addToWhitelistConfirm")
}

// AddToWhitelistConfirm is a paid mutator transaction binding the contract method 0x4c58ab1d.
//
// Solidity: function addToWhitelistConfirm() returns()
func (_Wallet *WalletSession) AddToWhitelistConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.AddToWhitelistConfirm(&_Wallet.TransactOpts)
}

// AddToWhitelistConfirm is a paid mutator transaction binding the contract method 0x4c58ab1d.
//
// Solidity: function addToWhitelistConfirm() returns()
func (_Wallet *WalletTransactorSession) AddToWhitelistConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.AddToWhitelistConfirm(&_Wallet.TransactOpts)
}

// InitializeDailyLimit is a paid mutator transaction binding the contract method 0xaaddce72.
//
// Solidity: function initializeDailyLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) InitializeDailyLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeDailyLimit", _amount)
}

// InitializeDailyLimit is a paid mutator transaction binding the contract method 0xaaddce72.
//
// Solidity: function initializeDailyLimit(_amount uint256) returns()
func (_Wallet *WalletSession) InitializeDailyLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeDailyLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeDailyLimit is a paid mutator transaction binding the contract method 0xaaddce72.
//
// Solidity: function initializeDailyLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) InitializeDailyLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeDailyLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeGasLimit is a paid mutator transaction binding the contract method 0x0a726ac5.
//
// Solidity: function initializeGasLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) InitializeGasLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeGasLimit", _amount)
}

// InitializeGasLimit is a paid mutator transaction binding the contract method 0x0a726ac5.
//
// Solidity: function initializeGasLimit(_amount uint256) returns()
func (_Wallet *WalletSession) InitializeGasLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeGasLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeGasLimit is a paid mutator transaction binding the contract method 0x0a726ac5.
//
// Solidity: function initializeGasLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) InitializeGasLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeGasLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeWhitelist is a paid mutator transaction binding the contract method 0xf4199bb8.
//
// Solidity: function initializeWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactor) InitializeWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeWhitelist", _addresses)
}

// InitializeWhitelist is a paid mutator transaction binding the contract method 0xf4199bb8.
//
// Solidity: function initializeWhitelist(_addresses address[]) returns()
func (_Wallet *WalletSession) InitializeWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeWhitelist(&_Wallet.TransactOpts, _addresses)
}

// InitializeWhitelist is a paid mutator transaction binding the contract method 0xf4199bb8.
//
// Solidity: function initializeWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactorSession) InitializeWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeWhitelist(&_Wallet.TransactOpts, _addresses)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(_account address) returns()
func (_Wallet *WalletTransactor) RemoveController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeController", _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(_account address) returns()
func (_Wallet *WalletSession) RemoveController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveController(&_Wallet.TransactOpts, _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(_account address) returns()
func (_Wallet *WalletTransactorSession) RemoveController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveController(&_Wallet.TransactOpts, _account)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeFromWhitelist", _addresses)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(_addresses address[]) returns()
func (_Wallet *WalletSession) RemoveFromWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveFromWhitelist(&_Wallet.TransactOpts, _addresses)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactorSession) RemoveFromWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveFromWhitelist(&_Wallet.TransactOpts, _addresses)
}

// RemoveFromWhitelistCancel is a paid mutator transaction binding the contract method 0xda97cb82.
//
// Solidity: function removeFromWhitelistCancel() returns()
func (_Wallet *WalletTransactor) RemoveFromWhitelistCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeFromWhitelistCancel")
}

// RemoveFromWhitelistCancel is a paid mutator transaction binding the contract method 0xda97cb82.
//
// Solidity: function removeFromWhitelistCancel() returns()
func (_Wallet *WalletSession) RemoveFromWhitelistCancel() (*types.Transaction, error) {
	return _Wallet.Contract.RemoveFromWhitelistCancel(&_Wallet.TransactOpts)
}

// RemoveFromWhitelistCancel is a paid mutator transaction binding the contract method 0xda97cb82.
//
// Solidity: function removeFromWhitelistCancel() returns()
func (_Wallet *WalletTransactorSession) RemoveFromWhitelistCancel() (*types.Transaction, error) {
	return _Wallet.Contract.RemoveFromWhitelistCancel(&_Wallet.TransactOpts)
}

// RemoveFromWhitelistConfirm is a paid mutator transaction binding the contract method 0xe30baedf.
//
// Solidity: function removeFromWhitelistConfirm() returns()
func (_Wallet *WalletTransactor) RemoveFromWhitelistConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeFromWhitelistConfirm")
}

// RemoveFromWhitelistConfirm is a paid mutator transaction binding the contract method 0xe30baedf.
//
// Solidity: function removeFromWhitelistConfirm() returns()
func (_Wallet *WalletSession) RemoveFromWhitelistConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.RemoveFromWhitelistConfirm(&_Wallet.TransactOpts)
}

// RemoveFromWhitelistConfirm is a paid mutator transaction binding the contract method 0xe30baedf.
//
// Solidity: function removeFromWhitelistConfirm() returns()
func (_Wallet *WalletTransactorSession) RemoveFromWhitelistConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.RemoveFromWhitelistConfirm(&_Wallet.TransactOpts)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) SetDailyLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setDailyLimit", _amount)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(_amount uint256) returns()
func (_Wallet *WalletSession) SetDailyLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetDailyLimit(&_Wallet.TransactOpts, _amount)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) SetDailyLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetDailyLimit(&_Wallet.TransactOpts, _amount)
}

// SetDailyLimitCancel is a paid mutator transaction binding the contract method 0xb43c120d.
//
// Solidity: function setDailyLimitCancel() returns()
func (_Wallet *WalletTransactor) SetDailyLimitCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setDailyLimitCancel")
}

// SetDailyLimitCancel is a paid mutator transaction binding the contract method 0xb43c120d.
//
// Solidity: function setDailyLimitCancel() returns()
func (_Wallet *WalletSession) SetDailyLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetDailyLimitCancel(&_Wallet.TransactOpts)
}

// SetDailyLimitCancel is a paid mutator transaction binding the contract method 0xb43c120d.
//
// Solidity: function setDailyLimitCancel() returns()
func (_Wallet *WalletTransactorSession) SetDailyLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetDailyLimitCancel(&_Wallet.TransactOpts)
}

// SetDailyLimitConfirm is a paid mutator transaction binding the contract method 0x9074d0ca.
//
// Solidity: function setDailyLimitConfirm() returns()
func (_Wallet *WalletTransactor) SetDailyLimitConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setDailyLimitConfirm")
}

// SetDailyLimitConfirm is a paid mutator transaction binding the contract method 0x9074d0ca.
//
// Solidity: function setDailyLimitConfirm() returns()
func (_Wallet *WalletSession) SetDailyLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetDailyLimitConfirm(&_Wallet.TransactOpts)
}

// SetDailyLimitConfirm is a paid mutator transaction binding the contract method 0x9074d0ca.
//
// Solidity: function setDailyLimitConfirm() returns()
func (_Wallet *WalletTransactorSession) SetDailyLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetDailyLimitConfirm(&_Wallet.TransactOpts)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xee7d72b4.
//
// Solidity: function setGasLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) SetGasLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setGasLimit", _amount)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xee7d72b4.
//
// Solidity: function setGasLimit(_amount uint256) returns()
func (_Wallet *WalletSession) SetGasLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetGasLimit(&_Wallet.TransactOpts, _amount)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xee7d72b4.
//
// Solidity: function setGasLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) SetGasLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetGasLimit(&_Wallet.TransactOpts, _amount)
}

// SetGasLimitCancel is a paid mutator transaction binding the contract method 0x441397a8.
//
// Solidity: function setGasLimitCancel() returns()
func (_Wallet *WalletTransactor) SetGasLimitCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setGasLimitCancel")
}

// SetGasLimitCancel is a paid mutator transaction binding the contract method 0x441397a8.
//
// Solidity: function setGasLimitCancel() returns()
func (_Wallet *WalletSession) SetGasLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetGasLimitCancel(&_Wallet.TransactOpts)
}

// SetGasLimitCancel is a paid mutator transaction binding the contract method 0x441397a8.
//
// Solidity: function setGasLimitCancel() returns()
func (_Wallet *WalletTransactorSession) SetGasLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetGasLimitCancel(&_Wallet.TransactOpts)
}

// SetGasLimitConfirm is a paid mutator transaction binding the contract method 0x1dafafbd.
//
// Solidity: function setGasLimitConfirm() returns()
func (_Wallet *WalletTransactor) SetGasLimitConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setGasLimitConfirm")
}

// SetGasLimitConfirm is a paid mutator transaction binding the contract method 0x1dafafbd.
//
// Solidity: function setGasLimitConfirm() returns()
func (_Wallet *WalletSession) SetGasLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetGasLimitConfirm(&_Wallet.TransactOpts)
}

// SetGasLimitConfirm is a paid mutator transaction binding the contract method 0x1dafafbd.
//
// Solidity: function setGasLimitConfirm() returns()
func (_Wallet *WalletTransactorSession) SetGasLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetGasLimitConfirm(&_Wallet.TransactOpts)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(_amount uint256) returns()
func (_Wallet *WalletTransactor) TopUpGas(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "topUpGas", _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(_amount uint256) returns()
func (_Wallet *WalletSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopUpGas(&_Wallet.TransactOpts, _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopUpGas(&_Wallet.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_to address, _asset address, _amount uint256) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_to address, _asset address, _amount uint256) returns()
func (_Wallet *WalletSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_to address, _asset address, _amount uint256) returns()
func (_Wallet *WalletTransactorSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _asset, _amount)
}

// WalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Wallet contract.
type WalletDepositIterator struct {
	Event *WalletDeposit // Event containing the contract specifics and raw log

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
func (it *WalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletDeposit)
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
		it.Event = new(WalletDeposit)
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
func (it *WalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletDeposit represents a Deposit event raised by the Wallet contract.
type WalletDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(_from address, _amount uint256)
func (_Wallet *WalletFilterer) FilterDeposit(opts *bind.FilterOpts) (*WalletDepositIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &WalletDepositIterator{contract: _Wallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(_from address, _amount uint256)
func (_Wallet *WalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WalletDeposit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletDeposit)
				if err := _Wallet.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// WalletSetDailyLimitIterator is returned from FilterSetDailyLimit and is used to iterate over the raw logs and unpacked data for SetDailyLimit events raised by the Wallet contract.
type WalletSetDailyLimitIterator struct {
	Event *WalletSetDailyLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetDailyLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetDailyLimit)
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
		it.Event = new(WalletSetDailyLimit)
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
func (it *WalletSetDailyLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetDailyLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetDailyLimit represents a SetDailyLimit event raised by the Wallet contract.
type WalletSetDailyLimit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetDailyLimit is a free log retrieval operation binding the contract event 0x4c6fb4d469797cfdfb4ac382c2fe84a494711f5b9676334ff9fac5f7601aef70.
//
// Solidity: event SetDailyLimit(_amount uint256)
func (_Wallet *WalletFilterer) FilterSetDailyLimit(opts *bind.FilterOpts) (*WalletSetDailyLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetDailyLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetDailyLimitIterator{contract: _Wallet.contract, event: "SetDailyLimit", logs: logs, sub: sub}, nil
}

// WatchSetDailyLimit is a free log subscription operation binding the contract event 0x4c6fb4d469797cfdfb4ac382c2fe84a494711f5b9676334ff9fac5f7601aef70.
//
// Solidity: event SetDailyLimit(_amount uint256)
func (_Wallet *WalletFilterer) WatchSetDailyLimit(opts *bind.WatchOpts, sink chan<- *WalletSetDailyLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetDailyLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetDailyLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetDailyLimit", log); err != nil {
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

// WalletSetGasLimitIterator is returned from FilterSetGasLimit and is used to iterate over the raw logs and unpacked data for SetGasLimit events raised by the Wallet contract.
type WalletSetGasLimitIterator struct {
	Event *WalletSetGasLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetGasLimit)
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
		it.Event = new(WalletSetGasLimit)
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
func (it *WalletSetGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetGasLimit represents a SetGasLimit event raised by the Wallet contract.
type WalletSetGasLimit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetGasLimit is a free log retrieval operation binding the contract event 0xf5bdeca176beddded5a1132996e4edf0d16be5100214b17d8bada54bf8676297.
//
// Solidity: event SetGasLimit(_amount uint256)
func (_Wallet *WalletFilterer) FilterSetGasLimit(opts *bind.FilterOpts) (*WalletSetGasLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetGasLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetGasLimitIterator{contract: _Wallet.contract, event: "SetGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetGasLimit is a free log subscription operation binding the contract event 0xf5bdeca176beddded5a1132996e4edf0d16be5100214b17d8bada54bf8676297.
//
// Solidity: event SetGasLimit(_amount uint256)
func (_Wallet *WalletFilterer) WatchSetGasLimit(opts *bind.WatchOpts, sink chan<- *WalletSetGasLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetGasLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetGasLimit", log); err != nil {
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

// WalletTopUpGasIterator is returned from FilterTopUpGas and is used to iterate over the raw logs and unpacked data for TopUpGas events raised by the Wallet contract.
type WalletTopUpGasIterator struct {
	Event *WalletTopUpGas // Event containing the contract specifics and raw log

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
func (it *WalletTopUpGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTopUpGas)
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
		it.Event = new(WalletTopUpGas)
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
func (it *WalletTopUpGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTopUpGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTopUpGas represents a TopUpGas event raised by the Wallet contract.
type WalletTopUpGas struct {
	Sender common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopUpGas is a free log retrieval operation binding the contract event 0x2235de9f3363e464311d990c51aeef966703c87d1c77e80737831d6944d87c86.
//
// Solidity: event TopUpGas(_sender address, _owner address, _amount uint256)
func (_Wallet *WalletFilterer) FilterTopUpGas(opts *bind.FilterOpts) (*WalletTopUpGasIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "TopUpGas")
	if err != nil {
		return nil, err
	}
	return &WalletTopUpGasIterator{contract: _Wallet.contract, event: "TopUpGas", logs: logs, sub: sub}, nil
}

// WatchTopUpGas is a free log subscription operation binding the contract event 0x2235de9f3363e464311d990c51aeef966703c87d1c77e80737831d6944d87c86.
//
// Solidity: event TopUpGas(_sender address, _owner address, _amount uint256)
func (_Wallet *WalletFilterer) WatchTopUpGas(opts *bind.WatchOpts, sink chan<- *WalletTopUpGas) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "TopUpGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTopUpGas)
				if err := _Wallet.contract.UnpackLog(event, "TopUpGas", log); err != nil {
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

// WalletTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wallet contract.
type WalletTransferIterator struct {
	Event *WalletTransfer // Event containing the contract specifics and raw log

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
func (it *WalletTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTransfer)
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
		it.Event = new(WalletTransfer)
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
func (it *WalletTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTransfer represents a Transfer event raised by the Wallet contract.
type WalletTransfer struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_to address, _asset address, _amount uint256)
func (_Wallet *WalletFilterer) FilterTransfer(opts *bind.FilterOpts) (*WalletTransferIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &WalletTransferIterator{contract: _Wallet.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_to address, _asset address, _amount uint256)
func (_Wallet *WalletFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WalletTransfer) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTransfer)
				if err := _Wallet.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// WalletWhitelistAdditionIterator is returned from FilterWhitelistAddition and is used to iterate over the raw logs and unpacked data for WhitelistAddition events raised by the Wallet contract.
type WalletWhitelistAdditionIterator struct {
	Event *WalletWhitelistAddition // Event containing the contract specifics and raw log

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
func (it *WalletWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletWhitelistAddition)
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
		it.Event = new(WalletWhitelistAddition)
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
func (it *WalletWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletWhitelistAddition represents a WhitelistAddition event raised by the Wallet contract.
type WalletWhitelistAddition struct {
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAddition is a free log retrieval operation binding the contract event 0xc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434.
//
// Solidity: event WhitelistAddition(_addresses address[])
func (_Wallet *WalletFilterer) FilterWhitelistAddition(opts *bind.FilterOpts) (*WalletWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "WhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletWhitelistAdditionIterator{contract: _Wallet.contract, event: "WhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchWhitelistAddition is a free log subscription operation binding the contract event 0xc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434.
//
// Solidity: event WhitelistAddition(_addresses address[])
func (_Wallet *WalletFilterer) WatchWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "WhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletWhitelistAddition)
				if err := _Wallet.contract.UnpackLog(event, "WhitelistAddition", log); err != nil {
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

// WalletWhitelistRemovalIterator is returned from FilterWhitelistRemoval and is used to iterate over the raw logs and unpacked data for WhitelistRemoval events raised by the Wallet contract.
type WalletWhitelistRemovalIterator struct {
	Event *WalletWhitelistRemoval // Event containing the contract specifics and raw log

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
func (it *WalletWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletWhitelistRemoval)
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
		it.Event = new(WalletWhitelistRemoval)
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
func (it *WalletWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletWhitelistRemoval represents a WhitelistRemoval event raised by the Wallet contract.
type WalletWhitelistRemoval struct {
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoval is a free log retrieval operation binding the contract event 0x4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94.
//
// Solidity: event WhitelistRemoval(_addresses address[])
func (_Wallet *WalletFilterer) FilterWhitelistRemoval(opts *bind.FilterOpts) (*WalletWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "WhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletWhitelistRemovalIterator{contract: _Wallet.contract, event: "WhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoval is a free log subscription operation binding the contract event 0x4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94.
//
// Solidity: event WhitelistRemoval(_addresses address[])
func (_Wallet *WalletFilterer) WatchWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "WhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletWhitelistRemoval)
				if err := _Wallet.contract.UnpackLog(event, "WhitelistRemoval", log); err != nil {
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
