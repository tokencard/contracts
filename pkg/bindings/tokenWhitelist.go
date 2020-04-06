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
const TokenWhitelistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_oracleNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_stablecoinAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenLoadable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenRedeemable\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"_redeemable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getERC20RecipientAndAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemableCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemableTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"setTokenLoadable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"setTokenRedeemable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenWhitelistBin is the compiled bytecode used for deploying new contracts.
var TokenWhitelistBin = "0x60806040527f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d589369760025534801561003457600080fd5b506040516124643803806124648339818101604052608081101561005757600080fd5b50805160208201516040830151606090930151600180546001600160a01b038086166001600160a01b03199283161792839055600080549092169216919091179055919290918180156100aa5760028190555b5060089290925550600780546001600160a01b039092166001600160a01b03199092169190911790555060046020527fa329885c08741397fd3c8a6391655875994b8b3f4267d51f002324c79dfda7c7805460ff1990811660019081179092557f272caa52cdef72c610b5313d50ca07262761e093a56fa2fa7114c2b7a13aca6180548216831790557ffe246a62db334be0c21bf6bcd2dda5f5c4dd84ad286b6c507001745ea44cfc4c80548216831790556323b872dd60e01b6000527f64f10a4788acb6e7f7bb8cb159bb306107fb979e4655d590b30c3f1bbc4b1a5e805490911690911790556122c3806101a16000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80636c3824ef116100a2578063996cba6811610071578063996cba6814610584578063afc72e93146105ba578063d545782e1461065b578063e2b4ce971461068d578063e9cbd8221461069557610116565b80636c3824ef146104bc5780637d73b2311461052a57806380cc0dda1461054e578063872773061461055657610116565b80633efec5e9116100e95780633efec5e9146103f0578063443dd2a4146103f857806344b049bc146104505780635d793a7d146104585780636a1744dc1461048657610116565b806313d5e8461461011b5780631d3a069f146101355780631f69565f1461017057806334c73edc14610240575b600080fd5b61012361069d565b60408051918252519081900360200190f35b61015c6004803603602081101561014b57600080fd5b50356001600160e01b0319166106a4565b604080519115158252519081900360200190f35b6101966004803603602081101561018657600080fd5b50356001600160a01b03166106c7565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b838110156101ff5781810151838201526020016101e7565b50505050905090810190601f16801561022c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103ee600480360360c081101561025657600080fd5b810190602081018135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460208302840111600160201b831117156102a357600080fd5b919390929091602081019035600160201b8111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460208302840111600160201b831117156102f357600080fd5b919390929091602081019035600160201b81111561031057600080fd5b82018360208201111561032257600080fd5b803590602001918460208302840111600160201b8311171561034357600080fd5b919390929091602081019035600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460208302840111600160201b8311171561039357600080fd5b919390929091602081019035600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460208302840111600160201b831117156103e357600080fd5b9193509150356107c2565b005b610196610c54565b610400610d4f565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561043c578181015183820152602001610424565b505050509050019250505060405180910390f35b610400610db1565b6103ee6004803603604081101561046e57600080fd5b506001600160a01b0381351690602001351515610e7e565b61015c6004803603604081101561049c57600080fd5b5080356001600160a01b031690602001356001600160e01b031916610fb6565b6103ee600480360360208110156104d257600080fd5b810190602081018135600160201b8111156104ec57600080fd5b8201836020820111156104fe57600080fd5b803590602001918460208302840111600160201b8311171561051f57600080fd5b50909250905061103d565b610532611300565b604080516001600160a01b039092168252519081900360200190f35b61012361130f565b6103ee6004803603604081101561056c57600080fd5b506001600160a01b0381351690602001351515611315565b6103ee6004803603606081101561059a57600080fd5b506001600160a01b0381358116916020810135909116906040013561144b565b610638600480360360408110156105d057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b5090925090506114f8565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103ee6004803603606081101561067157600080fd5b506001600160a01b03813516906020810135906040013561183c565b610123611987565b61053261198d565b6006545b90565b6001600160e01b0319811660009081526004602052604090205460ff165b919050565b6001600160a01b0381166000908152600360208181526040808420600180820154600280840154968401546004850154855487516101009682161587026000190190911693909304601f8101899004890284018901909752868352606099988998899889988998899891978897929660ff8085169693850481169562010000909504169390928991908301828280156107a15780601f10610776576101008083540402835291602001916107a1565b820191906000526020600020905b81548152906001019060200180831161078457829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107cb3361199c565b610815576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561082357508986145b801561082e57508984145b801561083957508984145b61088a576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4657600360008d8d848181106108a557fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610920576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606109456109408c8c8581811061093457fe5b90506020020135611a30565b611a5c565b90506040518060e001604052808281526020018a8a8581811061096457fe5b9050602002013581526020016000815260200160011515815260200188888581811061098c57fe5b905060200201351515151581526020018686858181106109a857fe5b9050602002013515151515815260200184815250600360008f8f868181106109cc57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000019080519060200190610a18929190612145565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560058d8d84818110610a8f57fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b039590920293909301359390931692909217905550848483818110610ad757fe5b9050602002013515610afb57600654610af790600163ffffffff611ab316565b6006555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2957fe5b905060200201356001600160a01b0316838c8c87818110610b4657fe5b905060200201358b8b88818110610b5957fe5b9050602002013515158a8a89818110610b6e57fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610bfe578181015183820152602001610be6565b50505050905090810190601f168015610c2b5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060010161088d565b505050505050505050505050565b6007546001600160a01b03166000908152600360208181526040808420600180820154600280840154968401546004850154855487516101009682161587026000190190911693909304601f8101899004890284018901909752868352606099988998899889988998899891978897929660ff808516969385048116956201000090950416939092899190830182828015610d305780601f10610d0557610100808354040283529160200191610d30565b820191906000526020600020905b815481529060010190602001808311610d1357829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606005805480602002602001604051908101604052809291908181526020018280548015610da757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d89575b5050505050905090565b606080600654604051908082528060200260200182016040528015610de0578160200160208202803683370190505b5090506000805b600554811015610e7657600060058281548110610e0057fe5b60009182526020808320909101546001600160a01b03168083526003918290526040909220015490915060ff620100009091041615610e6d5780848481518110610e4657fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610de7565b509091505090565b610e873361199c565b610ed1576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b0382166000908152600360208190526040909120015460ff16610f3b576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b038216600081815260036020818152604092839020909101805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526003602081905260408220015460ff1661101b576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526004602052604090205460ff16919050565b6110463361199c565b611090576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156112fb5760008383838181106110a957fe5b602090810292909201356001600160a01b0316600081815260039384905260409020909201549192505060ff16611120576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b0381166000908152600360208190526040909120015462010000900460ff16156111635760065461115f90600163ffffffff611b1416565b6006555b6001600160a01b03811660009081526003602052604081209061118682826121c3565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b6005546111c290600163ffffffff611b1416565b81101561127b57816001600160a01b0316600582815481106111e057fe5b6000918252602090912001546001600160a01b03161415611273576005805461121090600163ffffffff611b1416565b8154811061121a57fe5b600091825260209091200154600580546001600160a01b03909216918390811061124057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061127b565b6001016111ae565b50600580548061128757fe5b6000828152602090819020820160001990810180546001600160a01b0319169055909101909155604080513381526001600160a01b0384169281019290925280517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad09281900390910190a150600101611093565b505050565b6001546001600160a01b031690565b60085490565b61131e3361199c565b611368576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b0382166000908152600360208190526040909120015460ff166113d2576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b0382166000818152600360208181526040928390209091018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b6114543361199c565b61149e576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6114a9838383611b56565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6000806024831015611551576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b600061159d600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611bba169050565b90506115a98682610fb6565b6115ef576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b141561165d5785611653600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c39169050565b9250925050611834565b6001600160e01b031981166323b872dd60e01b14156117605760648410156116cc576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b611716603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c4c169050565b611653604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c39169050565b60448410156117a05760405162461bcd60e51b815260040180806020018281038252602581526020018061223f6025913960400191505060405180910390fd5b6117ea601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c4c169050565b611653602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c39169050565b935093915050565b6000611849600854611cce565b90506118543361199c565b806118675750336001600160a01b038216145b6118b1576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b0384166000908152600360208190526040909120015460ff1661191b576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03841660008181526003602090815260409182902060028101879055600401859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60025490565b6007546001600160a01b031690565b60006119a9600254611cce565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156119fe57600080fd5b505afa158015611a12573d6000803e3d6000fd5b505050506040513d6020811015611a2857600080fd5b505192915050565b611a3861220a565b6040516020810160405282815280602083015250611a5582611d90565b8152919050565b60608082600001516040519080825280601f01601f191660200182016040528015611a8e576020820181803683370190505b5090506000602082019050611aac8185602001518660000151611e29565b5092915050565b600082820183811015611b0d576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000611b0d83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611e67565b6001600160a01b038216611ba0576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015611b9a573d6000803e3d6000fd5b506112fb565b6112fb6001600160a01b038316848363ffffffff611efe16565b6000611bcd82600463ffffffff611ab316565b83511015611c19576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611c2d84602063ffffffff611ab316565b90940151949350505050565b6000611bcd82602063ffffffff611ab316565b6000611c5f82601463ffffffff611ab316565b83511015611cab576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611cbf84602063ffffffff611ab316565b9094015160601c949350505050565b6000805460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611d1b57600080fd5b505afa158015611d2f573d6000803e3d6000fd5b505050506040513d6020811015611d4557600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b1580156119fe57600080fd5b60008082611da25760009150506106c2565b6fffffffffffffffffffffffffffffffff8316611dc657601001600160801b830492505b67ffffffffffffffff8316611de75760080168010000000000000000830492505b63ffffffff8316611dff57600401600160201b830492505b61ffff8316611e145760020162010000830492505b60ff8316611e20576001015b60200392915050565b5b60208110611e49578151835260209283019290910190601f1901611e2a565b905182516020929092036101000a6000190180199091169116179052565b60008184841115611ef65760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ebb578181015183820152602001611ea3565b50505050905090810190601f168015611ee85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526112fb908490611f5d826001600160a01b0316612109565b611fae576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310611fec5780518252601f199092019160209182019101611fcd565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461204e576040519150601f19603f3d011682016040523d82523d6000602084013e612053565b606091505b5091509150816120aa576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115612103578080602001905160208110156120c657600080fd5b50516121035760405162461bcd60e51b815260040180806020018281038252602a815260200180612264602a913960400191505060405180910390fd5b50505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061213d57508115155b949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061218657805160ff19168380011785556121b3565b828001600101855582156121b3579182015b828111156121b3578251825591602001919060010190612198565b506121bf929150612224565b5090565b50805460018160011615610100020316600290046000825580601f106121e95750612207565b601f0160209004906000526020600020908101906122079190612224565b50565b604051806040016040528060008152602001600081525090565b6106a191905b808211156121bf576000815560010161222a56fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f76655361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220eb09b3a841e0569cabd12df8413a9476e4ef49e0d3e16ce8b7f4427c84e24ffd64736f6c63430006040033"

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

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "claim", _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Claim(&_TokenWhitelist.TransactOpts, _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Claim(&_TokenWhitelist.TransactOpts, _to, _asset, _amount)
}

// ControllerNode is a paid mutator transaction binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() returns(bytes32)
func (_TokenWhitelist *TokenWhitelistTransactor) ControllerNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "controllerNode")
}

// ControllerNode is a paid mutator transaction binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) ControllerNode() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.TransactOpts)
}

// ControllerNode is a paid mutator transaction binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() returns(bytes32)
func (_TokenWhitelist *TokenWhitelistTransactorSession) ControllerNode() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.TransactOpts)
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_TokenWhitelist *TokenWhitelistTransactor) EnsRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "ensRegistry")
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_TokenWhitelist *TokenWhitelistSession) EnsRegistry() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.TransactOpts)
}

// EnsRegistry is a paid mutator transaction binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() returns(address)
func (_TokenWhitelist *TokenWhitelistTransactorSession) EnsRegistry() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.TransactOpts)
}

// GetERC20RecipientAndAmount is a paid mutator transaction binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistTransactor) GetERC20RecipientAndAmount(opts *bind.TransactOpts, _token common.Address, _data []byte) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "getERC20RecipientAndAmount", _token, _data)
}

// GetERC20RecipientAndAmount is a paid mutator transaction binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetERC20RecipientAndAmount(_token common.Address, _data []byte) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.GetERC20RecipientAndAmount(&_TokenWhitelist.TransactOpts, _token, _data)
}

// GetERC20RecipientAndAmount is a paid mutator transaction binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistTransactorSession) GetERC20RecipientAndAmount(_token common.Address, _data []byte) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.GetERC20RecipientAndAmount(&_TokenWhitelist.TransactOpts, _token, _data)
}

// GetStablecoinInfo is a paid mutator transaction binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistTransactor) GetStablecoinInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "getStablecoinInfo")
}

// GetStablecoinInfo is a paid mutator transaction binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetStablecoinInfo() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.TransactOpts)
}

// GetStablecoinInfo is a paid mutator transaction binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistTransactorSession) GetStablecoinInfo() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.TransactOpts)
}

// GetTokenInfo is a paid mutator transaction binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistTransactor) GetTokenInfo(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "getTokenInfo", _a)
}

// GetTokenInfo is a paid mutator transaction binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetTokenInfo(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.TransactOpts, _a)
}

// GetTokenInfo is a paid mutator transaction binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistTransactorSession) GetTokenInfo(_a common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.TransactOpts, _a)
}

// IsERC20MethodSupported is a paid mutator transaction binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) returns(bool)
func (_TokenWhitelist *TokenWhitelistTransactor) IsERC20MethodSupported(opts *bind.TransactOpts, _token common.Address, _methodId [4]byte) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "isERC20MethodSupported", _token, _methodId)
}

// IsERC20MethodSupported is a paid mutator transaction binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsERC20MethodSupported(_token common.Address, _methodId [4]byte) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.IsERC20MethodSupported(&_TokenWhitelist.TransactOpts, _token, _methodId)
}

// IsERC20MethodSupported is a paid mutator transaction binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) returns(bool)
func (_TokenWhitelist *TokenWhitelistTransactorSession) IsERC20MethodSupported(_token common.Address, _methodId [4]byte) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.IsERC20MethodSupported(&_TokenWhitelist.TransactOpts, _token, _methodId)
}

// IsERC20MethodWhitelisted is a paid mutator transaction binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) returns(bool)
func (_TokenWhitelist *TokenWhitelistTransactor) IsERC20MethodWhitelisted(opts *bind.TransactOpts, _methodId [4]byte) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "isERC20MethodWhitelisted", _methodId)
}

// IsERC20MethodWhitelisted is a paid mutator transaction binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsERC20MethodWhitelisted(_methodId [4]byte) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.IsERC20MethodWhitelisted(&_TokenWhitelist.TransactOpts, _methodId)
}

// IsERC20MethodWhitelisted is a paid mutator transaction binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) returns(bool)
func (_TokenWhitelist *TokenWhitelistTransactorSession) IsERC20MethodWhitelisted(_methodId [4]byte) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.IsERC20MethodWhitelisted(&_TokenWhitelist.TransactOpts, _methodId)
}

// OracleNode is a paid mutator transaction binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() returns(bytes32)
func (_TokenWhitelist *TokenWhitelistTransactor) OracleNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "oracleNode")
}

// OracleNode is a paid mutator transaction binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) OracleNode() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.TransactOpts)
}

// OracleNode is a paid mutator transaction binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() returns(bytes32)
func (_TokenWhitelist *TokenWhitelistTransactorSession) OracleNode() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.TransactOpts)
}

// RedeemableCounter is a paid mutator transaction binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() returns(uint256)
func (_TokenWhitelist *TokenWhitelistTransactor) RedeemableCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "redeemableCounter")
}

// RedeemableCounter is a paid mutator transaction binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() returns(uint256)
func (_TokenWhitelist *TokenWhitelistSession) RedeemableCounter() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RedeemableCounter(&_TokenWhitelist.TransactOpts)
}

// RedeemableCounter is a paid mutator transaction binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() returns(uint256)
func (_TokenWhitelist *TokenWhitelistTransactorSession) RedeemableCounter() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RedeemableCounter(&_TokenWhitelist.TransactOpts)
}

// RedeemableTokens is a paid mutator transaction binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() returns(address[])
func (_TokenWhitelist *TokenWhitelistTransactor) RedeemableTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "redeemableTokens")
}

// RedeemableTokens is a paid mutator transaction binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) RedeemableTokens() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RedeemableTokens(&_TokenWhitelist.TransactOpts)
}

// RedeemableTokens is a paid mutator transaction binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() returns(address[])
func (_TokenWhitelist *TokenWhitelistTransactorSession) RedeemableTokens() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RedeemableTokens(&_TokenWhitelist.TransactOpts)
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

// Stablecoin is a paid mutator transaction binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() returns(address)
func (_TokenWhitelist *TokenWhitelistTransactor) Stablecoin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "stablecoin")
}

// Stablecoin is a paid mutator transaction binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() returns(address)
func (_TokenWhitelist *TokenWhitelistSession) Stablecoin() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.TransactOpts)
}

// Stablecoin is a paid mutator transaction binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() returns(address)
func (_TokenWhitelist *TokenWhitelistTransactorSession) Stablecoin() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.TransactOpts)
}

// TokenAddressArray is a paid mutator transaction binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() returns(address[])
func (_TokenWhitelist *TokenWhitelistTransactor) TokenAddressArray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "tokenAddressArray")
}

// TokenAddressArray is a paid mutator transaction binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) TokenAddressArray() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.TransactOpts)
}

// TokenAddressArray is a paid mutator transaction binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() returns(address[])
func (_TokenWhitelist *TokenWhitelistTransactorSession) TokenAddressArray() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.TransactOpts)
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

// TokenWhitelistClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the TokenWhitelist contract.
type TokenWhitelistClaimedIterator struct {
	Event *TokenWhitelistClaimed // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistClaimed)
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
		it.Event = new(TokenWhitelistClaimed)
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
func (it *TokenWhitelistClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistClaimed represents a Claimed event raised by the TokenWhitelist contract.
type TokenWhitelistClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterClaimed(opts *bind.FilterOpts) (*TokenWhitelistClaimedIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistClaimedIterator{contract: _TokenWhitelist.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *TokenWhitelistClaimed) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistClaimed)
				if err := _TokenWhitelist.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseClaimed(log types.Log) (*TokenWhitelistClaimed, error) {
	event := new(TokenWhitelistClaimed)
	if err := _TokenWhitelist.contract.UnpackLog(event, "Claimed", log); err != nil {
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
