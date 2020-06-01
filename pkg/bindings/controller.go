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

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AddedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"AddedSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ensRegistry\",\"type\":\"address\"}],\"name\":\"ENSSetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutedTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_gstNode\",\"type\":\"bytes32\"}],\"name\":\"GSTSetNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"RemovedSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Started\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adminCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ensRegistry\",\"type\":\"address\"}],\"name\":\"ensSetRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasRefund\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_gstNode\",\"type\":\"bytes32\"}],\"name\":\"gstSetNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isController\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ControllerBin is the compiled bytecode used for deploying new contracts.
var ControllerBin = "0x6080604052600180546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f09eda238d4d37e8ac546190722df083c57c5adeae8a32c990781b796ca0886c060025534801561005557600080fd5b506040516118783803806118788339818101604052602081101561007857600080fd5b5051600081816001600160a01b0382166100c35760405162461bcd60e51b81526004018080602001828103825260238152602001806118556023913960400191505060405180910390fd5b6001600160a01b03821630141561010b5760405162461bcd60e51b81526004018080602001828103825260248152602001806118316024913960400191505060405180910390fd5b60008054821515600160a01b810260ff60a01b196001600160a01b0387166001600160a01b031990941684171617835560408051938452602084019290925282820152517fbb069dd89e950c9b90877e32a7a0243158679d54b7d1cfc9ae316b6d10c0f8a59181900360600190a1505080156101ba5760028190556040805182815290517f7e458cfd14cd2f994e754780400d40cebeea78297d9417376df13c79690dbc8e9181900360200190a15b5050306000908152600760205260409020805460ff19166001908117909155600880549091019055611640806101f16000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806370480275116100b85780638da5cb5b1161007c5780638da5cb5b14610320578063b242e53414610328578063b429afeb14610356578063be9a65551461037c578063e0f2d61b14610384578063eb12d61e146103aa57610137565b806370480275146102a0578063715018a6146102c65780637ca548c6146102ce5780637d73b231146102d65780637df73e27146102fa57610137565b80632121dc75116100ff5780632121dc751461023157806324d7806c1461024d5780632a5831cd146102735780632b7832b3146102905780633f683b6a1461029857610137565b806306a41d091461013c57806307da68f5146101c35780630e316ab7146101cb57806315b9a8b8146101f15780631785f53c1461020b575b600080fd5b6101c16004803603608081101561015257600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561018257600080fd5b82018360208201111561019457600080fd5b803590602001918460018302840111640100000000831117156101b657600080fd5b9193509150356103d0565b005b6101c1610637565b6101c1600480360360208110156101e157600080fd5b50356001600160a01b03166106ed565b6101f961076d565b60408051918252519081900360200190f35b6101c16004803603602081101561022157600080fd5b50356001600160a01b0316610773565b6102396107c4565b604080519115158252519081900360200190f35b6102396004803603602081101561026357600080fd5b50356001600160a01b03166107d4565b6101c16004803603602081101561028957600080fd5b50356107f2565b6101f9610843565b610239610849565b6101c1600480360360208110156102b657600080fd5b50356001600160a01b0316610852565b6101c16108a3565b6101f96109b4565b6102de6109ba565b604080516001600160a01b039092168252519081900360200190f35b6102396004803603602081101561031057600080fd5b50356001600160a01b03166109c9565b6102de6109e7565b6101c16004803603604081101561033e57600080fd5b506001600160a01b03813516906020013515156109f6565b6102396004803603602081101561036c57600080fd5b50356001600160a01b0316610bb6565b6101c1610bd4565b6101c16004803603602081101561039a57600080fd5b50356001600160a01b0316610c5b565b6101c1600480360360208110156103c057600080fd5b50356001600160a01b0316610cac565b3360009081526005602052604090205460ff1661042d576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba10309039b4b3b732b960511b604482015290519081900360640190fd5b60095460ff161561047d576040805162461bcd60e51b815260206004820152601560248201527418dbdb9d1c9bdb1b195c881a5cc81cdd1bdc1c1959605a1b604482015290519081900360640190fd5b801561048c5761048c81610d29565b60006060866001600160a01b0316868686604051808383808284376040519201945060009350909150508083038185875af1925050503d80600081146104ee576040519150601f19603f3d011682016040523d82523d6000602084013e6104f3565b606091505b509150915081610541576040805162461bcd60e51b8152602060048201526014602482015273195e1d195c9b985b0818d85b1b0819985a5b195960621b604482015290519081900360640190fd5b7ff77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613878787878560405180866001600160a01b03166001600160a01b03168152602001858152602001806020018060200183810383528686828181526020019250808284376000838201819052601f909101601f191690920185810384528651815286516020918201939188019250908190849084905b838110156105ef5781810151838201526020016105d7565b50505050905090810190601f16801561061c5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150505050505050565b61064033610da8565b8061065a57503360009081526003602052604090205460ff165b6106ab576040805162461bcd60e51b815260206004820152601c60248201527f73656e646572206973206e6f742061646d696e206f72206f776e657200000000604482015290519081900360640190fd5b6009805460ff191660011790556040805133815290517f55c4adf1f68f084b809304657594a92ba835ada8d3b5340955bf05746723c05b9181900360200190a1565b6106f633610da8565b8061071057503360009081526003602052604090205460ff165b610761576040805162461bcd60e51b815260206004820152601c60248201527f73656e646572206973206e6f742061646d696e206f72206f776e657200000000604482015290519081900360640190fd5b61076a81610dbc565b50565b60085490565b61077c33610da8565b6107bb576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b61076a81610e92565b600054600160a01b900460ff1690565b6001600160a01b031660009081526003602052604090205460ff1690565b6107fb33610da8565b61083a576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b61076a81610f68565b60045490565b60095460ff1690565b61085b33610da8565b61089a576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b61076a81610fa3565b6108ac33610da8565b6108eb576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b600054600160a01b900460ff16610949576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6000805460ff60a01b1981168255604080516001600160a01b039092168252602082018390528181019290925290517fbb069dd89e950c9b90877e32a7a0243158679d54b7d1cfc9ae316b6d10c0f8a59181900360600190a1600080546001600160a01b0319169055565b60065490565b6001546001600160a01b031690565b6001600160a01b031660009081526005602052604090205460ff1690565b6000546001600160a01b031690565b6109ff33610da8565b610a3e576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b600054600160a01b900460ff16610a9c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b038216610ae15760405162461bcd60e51b81526004018080602001828103825260238152602001806115796023913960400191505060405180910390fd5b6001600160a01b038216301415610b295760405162461bcd60e51b81526004018080602001828103825260248152602001806114cb6024913960400191505060405180910390fd5b60008054821515600160a01b810260ff60a01b199092169190911791829055604080516001600160a01b039384168152928516602084015282810191909152517fbb069dd89e950c9b90877e32a7a0243158679d54b7d1cfc9ae316b6d10c0f8a59181900360600190a150600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b031660009081526007602052604090205460ff1690565b610bdd33610da8565b610c1c576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b6009805460ff191690556040805133815290517f27029695aa5f602a4ee81f4c32dfa86e562f200a17966496f3a7c3f2ec0f94179181900360200190a1565b610c6433610da8565b610ca3576040805162461bcd60e51b81526020600482015260166024820152600080516020611538833981519152604482015290519081900360640190fd5b61076a81611193565b610cb533610da8565b80610ccf57503360009081526003602052604090205460ff165b610d20576040805162461bcd60e51b815260206004820152601c60248201527f73656e646572206973206e6f742061646d696e206f72206f776e657200000000604482015290519081900360640190fd5b61076a816111e7565b610d346002546113d7565b6001600160a01b031663d8ccd0f3826040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015610d7957600080fd5b505af1158015610d8d573d6000803e3d6000fd5b505050506040513d6020811015610da357600080fd5b505050565b6000546001600160a01b0390811691161490565b6001600160a01b03811660009081526005602052604090205460ff16610e29576040805162461bcd60e51b815260206004820181905260248201527f70726f7669646564206163636f756e74206973206e6f742061207369676e6572604482015290519081900360640190fd5b6001600160a01b038116600081815260056020908152604091829020805460ff191690556006805460001901905581513381529081019290925280517fe821e394d9f830addcff910c44eede8797692425b5bd9b0cd8aafe4009ddce8e9281900390910190a150565b6001600160a01b03811660009081526003602052604090205460ff16610eff576040805162461bcd60e51b815260206004820181905260248201527f70726f7669646564206163636f756e74206973206e6f7420616e2061646d696e604482015290519081900360640190fd5b6001600160a01b038116600081815260036020908152604091829020805460ff191690556004805460001901905581513381529081019290925280517f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e9281900390910190a150565b60028190556040805182815290517f7e458cfd14cd2f994e754780400d40cebeea78297d9417376df13c79690dbc8e9181900360200190a150565b6001600160a01b03811660009081526003602052604090205460ff1615610ffb5760405162461bcd60e51b81526004018080602001828103825260248152602001806115146024913960400191505060405180910390fd5b6001600160a01b03811660009081526005602052604090205460ff16156110535760405162461bcd60e51b81526004018080602001828103825260288152602001806115e46028913960400191505060405180910390fd5b61105c81610da8565b156110985760405162461bcd60e51b81526004018080602001828103825260258152602001806114ef6025913960400191505060405180910390fd5b6001600160a01b0381166110dd5760405162461bcd60e51b81526004018080602001828103825260248152602001806115c06024913960400191505060405180910390fd5b6001600160a01b0381163014156111255760405162461bcd60e51b81526004018080602001828103825260218152602001806115586021913960400191505060405180910390fd5b6001600160a01b038116600081815260036020908152604091829020805460ff1916600190811790915560048054909101905581513381529081019290925280517fc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a9281900390910190a150565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda907739181900360200190a150565b6001600160a01b03811660009081526003602052604090205460ff161561123f5760405162461bcd60e51b81526004018080602001828103825260248152602001806115146024913960400191505060405180910390fd5b6001600160a01b03811660009081526005602052604090205460ff16156112975760405162461bcd60e51b815260040180806020018281038252602481526020018061159c6024913960400191505060405180910390fd5b6112a081610da8565b156112dc5760405162461bcd60e51b81526004018080602001828103825260258152602001806114ef6025913960400191505060405180910390fd5b6001600160a01b0381166113215760405162461bcd60e51b81526004018080602001828103825260248152602001806115c06024913960400191505060405180910390fd5b6001600160a01b0381163014156113695760405162461bcd60e51b81526004018080602001828103825260218152602001806115586021913960400191505060405180910390fd5b6001600160a01b038116600081815260056020908152604091829020805460ff1916600190811790915560068054909101905581513381529081019290925280517f05c1409cc15c342fda5d547189b22d4c0712ed35dc687a8ab325eaf697f01eb79281900390910190a150565b60015460408051630178b8bf60e01b81526004810184905290516000926001600160a01b031691630178b8bf916024808301926020929190829003018186803b15801561142357600080fd5b505afa158015611437573d6000803e3d6000fd5b505050506040513d602081101561144d57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561149857600080fd5b505afa1580156114ac573d6000803e3d6000fd5b505050506040513d60208110156114c257600080fd5b50519291505056fe6f776e65722063616e6e6f742062652073657420746f207468697320636f6e747261637470726f7669646564206163636f756e7420697320616c726561647920746865206f776e657270726f7669646564206163636f756e7420697320616c726561647920616e2061646d696e73656e646572206973206e6f7420616e206f776e65720000000000000000000070726f7669646564206163636f756e74206973207468697320636f6e74726163746f776e65722063616e6e6f742062652073657420746f207a65726f206164647265737370726f7669646564206163636f756e7420697320616c72656164792061207369676e657270726f7669646564206163636f756e7420697320746865207a65726f206164647265737370726f7669646564206163636f756e7420697320616c7265616479206120636f6e74726f6c6c6572a265627a7a7231582083926dd2ee49a9f478b1ca915f1479ec6390838f36b41cb5d4e70534e0bf867564736f6c634300051100326f776e65722063616e6e6f742062652073657420746f207468697320636f6e74726163746f776e65722063616e6e6f742062652073657420746f207a65726f2061646472657373"

// DeployController deploys a new Ethereum contract, binding an instance of Controller to it.
func DeployController(auth *bind.TransactOpts, backend bind.ContractBackend, _ownerAddress common.Address) (common.Address, *types.Transaction, *Controller, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControllerBin), backend, _ownerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_Controller *ControllerCaller) AdminCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "adminCount")
	return *ret0, err
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_Controller *ControllerSession) AdminCount() (*big.Int, error) {
	return _Controller.Contract.AdminCount(&_Controller.CallOpts)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_Controller *ControllerCallerSession) AdminCount() (*big.Int, error) {
	return _Controller.Contract.AdminCount(&_Controller.CallOpts)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Controller *ControllerCaller) ControllerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "controllerCount")
	return *ret0, err
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Controller *ControllerSession) ControllerCount() (*big.Int, error) {
	return _Controller.Contract.ControllerCount(&_Controller.CallOpts)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Controller *ControllerCallerSession) ControllerCount() (*big.Int, error) {
	return _Controller.Contract.ControllerCount(&_Controller.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Controller *ControllerCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Controller *ControllerSession) EnsRegistry() (common.Address, error) {
	return _Controller.Contract.EnsRegistry(&_Controller.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_Controller *ControllerCallerSession) EnsRegistry() (common.Address, error) {
	return _Controller.Contract.EnsRegistry(&_Controller.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) constant returns(bool)
func (_Controller *ControllerCaller) IsAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isAdmin", _account)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) constant returns(bool)
func (_Controller *ControllerSession) IsAdmin(_account common.Address) (bool, error) {
	return _Controller.Contract.IsAdmin(&_Controller.CallOpts, _account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) constant returns(bool)
func (_Controller *ControllerCallerSession) IsAdmin(_account common.Address) (bool, error) {
	return _Controller.Contract.IsAdmin(&_Controller.CallOpts, _account)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController(address _account) constant returns(bool)
func (_Controller *ControllerCaller) IsController(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isController", _account)
	return *ret0, err
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController(address _account) constant returns(bool)
func (_Controller *ControllerSession) IsController(_account common.Address) (bool, error) {
	return _Controller.Contract.IsController(&_Controller.CallOpts, _account)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController(address _account) constant returns(bool)
func (_Controller *ControllerCallerSession) IsController(_account common.Address) (bool, error) {
	return _Controller.Contract.IsController(&_Controller.CallOpts, _account)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _account) constant returns(bool)
func (_Controller *ControllerCaller) IsSigner(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isSigner", _account)
	return *ret0, err
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _account) constant returns(bool)
func (_Controller *ControllerSession) IsSigner(_account common.Address) (bool, error) {
	return _Controller.Contract.IsSigner(&_Controller.CallOpts, _account)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _account) constant returns(bool)
func (_Controller *ControllerCallerSession) IsSigner(_account common.Address) (bool, error) {
	return _Controller.Contract.IsSigner(&_Controller.CallOpts, _account)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() constant returns(bool)
func (_Controller *ControllerCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isStopped")
	return *ret0, err
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() constant returns(bool)
func (_Controller *ControllerSession) IsStopped() (bool, error) {
	return _Controller.Contract.IsStopped(&_Controller.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() constant returns(bool)
func (_Controller *ControllerCallerSession) IsStopped() (bool, error) {
	return _Controller.Contract.IsStopped(&_Controller.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Controller *ControllerCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Controller *ControllerSession) IsTransferable() (bool, error) {
	return _Controller.Contract.IsTransferable(&_Controller.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Controller *ControllerCallerSession) IsTransferable() (bool, error) {
	return _Controller.Contract.IsTransferable(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// SignerCount is a free data retrieval call binding the contract method 0x7ca548c6.
//
// Solidity: function signerCount() constant returns(uint256)
func (_Controller *ControllerCaller) SignerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "signerCount")
	return *ret0, err
}

// SignerCount is a free data retrieval call binding the contract method 0x7ca548c6.
//
// Solidity: function signerCount() constant returns(uint256)
func (_Controller *ControllerSession) SignerCount() (*big.Int, error) {
	return _Controller.Contract.SignerCount(&_Controller.CallOpts)
}

// SignerCount is a free data retrieval call binding the contract method 0x7ca548c6.
//
// Solidity: function signerCount() constant returns(uint256)
func (_Controller *ControllerCallerSession) SignerCount() (*big.Int, error) {
	return _Controller.Contract.SignerCount(&_Controller.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_Controller *ControllerTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_Controller *ControllerSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddAdmin(&_Controller.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_Controller *ControllerTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddAdmin(&_Controller.TransactOpts, _account)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _account) returns()
func (_Controller *ControllerTransactor) AddSigner(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addSigner", _account)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _account) returns()
func (_Controller *ControllerSession) AddSigner(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddSigner(&_Controller.TransactOpts, _account)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _account) returns()
func (_Controller *ControllerTransactorSession) AddSigner(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddSigner(&_Controller.TransactOpts, _account)
}

// EnsSetRegistry is a paid mutator transaction binding the contract method 0xe0f2d61b.
//
// Solidity: function ensSetRegistry(address _ensRegistry) returns()
func (_Controller *ControllerTransactor) EnsSetRegistry(opts *bind.TransactOpts, _ensRegistry common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "ensSetRegistry", _ensRegistry)
}

// EnsSetRegistry is a paid mutator transaction binding the contract method 0xe0f2d61b.
//
// Solidity: function ensSetRegistry(address _ensRegistry) returns()
func (_Controller *ControllerSession) EnsSetRegistry(_ensRegistry common.Address) (*types.Transaction, error) {
	return _Controller.Contract.EnsSetRegistry(&_Controller.TransactOpts, _ensRegistry)
}

// EnsSetRegistry is a paid mutator transaction binding the contract method 0xe0f2d61b.
//
// Solidity: function ensSetRegistry(address _ensRegistry) returns()
func (_Controller *ControllerTransactorSession) EnsSetRegistry(_ensRegistry common.Address) (*types.Transaction, error) {
	return _Controller.Contract.EnsSetRegistry(&_Controller.TransactOpts, _ensRegistry)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x06a41d09.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data, uint256 _gasRefund) returns()
func (_Controller *ControllerTransactor) ExecuteTransaction(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte, _gasRefund *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "executeTransaction", _destination, _value, _data, _gasRefund)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x06a41d09.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data, uint256 _gasRefund) returns()
func (_Controller *ControllerSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte, _gasRefund *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.ExecuteTransaction(&_Controller.TransactOpts, _destination, _value, _data, _gasRefund)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x06a41d09.
//
// Solidity: function executeTransaction(address _destination, uint256 _value, bytes _data, uint256 _gasRefund) returns()
func (_Controller *ControllerTransactorSession) ExecuteTransaction(_destination common.Address, _value *big.Int, _data []byte, _gasRefund *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.ExecuteTransaction(&_Controller.TransactOpts, _destination, _value, _data, _gasRefund)
}

// GstSetNode is a paid mutator transaction binding the contract method 0x2a5831cd.
//
// Solidity: function gstSetNode(bytes32 _gstNode) returns()
func (_Controller *ControllerTransactor) GstSetNode(opts *bind.TransactOpts, _gstNode [32]byte) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "gstSetNode", _gstNode)
}

// GstSetNode is a paid mutator transaction binding the contract method 0x2a5831cd.
//
// Solidity: function gstSetNode(bytes32 _gstNode) returns()
func (_Controller *ControllerSession) GstSetNode(_gstNode [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.GstSetNode(&_Controller.TransactOpts, _gstNode)
}

// GstSetNode is a paid mutator transaction binding the contract method 0x2a5831cd.
//
// Solidity: function gstSetNode(bytes32 _gstNode) returns()
func (_Controller *ControllerTransactorSession) GstSetNode(_gstNode [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.GstSetNode(&_Controller.TransactOpts, _gstNode)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_Controller *ControllerTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_Controller *ControllerSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveAdmin(&_Controller.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_Controller *ControllerTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveAdmin(&_Controller.TransactOpts, _account)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address _account) returns()
func (_Controller *ControllerTransactor) RemoveSigner(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeSigner", _account)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address _account) returns()
func (_Controller *ControllerSession) RemoveSigner(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveSigner(&_Controller.TransactOpts, _account)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address _account) returns()
func (_Controller *ControllerTransactorSession) RemoveSigner(_account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveSigner(&_Controller.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Controller *ControllerTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Controller *ControllerSession) Start() (*types.Transaction, error) {
	return _Controller.Contract.Start(&_Controller.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Controller *ControllerTransactorSession) Start() (*types.Transaction, error) {
	return _Controller.Contract.Start(&_Controller.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Controller *ControllerTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Controller *ControllerSession) Stop() (*types.Transaction, error) {
	return _Controller.Contract.Stop(&_Controller.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Controller *ControllerTransactorSession) Stop() (*types.Transaction, error) {
	return _Controller.Contract.Stop(&_Controller.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Controller *ControllerSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Controller *ControllerTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, _account, _transferable)
}

// ControllerAddedAdminIterator is returned from FilterAddedAdmin and is used to iterate over the raw logs and unpacked data for AddedAdmin events raised by the Controller contract.
type ControllerAddedAdminIterator struct {
	Event *ControllerAddedAdmin // Event containing the contract specifics and raw log

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
func (it *ControllerAddedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerAddedAdmin)
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
		it.Event = new(ControllerAddedAdmin)
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
func (it *ControllerAddedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerAddedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerAddedAdmin represents a AddedAdmin event raised by the Controller contract.
type ControllerAddedAdmin struct {
	Sender common.Address
	Admin  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedAdmin is a free log retrieval operation binding the contract event 0xc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a.
//
// Solidity: event AddedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) FilterAddedAdmin(opts *bind.FilterOpts) (*ControllerAddedAdminIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "AddedAdmin")
	if err != nil {
		return nil, err
	}
	return &ControllerAddedAdminIterator{contract: _Controller.contract, event: "AddedAdmin", logs: logs, sub: sub}, nil
}

// WatchAddedAdmin is a free log subscription operation binding the contract event 0xc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a.
//
// Solidity: event AddedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) WatchAddedAdmin(opts *bind.WatchOpts, sink chan<- *ControllerAddedAdmin) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "AddedAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerAddedAdmin)
				if err := _Controller.contract.UnpackLog(event, "AddedAdmin", log); err != nil {
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

// ParseAddedAdmin is a log parse operation binding the contract event 0xc58b647b8ba5a8cab2f11f32673636cc1061324240972ed05e8cc005b81a4b7a.
//
// Solidity: event AddedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) ParseAddedAdmin(log types.Log) (*ControllerAddedAdmin, error) {
	event := new(ControllerAddedAdmin)
	if err := _Controller.contract.UnpackLog(event, "AddedAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerAddedSignerIterator is returned from FilterAddedSigner and is used to iterate over the raw logs and unpacked data for AddedSigner events raised by the Controller contract.
type ControllerAddedSignerIterator struct {
	Event *ControllerAddedSigner // Event containing the contract specifics and raw log

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
func (it *ControllerAddedSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerAddedSigner)
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
		it.Event = new(ControllerAddedSigner)
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
func (it *ControllerAddedSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerAddedSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerAddedSigner represents a AddedSigner event raised by the Controller contract.
type ControllerAddedSigner struct {
	Sender common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedSigner is a free log retrieval operation binding the contract event 0x05c1409cc15c342fda5d547189b22d4c0712ed35dc687a8ab325eaf697f01eb7.
//
// Solidity: event AddedSigner(address _sender, address _signer)
func (_Controller *ControllerFilterer) FilterAddedSigner(opts *bind.FilterOpts) (*ControllerAddedSignerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "AddedSigner")
	if err != nil {
		return nil, err
	}
	return &ControllerAddedSignerIterator{contract: _Controller.contract, event: "AddedSigner", logs: logs, sub: sub}, nil
}

// WatchAddedSigner is a free log subscription operation binding the contract event 0x05c1409cc15c342fda5d547189b22d4c0712ed35dc687a8ab325eaf697f01eb7.
//
// Solidity: event AddedSigner(address _sender, address _signer)
func (_Controller *ControllerFilterer) WatchAddedSigner(opts *bind.WatchOpts, sink chan<- *ControllerAddedSigner) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "AddedSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerAddedSigner)
				if err := _Controller.contract.UnpackLog(event, "AddedSigner", log); err != nil {
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

// ParseAddedSigner is a log parse operation binding the contract event 0x05c1409cc15c342fda5d547189b22d4c0712ed35dc687a8ab325eaf697f01eb7.
//
// Solidity: event AddedSigner(address _sender, address _signer)
func (_Controller *ControllerFilterer) ParseAddedSigner(log types.Log) (*ControllerAddedSigner, error) {
	event := new(ControllerAddedSigner)
	if err := _Controller.contract.UnpackLog(event, "AddedSigner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerENSSetRegistryIterator is returned from FilterENSSetRegistry and is used to iterate over the raw logs and unpacked data for ENSSetRegistry events raised by the Controller contract.
type ControllerENSSetRegistryIterator struct {
	Event *ControllerENSSetRegistry // Event containing the contract specifics and raw log

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
func (it *ControllerENSSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerENSSetRegistry)
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
		it.Event = new(ControllerENSSetRegistry)
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
func (it *ControllerENSSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerENSSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerENSSetRegistry represents a ENSSetRegistry event raised by the Controller contract.
type ControllerENSSetRegistry struct {
	EnsRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterENSSetRegistry is a free log retrieval operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_Controller *ControllerFilterer) FilterENSSetRegistry(opts *bind.FilterOpts) (*ControllerENSSetRegistryIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return &ControllerENSSetRegistryIterator{contract: _Controller.contract, event: "ENSSetRegistry", logs: logs, sub: sub}, nil
}

// WatchENSSetRegistry is a free log subscription operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_Controller *ControllerFilterer) WatchENSSetRegistry(opts *bind.WatchOpts, sink chan<- *ControllerENSSetRegistry) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerENSSetRegistry)
				if err := _Controller.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
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

// ParseENSSetRegistry is a log parse operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_Controller *ControllerFilterer) ParseENSSetRegistry(log types.Log) (*ControllerENSSetRegistry, error) {
	event := new(ControllerENSSetRegistry)
	if err := _Controller.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerExecutedTransactionIterator is returned from FilterExecutedTransaction and is used to iterate over the raw logs and unpacked data for ExecutedTransaction events raised by the Controller contract.
type ControllerExecutedTransactionIterator struct {
	Event *ControllerExecutedTransaction // Event containing the contract specifics and raw log

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
func (it *ControllerExecutedTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerExecutedTransaction)
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
		it.Event = new(ControllerExecutedTransaction)
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
func (it *ControllerExecutedTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerExecutedTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerExecutedTransaction represents a ExecutedTransaction event raised by the Controller contract.
type ControllerExecutedTransaction struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	ReturnData  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedTransaction is a free log retrieval operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_Controller *ControllerFilterer) FilterExecutedTransaction(opts *bind.FilterOpts) (*ControllerExecutedTransactionIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return &ControllerExecutedTransactionIterator{contract: _Controller.contract, event: "ExecutedTransaction", logs: logs, sub: sub}, nil
}

// WatchExecutedTransaction is a free log subscription operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_Controller *ControllerFilterer) WatchExecutedTransaction(opts *bind.WatchOpts, sink chan<- *ControllerExecutedTransaction) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "ExecutedTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerExecutedTransaction)
				if err := _Controller.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
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

// ParseExecutedTransaction is a log parse operation binding the contract event 0xf77753fab406ecfff96d6ff2476c64a838fa9f6d37b1bf190f8546e395e3b613.
//
// Solidity: event ExecutedTransaction(address _destination, uint256 _value, bytes _data, bytes _returnData)
func (_Controller *ControllerFilterer) ParseExecutedTransaction(log types.Log) (*ControllerExecutedTransaction, error) {
	event := new(ControllerExecutedTransaction)
	if err := _Controller.contract.UnpackLog(event, "ExecutedTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerGSTSetNodeIterator is returned from FilterGSTSetNode and is used to iterate over the raw logs and unpacked data for GSTSetNode events raised by the Controller contract.
type ControllerGSTSetNodeIterator struct {
	Event *ControllerGSTSetNode // Event containing the contract specifics and raw log

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
func (it *ControllerGSTSetNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerGSTSetNode)
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
		it.Event = new(ControllerGSTSetNode)
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
func (it *ControllerGSTSetNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerGSTSetNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerGSTSetNode represents a GSTSetNode event raised by the Controller contract.
type ControllerGSTSetNode struct {
	GstNode [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGSTSetNode is a free log retrieval operation binding the contract event 0x7e458cfd14cd2f994e754780400d40cebeea78297d9417376df13c79690dbc8e.
//
// Solidity: event GSTSetNode(bytes32 _gstNode)
func (_Controller *ControllerFilterer) FilterGSTSetNode(opts *bind.FilterOpts) (*ControllerGSTSetNodeIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "GSTSetNode")
	if err != nil {
		return nil, err
	}
	return &ControllerGSTSetNodeIterator{contract: _Controller.contract, event: "GSTSetNode", logs: logs, sub: sub}, nil
}

// WatchGSTSetNode is a free log subscription operation binding the contract event 0x7e458cfd14cd2f994e754780400d40cebeea78297d9417376df13c79690dbc8e.
//
// Solidity: event GSTSetNode(bytes32 _gstNode)
func (_Controller *ControllerFilterer) WatchGSTSetNode(opts *bind.WatchOpts, sink chan<- *ControllerGSTSetNode) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "GSTSetNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerGSTSetNode)
				if err := _Controller.contract.UnpackLog(event, "GSTSetNode", log); err != nil {
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

// ParseGSTSetNode is a log parse operation binding the contract event 0x7e458cfd14cd2f994e754780400d40cebeea78297d9417376df13c79690dbc8e.
//
// Solidity: event GSTSetNode(bytes32 _gstNode)
func (_Controller *ControllerFilterer) ParseGSTSetNode(log types.Log) (*ControllerGSTSetNode, error) {
	event := new(ControllerGSTSetNode)
	if err := _Controller.contract.UnpackLog(event, "GSTSetNode", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the Controller contract.
type ControllerRemovedAdminIterator struct {
	Event *ControllerRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *ControllerRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRemovedAdmin)
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
		it.Event = new(ControllerRemovedAdmin)
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
func (it *ControllerRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRemovedAdmin represents a RemovedAdmin event raised by the Controller contract.
type ControllerRemovedAdmin struct {
	Sender common.Address
	Admin  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) FilterRemovedAdmin(opts *bind.FilterOpts) (*ControllerRemovedAdminIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RemovedAdmin")
	if err != nil {
		return nil, err
	}
	return &ControllerRemovedAdminIterator{contract: _Controller.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *ControllerRemovedAdmin) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RemovedAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRemovedAdmin)
				if err := _Controller.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address _sender, address _admin)
func (_Controller *ControllerFilterer) ParseRemovedAdmin(log types.Log) (*ControllerRemovedAdmin, error) {
	event := new(ControllerRemovedAdmin)
	if err := _Controller.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerRemovedSignerIterator is returned from FilterRemovedSigner and is used to iterate over the raw logs and unpacked data for RemovedSigner events raised by the Controller contract.
type ControllerRemovedSignerIterator struct {
	Event *ControllerRemovedSigner // Event containing the contract specifics and raw log

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
func (it *ControllerRemovedSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRemovedSigner)
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
		it.Event = new(ControllerRemovedSigner)
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
func (it *ControllerRemovedSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRemovedSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRemovedSigner represents a RemovedSigner event raised by the Controller contract.
type ControllerRemovedSigner struct {
	Sender common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedSigner is a free log retrieval operation binding the contract event 0xe821e394d9f830addcff910c44eede8797692425b5bd9b0cd8aafe4009ddce8e.
//
// Solidity: event RemovedSigner(address _sender, address _signer)
func (_Controller *ControllerFilterer) FilterRemovedSigner(opts *bind.FilterOpts) (*ControllerRemovedSignerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RemovedSigner")
	if err != nil {
		return nil, err
	}
	return &ControllerRemovedSignerIterator{contract: _Controller.contract, event: "RemovedSigner", logs: logs, sub: sub}, nil
}

// WatchRemovedSigner is a free log subscription operation binding the contract event 0xe821e394d9f830addcff910c44eede8797692425b5bd9b0cd8aafe4009ddce8e.
//
// Solidity: event RemovedSigner(address _sender, address _signer)
func (_Controller *ControllerFilterer) WatchRemovedSigner(opts *bind.WatchOpts, sink chan<- *ControllerRemovedSigner) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RemovedSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRemovedSigner)
				if err := _Controller.contract.UnpackLog(event, "RemovedSigner", log); err != nil {
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

// ParseRemovedSigner is a log parse operation binding the contract event 0xe821e394d9f830addcff910c44eede8797692425b5bd9b0cd8aafe4009ddce8e.
//
// Solidity: event RemovedSigner(address _sender, address _signer)
func (_Controller *ControllerFilterer) ParseRemovedSigner(log types.Log) (*ControllerRemovedSigner, error) {
	event := new(ControllerRemovedSigner)
	if err := _Controller.contract.UnpackLog(event, "RemovedSigner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerStartedIterator is returned from FilterStarted and is used to iterate over the raw logs and unpacked data for Started events raised by the Controller contract.
type ControllerStartedIterator struct {
	Event *ControllerStarted // Event containing the contract specifics and raw log

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
func (it *ControllerStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerStarted)
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
		it.Event = new(ControllerStarted)
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
func (it *ControllerStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerStarted represents a Started event raised by the Controller contract.
type ControllerStarted struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStarted is a free log retrieval operation binding the contract event 0x27029695aa5f602a4ee81f4c32dfa86e562f200a17966496f3a7c3f2ec0f9417.
//
// Solidity: event Started(address _sender)
func (_Controller *ControllerFilterer) FilterStarted(opts *bind.FilterOpts) (*ControllerStartedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Started")
	if err != nil {
		return nil, err
	}
	return &ControllerStartedIterator{contract: _Controller.contract, event: "Started", logs: logs, sub: sub}, nil
}

// WatchStarted is a free log subscription operation binding the contract event 0x27029695aa5f602a4ee81f4c32dfa86e562f200a17966496f3a7c3f2ec0f9417.
//
// Solidity: event Started(address _sender)
func (_Controller *ControllerFilterer) WatchStarted(opts *bind.WatchOpts, sink chan<- *ControllerStarted) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Started")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerStarted)
				if err := _Controller.contract.UnpackLog(event, "Started", log); err != nil {
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

// ParseStarted is a log parse operation binding the contract event 0x27029695aa5f602a4ee81f4c32dfa86e562f200a17966496f3a7c3f2ec0f9417.
//
// Solidity: event Started(address _sender)
func (_Controller *ControllerFilterer) ParseStarted(log types.Log) (*ControllerStarted, error) {
	event := new(ControllerStarted)
	if err := _Controller.contract.UnpackLog(event, "Started", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the Controller contract.
type ControllerStoppedIterator struct {
	Event *ControllerStopped // Event containing the contract specifics and raw log

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
func (it *ControllerStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerStopped)
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
		it.Event = new(ControllerStopped)
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
func (it *ControllerStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerStopped represents a Stopped event raised by the Controller contract.
type ControllerStopped struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x55c4adf1f68f084b809304657594a92ba835ada8d3b5340955bf05746723c05b.
//
// Solidity: event Stopped(address _sender)
func (_Controller *ControllerFilterer) FilterStopped(opts *bind.FilterOpts) (*ControllerStoppedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &ControllerStoppedIterator{contract: _Controller.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x55c4adf1f68f084b809304657594a92ba835ada8d3b5340955bf05746723c05b.
//
// Solidity: event Stopped(address _sender)
func (_Controller *ControllerFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *ControllerStopped) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerStopped)
				if err := _Controller.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x55c4adf1f68f084b809304657594a92ba835ada8d3b5340955bf05746723c05b.
//
// Solidity: event Stopped(address _sender)
func (_Controller *ControllerFilterer) ParseStopped(log types.Log) (*ControllerStopped, error) {
	event := new(ControllerStopped)
	if err := _Controller.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ControllerTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Controller contract.
type ControllerTransferredOwnershipIterator struct {
	Event *ControllerTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *ControllerTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTransferredOwnership)
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
		it.Event = new(ControllerTransferredOwnership)
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
func (it *ControllerTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTransferredOwnership represents a TransferredOwnership event raised by the Controller contract.
type ControllerTransferredOwnership struct {
	From         common.Address
	To           common.Address
	Transferable bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0xbb069dd89e950c9b90877e32a7a0243158679d54b7d1cfc9ae316b6d10c0f8a5.
//
// Solidity: event TransferredOwnership(address _from, address _to, bool _transferable)
func (_Controller *ControllerFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*ControllerTransferredOwnershipIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &ControllerTransferredOwnershipIterator{contract: _Controller.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0xbb069dd89e950c9b90877e32a7a0243158679d54b7d1cfc9ae316b6d10c0f8a5.
//
// Solidity: event TransferredOwnership(address _from, address _to, bool _transferable)
func (_Controller *ControllerFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *ControllerTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTransferredOwnership)
				if err := _Controller.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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

// ParseTransferredOwnership is a log parse operation binding the contract event 0xbb069dd89e950c9b90877e32a7a0243158679d54b7d1cfc9ae316b6d10c0f8a5.
//
// Solidity: event TransferredOwnership(address _from, address _to, bool _transferable)
func (_Controller *ControllerFilterer) ParseTransferredOwnership(log types.Log) (*ControllerTransferredOwnership, error) {
	event := new(ControllerTransferredOwnership)
	if err := _Controller.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
		return nil, err
	}
	return event, nil
}
