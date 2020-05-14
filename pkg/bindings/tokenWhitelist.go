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
const TokenWhitelistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_oracleNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_stablecoinAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenLoadable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenRedeemable\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"_redeemable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getERC20RecipientAndAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loadableCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"redeemableCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"redeemableTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"setTokenLoadable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"setTokenRedeemable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenWhitelistBin is the compiled bytecode used for deploying new contracts.
var TokenWhitelistBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d589369760015534801561003457600080fd5b506040516122393803806122398339818101604052608081101561005757600080fd5b50805160208201516040830151606090930151600080546001600160a01b0319166001600160a01b038516179055919290918180156100965760018190555b5060099290925550600880546001600160a01b039092166001600160a01b03199092169190911790555060036020527fcdb433fee9fcb0c66fb3017cd7ac58de1aaf0f6312cb21f38875538e20afd62b805460ff1990811660019081179092557fe807f63c4c8c5262d3389f70cc5810be9d2b6ff67ec0f02f24406aa89b2d508e80548216831790557f55fb275ebfd19bd09696346ba9fee327c970084fe514a163b9376f5a6c99b86380548216831790556323b872dd60e01b6000527feb1525fd03aeae976171a87d04a6b7cfac27c30b47b207519998e7ec6c02c8ec805490911690911790556120ac8061018d6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80636c3824ef116100ad578063d082e38111610071578063d082e38114610630578063d545782e14610638578063e2b4ce971461066a578063e9cbd82214610672578063fdb780ed1461067a57610121565b80636c3824ef146104c75780637d73b2311461053557806380cc0dda146105595780638727730614610561578063afc72e931461058f57610121565b80633efec5e9116100f45780633efec5e9146103fb578063443dd2a41461040357806344b049bc1461045b5780635d793a7d146104635780636a1744dc1461049157610121565b806313d5e846146101265780631d3a069f146101405780631f69565f1461017b57806334c73edc1461024b575b600080fd5b61012e610682565b60408051918252519081900360200190f35b6101676004803603602081101561015657600080fd5b50356001600160e01b031916610689565b604080519115158252519081900360200190f35b6101a16004803603602081101561019157600080fd5b50356001600160a01b03166106ac565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b8381101561020a5781810151838201526020016101f2565b50505050905090810190601f1680156102375780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103f9600480360360c081101561026157600080fd5b810190602081018135600160201b81111561027b57600080fd5b82018360208201111561028d57600080fd5b803590602001918460208302840111600160201b831117156102ae57600080fd5b919390929091602081019035600160201b8111156102cb57600080fd5b8201836020820111156102dd57600080fd5b803590602001918460208302840111600160201b831117156102fe57600080fd5b919390929091602081019035600160201b81111561031b57600080fd5b82018360208201111561032d57600080fd5b803590602001918460208302840111600160201b8311171561034e57600080fd5b919390929091602081019035600160201b81111561036b57600080fd5b82018360208201111561037d57600080fd5b803590602001918460208302840111600160201b8311171561039e57600080fd5b919390929091602081019035600160201b8111156103bb57600080fd5b8201836020820111156103cd57600080fd5b803590602001918460208302840111600160201b831117156103ee57600080fd5b9193509150356107a9565b005b6101a1610c56565b61040b610d53565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561044757818101518382015260200161042f565b505050509050019250505060405180910390f35b61040b610db5565b6103f96004803603604081101561047957600080fd5b506001600160a01b0381351690602001351515610e83565b610167600480360360408110156104a757600080fd5b5080356001600160a01b031690602001356001600160e01b031916611079565b6103f9600480360360208110156104dd57600080fd5b810190602081018135600160201b8111156104f757600080fd5b82018360208201111561050957600080fd5b803590602001918460208302840111600160201b8311171561052a57600080fd5b509092509050611100565b61053d6113fb565b604080516001600160a01b039092168252519081900360200190f35b61012e61140a565b6103f96004803603604081101561057757600080fd5b506001600160a01b0381351690602001351515611410565b61060d600480360360408110156105a557600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105cf57600080fd5b8201836020820111156105e157600080fd5b803590602001918460018302840111600160201b8311171561060257600080fd5b509092509050611603565b604080516001600160a01b03909316835260208301919091528051918290030190f35b61012e611947565b6103f96004803603606081101561064e57600080fd5b506001600160a01b03813516906020810135906040013561194d565b61012e611a98565b61053d611a9e565b61012e611aad565b6007545b90565b6001600160e01b0319811660009081526003602052604090205460ff165b919050565b6001600160a01b03811660009081526002602081815260408084206001808201548286015460038401546004850154855487516101009682161587026000190190911699909904601f810189900489028a0189019097528689526060999889988998899889988998919788979296909560ff808516969385048116956201000090950416939192909189918301828280156107885780601f1061075d57610100808354040283529160200191610788565b820191906000526020600020905b81548152906001019060200180831161076b57829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107b233611ab3565b6107fc576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561080a57508986145b801561081557508984145b801561082057508984145b610871576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4857600260008d8d8481811061088c57fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610907576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b606061092c6109278c8c8581811061091b57fe5b90506020020135611b47565b611b73565b90506040518060e001604052808281526020018a8a8581811061094b57fe5b9050602002013581526020016000815260200160011515815260200188888581811061097357fe5b9050602002013515151515815260200186868581811061098f57fe5b9050602002013515151515815260200184815250600260008f8f868181106109b357fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000190805190602001906109ff929190611f35565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c0909101516004918201558d8d84818110610a7457fe5b835460018082018655600095865260209586902090910180549290950293909301356001600160a01b03166001600160a01b03199091161790925560058054909101905550868683818110610ac557fe5b9050602002013515610adb576006805460010190555b848483818110610ae757fe5b9050602002013515610afd576007805460010190555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2b57fe5b905060200201356001600160a01b0316838c8c87818110610b4857fe5b905060200201358b8b88818110610b5b57fe5b9050602002013515158a8a89818110610b7057fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610c00578181015183820152602001610be8565b50505050905090810190601f168015610c2d5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610874565b505050505050505050505050565b6008546001600160a01b031660009081526002602081815260408084206001808201548286015460038401546004850154855487516101009682161587026000190190911699909904601f810189900489028a0189019097528689526060999889988998899889988998919788979296909560ff80851696938504811695620100009095041693919290918991830182828015610d345780601f10610d0957610100808354040283529160200191610d34565b820191906000526020600020905b815481529060010190602001808311610d1757829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606004805480602002602001604051908101604052809291908181526020018280548015610dab57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d8d575b5050505050905090565b606080600754604051908082528060200260200182016040528015610de4578160200160208202803883390190505b5090506000805b600454811015610e7b57600060048281548110610e0457fe5b60009182526020808320909101546001600160a01b0316808352600290915260409091206003015490915060ff620100009091041615610e725780848481518110610e4b57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610deb565b509091505090565b610e8c33611ab3565b610ed6576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090206003015460ff16610f46576040805162461bcd60e51b815260206004820152601f60248201527f72656465656d61626c653a20746f6b656e206e6f7420617661696c61626c6500604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090206003015460ff620100009091041615158115151415610fc4576040805162461bcd60e51b815260206004820152601b60248201527f72656465656d61626c653a206e6f207374617465206368616e67650000000000604482015290519081900360640190fd5b8015610fe657600754610fde90600163ffffffff611bca16565b600755610ffe565b600754610ffa90600163ffffffff611c2b16565b6007555b6001600160a01b038216600081815260026020908152604091829020600301805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526002602052604081206003015460ff166110de576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526003602052604090205460ff16919050565b61110933611ab3565b611153576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156113f657600083838381811061116c57fe5b602090810292909201356001600160a01b0316600081815260029093526040909220600301549192505060ff166111e3576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6005546111f790600163ffffffff611c2b16565b6005556001600160a01b038116600090815260026020526040902060030154610100900460ff161561123b5760065461123790600163ffffffff611c2b16565b6006555b6001600160a01b03811660009081526002602052604090206003015462010000900460ff161561127d5760075461127990600163ffffffff611c2b16565b6007555b6001600160a01b0381166000908152600260205260408120906112a08282611fb3565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b6004546112dc90600163ffffffff611c2b16565b81101561139557816001600160a01b0316600482815481106112fa57fe5b6000918252602090912001546001600160a01b0316141561138d576004805461132a90600163ffffffff611c2b16565b8154811061133457fe5b600091825260209091200154600480546001600160a01b03909216918390811061135a57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550611395565b6001016112c8565b5060048054906113a9906000198301611ffa565b50604080513381526001600160a01b038316602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a150600101611156565b505050565b6000546001600160a01b031690565b60095490565b61141933611ab3565b611463576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090206003015460ff166114d3576040805162461bcd60e51b815260206004820152601d60248201527f6c6f616461626c653a20746f6b656e206e6f7420617661696c61626c65000000604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090206003015460ff6101009091041615158115151415611550576040805162461bcd60e51b815260206004820152601960248201527f6c6f616461626c653a206e6f207374617465206368616e676500000000000000604482015290519081900360640190fd5b80156115725760065461156a90600163ffffffff611bca16565b60065561158a565b60065461158690600163ffffffff611c2b16565b6006555b6001600160a01b0382166000818152600260209081526040918290206003018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b600080602483101561165c576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b60006116a8600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c88169050565b90506116b48682611079565b6116fa576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b1415611768578561175e600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d07169050565b925092505061193f565b6001600160e01b031981166323b872dd60e01b141561186b5760648410156117d7576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b611821603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d1a169050565b61175e604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d07169050565b60448410156118ab5760405162461bcd60e51b81526004018080602001828103825260258152602001806120536025913960400191505060405180910390fd5b6118f5601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d1a169050565b61175e602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d07169050565b935093915050565b60055490565b600061195a600954611d9c565b905061196533611ab3565b806119785750336001600160a01b038216145b6119c2576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03841660009081526002602052604090206003015460ff16611a2b576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b0384166000818152600260208181526040928390209182018790556004909101859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60015490565b6008546001600160a01b031690565b60065490565b6000611ac0600154611d9c565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611b1557600080fd5b505afa158015611b29573d6000803e3d6000fd5b505050506040513d6020811015611b3f57600080fd5b505192915050565b611b4f61201e565b6040516020810160405282815280602083015250611b6c82611e5e565b8152919050565b60608082600001516040519080825280601f01601f191660200182016040528015611ba5576020820181803883390190505b5090506000602082019050611bc38185602001518660000151611ef7565b5092915050565b600082820183811015611c24576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600082821115611c82576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000611c9b82600463ffffffff611bca16565b83511015611ce7576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611cfb84602063ffffffff611bca16565b90940151949350505050565b6000611c9b82602063ffffffff611bca16565b6000611d2d82601463ffffffff611bca16565b83511015611d79576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d8d84602063ffffffff611bca16565b9094015160601c949350505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611de957600080fd5b505afa158015611dfd573d6000803e3d6000fd5b505050506040513d6020811015611e1357600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015611b1557600080fd5b60008082611e705760009150506106a7565b6fffffffffffffffffffffffffffffffff8316611e9457601001600160801b830492505b67ffffffffffffffff8316611eb55760080168010000000000000000830492505b63ffffffff8316611ecd57600401600160201b830492505b61ffff8316611ee25760020162010000830492505b60ff8316611eee576001015b60200392915050565b5b60208110611f17578151835260209283019290910190601f1901611ef8565b905182516020929092036101000a6000190180199091169116179052565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611f7657805160ff1916838001178555611fa3565b82800160010185558215611fa3579182015b82811115611fa3578251825591602001919060010190611f88565b50611faf929150612038565b5090565b50805460018160011615610100020316600290046000825580601f10611fd95750611ff7565b601f016020900490600052602060002090810190611ff79190612038565b50565b8154818355818111156113f6576000838152602090206113f6918101908301612038565b604051806040016040528060008152602001600081525090565b61068691905b80821115611faf576000815560010161203e56fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f7665a265627a7a7231582006083ec7cce0ee7af1963aa16c5b717b775646a163e5fa48c727f84e01e0bdb064736f6c63430005110032"

// DeployTokenWhitelist deploys a new Ethereum contract, binding an instance of TokenWhitelist to it.
func DeployTokenWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _oracleNode_ [32]byte, _controllerNode_ [32]byte, _stablecoinAddress_ common.Address) (common.Address, *types.Transaction, *TokenWhitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistBin), backend, _ens_, _oracleNode_, _controllerNode_, _stablecoinAddress_)
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

// GetERC20RecipientAndAmount is a free data retrieval call binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) constant returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetERC20RecipientAndAmount(opts *bind.CallOpts, _token common.Address, _data []byte) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getERC20RecipientAndAmount", _token, _data)
	return *ret0, *ret1, err
}

// GetERC20RecipientAndAmount is a free data retrieval call binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) constant returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetERC20RecipientAndAmount(_token common.Address, _data []byte) (common.Address, *big.Int, error) {
	return _TokenWhitelist.Contract.GetERC20RecipientAndAmount(&_TokenWhitelist.CallOpts, _token, _data)
}

// GetERC20RecipientAndAmount is a free data retrieval call binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) constant returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetERC20RecipientAndAmount(_token common.Address, _data []byte) (common.Address, *big.Int, error) {
	return _TokenWhitelist.Contract.GetERC20RecipientAndAmount(&_TokenWhitelist.CallOpts, _token, _data)
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

// IsERC20MethodSupported is a free data retrieval call binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCaller) IsERC20MethodSupported(opts *bind.CallOpts, _token common.Address, _methodId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "isERC20MethodSupported", _token, _methodId)
	return *ret0, err
}

// IsERC20MethodSupported is a free data retrieval call binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsERC20MethodSupported(_token common.Address, _methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodSupported(&_TokenWhitelist.CallOpts, _token, _methodId)
}

// IsERC20MethodSupported is a free data retrieval call binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCallerSession) IsERC20MethodSupported(_token common.Address, _methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodSupported(&_TokenWhitelist.CallOpts, _token, _methodId)
}

// IsERC20MethodWhitelisted is a free data retrieval call binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCaller) IsERC20MethodWhitelisted(opts *bind.CallOpts, _methodId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "isERC20MethodWhitelisted", _methodId)
	return *ret0, err
}

// IsERC20MethodWhitelisted is a free data retrieval call binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsERC20MethodWhitelisted(_methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodWhitelisted(&_TokenWhitelist.CallOpts, _methodId)
}

// IsERC20MethodWhitelisted is a free data retrieval call binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCallerSession) IsERC20MethodWhitelisted(_methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodWhitelisted(&_TokenWhitelist.CallOpts, _methodId)
}

// LoadableCounter is a free data retrieval call binding the contract method 0xfdb780ed.
//
// Solidity: function loadableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCaller) LoadableCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "loadableCounter")
	return *ret0, err
}

// LoadableCounter is a free data retrieval call binding the contract method 0xfdb780ed.
//
// Solidity: function loadableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistSession) LoadableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.LoadableCounter(&_TokenWhitelist.CallOpts)
}

// LoadableCounter is a free data retrieval call binding the contract method 0xfdb780ed.
//
// Solidity: function loadableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) LoadableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.LoadableCounter(&_TokenWhitelist.CallOpts)
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

// RedeemableCounter is a free data retrieval call binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCaller) RedeemableCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "redeemableCounter")
	return *ret0, err
}

// RedeemableCounter is a free data retrieval call binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistSession) RedeemableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.RedeemableCounter(&_TokenWhitelist.CallOpts)
}

// RedeemableCounter is a free data retrieval call binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) RedeemableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.RedeemableCounter(&_TokenWhitelist.CallOpts)
}

// RedeemableTokens is a free data retrieval call binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) RedeemableTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "redeemableTokens")
	return *ret0, err
}

// RedeemableTokens is a free data retrieval call binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) RedeemableTokens() ([]common.Address, error) {
	return _TokenWhitelist.Contract.RedeemableTokens(&_TokenWhitelist.CallOpts)
}

// RedeemableTokens is a free data retrieval call binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) RedeemableTokens() ([]common.Address, error) {
	return _TokenWhitelist.Contract.RedeemableTokens(&_TokenWhitelist.CallOpts)
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

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) TokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "tokenAddressArray")
	return *ret0, err
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.CallOpts)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCaller) TokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "tokenCounter")
	return *ret0, err
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistSession) TokenCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.TokenCounter(&_TokenWhitelist.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) TokenCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.TokenCounter(&_TokenWhitelist.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _redeemable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _redeemable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "addTokens", _tokens, _symbols, _magnitude, _loadable, _redeemable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _redeemable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _redeemable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _redeemable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _redeemable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _redeemable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _redeemable, _lastUpdate)
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

// SetTokenLoadable is a paid mutator transaction binding the contract method 0x87277306.
//
// Solidity: function setTokenLoadable(address _token, bool _loadable) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) SetTokenLoadable(opts *bind.TransactOpts, _token common.Address, _loadable bool) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "setTokenLoadable", _token, _loadable)
}

// SetTokenLoadable is a paid mutator transaction binding the contract method 0x87277306.
//
// Solidity: function setTokenLoadable(address _token, bool _loadable) returns()
func (_TokenWhitelist *TokenWhitelistSession) SetTokenLoadable(_token common.Address, _loadable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenLoadable(&_TokenWhitelist.TransactOpts, _token, _loadable)
}

// SetTokenLoadable is a paid mutator transaction binding the contract method 0x87277306.
//
// Solidity: function setTokenLoadable(address _token, bool _loadable) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) SetTokenLoadable(_token common.Address, _loadable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenLoadable(&_TokenWhitelist.TransactOpts, _token, _loadable)
}

// SetTokenRedeemable is a paid mutator transaction binding the contract method 0x5d793a7d.
//
// Solidity: function setTokenRedeemable(address _token, bool _redeemable) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) SetTokenRedeemable(opts *bind.TransactOpts, _token common.Address, _redeemable bool) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "setTokenRedeemable", _token, _redeemable)
}

// SetTokenRedeemable is a paid mutator transaction binding the contract method 0x5d793a7d.
//
// Solidity: function setTokenRedeemable(address _token, bool _redeemable) returns()
func (_TokenWhitelist *TokenWhitelistSession) SetTokenRedeemable(_token common.Address, _redeemable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenRedeemable(&_TokenWhitelist.TransactOpts, _token, _redeemable)
}

// SetTokenRedeemable is a paid mutator transaction binding the contract method 0x5d793a7d.
//
// Solidity: function setTokenRedeemable(address _token, bool _redeemable) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) SetTokenRedeemable(_token common.Address, _redeemable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenRedeemable(&_TokenWhitelist.TransactOpts, _token, _redeemable)
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

// TokenWhitelistAddedExclusiveMethodIterator is returned from FilterAddedExclusiveMethod and is used to iterate over the raw logs and unpacked data for AddedExclusiveMethod events raised by the TokenWhitelist contract.
type TokenWhitelistAddedExclusiveMethodIterator struct {
	Event *TokenWhitelistAddedExclusiveMethod // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAddedExclusiveMethodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedExclusiveMethod)
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
		it.Event = new(TokenWhitelistAddedExclusiveMethod)
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
func (it *TokenWhitelistAddedExclusiveMethodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedExclusiveMethodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedExclusiveMethod represents a AddedExclusiveMethod event raised by the TokenWhitelist contract.
type TokenWhitelistAddedExclusiveMethod struct {
	Token    common.Address
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddedExclusiveMethod is a free log retrieval operation binding the contract event 0xfb181256b03ef9051c59b29b98e8ef8dc1161e61d9062e1192ddd073806b0876.
//
// Solidity: event AddedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedExclusiveMethod(opts *bind.FilterOpts) (*TokenWhitelistAddedExclusiveMethodIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedExclusiveMethodIterator{contract: _TokenWhitelist.contract, event: "AddedExclusiveMethod", logs: logs, sub: sub}, nil
}

// WatchAddedExclusiveMethod is a free log subscription operation binding the contract event 0xfb181256b03ef9051c59b29b98e8ef8dc1161e61d9062e1192ddd073806b0876.
//
// Solidity: event AddedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedExclusiveMethod(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedExclusiveMethod) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedExclusiveMethod)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedExclusiveMethod", log); err != nil {
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

// ParseAddedExclusiveMethod is a log parse operation binding the contract event 0xfb181256b03ef9051c59b29b98e8ef8dc1161e61d9062e1192ddd073806b0876.
//
// Solidity: event AddedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseAddedExclusiveMethod(log types.Log) (*TokenWhitelistAddedExclusiveMethod, error) {
	event := new(TokenWhitelistAddedExclusiveMethod)
	if err := _TokenWhitelist.contract.UnpackLog(event, "AddedExclusiveMethod", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistAddedMethodIdIterator is returned from FilterAddedMethodId and is used to iterate over the raw logs and unpacked data for AddedMethodId events raised by the TokenWhitelist contract.
type TokenWhitelistAddedMethodIdIterator struct {
	Event *TokenWhitelistAddedMethodId // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAddedMethodIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedMethodId)
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
		it.Event = new(TokenWhitelistAddedMethodId)
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
func (it *TokenWhitelistAddedMethodIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedMethodIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedMethodId represents a AddedMethodId event raised by the TokenWhitelist contract.
type TokenWhitelistAddedMethodId struct {
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddedMethodId is a free log retrieval operation binding the contract event 0xcad8cc4e064e022264c8f21f5293f8b3c267eaa6895ee7c9e0b34689726eae71.
//
// Solidity: event AddedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedMethodId(opts *bind.FilterOpts) (*TokenWhitelistAddedMethodIdIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedMethodId")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedMethodIdIterator{contract: _TokenWhitelist.contract, event: "AddedMethodId", logs: logs, sub: sub}, nil
}

// WatchAddedMethodId is a free log subscription operation binding the contract event 0xcad8cc4e064e022264c8f21f5293f8b3c267eaa6895ee7c9e0b34689726eae71.
//
// Solidity: event AddedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedMethodId(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedMethodId) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedMethodId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedMethodId)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedMethodId", log); err != nil {
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

// ParseAddedMethodId is a log parse operation binding the contract event 0xcad8cc4e064e022264c8f21f5293f8b3c267eaa6895ee7c9e0b34689726eae71.
//
// Solidity: event AddedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseAddedMethodId(log types.Log) (*TokenWhitelistAddedMethodId, error) {
	event := new(TokenWhitelistAddedMethodId)
	if err := _TokenWhitelist.contract.UnpackLog(event, "AddedMethodId", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Sender     common.Address
	Token      common.Address
	Symbol     string
	Magnitude  *big.Int
	Loadable   bool
	Redeemable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedToken(opts *bind.FilterOpts) (*TokenWhitelistAddedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedTokenIterator{contract: _TokenWhitelist.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable)
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

// ParseAddedToken is a log parse operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseAddedToken(log types.Log) (*TokenWhitelistAddedToken, error) {
	event := new(TokenWhitelistAddedToken)
	if err := _TokenWhitelist.contract.UnpackLog(event, "AddedToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistRemovedExclusiveMethodIterator is returned from FilterRemovedExclusiveMethod and is used to iterate over the raw logs and unpacked data for RemovedExclusiveMethod events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedExclusiveMethodIterator struct {
	Event *TokenWhitelistRemovedExclusiveMethod // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistRemovedExclusiveMethodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedExclusiveMethod)
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
		it.Event = new(TokenWhitelistRemovedExclusiveMethod)
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
func (it *TokenWhitelistRemovedExclusiveMethodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedExclusiveMethodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedExclusiveMethod represents a RemovedExclusiveMethod event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedExclusiveMethod struct {
	Token    common.Address
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovedExclusiveMethod is a free log retrieval operation binding the contract event 0xe01bc5ecc4d7ff06fdb26bad9a3601ef089d9e5aa6f7dd03dc713b468eec117a.
//
// Solidity: event RemovedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedExclusiveMethod(opts *bind.FilterOpts) (*TokenWhitelistRemovedExclusiveMethodIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedExclusiveMethodIterator{contract: _TokenWhitelist.contract, event: "RemovedExclusiveMethod", logs: logs, sub: sub}, nil
}

// WatchRemovedExclusiveMethod is a free log subscription operation binding the contract event 0xe01bc5ecc4d7ff06fdb26bad9a3601ef089d9e5aa6f7dd03dc713b468eec117a.
//
// Solidity: event RemovedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedExclusiveMethod(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedExclusiveMethod) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedExclusiveMethod)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedExclusiveMethod", log); err != nil {
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

// ParseRemovedExclusiveMethod is a log parse operation binding the contract event 0xe01bc5ecc4d7ff06fdb26bad9a3601ef089d9e5aa6f7dd03dc713b468eec117a.
//
// Solidity: event RemovedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseRemovedExclusiveMethod(log types.Log) (*TokenWhitelistRemovedExclusiveMethod, error) {
	event := new(TokenWhitelistRemovedExclusiveMethod)
	if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedExclusiveMethod", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistRemovedMethodIdIterator is returned from FilterRemovedMethodId and is used to iterate over the raw logs and unpacked data for RemovedMethodId events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedMethodIdIterator struct {
	Event *TokenWhitelistRemovedMethodId // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistRemovedMethodIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedMethodId)
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
		it.Event = new(TokenWhitelistRemovedMethodId)
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
func (it *TokenWhitelistRemovedMethodIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedMethodIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedMethodId represents a RemovedMethodId event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedMethodId struct {
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovedMethodId is a free log retrieval operation binding the contract event 0x006dd38caa262b48ea0824b897ee1c4f238521632ad2c5d12f3f0225a1378d1d.
//
// Solidity: event RemovedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedMethodId(opts *bind.FilterOpts) (*TokenWhitelistRemovedMethodIdIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedMethodId")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedMethodIdIterator{contract: _TokenWhitelist.contract, event: "RemovedMethodId", logs: logs, sub: sub}, nil
}

// WatchRemovedMethodId is a free log subscription operation binding the contract event 0x006dd38caa262b48ea0824b897ee1c4f238521632ad2c5d12f3f0225a1378d1d.
//
// Solidity: event RemovedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedMethodId(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedMethodId) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedMethodId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedMethodId)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedMethodId", log); err != nil {
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

// ParseRemovedMethodId is a log parse operation binding the contract event 0x006dd38caa262b48ea0824b897ee1c4f238521632ad2c5d12f3f0225a1378d1d.
//
// Solidity: event RemovedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseRemovedMethodId(log types.Log) (*TokenWhitelistRemovedMethodId, error) {
	event := new(TokenWhitelistRemovedMethodId)
	if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedMethodId", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseRemovedToken is a log parse operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseRemovedToken(log types.Log) (*TokenWhitelistRemovedToken, error) {
	event := new(TokenWhitelistRemovedToken)
	if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistUpdatedTokenLoadableIterator is returned from FilterUpdatedTokenLoadable and is used to iterate over the raw logs and unpacked data for UpdatedTokenLoadable events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenLoadableIterator struct {
	Event *TokenWhitelistUpdatedTokenLoadable // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistUpdatedTokenLoadableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenLoadable)
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
		it.Event = new(TokenWhitelistUpdatedTokenLoadable)
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
func (it *TokenWhitelistUpdatedTokenLoadableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenLoadableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenLoadable represents a UpdatedTokenLoadable event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenLoadable struct {
	Sender   common.Address
	Token    common.Address
	Loadable bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenLoadable is a free log retrieval operation binding the contract event 0x0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec.
//
// Solidity: event UpdatedTokenLoadable(address _sender, address _token, bool _loadable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenLoadable(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenLoadableIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenLoadable")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenLoadableIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenLoadable", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenLoadable is a free log subscription operation binding the contract event 0x0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec.
//
// Solidity: event UpdatedTokenLoadable(address _sender, address _token, bool _loadable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenLoadable(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenLoadable) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenLoadable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenLoadable)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenLoadable", log); err != nil {
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

// ParseUpdatedTokenLoadable is a log parse operation binding the contract event 0x0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec.
//
// Solidity: event UpdatedTokenLoadable(address _sender, address _token, bool _loadable)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseUpdatedTokenLoadable(log types.Log) (*TokenWhitelistUpdatedTokenLoadable, error) {
	event := new(TokenWhitelistUpdatedTokenLoadable)
	if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenLoadable", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseUpdatedTokenRate is a log parse operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseUpdatedTokenRate(log types.Log) (*TokenWhitelistUpdatedTokenRate, error) {
	event := new(TokenWhitelistUpdatedTokenRate)
	if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistUpdatedTokenRedeemableIterator is returned from FilterUpdatedTokenRedeemable and is used to iterate over the raw logs and unpacked data for UpdatedTokenRedeemable events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRedeemableIterator struct {
	Event *TokenWhitelistUpdatedTokenRedeemable // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistUpdatedTokenRedeemableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenRedeemable)
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
		it.Event = new(TokenWhitelistUpdatedTokenRedeemable)
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
func (it *TokenWhitelistUpdatedTokenRedeemableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenRedeemableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenRedeemable represents a UpdatedTokenRedeemable event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRedeemable struct {
	Sender     common.Address
	Token      common.Address
	Redeemable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenRedeemable is a free log retrieval operation binding the contract event 0xcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d.
//
// Solidity: event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenRedeemable(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenRedeemableIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenRedeemable")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenRedeemableIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenRedeemable", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenRedeemable is a free log subscription operation binding the contract event 0xcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d.
//
// Solidity: event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenRedeemable(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenRedeemable) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenRedeemable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenRedeemable)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRedeemable", log); err != nil {
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

// ParseUpdatedTokenRedeemable is a log parse operation binding the contract event 0xcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d.
//
// Solidity: event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseUpdatedTokenRedeemable(log types.Log) (*TokenWhitelistUpdatedTokenRedeemable, error) {
	event := new(TokenWhitelistUpdatedTokenRedeemable)
	if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRedeemable", log); err != nil {
		return nil, err
	}
	return event, nil
}
