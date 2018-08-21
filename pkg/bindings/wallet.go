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
const WalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"pendingTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingRemoval\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setTopupLimitCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addToWhitelistConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initializedSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"topupAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initializedWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initializedTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedAddition\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"topupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addToWhitelistCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeFromWhitelistCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAddition\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeFromWhitelistConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topupGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"initializeWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setTopupLimitConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedRemoval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setSpendLimitConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setSpendLimitCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_controllers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetTopupLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TopupGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"AddController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"RemoveController\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `0x60806040523480156200001157600080fd5b5060405162002cc238038062002cc2833981018060405281019080805190602001909291908051906020019092919080518201929190505050828282600068056bc75e2d631000006007819055504260088190555060075460098190555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600b60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600090505b8151811015620002765760008083838151811015156200011257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156200026857600160008084848151811015156200017f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055507fc187ac4977e802e627f1d633bbe3561806983fdbbde5904319f329bc0516365a82828151811015156200021c57fe5b90602001906020020151604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b8080600101915050620000f6565b505050506706f05b59d3b20000600c81905550600c54600e81905550505050612a1d80620002a56000396000f3006080604052600436106101ee576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630978f5a01461026557806315b9a8b8146102905780632ea72b52146102bb578063333ebf8a146103275780633671bb0d1461033e5780633af32abf1461036d5780633c672eb7146103c85780634c58ab1d146103f5578063548db1741461040c5780635658eff01461047257806358453569146104a15780636fc95b89146104ce5780637dc0d1d0146104f95780637f649783146105505780638da5cb5b146105b65780638ff3bf331461060d5780639fed03861461063a578063a7fc7a0714610667578063aceaf92d146106aa578063afa0fd9b146106d5578063b429afeb14610704578063beabacc81461075f578063c8ecaddb146107cc578063cd96895c146107f7578063d2f9a63914610826578063d566659014610855578063d87d2a3914610884578063d9e98d45146108af578063da97cb82146108c6578063dae37fac146108dd578063dd0e917014610908578063e30baedf14610974578063e3d670d71461098b578063eca71792146109e2578063f4199bb814610a0f578063f6a74ed714610a75578063f88b828c14610ab8578063f9c26de414610acf578063f9f5e8af14610afe578063fb02a11714610b15575b6000341115610263577fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b005b34801561027157600080fd5b5061027a610b2c565b6040518082815260200191505060405180910390f35b34801561029c57600080fd5b506102a5610b32565b6040518082815260200191505060405180910390f35b3480156102c757600080fd5b506102d0610b38565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156103135780820151818401526020810190506102f8565b505050509050019250505060405180910390f35b34801561033357600080fd5b5061033c610bc6565b005b34801561034a57600080fd5b50610353610c42565b604051808215151515815260200191505060405180910390f35b34801561037957600080fd5b506103ae600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c55565b604051808215151515815260200191505060405180910390f35b3480156103d457600080fd5b506103f360048036038101908080359060200190929190505050610c75565b005b34801561040157600080fd5b5061040a610d44565b005b34801561041857600080fd5b5061047060048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610fd4565b005b34801561047e57600080fd5b50610487611118565b604051808215151515815260200191505060405180910390f35b3480156104ad57600080fd5b506104cc6004803603810190808035906020019092919050505061112b565b005b3480156104da57600080fd5b506104e3611201565b6040518082815260200191505060405180910390f35b34801561050557600080fd5b5061050e611225565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561055c57600080fd5b506105b46004803603810190808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061124b565b005b3480156105c257600080fd5b506105cb61133d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561061957600080fd5b5061063860048036038101908080359060200190929190505050611363565b005b34801561064657600080fd5b5061066560048036038101908080359060200190929190505050611463565b005b34801561067357600080fd5b506106a8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061155c565b005b3480156106b657600080fd5b506106bf6116da565b6040518082815260200191505060405180910390f35b3480156106e157600080fd5b506106ea6116e0565b604051808215151515815260200191505060405180910390f35b34801561071057600080fd5b50610745600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116f3565b604051808215151515815260200191505060405180910390f35b34801561076b57600080fd5b506107ca600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611713565b005b3480156107d857600080fd5b506107e1611b26565b6040518082815260200191505060405180910390f35b34801561080357600080fd5b5061080c611b2c565b604051808215151515815260200191505060405180910390f35b34801561083257600080fd5b5061083b611b3f565b604051808215151515815260200191505060405180910390f35b34801561086157600080fd5b5061086a611b52565b604051808215151515815260200191505060405180910390f35b34801561089057600080fd5b50610899611b65565b6040518082815260200191505060405180910390f35b3480156108bb57600080fd5b506108c4611b6b565b005b3480156108d257600080fd5b506108db611bed565b005b3480156108e957600080fd5b506108f2611c6f565b6040518082815260200191505060405180910390f35b34801561091457600080fd5b5061091d611c93565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610960578082015181840152602081019050610945565b505050509050019250505060405180910390f35b34801561098057600080fd5b50610989611d21565b005b34801561099757600080fd5b506109cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f20565b6040518082815260200191505060405180910390f35b3480156109ee57600080fd5b50610a0d6004803603810190808035906020019092919050505061203e565b005b348015610a1b57600080fd5b50610a7360048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050612265565b005b348015610a8157600080fd5b50610ab6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612413565b005b348015610ac457600080fd5b50610acd61259e565b005b348015610adb57600080fd5b50610ae461269c565b604051808215151515815260200191505060405180910390f35b348015610b0a57600080fd5b50610b136126af565b005b348015610b2157600080fd5b50610b2a612782565b005b600f5481565b60015481565b60606005805480602002602001604051908101604052809291908181526020018280548015610bbc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610b72575b5050505050905090565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610c1d57600080fd5b6000600f819055506000601060006101000a81548160ff021916908315150217905550565b601060009054906101000a900460ff1681565b60036020528060005260406000206000915054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610cd157600080fd5b600b60009054906101000a900460ff16151515610ced57600080fd5b80600a819055506001600b60006101000a81548160ff021916908315150217905550600b60019054906101000a900460ff161515610d41576001600b60016101000a81548160ff0219169083151502179055505b50565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610d9d57600080fd5b6000600480549050118015610dbe5750600660009054906101000a900460ff165b1515610dc957600080fd5b600090505b600480549050811015610f0a57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600482815481101515610e2357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610efd57600160036000600484815481101515610e8057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8080600101915050610dce565b6000600660006101000a81548160ff0219169083151502179055507fc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434600460405180806020018281038252838181548152602001915080548015610fc357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610f79575b50509250505060405180910390a150565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561103257600080fd5b81601481511115151561104457600080fd5b600660019054906101000a900460ff1615151561106057600080fd5b600091505b82518210156110f8576005838381518110151561107e57fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508180600101925050611065565b6001600660016101000a81548160ff021916908315150217905550505050565b600b60019054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561118757600080fd5b600b60019054906101000a900460ff161515156111a357600080fd5b6111ac816127fe565b6001600b60016101000a81548160ff0219169083151502179055507f21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253816040518082815260200191505060405180910390a150565b600062015180600d540142111561121c57600c549050611222565b600e5490505b90565b600b60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112a757600080fd5b8060148151111515156112b957600080fd5b600660009054906101000a900460ff161515156112d557600080fd5b81600490805190602001906112eb9291906128de565b506001600660006101000a81548160ff021916908315150217905550600660029054906101000a900460ff161515611339576001600660026101000a81548160ff0219169083151502179055505b5050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113bf57600080fd5b601060019054906101000a900460ff161515156113db57600080fd5b8066038d7ea4c68000111580156113fa57506706f05b59d3b200008111155b151561140557600080fd5b61140e81612826565b6001601060016101000a81548160ff0219169083151502179055507f19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0816040518082815260200191505060405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156114bf57600080fd5b601060009054906101000a900460ff161515156114db57600080fd5b8066038d7ea4c68000111580156114fa57506706f05b59d3b200008111155b151561150557600080fd5b80600f819055506001601060006101000a81548160ff021916908315150217905550601060019054906101000a900460ff161515611559576001601060016101000a81548160ff0219169083151502179055505b50565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156115b357600080fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151561160b57600080fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055507fc187ac4977e802e627f1d633bbe3561806983fdbbde5904319f329bc0516365a81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60075481565b600660029054906101000a900460ff1681565b60006020528060005260406000206000915054906101000a900460ff1681565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561177157600080fd5b816000811415151561178257600080fd5b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561192a576117dd61284e565b60008473ffffffffffffffffffffffffffffffffffffffff1614151561190457600b60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166367c6e39c85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156118c257600080fd5b505af11580156118d6573d6000803e3d6000fd5b505050506040513d60208110156118ec57600080fd5b81019080805190602001909291905050509150611908565b8291505b600954821115151561191957600080fd5b816009600082825403925050819055505b60008473ffffffffffffffffffffffffffffffffffffffff16141515611a38578373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156119ed57600080fd5b505af1158015611a01573d6000803e3d6000fd5b505050506040513d6020811015611a1757600080fd5b81019080805190602001909291905050501515611a3357600080fd5b611a80565b8473ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015611a7e573d6000803e3d6000fd5b505b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef858585604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050505050565b600a5481565b601060019054906101000a900460ff1681565b600660009054906101000a900460ff1681565b600b60009054906101000a900460ff1681565b600c5481565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611bc257600080fd5b60046000611bd09190612968565b6000600660006101000a81548160ff021916908315150217905550565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611c4457600080fd5b60056000611c529190612968565b6000600660016101000a81548160ff021916908315150217905550565b60006201518060085401421115611c8a576007549050611c90565b60095490505b90565b60606004805480602002602001604051908101604052809291908181526020018280548015611d1757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611ccd575b5050505050905090565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611d7a57600080fd5b6000600580549050118015611d9b5750600660019054906101000a900460ff165b1515611da657600080fd5b600090505b600580549050811015611e5657600060036000600584815481101515611dcd57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050611dab565b6000600660016101000a81548160ff0219169083151502179055507f4b089aff1cdd9a6984aa832d4a013996b3acd3d6244ce3de5e07e6ab050d2b94600560405180806020018281038252838181548152602001915080548015611f0f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611ec5575b50509250505060405180910390a150565b6000808273ffffffffffffffffffffffffffffffffffffffff1614151561201e578173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611fdc57600080fd5b505af1158015611ff0573d6000803e3d6000fd5b505050506040513d602081101561200657600080fd5b81019080805190602001909291905050509050612039565b3073ffffffffffffffffffffffffffffffffffffffff163190505b919050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806120e25750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b15156120ed57600080fd5b80600081141515156120fe57600080fd5b612106612896565b6000600e5411151561211757600080fd5b600e5482111561212757600e5491505b81600e60008282540392505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561219f573d6000803e3d6000fd5b507f11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c674232600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156122c357600080fd5b8160148151111515156122d557600080fd5b600660029054906101000a900460ff161515156122f157600080fd5b600091505b825182101561237b57600160036000858581518110151561231357fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555081806001019250506122f6565b6001600660026101000a81548160ff0219169083151502179055507fc76dd62bd7d0b2212e0d3445c1703a522dd816a749fe499b3bcb0f51b2500434836040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156123fb5780820151818401526020810190506123e0565b505050509050019250505060405180910390a1505050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561246a57600080fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680156124c3575060018054115b15156124ce57600080fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055507f84c8c1572fb6a01167de9135b17f8c2dd31ef2d755fa7d2b2d8f2b05c369868881604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156125f557600080fd5b601060009054906101000a900460ff16151561261057600080fd5b600f5466038d7ea4c680001115801561263357506706f05b59d3b20000600f5411155b151561263b57fe5b612646600f54612826565b6000601060006101000a81548160ff0219169083151502179055507f19ec72a595b8aab321636cc55d51478ac78e93a69b5c4a07ae548eb29e40c0a0600f546040518082815260200191505060405180910390a1565b600660019054906101000a900460ff1681565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561270657600080fd5b600b60009054906101000a900460ff16151561272157600080fd5b61272c600a546127fe565b6000600b60006101000a81548160ff0219169083151502179055507f21e1049325acc99b4f885709c6ca1a70281b586f585ef03485b62f7ad0a1e253600a546040518082815260200191505060405180910390a1565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156127d957600080fd5b6000600a819055506000600b60006101000a81548160ff021916908315150217905550565b61280661284e565b806007819055506007546009541115612823576007546009819055505b50565b61282e612896565b80600c81905550600c54600e54111561284b57600c54600e819055505b50565b600062015180600854014211156128935762015180600854420381151561287157fe5b0490506201518081026008600082825401925050819055506007546009819055505b50565b600062015180600d54014211156128db5762015180600d5442038115156128b957fe5b049050620151808102600d60008282540192505081905550600c54600e819055505b50565b828054828255906000526020600020908101928215612957579160200282015b828111156129565782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906128fe565b5b5090506129649190612989565b5090565b508054600082559060005260206000209081019061298691906129cc565b50565b6129c991905b808211156129c557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010161298f565b5090565b90565b6129ee91905b808211156129ea5760008160009055506001016129d2565b5090565b905600a165627a7a72305820aff657d329306657e336fc34577d84bbe953750055b7f78c80ed5232ba436b000029`

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

// InitializedSpendLimit is a free data retrieval call binding the contract method 0x5658eff0.
//
// Solidity: function initializedSpendLimit() constant returns(bool)
func (_Wallet *WalletCaller) InitializedSpendLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "initializedSpendLimit")
	return *ret0, err
}

// InitializedSpendLimit is a free data retrieval call binding the contract method 0x5658eff0.
//
// Solidity: function initializedSpendLimit() constant returns(bool)
func (_Wallet *WalletSession) InitializedSpendLimit() (bool, error) {
	return _Wallet.Contract.InitializedSpendLimit(&_Wallet.CallOpts)
}

// InitializedSpendLimit is a free data retrieval call binding the contract method 0x5658eff0.
//
// Solidity: function initializedSpendLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) InitializedSpendLimit() (bool, error) {
	return _Wallet.Contract.InitializedSpendLimit(&_Wallet.CallOpts)
}

// InitializedTopupLimit is a free data retrieval call binding the contract method 0xcd96895c.
//
// Solidity: function initializedTopupLimit() constant returns(bool)
func (_Wallet *WalletCaller) InitializedTopupLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "initializedTopupLimit")
	return *ret0, err
}

// InitializedTopupLimit is a free data retrieval call binding the contract method 0xcd96895c.
//
// Solidity: function initializedTopupLimit() constant returns(bool)
func (_Wallet *WalletSession) InitializedTopupLimit() (bool, error) {
	return _Wallet.Contract.InitializedTopupLimit(&_Wallet.CallOpts)
}

// InitializedTopupLimit is a free data retrieval call binding the contract method 0xcd96895c.
//
// Solidity: function initializedTopupLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) InitializedTopupLimit() (bool, error) {
	return _Wallet.Contract.InitializedTopupLimit(&_Wallet.CallOpts)
}

// InitializedWhitelist is a free data retrieval call binding the contract method 0xafa0fd9b.
//
// Solidity: function initializedWhitelist() constant returns(bool)
func (_Wallet *WalletCaller) InitializedWhitelist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "initializedWhitelist")
	return *ret0, err
}

// InitializedWhitelist is a free data retrieval call binding the contract method 0xafa0fd9b.
//
// Solidity: function initializedWhitelist() constant returns(bool)
func (_Wallet *WalletSession) InitializedWhitelist() (bool, error) {
	return _Wallet.Contract.InitializedWhitelist(&_Wallet.CallOpts)
}

// InitializedWhitelist is a free data retrieval call binding the contract method 0xafa0fd9b.
//
// Solidity: function initializedWhitelist() constant returns(bool)
func (_Wallet *WalletCallerSession) InitializedWhitelist() (bool, error) {
	return _Wallet.Contract.InitializedWhitelist(&_Wallet.CallOpts)
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

// SubmittedAddition is a free data retrieval call binding the contract method 0xd2f9a639.
//
// Solidity: function submittedAddition() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedAddition(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedAddition")
	return *ret0, err
}

// SubmittedAddition is a free data retrieval call binding the contract method 0xd2f9a639.
//
// Solidity: function submittedAddition() constant returns(bool)
func (_Wallet *WalletSession) SubmittedAddition() (bool, error) {
	return _Wallet.Contract.SubmittedAddition(&_Wallet.CallOpts)
}

// SubmittedAddition is a free data retrieval call binding the contract method 0xd2f9a639.
//
// Solidity: function submittedAddition() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedAddition() (bool, error) {
	return _Wallet.Contract.SubmittedAddition(&_Wallet.CallOpts)
}

// SubmittedRemoval is a free data retrieval call binding the contract method 0xf9c26de4.
//
// Solidity: function submittedRemoval() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedRemoval(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedRemoval")
	return *ret0, err
}

// SubmittedRemoval is a free data retrieval call binding the contract method 0xf9c26de4.
//
// Solidity: function submittedRemoval() constant returns(bool)
func (_Wallet *WalletSession) SubmittedRemoval() (bool, error) {
	return _Wallet.Contract.SubmittedRemoval(&_Wallet.CallOpts)
}

// SubmittedRemoval is a free data retrieval call binding the contract method 0xf9c26de4.
//
// Solidity: function submittedRemoval() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedRemoval() (bool, error) {
	return _Wallet.Contract.SubmittedRemoval(&_Wallet.CallOpts)
}

// SubmittedSpendLimit is a free data retrieval call binding the contract method 0xd5666590.
//
// Solidity: function submittedSpendLimit() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedSpendLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedSpendLimit")
	return *ret0, err
}

// SubmittedSpendLimit is a free data retrieval call binding the contract method 0xd5666590.
//
// Solidity: function submittedSpendLimit() constant returns(bool)
func (_Wallet *WalletSession) SubmittedSpendLimit() (bool, error) {
	return _Wallet.Contract.SubmittedSpendLimit(&_Wallet.CallOpts)
}

// SubmittedSpendLimit is a free data retrieval call binding the contract method 0xd5666590.
//
// Solidity: function submittedSpendLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedSpendLimit() (bool, error) {
	return _Wallet.Contract.SubmittedSpendLimit(&_Wallet.CallOpts)
}

// SubmittedTopupLimit is a free data retrieval call binding the contract method 0x3671bb0d.
//
// Solidity: function submittedTopupLimit() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedTopupLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedTopupLimit")
	return *ret0, err
}

// SubmittedTopupLimit is a free data retrieval call binding the contract method 0x3671bb0d.
//
// Solidity: function submittedTopupLimit() constant returns(bool)
func (_Wallet *WalletSession) SubmittedTopupLimit() (bool, error) {
	return _Wallet.Contract.SubmittedTopupLimit(&_Wallet.CallOpts)
}

// SubmittedTopupLimit is a free data retrieval call binding the contract method 0x3671bb0d.
//
// Solidity: function submittedTopupLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedTopupLimit() (bool, error) {
	return _Wallet.Contract.SubmittedTopupLimit(&_Wallet.CallOpts)
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
