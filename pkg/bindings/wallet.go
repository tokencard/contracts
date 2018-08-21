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
const WalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"pendingTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingRemoval\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setTopupLimitCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addToWhitelistConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"topupAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"topupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addToWhitelistCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeFromWhitelistCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAddition\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeFromWhitelistConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topupGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"initializeWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setTopupLimitConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setSpendLimitConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setSpendLimitCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_controllers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetTopupLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TopupGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"AddController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"RemoveController\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `0x60806040523480156200001157600080fd5b5060405162002a9038038062002a90833981018060405281019080805190602001909291908051906020019092919080518201929190505050828282600068056bc75e2d631000006007819055504260088190555060075460098190555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600b60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600090505b8151811015620002765760008083838151811015156200011257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156200026857600160008084848151811015156200017f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055507fc187ac4977e802e627f1d633bbe3561806983fdbbde5904319f329bc0516365a82828151811015156200021c57fe5b90602001906020020151604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b8080600101915050620000f6565b505050506706f05b59d3b20000600c81905550600c54600e819055505050506127eb80620002a56000396000f3006080604052600436106101a1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630978f5a01461021857806315b9a8b8146102435780632ea72b521461026e578063333ebf8a146102da5780633af32abf146102f15780633c672eb71461034c5780634c58ab1d14610379578063548db1741461039057806358453569146103f65780636fc95b89146104235780637dc0d1d01461044e5780637f649783146104a55780638da5cb5b1461050b5780638ff3bf33146105625780639fed03861461058f578063a7fc7a07146105bc578063aceaf92d146105ff578063b429afeb1461062a578063beabacc814610685578063c8ecaddb146106f2578063d87d2a391461071d578063d9e98d4514610748578063da97cb821461075f578063dae37fac14610776578063dd0e9170146107a1578063e30baedf1461080d578063e3d670d714610824578063eca717921461087b578063f4199bb8146108a8578063f6a74ed71461090e578063f88b828c14610951578063f9f5e8af14610968578063fb02a1171461097f575b6000341115610216577fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b005b34801561022457600080fd5b5061022d610996565b6040518082815260200191505060405180910390f35b34801561024f57600080fd5b5061025861099c565b6040518082815260200191505060405180910390f35b34801561027a57600080fd5b506102836109a2565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102c65780820151818401526020810190506102ab565b505050509050019250505060405180910390f35b3480156102e657600080fd5b506102ef610a30565b005b3480156102fd57600080fd5b50610332600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610aac565b604051808215151515815260200191505060405180910390f35b34801561035857600080fd5b5061037760048036038101908080359060200190929190505050610acc565b005b34801561038557600080fd5b5061038e610b69565b005b34801561039c57600080fd5b506103f460048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610e2b565b005b34801561040257600080fd5b5061042160048036038101908080359060200190929190505050610f6d565b005b34801561042f57600080fd5b50610438611043565b6040518082815260200191505060405180910390f35b34801561045a57600080fd5b50610463611067565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104b157600080fd5b506105096004803603810190808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061108d565b005b34801561051757600080fd5b5061052061114b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561056e57600080fd5b5061058d60048036038101908080359060200190929190505050611171565b005b34801561059b57600080fd5b506105ba60048036038101908080359060200190929190505050611271565b005b3480156105c857600080fd5b506105fd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611338565b005b34801561060b57600080fd5b506106146114b6565b6040518082815260200191505060405180910390f35b34801561063657600080fd5b5061066b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114bc565b604051808215151515815260200191505060405180910390f35b34801561069157600080fd5b506106f0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506114dc565b005b3480156106fe57600080fd5b506107076118ef565b6040518082815260200191505060405180910390f35b34801561072957600080fd5b506107326118f5565b6040518082815260200191505060405180910390f35b34801561075457600080fd5b5061075d6118fb565b005b34801561076b57600080fd5b5061077461197d565b005b34801561078257600080fd5b5061078b6119ff565b6040518082815260200191505060405180910390f35b3480156107ad57600080fd5b506107b6611a23565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156107f95780820151818401526020810190506107de565b505050509050019250505060405180910390f35b34801561081957600080fd5b50610822611ab1565b005b34801561083057600080fd5b50610865600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611cb0565b6040518082815260200191505060405180910390f35b34801561088757600080fd5b506108a660048036038101908080359060200190929190505050611dce565b005b3480156108b457600080fd5b5061090c60048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611ff5565b005b34801561091a57600080fd5b5061094f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612190565b005b34801561095d57600080fd5b5061096661231b565b005b34801561097457600080fd5b5061097d61244b565b005b34801561098b57600080fd5b50610994612550565b005b600f5481565b60015481565b60606005805480602002602001604051908101604052809291908181526020018280548015610a2657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116109dc575b5050505050905090565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610a8757600080fd5b6000600f819055506000601060006101000a81548160ff021916908315150217905550565b60036020528060005260406000206000915054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b2857600080fd5b600b60009054906101000a900460ff16151515610b4457600080fd5b80600a819055506001600b60006101000a81548160ff02191690831515021790555050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610bc257600080fd5b6000600480549050118015610be35750600660009054906101000a900460ff165b1515610bee57600080fd5b600090505b600480549050811015610d2f57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600482815481101515610c4857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610d2257600160036000600484815481101515610ca557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8080600101915050610bf3565b6000600660006101000a81548160ff021916908315150217905550600660029054906101000a900460ff161515610d7c576001600660026101000a81548160ff0219169083151502179055505b7fc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434600460405180806020018281038252838181548152602001915080548015610e1a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610dd0575b50509250505060405180910390a150565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e8957600080fd5b600660019054906101000a900460ff16151515610ea557600080fd5b6014825111151515610eb657600080fd5b600090505b8151811015610f4e5760058282815181101515610ed457fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508080600101915050610ebb565b6001600660016101000a81548160ff0219169083151502179055505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fc957600080fd5b600b60019054906101000a900460ff16151515610fe557600080fd5b610fee816125cc565b6001600b60016101000a81548160ff0219169083151502179055507f21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253816040518082815260200191505060405180910390a150565b600062015180600d540142111561105e57600c549050611064565b600e5490505b90565b600b60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110e957600080fd5b600660009054906101000a900460ff1615151561110557600080fd5b601481511115151561111657600080fd5b806004908051906020019061112c9291906126ac565b506001600660006101000a81548160ff02191690831515021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156111cd57600080fd5b601060019054906101000a900460ff161515156111e957600080fd5b8066038d7ea4c680001115801561120857506706f05b59d3b200008111155b151561121357600080fd5b61121c816125f4565b6001601060016101000a81548160ff0219169083151502179055507f19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0816040518082815260200191505060405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112cd57600080fd5b601060009054906101000a900460ff161515156112e957600080fd5b8066038d7ea4c680001115801561130857506706f05b59d3b200008111155b151561131357600080fd5b80600f819055506001601060006101000a81548160ff02191690831515021790555050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561138f57600080fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515156113e757600080fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055507fc187ac4977e802e627f1d633bbe3561806983fdbbde5904319f329bc0516365a81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60075481565b60006020528060005260406000206000915054906101000a900460ff1681565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561153a57600080fd5b816000811415151561154b57600080fd5b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156116f3576115a661261c565b60008473ffffffffffffffffffffffffffffffffffffffff161415156116cd57600b60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166367c6e39c85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561168b57600080fd5b505af115801561169f573d6000803e3d6000fd5b505050506040513d60208110156116b557600080fd5b810190808051906020019092919050505091506116d1565b8291505b60095482111515156116e257600080fd5b816009600082825403925050819055505b60008473ffffffffffffffffffffffffffffffffffffffff16141515611801578373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156117b657600080fd5b505af11580156117ca573d6000803e3d6000fd5b505050506040513d60208110156117e057600080fd5b810190808051906020019092919050505015156117fc57600080fd5b611849565b8473ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015611847573d6000803e3d6000fd5b505b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef858585604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050505050565b600a5481565b600c5481565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561195257600080fd5b600460006119609190612736565b6000600660006101000a81548160ff021916908315150217905550565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156119d457600080fd5b600560006119e29190612736565b6000600660016101000a81548160ff021916908315150217905550565b60006201518060085401421115611a1a576007549050611a20565b60095490505b90565b60606004805480602002602001604051908101604052809291908181526020018280548015611aa757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611a5d575b5050505050905090565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611b0a57600080fd5b6000600580549050118015611b2b5750600660019054906101000a900460ff165b1515611b3657600080fd5b600090505b600580549050811015611be657600060036000600584815481101515611b5d57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050611b3b565b6000600660016101000a81548160ff0219169083151502179055507f4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94600560405180806020018281038252838181548152602001915080548015611c9f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611c55575b50509250505060405180910390a150565b6000808273ffffffffffffffffffffffffffffffffffffffff16141515611dae578173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611d6c57600080fd5b505af1158015611d80573d6000803e3d6000fd5b505050506040513d6020811015611d9657600080fd5b81019080805190602001909291905050509050611dc9565b3073ffffffffffffffffffffffffffffffffffffffff163190505b919050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680611e725750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611e7d57600080fd5b8060008114151515611e8e57600080fd5b611e96612664565b6000600e54111515611ea757600080fd5b600e54821115611eb757600e5491505b81600e60008282540392505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015611f2f573d6000803e3d6000fd5b507f11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c674232600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561205357600080fd5b600660029054906101000a900460ff1615151561206f57600080fd5b600090505b81518110156120f957600160036000848481518110151561209157fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050612074565b6001600660026101000a81548160ff0219169083151502179055507fc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434826040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561217957808201518184015260208101905061215e565b505050509050019250505060405180910390a15050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156121e757600080fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168015612240575060018054115b151561224b57600080fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055507f84c8c1572fb6a01167de9135b17f8c2dd31ef2d755fa7d2b2d8f2b05c369868881604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561237257600080fd5b601060009054906101000a900460ff16151561238d57600080fd5b600f5466038d7ea4c68000111580156123b057506706f05b59d3b20000600f5411155b15156123b857fe5b6123c3600f546125f4565b6000601060006101000a81548160ff021916908315150217905550601060019054906101000a900460ff161515612410576001601060016101000a81548160ff0219169083151502179055505b7f19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0600f546040518082815260200191505060405180910390a1565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156124a257600080fd5b600b60009054906101000a900460ff1615156124bd57600080fd5b6124c8600a546125cc565b6000600b60006101000a81548160ff021916908315150217905550600b60019054906101000a900460ff161515612515576001600b60016101000a81548160ff0219169083151502179055505b7f21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253600a546040518082815260200191505060405180910390a1565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156125a757600080fd5b6000600a819055506000600b60006101000a81548160ff021916908315150217905550565b6125d461261c565b8060078190555060075460095411156125f1576007546009819055505b50565b6125fc612664565b80600c81905550600c54600e54111561261957600c54600e819055505b50565b600062015180600854014211156126615762015180600854420381151561263f57fe5b0490506201518081026008600082825401925050819055506007546009819055505b50565b600062015180600d54014211156126a95762015180600d54420381151561268757fe5b049050620151808102600d60008282540192505081905550600c54600e819055505b50565b828054828255906000526020600020908101928215612725579160200282015b828111156127245782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906126cc565b5b5090506127329190612757565b5090565b5080546000825590600052602060002090810190612754919061279a565b50565b61279791905b8082111561279357600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010161275d565b5090565b90565b6127bc91905b808211156127b85760008160009055506001016127a0565b5090565b905600a165627a7a7230582094b38280494988ddb74fe3535492729a59c28944e664c2fd2b09f25603ff5b030029`

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

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Wallet *WalletCaller) ControllerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "controllerCount")
	return *ret0, err
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Wallet *WalletSession) ControllerCount() (*big.Int, error) {
	return _Wallet.Contract.ControllerCount(&_Wallet.CallOpts)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Wallet *WalletCallerSession) ControllerCount() (*big.Int, error) {
	return _Wallet.Contract.ControllerCount(&_Wallet.CallOpts)
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

// PendingSpendLimit is a free data retrieval call binding the contract method 0xc8ecaddb.
//
// Solidity: function pendingSpendLimit() constant returns(uint256)
func (_Wallet *WalletCaller) PendingSpendLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingSpendLimit")
	return *ret0, err
}

// PendingSpendLimit is a free data retrieval call binding the contract method 0xc8ecaddb.
//
// Solidity: function pendingSpendLimit() constant returns(uint256)
func (_Wallet *WalletSession) PendingSpendLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingSpendLimit(&_Wallet.CallOpts)
}

// PendingSpendLimit is a free data retrieval call binding the contract method 0xc8ecaddb.
//
// Solidity: function pendingSpendLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) PendingSpendLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingSpendLimit(&_Wallet.CallOpts)
}

// PendingTopupLimit is a free data retrieval call binding the contract method 0x0978f5a0.
//
// Solidity: function pendingTopupLimit() constant returns(uint256)
func (_Wallet *WalletCaller) PendingTopupLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingTopupLimit")
	return *ret0, err
}

// PendingTopupLimit is a free data retrieval call binding the contract method 0x0978f5a0.
//
// Solidity: function pendingTopupLimit() constant returns(uint256)
func (_Wallet *WalletSession) PendingTopupLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingTopupLimit(&_Wallet.CallOpts)
}

// PendingTopupLimit is a free data retrieval call binding the contract method 0x0978f5a0.
//
// Solidity: function pendingTopupLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) PendingTopupLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingTopupLimit(&_Wallet.CallOpts)
}

// SpendAvailable is a free data retrieval call binding the contract method 0xdae37fac.
//
// Solidity: function spendAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) SpendAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendAvailable")
	return *ret0, err
}

// SpendAvailable is a free data retrieval call binding the contract method 0xdae37fac.
//
// Solidity: function spendAvailable() constant returns(uint256)
func (_Wallet *WalletSession) SpendAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendAvailable(&_Wallet.CallOpts)
}

// SpendAvailable is a free data retrieval call binding the contract method 0xdae37fac.
//
// Solidity: function spendAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendAvailable(&_Wallet.CallOpts)
}

// SpendLimit is a free data retrieval call binding the contract method 0xaceaf92d.
//
// Solidity: function spendLimit() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimit")
	return *ret0, err
}

// SpendLimit is a free data retrieval call binding the contract method 0xaceaf92d.
//
// Solidity: function spendLimit() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimit() (*big.Int, error) {
	return _Wallet.Contract.SpendLimit(&_Wallet.CallOpts)
}

// SpendLimit is a free data retrieval call binding the contract method 0xaceaf92d.
//
// Solidity: function spendLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimit() (*big.Int, error) {
	return _Wallet.Contract.SpendLimit(&_Wallet.CallOpts)
}

// TopupAvailable is a free data retrieval call binding the contract method 0x6fc95b89.
//
// Solidity: function topupAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) TopupAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "topupAvailable")
	return *ret0, err
}

// TopupAvailable is a free data retrieval call binding the contract method 0x6fc95b89.
//
// Solidity: function topupAvailable() constant returns(uint256)
func (_Wallet *WalletSession) TopupAvailable() (*big.Int, error) {
	return _Wallet.Contract.TopupAvailable(&_Wallet.CallOpts)
}

// TopupAvailable is a free data retrieval call binding the contract method 0x6fc95b89.
//
// Solidity: function topupAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) TopupAvailable() (*big.Int, error) {
	return _Wallet.Contract.TopupAvailable(&_Wallet.CallOpts)
}

// TopupLimit is a free data retrieval call binding the contract method 0xd87d2a39.
//
// Solidity: function topupLimit() constant returns(uint256)
func (_Wallet *WalletCaller) TopupLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "topupLimit")
	return *ret0, err
}

// TopupLimit is a free data retrieval call binding the contract method 0xd87d2a39.
//
// Solidity: function topupLimit() constant returns(uint256)
func (_Wallet *WalletSession) TopupLimit() (*big.Int, error) {
	return _Wallet.Contract.TopupLimit(&_Wallet.CallOpts)
}

// TopupLimit is a free data retrieval call binding the contract method 0xd87d2a39.
//
// Solidity: function topupLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) TopupLimit() (*big.Int, error) {
	return _Wallet.Contract.TopupLimit(&_Wallet.CallOpts)
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

// InitializeSpendLimit is a paid mutator transaction binding the contract method 0x58453569.
//
// Solidity: function initializeSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) InitializeSpendLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeSpendLimit", _amount)
}

// InitializeSpendLimit is a paid mutator transaction binding the contract method 0x58453569.
//
// Solidity: function initializeSpendLimit(_amount uint256) returns()
func (_Wallet *WalletSession) InitializeSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeSpendLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeSpendLimit is a paid mutator transaction binding the contract method 0x58453569.
//
// Solidity: function initializeSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) InitializeSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeSpendLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeTopupLimit is a paid mutator transaction binding the contract method 0x8ff3bf33.
//
// Solidity: function initializeTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) InitializeTopupLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeTopupLimit", _amount)
}

// InitializeTopupLimit is a paid mutator transaction binding the contract method 0x8ff3bf33.
//
// Solidity: function initializeTopupLimit(_amount uint256) returns()
func (_Wallet *WalletSession) InitializeTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeTopupLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeTopupLimit is a paid mutator transaction binding the contract method 0x8ff3bf33.
//
// Solidity: function initializeTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) InitializeTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeTopupLimit(&_Wallet.TransactOpts, _amount)
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

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) SetSpendLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setSpendLimit", _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(_amount uint256) returns()
func (_Wallet *WalletSession) SetSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimit(&_Wallet.TransactOpts, _amount)
}

// SetSpendLimit is a paid mutator transaction binding the contract method 0x3c672eb7.
//
// Solidity: function setSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) SetSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimit(&_Wallet.TransactOpts, _amount)
}

// SetSpendLimitCancel is a paid mutator transaction binding the contract method 0xfb02a117.
//
// Solidity: function setSpendLimitCancel() returns()
func (_Wallet *WalletTransactor) SetSpendLimitCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setSpendLimitCancel")
}

// SetSpendLimitCancel is a paid mutator transaction binding the contract method 0xfb02a117.
//
// Solidity: function setSpendLimitCancel() returns()
func (_Wallet *WalletSession) SetSpendLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimitCancel(&_Wallet.TransactOpts)
}

// SetSpendLimitCancel is a paid mutator transaction binding the contract method 0xfb02a117.
//
// Solidity: function setSpendLimitCancel() returns()
func (_Wallet *WalletTransactorSession) SetSpendLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimitCancel(&_Wallet.TransactOpts)
}

// SetSpendLimitConfirm is a paid mutator transaction binding the contract method 0xf9f5e8af.
//
// Solidity: function setSpendLimitConfirm() returns()
func (_Wallet *WalletTransactor) SetSpendLimitConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setSpendLimitConfirm")
}

// SetSpendLimitConfirm is a paid mutator transaction binding the contract method 0xf9f5e8af.
//
// Solidity: function setSpendLimitConfirm() returns()
func (_Wallet *WalletSession) SetSpendLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimitConfirm(&_Wallet.TransactOpts)
}

// SetSpendLimitConfirm is a paid mutator transaction binding the contract method 0xf9f5e8af.
//
// Solidity: function setSpendLimitConfirm() returns()
func (_Wallet *WalletTransactorSession) SetSpendLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetSpendLimitConfirm(&_Wallet.TransactOpts)
}

// SetTopupLimit is a paid mutator transaction binding the contract method 0x9fed0386.
//
// Solidity: function setTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) SetTopupLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setTopupLimit", _amount)
}

// SetTopupLimit is a paid mutator transaction binding the contract method 0x9fed0386.
//
// Solidity: function setTopupLimit(_amount uint256) returns()
func (_Wallet *WalletSession) SetTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetTopupLimit(&_Wallet.TransactOpts, _amount)
}

// SetTopupLimit is a paid mutator transaction binding the contract method 0x9fed0386.
//
// Solidity: function setTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) SetTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SetTopupLimit(&_Wallet.TransactOpts, _amount)
}

// SetTopupLimitCancel is a paid mutator transaction binding the contract method 0x333ebf8a.
//
// Solidity: function setTopupLimitCancel() returns()
func (_Wallet *WalletTransactor) SetTopupLimitCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setTopupLimitCancel")
}

// SetTopupLimitCancel is a paid mutator transaction binding the contract method 0x333ebf8a.
//
// Solidity: function setTopupLimitCancel() returns()
func (_Wallet *WalletSession) SetTopupLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetTopupLimitCancel(&_Wallet.TransactOpts)
}

// SetTopupLimitCancel is a paid mutator transaction binding the contract method 0x333ebf8a.
//
// Solidity: function setTopupLimitCancel() returns()
func (_Wallet *WalletTransactorSession) SetTopupLimitCancel() (*types.Transaction, error) {
	return _Wallet.Contract.SetTopupLimitCancel(&_Wallet.TransactOpts)
}

// SetTopupLimitConfirm is a paid mutator transaction binding the contract method 0xf88b828c.
//
// Solidity: function setTopupLimitConfirm() returns()
func (_Wallet *WalletTransactor) SetTopupLimitConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setTopupLimitConfirm")
}

// SetTopupLimitConfirm is a paid mutator transaction binding the contract method 0xf88b828c.
//
// Solidity: function setTopupLimitConfirm() returns()
func (_Wallet *WalletSession) SetTopupLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetTopupLimitConfirm(&_Wallet.TransactOpts)
}

// SetTopupLimitConfirm is a paid mutator transaction binding the contract method 0xf88b828c.
//
// Solidity: function setTopupLimitConfirm() returns()
func (_Wallet *WalletTransactorSession) SetTopupLimitConfirm() (*types.Transaction, error) {
	return _Wallet.Contract.SetTopupLimitConfirm(&_Wallet.TransactOpts)
}

// TopupGas is a paid mutator transaction binding the contract method 0xeca71792.
//
// Solidity: function topupGas(_amount uint256) returns()
func (_Wallet *WalletTransactor) TopupGas(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "topupGas", _amount)
}

// TopupGas is a paid mutator transaction binding the contract method 0xeca71792.
//
// Solidity: function topupGas(_amount uint256) returns()
func (_Wallet *WalletSession) TopupGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopupGas(&_Wallet.TransactOpts, _amount)
}

// TopupGas is a paid mutator transaction binding the contract method 0xeca71792.
//
// Solidity: function topupGas(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) TopupGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopupGas(&_Wallet.TransactOpts, _amount)
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

// WalletAddControllerIterator is returned from FilterAddController and is used to iterate over the raw logs and unpacked data for AddController events raised by the Wallet contract.
type WalletAddControllerIterator struct {
	Event *WalletAddController // Event containing the contract specifics and raw log

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
func (it *WalletAddControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletAddController)
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
		it.Event = new(WalletAddController)
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
func (it *WalletAddControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletAddControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletAddController represents a AddController event raised by the Wallet contract.
type WalletAddController struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddController is a free log retrieval operation binding the contract event 0xc187ac4977e802e627f1d633bbe3561806983fdbbde5904319f329bc0516365a.
//
// Solidity: e AddController(_account address)
func (_Wallet *WalletFilterer) FilterAddController(opts *bind.FilterOpts) (*WalletAddControllerIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "AddController")
	if err != nil {
		return nil, err
	}
	return &WalletAddControllerIterator{contract: _Wallet.contract, event: "AddController", logs: logs, sub: sub}, nil
}

// WatchAddController is a free log subscription operation binding the contract event 0xc187ac4977e802e627f1d633bbe3561806983fdbbde5904319f329bc0516365a.
//
// Solidity: e AddController(_account address)
func (_Wallet *WalletFilterer) WatchAddController(opts *bind.WatchOpts, sink chan<- *WalletAddController) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "AddController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletAddController)
				if err := _Wallet.contract.UnpackLog(event, "AddController", log); err != nil {
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
// Solidity: e Deposit(_from address, _amount uint256)
func (_Wallet *WalletFilterer) FilterDeposit(opts *bind.FilterOpts) (*WalletDepositIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &WalletDepositIterator{contract: _Wallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: e Deposit(_from address, _amount uint256)
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

// WalletRemoveControllerIterator is returned from FilterRemoveController and is used to iterate over the raw logs and unpacked data for RemoveController events raised by the Wallet contract.
type WalletRemoveControllerIterator struct {
	Event *WalletRemoveController // Event containing the contract specifics and raw log

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
func (it *WalletRemoveControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletRemoveController)
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
		it.Event = new(WalletRemoveController)
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
func (it *WalletRemoveControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletRemoveControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletRemoveController represents a RemoveController event raised by the Wallet contract.
type WalletRemoveController struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveController is a free log retrieval operation binding the contract event 0x84c8c1572fb6a01167de9135b17f8c2dd31ef2d755fa7d2b2d8f2b05c3698688.
//
// Solidity: e RemoveController(_account address)
func (_Wallet *WalletFilterer) FilterRemoveController(opts *bind.FilterOpts) (*WalletRemoveControllerIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "RemoveController")
	if err != nil {
		return nil, err
	}
	return &WalletRemoveControllerIterator{contract: _Wallet.contract, event: "RemoveController", logs: logs, sub: sub}, nil
}

// WatchRemoveController is a free log subscription operation binding the contract event 0x84c8c1572fb6a01167de9135b17f8c2dd31ef2d755fa7d2b2d8f2b05c3698688.
//
// Solidity: e RemoveController(_account address)
func (_Wallet *WalletFilterer) WatchRemoveController(opts *bind.WatchOpts, sink chan<- *WalletRemoveController) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "RemoveController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletRemoveController)
				if err := _Wallet.contract.UnpackLog(event, "RemoveController", log); err != nil {
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

// WalletSetSpendLimitIterator is returned from FilterSetSpendLimit and is used to iterate over the raw logs and unpacked data for SetSpendLimit events raised by the Wallet contract.
type WalletSetSpendLimitIterator struct {
	Event *WalletSetSpendLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetSpendLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetSpendLimit)
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
		it.Event = new(WalletSetSpendLimit)
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
func (it *WalletSetSpendLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetSpendLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetSpendLimit represents a SetSpendLimit event raised by the Wallet contract.
type WalletSetSpendLimit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSpendLimit is a free log retrieval operation binding the contract event 0x21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253.
//
// Solidity: e SetSpendLimit(_amount uint256)
func (_Wallet *WalletFilterer) FilterSetSpendLimit(opts *bind.FilterOpts) (*WalletSetSpendLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetSpendLimitIterator{contract: _Wallet.contract, event: "SetSpendLimit", logs: logs, sub: sub}, nil
}

// WatchSetSpendLimit is a free log subscription operation binding the contract event 0x21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253.
//
// Solidity: e SetSpendLimit(_amount uint256)
func (_Wallet *WalletFilterer) WatchSetSpendLimit(opts *bind.WatchOpts, sink chan<- *WalletSetSpendLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetSpendLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetSpendLimit", log); err != nil {
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

// WalletSetTopupLimitIterator is returned from FilterSetTopupLimit and is used to iterate over the raw logs and unpacked data for SetTopupLimit events raised by the Wallet contract.
type WalletSetTopupLimitIterator struct {
	Event *WalletSetTopupLimit // Event containing the contract specifics and raw log

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
func (it *WalletSetTopupLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetTopupLimit)
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
		it.Event = new(WalletSetTopupLimit)
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
func (it *WalletSetTopupLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetTopupLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetTopupLimit represents a SetTopupLimit event raised by the Wallet contract.
type WalletSetTopupLimit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTopupLimit is a free log retrieval operation binding the contract event 0x19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0.
//
// Solidity: e SetTopupLimit(_amount uint256)
func (_Wallet *WalletFilterer) FilterSetTopupLimit(opts *bind.FilterOpts) (*WalletSetTopupLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetTopupLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetTopupLimitIterator{contract: _Wallet.contract, event: "SetTopupLimit", logs: logs, sub: sub}, nil
}

// WatchSetTopupLimit is a free log subscription operation binding the contract event 0x19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0.
//
// Solidity: e SetTopupLimit(_amount uint256)
func (_Wallet *WalletFilterer) WatchSetTopupLimit(opts *bind.WatchOpts, sink chan<- *WalletSetTopupLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetTopupLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetTopupLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetTopupLimit", log); err != nil {
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

// WalletTopupGasIterator is returned from FilterTopupGas and is used to iterate over the raw logs and unpacked data for TopupGas events raised by the Wallet contract.
type WalletTopupGasIterator struct {
	Event *WalletTopupGas // Event containing the contract specifics and raw log

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
func (it *WalletTopupGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTopupGas)
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
		it.Event = new(WalletTopupGas)
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
func (it *WalletTopupGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTopupGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTopupGas represents a TopupGas event raised by the Wallet contract.
type WalletTopupGas struct {
	Sender common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopupGas is a free log retrieval operation binding the contract event 0x11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c6742.
//
// Solidity: e TopupGas(_sender address, _owner address, _amount uint256)
func (_Wallet *WalletFilterer) FilterTopupGas(opts *bind.FilterOpts) (*WalletTopupGasIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "TopupGas")
	if err != nil {
		return nil, err
	}
	return &WalletTopupGasIterator{contract: _Wallet.contract, event: "TopupGas", logs: logs, sub: sub}, nil
}

// WatchTopupGas is a free log subscription operation binding the contract event 0x11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c6742.
//
// Solidity: e TopupGas(_sender address, _owner address, _amount uint256)
func (_Wallet *WalletFilterer) WatchTopupGas(opts *bind.WatchOpts, sink chan<- *WalletTopupGas) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "TopupGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTopupGas)
				if err := _Wallet.contract.UnpackLog(event, "TopupGas", log); err != nil {
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
// Solidity: e Transfer(_to address, _asset address, _amount uint256)
func (_Wallet *WalletFilterer) FilterTransfer(opts *bind.FilterOpts) (*WalletTransferIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &WalletTransferIterator{contract: _Wallet.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_to address, _asset address, _amount uint256)
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
// Solidity: e WhitelistAddition(_addresses address[])
func (_Wallet *WalletFilterer) FilterWhitelistAddition(opts *bind.FilterOpts) (*WalletWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "WhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletWhitelistAdditionIterator{contract: _Wallet.contract, event: "WhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchWhitelistAddition is a free log subscription operation binding the contract event 0xc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434.
//
// Solidity: e WhitelistAddition(_addresses address[])
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
// Solidity: e WhitelistRemoval(_addresses address[])
func (_Wallet *WalletFilterer) FilterWhitelistRemoval(opts *bind.FilterOpts) (*WalletWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "WhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletWhitelistRemovalIterator{contract: _Wallet.contract, event: "WhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoval is a free log subscription operation binding the contract event 0x4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94.
//
// Solidity: e WhitelistRemoval(_addresses address[])
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
