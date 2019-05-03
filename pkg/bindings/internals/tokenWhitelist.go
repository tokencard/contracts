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
const TokenWhitelistABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"},{\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"name\":\"_burnable\",\"type\":\"bool[]\"},{\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens_\",\"type\":\"address\"},{\"name\":\"_oracleNameHash_\",\"type\":\"bytes32\"},{\"name\":\"_controllerNameHash_\",\"type\":\"bytes32\"},{\"name\":\"_owner_\",\"type\":\"address\"},{\"name\":\"_transferable_\",\"type\":\"bool\"},{\"name\":\"_stabelcoinAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_burnable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"}]"

// TokenWhitelistBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistBin = `608060405234801561001057600080fd5b5060405160c080611b5c833981018060405260c081101561003057600080fd5b508051602082015160408301516060840151608085015160a090950151600180546001600160a01b038088166001600160a01b031992831617928390556000805483169382169390931790925560028590556003805489151574010000000000000000000000000000000000000000908102600160a01b60ff021995881692909416919091179390931691909117908190559596949593949293849184910460ff1661011357604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506007949094555050600680546001600160a01b0319166001600160a01b0390931692909217909155506119c99050806101936000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637d73b23111610097578063b242e53411610066578063b242e534146104e8578063d545782e14610516578063e2b4ce9714610548578063e9cbd82214610550576100f5565b80637d73b2311461046c57806380cc0dda146104905780638da5cb5b146104aa578063996cba68146104b2576100f5565b80633efec5e9116100d35780633efec5e914610396578063443dd2a41461039e5780636c3824ef146103f6578063715018a614610464576100f5565b80631f69565f146100fa5780632121dc75146101ca57806334c73edc146101e6575b600080fd5b6101206004803603602081101561011057600080fd5b50356001600160a01b0316610558565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b83811015610189578181015183820152602001610171565b50505050905090810190601f1680156101b65780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6101d2610654565b604080519115158252519081900360200190f35b610394600480360360c08110156101fc57600080fd5b810190602081018135600160201b81111561021657600080fd5b82018360208201111561022857600080fd5b803590602001918460208302840111600160201b8311171561024957600080fd5b919390929091602081019035600160201b81111561026657600080fd5b82018360208201111561027857600080fd5b803590602001918460208302840111600160201b8311171561029957600080fd5b919390929091602081019035600160201b8111156102b657600080fd5b8201836020820111156102c857600080fd5b803590602001918460208302840111600160201b831117156102e957600080fd5b919390929091602081019035600160201b81111561030657600080fd5b82018360208201111561031857600080fd5b803590602001918460208302840111600160201b8311171561033957600080fd5b919390929091602081019035600160201b81111561035657600080fd5b82018360208201111561036857600080fd5b803590602001918460208302840111600160201b8311171561038957600080fd5b919350915035610665565b005b610120610ad7565b6103a6610bd3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103e25781810151838201526020016103ca565b505050509050019250505060405180910390f35b6103946004803603602081101561040c57600080fd5b810190602081018135600160201b81111561042657600080fd5b82018360208201111561043857600080fd5b803590602001918460208302840111600160201b8311171561045957600080fd5b509092509050610c35565b610394610ec4565b610474610fcf565b604080516001600160a01b039092168252519081900360200190f35b610498610fde565b60408051918252519081900360200190f35b610474610fe4565b610394600480360360608110156104c857600080fd5b506001600160a01b03813581169160208101359091169060400135610ff3565b610394600480360360408110156104fe57600080fd5b506001600160a01b038135169060200135151561105b565b6103946004803603606081101561052c57600080fd5b506001600160a01b038135169060208101359060400135611236565b610498611394565b61047461139a565b6001600160a01b0381166000908152600460208181526040808420600180820154600280840154600385015497850154855487516101009682161587026000190190911693909304601f81018990048902840189019097528683526060999889988998899889988998919788979296929560ff80821696938204811695620100009092041693928991908301828280156106335780601f1061060857610100808354040283529160200191610633565b820191906000526020600020905b81548152906001019060200180831161061657829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b600354600160a01b900460ff165b90565b61066e336113a9565b6106c25760408051600160e51b62461bcd02815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b89881480156106d057508986145b80156106db57508984145b80156106e657508984145b61073a5760408051600160e51b62461bcd02815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610ac957600460008d8d8481811061075557fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff16156107d35760408051600160e51b62461bcd02815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606107f86107f38c8c858181106107e757fe5b9050602002013561143f565b61146b565b90506040518060e001604052808281526020018a8a8581811061081757fe5b9050602002013581526020016000815260200160011515815260200188888581811061083f57fe5b9050602002013515151515815260200186868581811061085b57fe5b9050602002013515151515815260200184815250600460008f8f8681811061087f57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000190805190602001906108cb929190611838565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560058d8d8481811061094257fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055507f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e858181106109ac57fe5b905060200201356001600160a01b0316838c8c878181106109c957fe5b905060200201358b8b888181106109dc57fe5b9050602002013515158a8a898181106109f157fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610a81578181015183820152602001610a69565b50505050905090810190601f168015610aae5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060010161073d565b505050505050505050505050565b6006546001600160a01b03166000908152600460208181526040808420600180820154600280840154600385015497850154855487516101009682161587026000190190911693909304601f81018990048902840189019097528683526060999889988998899889988998919788979296929560ff8082169693820481169562010000909204169392899190830182828015610bb45780601f10610b8957610100808354040283529160200191610bb4565b820191906000526020600020905b815481529060010190602001808311610b9757829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606005805480602002602001604051908101604052809291908181526020018280548015610c2b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610c0d575b5050505050905090565b610c3e336113a9565b610c925760408051600160e51b62461bcd02815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b60005b81811015610ebf5760046000848484818110610cad57fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff16610d2a5760408051600160e51b62461bcd02815260206004820152601660248201527f746f6b656e206973206e6f7420617661696c61626c6500000000000000000000604482015290519081900360640190fd5b6000838383818110610d3857fe5b602090810292909201356001600160a01b03166000818152600490935260408320909350919050610d6982826118b6565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b600554610da590600163ffffffff6114c216565b811015610e5e57816001600160a01b031660058281548110610dc357fe5b6000918252602090912001546001600160a01b03161415610e565760058054610df390600163ffffffff6114c216565b81548110610dfd57fe5b600091825260209091200154600580546001600160a01b039092169183908110610e2357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550610e5e565b600101610d91565b506005805490610e729060001983016118fd565b50604080513381526001600160a01b038316602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a150600101610c95565b505050565b610ecd33611522565b610f215760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b600354600160a01b900460ff16610f825760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b6001546001600160a01b031690565b60075490565b6003546001600160a01b031690565b610ffc33611522565b6110505760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b610ebf838383611536565b61106433611522565b6110b85760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b600354600160a01b900460ff166111195760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b03821661116157604051600160e51b62461bcd0281526004018080602001828103825260238152602001806119566023913960400191505060405180910390fd5b6003805474ff00000000000000000000000000000000000000001916600160a01b83151502179055806111cb57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000611243600754611699565b905061124e336113a9565b806112615750336001600160a01b038216145b6112b55760408051600160e51b62461bcd02815260206004820152601b60248201527f656974686572206f7261636c65206f7220636f6e74726f6c6c65720000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090206003015460ff166113285760408051600160e51b62461bcd02815260206004820152601660248201527f746f6b656e206973206e6f7420617661696c61626c6500000000000000000000604482015290519081900360640190fd5b6001600160a01b03841660008181526004602081815260409283902060028101889055909101859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60025490565b6006546001600160a01b031690565b60006113b6600254611699565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561140b57600080fd5b505afa15801561141f573d6000803e3d6000fd5b505050506040513d602081101561143557600080fd5b505190505b919050565b611447611921565b604051602081016040528281528060208301525061146482611761565b8152919050565b60608082600001516040519080825280601f01601f19166020018201604052801561149d576020820181803883390190505b50905060006020820190506114bb81856020015186600001516117fa565b5092915050565b60008282111561151c5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6003546001600160a01b0390811691161490565b6001600160a01b038216611580576040516001600160a01b0384169082156108fc029083906000818181858888f1935050505015801561157a573d6000803e3d6000fd5b5061164a565b816001600160a01b031663a9059cbb84836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156115e057600080fd5b505af11580156115f4573d6000803e3d6000fd5b505050506040513d602081101561160a57600080fd5b505161164a57604051600160e51b62461bcd0281526004018080602001828103825260258152602001806119796025913960400191505060405180910390fd5b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6000805460408051600160e01b630178b8bf0281526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156116e957600080fd5b505afa1580156116fd573d6000803e3d6000fd5b505050506040513d602081101561171357600080fd5b505160408051600160e11b631d9dabef0281526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561140b57600080fd5b6000808261177357600091505061143a565b6fffffffffffffffffffffffffffffffff831661179757601001600160801b830492505b67ffffffffffffffff83166117b85760080168010000000000000000830492505b63ffffffff83166117d057600401600160201b830492505b61ffff83166117e55760020162010000830492505b60ff83166117f1576001015b60200392915050565b5b6020811061181a578151835260209283019290910190601f19016117fb565b905182516020929092036101000a6000190180199091169116179052565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061187957805160ff19168380011785556118a6565b828001600101855582156118a6579182015b828111156118a657825182559160200191906001019061188b565b506118b292915061193b565b5090565b50805460018160011615610100020316600290046000825580601f106118dc57506118fa565b601f0160209004906000526020600020908101906118fa919061193b565b50565b815481835581811115610ebf57600083815260209020610ebf91810190830161193b565b604051806040016040528060008152602001600081525090565b61066291905b808211156118b2576000815560010161194156fe6f776e65722063616e6e6f742062652073657420746f207a65726f2061646472657373455243323020746f6b656e207472616e736665722077617320756e7375636365737366756ca165627a7a72305820d6077d3b9917a7a7d92e652ced742860eab420242d2dab52091d7e37a110d2790029`

// DeployTokenWhitelist deploys a new Ethereum contract, binding an instance of TokenWhitelist to it.
func DeployTokenWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _oracleNameHash_ [32]byte, _controllerNameHash_ [32]byte, _owner_ common.Address, _transferable_ bool, _stabelcoinAddress_ common.Address) (common.Address, *types.Transaction, *TokenWhitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistBin), backend, _ens_, _oracleNameHash_, _controllerNameHash_, _owner_, _transferable_, _stabelcoinAddress_)
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
