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

// LicenceABI is the input ABI used to generate the binding from.
const LicenceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"lockTKNContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTKN\",\"type\":\"address\"}],\"name\":\"updateTKNContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"load\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceDAO\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newHolder\",\"type\":\"address\"}],\"name\":\"updateHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockLicenceDAO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAmount\",\"type\":\"uint256\"}],\"name\":\"updateLicenceAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tknContractAddressLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"floatLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tknContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFloat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceAmountScaled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_AMOUNT_SCALE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockFloat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFloat\",\"type\":\"address\"}],\"name\":\"updateFloat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licenceDAOLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDAO\",\"type\":\"address\"}],\"name\":\"updateLicenceDAO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holderLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner_\",\"type\":\"address\"},{\"name\":\"_transferable_\",\"type\":\"bool\"},{\"name\":\"_licence_\",\"type\":\"uint256\"},{\"name\":\"_float_\",\"type\":\"address\"},{\"name\":\"_holder_\",\"type\":\"address\"},{\"name\":\"_tknAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newDAO\",\"type\":\"address\"}],\"name\":\"UpdatedLicenceDAO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newFloat\",\"type\":\"address\"}],\"name\":\"UpdatedCryptoFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newHolder\",\"type\":\"address\"}],\"name\":\"UpdatedTokenHolder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newTKN\",\"type\":\"address\"}],\"name\":\"UpdatedTKNContractAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedLicenceAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferredToTokenHolder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferredToCryptoFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogTokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"}]"

// LicenceBin is the compiled bytecode used for deploying new contracts.
const LicenceBin = `6080604052600180546001600160a01b03191673aaaf91d9b90df800df4f55c205fd6989c977e73a17905534801561003657600080fd5b5060405160c0806118e2833981018060405260c081101561005657600080fd5b508051602082015160408301516060840151608085015160a090950151600080546001600160a01b0319166001600160a01b03871617600160a01b60ff021916740100000000000000000000000000000000000000008615158102919091179182905595969495939492938791879160ff91041661010b57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506005849055600280546001600160a01b038086166001600160a01b0319928316179092556003805485841692169190911790558116156101a757600180546001600160a01b0319166001600160a01b0383161790555b505050505050611726806101bc6000396000f3fe6080604052600436106101665760003560e01c8063940b9c3b116100d1578063ca0e2e201161008a578063d1696b1611610064578063d1696b1614610479578063e30c5fa81461048e578063e3d80242146104a3578063f15ff455146104d657610166565b8063ca0e2e201461041c578063d08b4ecc14610431578063d0cddd671461044657610166565b8063940b9c3b14610338578063996cba681461034d57806399a5e1d014610390578063a036ba60146103a5578063ac904c63146103ba578063b242e534146103e157610166565b806342719faa1161012357806342719faa146102875780634ac22b3c146102ba57806368ce74e7146102cf578063715018a6146102f9578063837c70ef1461030e5780638da5cb5b1461032357610166565b80630bf25c91146101a25780630d42e82f146101b95780631b3c96b4146101ec5780632121dc75146102185780633a7afe0214610241578063420a83e714610272575b6040805133815234602082015281517f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874929181900390910190a1005b3480156101ae57600080fd5b506101b76104eb565b005b3480156101c557600080fd5b506101b7600480360360208110156101dc57600080fd5b50356001600160a01b031661054e565b6101b76004803603604081101561020257600080fd5b506001600160a01b03813516906020013561064a565b34801561022457600080fd5b5061022d61089d565b604080519115158252519081900360200190f35b34801561024d57600080fd5b506102566108ad565b604080516001600160a01b039092168252519081900360200190f35b34801561027e57600080fd5b506102566108bc565b34801561029357600080fd5b506101b7600480360360208110156102aa57600080fd5b50356001600160a01b03166108cb565b3480156102c657600080fd5b506101b76109c7565b3480156102db57600080fd5b506101b7600480360360208110156102f257600080fd5b5035610a2a565b34801561030557600080fd5b506101b7610b2e565b34801561031a57600080fd5b5061022d610c27565b34801561032f57600080fd5b50610256610c37565b34801561034457600080fd5b5061022d610c46565b34801561035957600080fd5b506101b76004803603606081101561037057600080fd5b506001600160a01b03813581169160208101359091169060400135610c56565b34801561039c57600080fd5b50610256610cb1565b3480156103b157600080fd5b50610256610cc0565b3480156103c657600080fd5b506103cf610ccf565b60408051918252519081900360200190f35b3480156103ed57600080fd5b506101b76004803603604081101561040457600080fd5b506001600160a01b0381351690602001351515610cd5565b34801561042857600080fd5b506103cf610e90565b34801561043d57600080fd5b506101b7610e96565b34801561045257600080fd5b506101b76004803603602081101561046957600080fd5b50356001600160a01b0316610ef9565b34801561048557600080fd5b506101b7610ff5565b34801561049a57600080fd5b5061022d611058565b3480156104af57600080fd5b506101b7600480360360208110156104c657600080fd5b50356001600160a01b0316611068565b3480156104e257600080fd5b5061022d611164565b6104f433611174565b6105365760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b60048054600160b81b60ff021916600160b81b179055565b61055733611174565b6105995760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b6105a1610c27565b156105f65760408051600160e51b62461bcd02815260206004820152600d60248201527f544b4e206973206c6f636b656400000000000000000000000000000000000000604482015290519081900360640190fd5b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f2aeed92123e61fe64748a447c2ba122c4bfc0201d1ed5149e9ce9ede5adda5459181900360200190a150565b60015481906001600160a01b038481169116141561068957600254610684906001600160a01b03858116913391168463ffffffff61118816565b610845565b6106b46103e8600554016106a86103e8856111eb90919063ffffffff16565b9063ffffffff61125016565b905060006106c8838363ffffffff6112bd16565b90506001600160a01b03841615610722576003546106fb906001600160a01b03868116913391168463ffffffff61118816565b60025461071d906001600160a01b03868116913391168563ffffffff61118816565b6107f0565b8234146107795760408051600160e51b62461bcd02815260206004820152601f60248201527f4554482073656e74206973206e6f7420657175616c20746f20616d6f756e7400604482015290519081900360640190fd5b6003546040516001600160a01b039091169082156108fc029083906000818181858888f193505050501580156107b3573d6000803e3d6000fd5b506002546040516001600160a01b039091169083156108fc029084906000818181858888f193505050501580156107ee573d6000803e3d6000fd5b505b600354604080513381526001600160a01b0392831660208201529186168282015260608201839052517fdd9dfad7b30d6b224e235f89565871419d3dec3b563a4e231f12d2cc97f9acfc9181900360800190a1505b600254604080513381526001600160a01b0392831660208201529185168282015260608201839052517fc8a7b0bd71097b47b2cad75e4e939d2aeb7fae88110e68f93b83fed08e9d3c389181900360800190a1505050565b600054600160a01b900460ff1690565b6004546001600160a01b031690565b6003546001600160a01b031690565b6108d433611174565b6109165760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b61091e611164565b156109735760408051600160e51b62461bcd02815260206004820152601960248201527f686f6c64657220636f6e7472616374206973206c6f636b656400000000000000604482015290519081900360640190fd5b600380546001600160a01b0383166001600160a01b0319909116811790915560408051918252517ffa6bae0f250db86534a013b1c7a6c4076aa8f8d1ac248771a1c73f4ba366922a9181900360200190a150565b6109d033611174565b610a125760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b60048054600160b01b60ff021916600160b01b179055565b6004546001600160a01b03163314610a8c5760408051600160e51b62461bcd02815260206004820152601860248201527f7468652073656e6465722069736e2774207468652044414f0000000000000000604482015290519081900360640190fd5b80600111158015610a9f57506103e88111155b610af35760408051600160e51b62461bcd02815260206004820152601b60248201527f6c6963656e636520616d6f756e74206f7574206f662072616e67650000000000604482015290519081900360640190fd5b60058190556040805182815290517f587b6068be8c555e2cddc6ad8a56df5e8dfb1533cc063d6703f79c791de151489181900360200190a150565b610b3733611174565b610b795760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b600054600160a01b900460ff16610bda5760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600080546001600160a01b031916815560408051828152602081019290925280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a1565b600454600160b81b900460ff1690565b6000546001600160a01b031690565b600454600160a01b900460ff1690565b610c5f33611174565b610ca15760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b610cac83838361131d565b505050565b6001546001600160a01b031690565b6002546001600160a01b031690565b60055490565b610cde33611174565b610d205760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b600054600160a01b900460ff16610d815760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b038216610dc957604051600160e51b62461bcd0281526004018080602001828103825260238152602001806116896023913960400191505060405180910390fd5b60008054600160a01b60ff021916600160a01b8315150217905580610e2557604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600054604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600080546001600160a01b0319166001600160a01b0392909216919091179055565b6103e881565b610e9f33611174565b610ee15760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b60048054600160a01b60ff021916600160a01b179055565b610f0233611174565b610f445760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b610f4c610c46565b15610fa15760408051600160e51b62461bcd02815260206004820152600f60248201527f666c6f6174206973206c6f636b65640000000000000000000000000000000000604482015290519081900360640190fd5b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f9af2841b0db134bda87280e2a9cababb156f95023c87023d708a677d61b4b6d89181900360200190a150565b610ffe33611174565b6110405760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b60048054600160a81b60ff021916600160a81b179055565b600454600160b01b900460ff1690565b61107133611174565b6110b35760408051600160e51b62461bcd0281526020600482015260166024820152600080516020611648833981519152604482015290519081900360640190fd5b6110bb611058565b156111105760408051600160e51b62461bcd02815260206004820152600d60248201527f44414f206973206c6f636b656400000000000000000000000000000000000000604482015290519081900360640190fd5b600480546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fd32c17b277c7e87842861153d758814a267634f4308ec2461f88756df7dd70689181900360200190a150565b600454600160a81b900460ff1690565b6000546001600160a01b0390811691161490565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b0316600160e01b6323b872dd021790526111e5908590611480565b50505050565b6000826111fa5750600061124a565b8282028284828161120757fe5b041461124757604051600160e51b62461bcd0281526004018080602001828103825260218152602001806116686021913960400191505060405180910390fd5b90505b92915050565b60008082116112a95760408051600160e51b62461bcd02815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b60008284816112b457fe5b04949350505050565b6000828211156113175760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b038216611367576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015611361573d6000803e3d6000fd5b50611431565b816001600160a01b031663a9059cbb84836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156113c757600080fd5b505af11580156113db573d6000803e3d6000fd5b505050506040513d60208110156113f157600080fd5b505161143157604051600160e51b62461bcd0281526004018080602001828103825260258152602001806116ac6025913960400191505060405180910390fd5b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b611492826001600160a01b0316611641565b6114e65760408051600160e51b62461bcd02815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b602083106115245780518252601f199092019160209182019101611505565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611586576040519150601f19603f3d011682016040523d82523d6000602084013e61158b565b606091505b5091509150816115e55760408051600160e51b62461bcd02815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b8051156111e55780806020019051602081101561160157600080fd5b50516111e557604051600160e51b62461bcd02815260040180806020018281038252602a8152602001806116d1602a913960400191505060405180910390fd5b3b15159056fe73656e646572206973206e6f7420616e206f776e657200000000000000000000536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f776e65722063616e6e6f742062652073657420746f207a65726f2061646472657373455243323020746f6b656e207472616e736665722077617320756e7375636365737366756c5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a165627a7a72305820e63a1bdbf285be2d53f20a797b4c3bca0afb6b85effa804b116737a462ffc82c0029`

// DeployLicence deploys a new Ethereum contract, binding an instance of Licence to it.
func DeployLicence(auth *bind.TransactOpts, backend bind.ContractBackend, _owner_ common.Address, _transferable_ bool, _licence_ *big.Int, _float_ common.Address, _holder_ common.Address, _tknAddress_ common.Address) (common.Address, *types.Transaction, *Licence, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LicenceBin), backend, _owner_, _transferable_, _licence_, _float_, _holder_, _tknAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Licence{LicenceCaller: LicenceCaller{contract: contract}, LicenceTransactor: LicenceTransactor{contract: contract}, LicenceFilterer: LicenceFilterer{contract: contract}}, nil
}

// Licence is an auto generated Go binding around an Ethereum contract.
type Licence struct {
	LicenceCaller     // Read-only binding to the contract
	LicenceTransactor // Write-only binding to the contract
	LicenceFilterer   // Log filterer for contract events
}

// LicenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LicenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LicenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LicenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LicenceSession struct {
	Contract     *Licence          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LicenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LicenceCallerSession struct {
	Contract *LicenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LicenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LicenceTransactorSession struct {
	Contract     *LicenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LicenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LicenceRaw struct {
	Contract *Licence // Generic contract binding to access the raw methods on
}

// LicenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LicenceCallerRaw struct {
	Contract *LicenceCaller // Generic read-only contract binding to access the raw methods on
}

// LicenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LicenceTransactorRaw struct {
	Contract *LicenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLicence creates a new instance of Licence, bound to a specific deployed contract.
func NewLicence(address common.Address, backend bind.ContractBackend) (*Licence, error) {
	contract, err := bindLicence(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Licence{LicenceCaller: LicenceCaller{contract: contract}, LicenceTransactor: LicenceTransactor{contract: contract}, LicenceFilterer: LicenceFilterer{contract: contract}}, nil
}

// NewLicenceCaller creates a new read-only instance of Licence, bound to a specific deployed contract.
func NewLicenceCaller(address common.Address, caller bind.ContractCaller) (*LicenceCaller, error) {
	contract, err := bindLicence(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LicenceCaller{contract: contract}, nil
}

// NewLicenceTransactor creates a new write-only instance of Licence, bound to a specific deployed contract.
func NewLicenceTransactor(address common.Address, transactor bind.ContractTransactor) (*LicenceTransactor, error) {
	contract, err := bindLicence(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LicenceTransactor{contract: contract}, nil
}

// NewLicenceFilterer creates a new log filterer instance of Licence, bound to a specific deployed contract.
func NewLicenceFilterer(address common.Address, filterer bind.ContractFilterer) (*LicenceFilterer, error) {
	contract, err := bindLicence(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LicenceFilterer{contract: contract}, nil
}

// bindLicence binds a generic wrapper to an already deployed contract.
func bindLicence(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Licence *LicenceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Licence.Contract.LicenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Licence *LicenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.Contract.LicenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Licence *LicenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Licence.Contract.LicenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Licence *LicenceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Licence.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Licence *LicenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Licence *LicenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Licence.Contract.contract.Transact(opts, method, params...)
}

// MAXAMOUNTSCALE is a free data retrieval call binding the contract method 0xca0e2e20.
//
// Solidity: function MAX_AMOUNT_SCALE() constant returns(uint256)
func (_Licence *LicenceCaller) MAXAMOUNTSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "MAX_AMOUNT_SCALE")
	return *ret0, err
}

// MAXAMOUNTSCALE is a free data retrieval call binding the contract method 0xca0e2e20.
//
// Solidity: function MAX_AMOUNT_SCALE() constant returns(uint256)
func (_Licence *LicenceSession) MAXAMOUNTSCALE() (*big.Int, error) {
	return _Licence.Contract.MAXAMOUNTSCALE(&_Licence.CallOpts)
}

// MAXAMOUNTSCALE is a free data retrieval call binding the contract method 0xca0e2e20.
//
// Solidity: function MAX_AMOUNT_SCALE() constant returns(uint256)
func (_Licence *LicenceCallerSession) MAXAMOUNTSCALE() (*big.Int, error) {
	return _Licence.Contract.MAXAMOUNTSCALE(&_Licence.CallOpts)
}

// CryptoFloat is a free data retrieval call binding the contract method 0xa036ba60.
//
// Solidity: function cryptoFloat() constant returns(address)
func (_Licence *LicenceCaller) CryptoFloat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "cryptoFloat")
	return *ret0, err
}

// CryptoFloat is a free data retrieval call binding the contract method 0xa036ba60.
//
// Solidity: function cryptoFloat() constant returns(address)
func (_Licence *LicenceSession) CryptoFloat() (common.Address, error) {
	return _Licence.Contract.CryptoFloat(&_Licence.CallOpts)
}

// CryptoFloat is a free data retrieval call binding the contract method 0xa036ba60.
//
// Solidity: function cryptoFloat() constant returns(address)
func (_Licence *LicenceCallerSession) CryptoFloat() (common.Address, error) {
	return _Licence.Contract.CryptoFloat(&_Licence.CallOpts)
}

// FloatLocked is a free data retrieval call binding the contract method 0x940b9c3b.
//
// Solidity: function floatLocked() constant returns(bool)
func (_Licence *LicenceCaller) FloatLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "floatLocked")
	return *ret0, err
}

// FloatLocked is a free data retrieval call binding the contract method 0x940b9c3b.
//
// Solidity: function floatLocked() constant returns(bool)
func (_Licence *LicenceSession) FloatLocked() (bool, error) {
	return _Licence.Contract.FloatLocked(&_Licence.CallOpts)
}

// FloatLocked is a free data retrieval call binding the contract method 0x940b9c3b.
//
// Solidity: function floatLocked() constant returns(bool)
func (_Licence *LicenceCallerSession) FloatLocked() (bool, error) {
	return _Licence.Contract.FloatLocked(&_Licence.CallOpts)
}

// HolderLocked is a free data retrieval call binding the contract method 0xf15ff455.
//
// Solidity: function holderLocked() constant returns(bool)
func (_Licence *LicenceCaller) HolderLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "holderLocked")
	return *ret0, err
}

// HolderLocked is a free data retrieval call binding the contract method 0xf15ff455.
//
// Solidity: function holderLocked() constant returns(bool)
func (_Licence *LicenceSession) HolderLocked() (bool, error) {
	return _Licence.Contract.HolderLocked(&_Licence.CallOpts)
}

// HolderLocked is a free data retrieval call binding the contract method 0xf15ff455.
//
// Solidity: function holderLocked() constant returns(bool)
func (_Licence *LicenceCallerSession) HolderLocked() (bool, error) {
	return _Licence.Contract.HolderLocked(&_Licence.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Licence *LicenceCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Licence *LicenceSession) IsTransferable() (bool, error) {
	return _Licence.Contract.IsTransferable(&_Licence.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Licence *LicenceCallerSession) IsTransferable() (bool, error) {
	return _Licence.Contract.IsTransferable(&_Licence.CallOpts)
}

// LicenceAmountScaled is a free data retrieval call binding the contract method 0xac904c63.
//
// Solidity: function licenceAmountScaled() constant returns(uint256)
func (_Licence *LicenceCaller) LicenceAmountScaled(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "licenceAmountScaled")
	return *ret0, err
}

// LicenceAmountScaled is a free data retrieval call binding the contract method 0xac904c63.
//
// Solidity: function licenceAmountScaled() constant returns(uint256)
func (_Licence *LicenceSession) LicenceAmountScaled() (*big.Int, error) {
	return _Licence.Contract.LicenceAmountScaled(&_Licence.CallOpts)
}

// LicenceAmountScaled is a free data retrieval call binding the contract method 0xac904c63.
//
// Solidity: function licenceAmountScaled() constant returns(uint256)
func (_Licence *LicenceCallerSession) LicenceAmountScaled() (*big.Int, error) {
	return _Licence.Contract.LicenceAmountScaled(&_Licence.CallOpts)
}

// LicenceDAO is a free data retrieval call binding the contract method 0x3a7afe02.
//
// Solidity: function licenceDAO() constant returns(address)
func (_Licence *LicenceCaller) LicenceDAO(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "licenceDAO")
	return *ret0, err
}

// LicenceDAO is a free data retrieval call binding the contract method 0x3a7afe02.
//
// Solidity: function licenceDAO() constant returns(address)
func (_Licence *LicenceSession) LicenceDAO() (common.Address, error) {
	return _Licence.Contract.LicenceDAO(&_Licence.CallOpts)
}

// LicenceDAO is a free data retrieval call binding the contract method 0x3a7afe02.
//
// Solidity: function licenceDAO() constant returns(address)
func (_Licence *LicenceCallerSession) LicenceDAO() (common.Address, error) {
	return _Licence.Contract.LicenceDAO(&_Licence.CallOpts)
}

// LicenceDAOLocked is a free data retrieval call binding the contract method 0xe30c5fa8.
//
// Solidity: function licenceDAOLocked() constant returns(bool)
func (_Licence *LicenceCaller) LicenceDAOLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "licenceDAOLocked")
	return *ret0, err
}

// LicenceDAOLocked is a free data retrieval call binding the contract method 0xe30c5fa8.
//
// Solidity: function licenceDAOLocked() constant returns(bool)
func (_Licence *LicenceSession) LicenceDAOLocked() (bool, error) {
	return _Licence.Contract.LicenceDAOLocked(&_Licence.CallOpts)
}

// LicenceDAOLocked is a free data retrieval call binding the contract method 0xe30c5fa8.
//
// Solidity: function licenceDAOLocked() constant returns(bool)
func (_Licence *LicenceCallerSession) LicenceDAOLocked() (bool, error) {
	return _Licence.Contract.LicenceDAOLocked(&_Licence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Licence *LicenceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Licence *LicenceSession) Owner() (common.Address, error) {
	return _Licence.Contract.Owner(&_Licence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Licence *LicenceCallerSession) Owner() (common.Address, error) {
	return _Licence.Contract.Owner(&_Licence.CallOpts)
}

// TknContractAddress is a free data retrieval call binding the contract method 0x99a5e1d0.
//
// Solidity: function tknContractAddress() constant returns(address)
func (_Licence *LicenceCaller) TknContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "tknContractAddress")
	return *ret0, err
}

// TknContractAddress is a free data retrieval call binding the contract method 0x99a5e1d0.
//
// Solidity: function tknContractAddress() constant returns(address)
func (_Licence *LicenceSession) TknContractAddress() (common.Address, error) {
	return _Licence.Contract.TknContractAddress(&_Licence.CallOpts)
}

// TknContractAddress is a free data retrieval call binding the contract method 0x99a5e1d0.
//
// Solidity: function tknContractAddress() constant returns(address)
func (_Licence *LicenceCallerSession) TknContractAddress() (common.Address, error) {
	return _Licence.Contract.TknContractAddress(&_Licence.CallOpts)
}

// TknContractAddressLocked is a free data retrieval call binding the contract method 0x837c70ef.
//
// Solidity: function tknContractAddressLocked() constant returns(bool)
func (_Licence *LicenceCaller) TknContractAddressLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "tknContractAddressLocked")
	return *ret0, err
}

// TknContractAddressLocked is a free data retrieval call binding the contract method 0x837c70ef.
//
// Solidity: function tknContractAddressLocked() constant returns(bool)
func (_Licence *LicenceSession) TknContractAddressLocked() (bool, error) {
	return _Licence.Contract.TknContractAddressLocked(&_Licence.CallOpts)
}

// TknContractAddressLocked is a free data retrieval call binding the contract method 0x837c70ef.
//
// Solidity: function tknContractAddressLocked() constant returns(bool)
func (_Licence *LicenceCallerSession) TknContractAddressLocked() (bool, error) {
	return _Licence.Contract.TknContractAddressLocked(&_Licence.CallOpts)
}

// TokenHolder is a free data retrieval call binding the contract method 0x420a83e7.
//
// Solidity: function tokenHolder() constant returns(address)
func (_Licence *LicenceCaller) TokenHolder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Licence.contract.Call(opts, out, "tokenHolder")
	return *ret0, err
}

// TokenHolder is a free data retrieval call binding the contract method 0x420a83e7.
//
// Solidity: function tokenHolder() constant returns(address)
func (_Licence *LicenceSession) TokenHolder() (common.Address, error) {
	return _Licence.Contract.TokenHolder(&_Licence.CallOpts)
}

// TokenHolder is a free data retrieval call binding the contract method 0x420a83e7.
//
// Solidity: function tokenHolder() constant returns(address)
func (_Licence *LicenceCallerSession) TokenHolder() (common.Address, error) {
	return _Licence.Contract.TokenHolder(&_Licence.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_Licence *LicenceTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "claim", _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_Licence *LicenceSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.Claim(&_Licence.TransactOpts, _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_Licence *LicenceTransactorSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.Claim(&_Licence.TransactOpts, _to, _asset, _amount)
}

// Load is a paid mutator transaction binding the contract method 0x1b3c96b4.
//
// Solidity: function load(address _asset, uint256 _amount) returns()
func (_Licence *LicenceTransactor) Load(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "load", _asset, _amount)
}

// Load is a paid mutator transaction binding the contract method 0x1b3c96b4.
//
// Solidity: function load(address _asset, uint256 _amount) returns()
func (_Licence *LicenceSession) Load(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.Load(&_Licence.TransactOpts, _asset, _amount)
}

// Load is a paid mutator transaction binding the contract method 0x1b3c96b4.
//
// Solidity: function load(address _asset, uint256 _amount) returns()
func (_Licence *LicenceTransactorSession) Load(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.Load(&_Licence.TransactOpts, _asset, _amount)
}

// LockFloat is a paid mutator transaction binding the contract method 0xd08b4ecc.
//
// Solidity: function lockFloat() returns()
func (_Licence *LicenceTransactor) LockFloat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "lockFloat")
}

// LockFloat is a paid mutator transaction binding the contract method 0xd08b4ecc.
//
// Solidity: function lockFloat() returns()
func (_Licence *LicenceSession) LockFloat() (*types.Transaction, error) {
	return _Licence.Contract.LockFloat(&_Licence.TransactOpts)
}

// LockFloat is a paid mutator transaction binding the contract method 0xd08b4ecc.
//
// Solidity: function lockFloat() returns()
func (_Licence *LicenceTransactorSession) LockFloat() (*types.Transaction, error) {
	return _Licence.Contract.LockFloat(&_Licence.TransactOpts)
}

// LockHolder is a paid mutator transaction binding the contract method 0xd1696b16.
//
// Solidity: function lockHolder() returns()
func (_Licence *LicenceTransactor) LockHolder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "lockHolder")
}

// LockHolder is a paid mutator transaction binding the contract method 0xd1696b16.
//
// Solidity: function lockHolder() returns()
func (_Licence *LicenceSession) LockHolder() (*types.Transaction, error) {
	return _Licence.Contract.LockHolder(&_Licence.TransactOpts)
}

// LockHolder is a paid mutator transaction binding the contract method 0xd1696b16.
//
// Solidity: function lockHolder() returns()
func (_Licence *LicenceTransactorSession) LockHolder() (*types.Transaction, error) {
	return _Licence.Contract.LockHolder(&_Licence.TransactOpts)
}

// LockLicenceDAO is a paid mutator transaction binding the contract method 0x4ac22b3c.
//
// Solidity: function lockLicenceDAO() returns()
func (_Licence *LicenceTransactor) LockLicenceDAO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "lockLicenceDAO")
}

// LockLicenceDAO is a paid mutator transaction binding the contract method 0x4ac22b3c.
//
// Solidity: function lockLicenceDAO() returns()
func (_Licence *LicenceSession) LockLicenceDAO() (*types.Transaction, error) {
	return _Licence.Contract.LockLicenceDAO(&_Licence.TransactOpts)
}

// LockLicenceDAO is a paid mutator transaction binding the contract method 0x4ac22b3c.
//
// Solidity: function lockLicenceDAO() returns()
func (_Licence *LicenceTransactorSession) LockLicenceDAO() (*types.Transaction, error) {
	return _Licence.Contract.LockLicenceDAO(&_Licence.TransactOpts)
}

// LockTKNContractAddress is a paid mutator transaction binding the contract method 0x0bf25c91.
//
// Solidity: function lockTKNContractAddress() returns()
func (_Licence *LicenceTransactor) LockTKNContractAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "lockTKNContractAddress")
}

// LockTKNContractAddress is a paid mutator transaction binding the contract method 0x0bf25c91.
//
// Solidity: function lockTKNContractAddress() returns()
func (_Licence *LicenceSession) LockTKNContractAddress() (*types.Transaction, error) {
	return _Licence.Contract.LockTKNContractAddress(&_Licence.TransactOpts)
}

// LockTKNContractAddress is a paid mutator transaction binding the contract method 0x0bf25c91.
//
// Solidity: function lockTKNContractAddress() returns()
func (_Licence *LicenceTransactorSession) LockTKNContractAddress() (*types.Transaction, error) {
	return _Licence.Contract.LockTKNContractAddress(&_Licence.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Licence *LicenceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Licence *LicenceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Licence.Contract.RenounceOwnership(&_Licence.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Licence *LicenceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Licence.Contract.RenounceOwnership(&_Licence.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Licence *LicenceTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Licence *LicenceSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Licence.Contract.TransferOwnership(&_Licence.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Licence *LicenceTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Licence.Contract.TransferOwnership(&_Licence.TransactOpts, _account, _transferable)
}

// UpdateFloat is a paid mutator transaction binding the contract method 0xd0cddd67.
//
// Solidity: function updateFloat(address _newFloat) returns()
func (_Licence *LicenceTransactor) UpdateFloat(opts *bind.TransactOpts, _newFloat common.Address) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateFloat", _newFloat)
}

// UpdateFloat is a paid mutator transaction binding the contract method 0xd0cddd67.
//
// Solidity: function updateFloat(address _newFloat) returns()
func (_Licence *LicenceSession) UpdateFloat(_newFloat common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateFloat(&_Licence.TransactOpts, _newFloat)
}

// UpdateFloat is a paid mutator transaction binding the contract method 0xd0cddd67.
//
// Solidity: function updateFloat(address _newFloat) returns()
func (_Licence *LicenceTransactorSession) UpdateFloat(_newFloat common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateFloat(&_Licence.TransactOpts, _newFloat)
}

// UpdateHolder is a paid mutator transaction binding the contract method 0x42719faa.
//
// Solidity: function updateHolder(address _newHolder) returns()
func (_Licence *LicenceTransactor) UpdateHolder(opts *bind.TransactOpts, _newHolder common.Address) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateHolder", _newHolder)
}

// UpdateHolder is a paid mutator transaction binding the contract method 0x42719faa.
//
// Solidity: function updateHolder(address _newHolder) returns()
func (_Licence *LicenceSession) UpdateHolder(_newHolder common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateHolder(&_Licence.TransactOpts, _newHolder)
}

// UpdateHolder is a paid mutator transaction binding the contract method 0x42719faa.
//
// Solidity: function updateHolder(address _newHolder) returns()
func (_Licence *LicenceTransactorSession) UpdateHolder(_newHolder common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateHolder(&_Licence.TransactOpts, _newHolder)
}

// UpdateLicenceAmount is a paid mutator transaction binding the contract method 0x68ce74e7.
//
// Solidity: function updateLicenceAmount(uint256 _newAmount) returns()
func (_Licence *LicenceTransactor) UpdateLicenceAmount(opts *bind.TransactOpts, _newAmount *big.Int) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateLicenceAmount", _newAmount)
}

// UpdateLicenceAmount is a paid mutator transaction binding the contract method 0x68ce74e7.
//
// Solidity: function updateLicenceAmount(uint256 _newAmount) returns()
func (_Licence *LicenceSession) UpdateLicenceAmount(_newAmount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.UpdateLicenceAmount(&_Licence.TransactOpts, _newAmount)
}

// UpdateLicenceAmount is a paid mutator transaction binding the contract method 0x68ce74e7.
//
// Solidity: function updateLicenceAmount(uint256 _newAmount) returns()
func (_Licence *LicenceTransactorSession) UpdateLicenceAmount(_newAmount *big.Int) (*types.Transaction, error) {
	return _Licence.Contract.UpdateLicenceAmount(&_Licence.TransactOpts, _newAmount)
}

// UpdateLicenceDAO is a paid mutator transaction binding the contract method 0xe3d80242.
//
// Solidity: function updateLicenceDAO(address _newDAO) returns()
func (_Licence *LicenceTransactor) UpdateLicenceDAO(opts *bind.TransactOpts, _newDAO common.Address) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateLicenceDAO", _newDAO)
}

// UpdateLicenceDAO is a paid mutator transaction binding the contract method 0xe3d80242.
//
// Solidity: function updateLicenceDAO(address _newDAO) returns()
func (_Licence *LicenceSession) UpdateLicenceDAO(_newDAO common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateLicenceDAO(&_Licence.TransactOpts, _newDAO)
}

// UpdateLicenceDAO is a paid mutator transaction binding the contract method 0xe3d80242.
//
// Solidity: function updateLicenceDAO(address _newDAO) returns()
func (_Licence *LicenceTransactorSession) UpdateLicenceDAO(_newDAO common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateLicenceDAO(&_Licence.TransactOpts, _newDAO)
}

// UpdateTKNContractAddress is a paid mutator transaction binding the contract method 0x0d42e82f.
//
// Solidity: function updateTKNContractAddress(address _newTKN) returns()
func (_Licence *LicenceTransactor) UpdateTKNContractAddress(opts *bind.TransactOpts, _newTKN common.Address) (*types.Transaction, error) {
	return _Licence.contract.Transact(opts, "updateTKNContractAddress", _newTKN)
}

// UpdateTKNContractAddress is a paid mutator transaction binding the contract method 0x0d42e82f.
//
// Solidity: function updateTKNContractAddress(address _newTKN) returns()
func (_Licence *LicenceSession) UpdateTKNContractAddress(_newTKN common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateTKNContractAddress(&_Licence.TransactOpts, _newTKN)
}

// UpdateTKNContractAddress is a paid mutator transaction binding the contract method 0x0d42e82f.
//
// Solidity: function updateTKNContractAddress(address _newTKN) returns()
func (_Licence *LicenceTransactorSession) UpdateTKNContractAddress(_newTKN common.Address) (*types.Transaction, error) {
	return _Licence.Contract.UpdateTKNContractAddress(&_Licence.TransactOpts, _newTKN)
}

// LicenceClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Licence contract.
type LicenceClaimedIterator struct {
	Event *LicenceClaimed // Event containing the contract specifics and raw log

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
func (it *LicenceClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceClaimed)
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
		it.Event = new(LicenceClaimed)
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
func (it *LicenceClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceClaimed represents a Claimed event raised by the Licence contract.
type LicenceClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Licence *LicenceFilterer) FilterClaimed(opts *bind.FilterOpts) (*LicenceClaimedIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &LicenceClaimedIterator{contract: _Licence.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_Licence *LicenceFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *LicenceClaimed) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceClaimed)
				if err := _Licence.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// LicenceLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the Licence contract.
type LicenceLockedOwnershipIterator struct {
	Event *LicenceLockedOwnership // Event containing the contract specifics and raw log

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
func (it *LicenceLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceLockedOwnership)
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
		it.Event = new(LicenceLockedOwnership)
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
func (it *LicenceLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceLockedOwnership represents a LockedOwnership event raised by the Licence contract.
type LicenceLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Licence *LicenceFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*LicenceLockedOwnershipIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &LicenceLockedOwnershipIterator{contract: _Licence.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Licence *LicenceFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *LicenceLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceLockedOwnership)
				if err := _Licence.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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

// LicenceLogTokenTransferIterator is returned from FilterLogTokenTransfer and is used to iterate over the raw logs and unpacked data for LogTokenTransfer events raised by the Licence contract.
type LicenceLogTokenTransferIterator struct {
	Event *LicenceLogTokenTransfer // Event containing the contract specifics and raw log

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
func (it *LicenceLogTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceLogTokenTransfer)
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
		it.Event = new(LicenceLogTokenTransfer)
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
func (it *LicenceLogTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceLogTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceLogTokenTransfer represents a LogTokenTransfer event raised by the Licence contract.
type LicenceLogTokenTransfer struct {
	Asset  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTokenTransfer is a free log retrieval operation binding the contract event 0x4bf0cfb200c294b3cbb11e37d57eab8dfbd930d7390c3b97ecae8ef827f86884.
//
// Solidity: event LogTokenTransfer(address _asset, address _to, uint256 _amount)
func (_Licence *LicenceFilterer) FilterLogTokenTransfer(opts *bind.FilterOpts) (*LicenceLogTokenTransferIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "LogTokenTransfer")
	if err != nil {
		return nil, err
	}
	return &LicenceLogTokenTransferIterator{contract: _Licence.contract, event: "LogTokenTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTokenTransfer is a free log subscription operation binding the contract event 0x4bf0cfb200c294b3cbb11e37d57eab8dfbd930d7390c3b97ecae8ef827f86884.
//
// Solidity: event LogTokenTransfer(address _asset, address _to, uint256 _amount)
func (_Licence *LicenceFilterer) WatchLogTokenTransfer(opts *bind.WatchOpts, sink chan<- *LicenceLogTokenTransfer) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "LogTokenTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceLogTokenTransfer)
				if err := _Licence.contract.UnpackLog(event, "LogTokenTransfer", log); err != nil {
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

// LicenceReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Licence contract.
type LicenceReceivedIterator struct {
	Event *LicenceReceived // Event containing the contract specifics and raw log

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
func (it *LicenceReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceReceived)
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
		it.Event = new(LicenceReceived)
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
func (it *LicenceReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceReceived represents a Received event raised by the Licence contract.
type LicenceReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Licence *LicenceFilterer) FilterReceived(opts *bind.FilterOpts) (*LicenceReceivedIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &LicenceReceivedIterator{contract: _Licence.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _from, uint256 _amount)
func (_Licence *LicenceFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *LicenceReceived) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceReceived)
				if err := _Licence.contract.UnpackLog(event, "Received", log); err != nil {
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

// LicenceTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Licence contract.
type LicenceTransferredOwnershipIterator struct {
	Event *LicenceTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *LicenceTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceTransferredOwnership)
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
		it.Event = new(LicenceTransferredOwnership)
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
func (it *LicenceTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceTransferredOwnership represents a TransferredOwnership event raised by the Licence contract.
type LicenceTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Licence *LicenceFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*LicenceTransferredOwnershipIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &LicenceTransferredOwnershipIterator{contract: _Licence.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Licence *LicenceFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *LicenceTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceTransferredOwnership)
				if err := _Licence.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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

// LicenceTransferredToCryptoFloatIterator is returned from FilterTransferredToCryptoFloat and is used to iterate over the raw logs and unpacked data for TransferredToCryptoFloat events raised by the Licence contract.
type LicenceTransferredToCryptoFloatIterator struct {
	Event *LicenceTransferredToCryptoFloat // Event containing the contract specifics and raw log

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
func (it *LicenceTransferredToCryptoFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceTransferredToCryptoFloat)
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
		it.Event = new(LicenceTransferredToCryptoFloat)
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
func (it *LicenceTransferredToCryptoFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceTransferredToCryptoFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceTransferredToCryptoFloat represents a TransferredToCryptoFloat event raised by the Licence contract.
type LicenceTransferredToCryptoFloat struct {
	From   common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredToCryptoFloat is a free log retrieval operation binding the contract event 0xc8a7b0bd71097b47b2cad75e4e939d2aeb7fae88110e68f93b83fed08e9d3c38.
//
// Solidity: event TransferredToCryptoFloat(address _from, address _to, address _asset, uint256 _amount)
func (_Licence *LicenceFilterer) FilterTransferredToCryptoFloat(opts *bind.FilterOpts) (*LicenceTransferredToCryptoFloatIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "TransferredToCryptoFloat")
	if err != nil {
		return nil, err
	}
	return &LicenceTransferredToCryptoFloatIterator{contract: _Licence.contract, event: "TransferredToCryptoFloat", logs: logs, sub: sub}, nil
}

// WatchTransferredToCryptoFloat is a free log subscription operation binding the contract event 0xc8a7b0bd71097b47b2cad75e4e939d2aeb7fae88110e68f93b83fed08e9d3c38.
//
// Solidity: event TransferredToCryptoFloat(address _from, address _to, address _asset, uint256 _amount)
func (_Licence *LicenceFilterer) WatchTransferredToCryptoFloat(opts *bind.WatchOpts, sink chan<- *LicenceTransferredToCryptoFloat) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "TransferredToCryptoFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceTransferredToCryptoFloat)
				if err := _Licence.contract.UnpackLog(event, "TransferredToCryptoFloat", log); err != nil {
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

// LicenceTransferredToTokenHolderIterator is returned from FilterTransferredToTokenHolder and is used to iterate over the raw logs and unpacked data for TransferredToTokenHolder events raised by the Licence contract.
type LicenceTransferredToTokenHolderIterator struct {
	Event *LicenceTransferredToTokenHolder // Event containing the contract specifics and raw log

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
func (it *LicenceTransferredToTokenHolderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceTransferredToTokenHolder)
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
		it.Event = new(LicenceTransferredToTokenHolder)
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
func (it *LicenceTransferredToTokenHolderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceTransferredToTokenHolderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceTransferredToTokenHolder represents a TransferredToTokenHolder event raised by the Licence contract.
type LicenceTransferredToTokenHolder struct {
	From   common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredToTokenHolder is a free log retrieval operation binding the contract event 0xdd9dfad7b30d6b224e235f89565871419d3dec3b563a4e231f12d2cc97f9acfc.
//
// Solidity: event TransferredToTokenHolder(address _from, address _to, address _asset, uint256 _amount)
func (_Licence *LicenceFilterer) FilterTransferredToTokenHolder(opts *bind.FilterOpts) (*LicenceTransferredToTokenHolderIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "TransferredToTokenHolder")
	if err != nil {
		return nil, err
	}
	return &LicenceTransferredToTokenHolderIterator{contract: _Licence.contract, event: "TransferredToTokenHolder", logs: logs, sub: sub}, nil
}

// WatchTransferredToTokenHolder is a free log subscription operation binding the contract event 0xdd9dfad7b30d6b224e235f89565871419d3dec3b563a4e231f12d2cc97f9acfc.
//
// Solidity: event TransferredToTokenHolder(address _from, address _to, address _asset, uint256 _amount)
func (_Licence *LicenceFilterer) WatchTransferredToTokenHolder(opts *bind.WatchOpts, sink chan<- *LicenceTransferredToTokenHolder) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "TransferredToTokenHolder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceTransferredToTokenHolder)
				if err := _Licence.contract.UnpackLog(event, "TransferredToTokenHolder", log); err != nil {
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

// LicenceUpdatedCryptoFloatIterator is returned from FilterUpdatedCryptoFloat and is used to iterate over the raw logs and unpacked data for UpdatedCryptoFloat events raised by the Licence contract.
type LicenceUpdatedCryptoFloatIterator struct {
	Event *LicenceUpdatedCryptoFloat // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedCryptoFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedCryptoFloat)
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
		it.Event = new(LicenceUpdatedCryptoFloat)
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
func (it *LicenceUpdatedCryptoFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedCryptoFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedCryptoFloat represents a UpdatedCryptoFloat event raised by the Licence contract.
type LicenceUpdatedCryptoFloat struct {
	NewFloat common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCryptoFloat is a free log retrieval operation binding the contract event 0x9af2841b0db134bda87280e2a9cababb156f95023c87023d708a677d61b4b6d8.
//
// Solidity: event UpdatedCryptoFloat(address _newFloat)
func (_Licence *LicenceFilterer) FilterUpdatedCryptoFloat(opts *bind.FilterOpts) (*LicenceUpdatedCryptoFloatIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedCryptoFloat")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedCryptoFloatIterator{contract: _Licence.contract, event: "UpdatedCryptoFloat", logs: logs, sub: sub}, nil
}

// WatchUpdatedCryptoFloat is a free log subscription operation binding the contract event 0x9af2841b0db134bda87280e2a9cababb156f95023c87023d708a677d61b4b6d8.
//
// Solidity: event UpdatedCryptoFloat(address _newFloat)
func (_Licence *LicenceFilterer) WatchUpdatedCryptoFloat(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedCryptoFloat) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedCryptoFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedCryptoFloat)
				if err := _Licence.contract.UnpackLog(event, "UpdatedCryptoFloat", log); err != nil {
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

// LicenceUpdatedLicenceAmountIterator is returned from FilterUpdatedLicenceAmount and is used to iterate over the raw logs and unpacked data for UpdatedLicenceAmount events raised by the Licence contract.
type LicenceUpdatedLicenceAmountIterator struct {
	Event *LicenceUpdatedLicenceAmount // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedLicenceAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedLicenceAmount)
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
		it.Event = new(LicenceUpdatedLicenceAmount)
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
func (it *LicenceUpdatedLicenceAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedLicenceAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedLicenceAmount represents a UpdatedLicenceAmount event raised by the Licence contract.
type LicenceUpdatedLicenceAmount struct {
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedLicenceAmount is a free log retrieval operation binding the contract event 0x587b6068be8c555e2cddc6ad8a56df5e8dfb1533cc063d6703f79c791de15148.
//
// Solidity: event UpdatedLicenceAmount(uint256 _newAmount)
func (_Licence *LicenceFilterer) FilterUpdatedLicenceAmount(opts *bind.FilterOpts) (*LicenceUpdatedLicenceAmountIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedLicenceAmount")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedLicenceAmountIterator{contract: _Licence.contract, event: "UpdatedLicenceAmount", logs: logs, sub: sub}, nil
}

// WatchUpdatedLicenceAmount is a free log subscription operation binding the contract event 0x587b6068be8c555e2cddc6ad8a56df5e8dfb1533cc063d6703f79c791de15148.
//
// Solidity: event UpdatedLicenceAmount(uint256 _newAmount)
func (_Licence *LicenceFilterer) WatchUpdatedLicenceAmount(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedLicenceAmount) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedLicenceAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedLicenceAmount)
				if err := _Licence.contract.UnpackLog(event, "UpdatedLicenceAmount", log); err != nil {
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

// LicenceUpdatedLicenceDAOIterator is returned from FilterUpdatedLicenceDAO and is used to iterate over the raw logs and unpacked data for UpdatedLicenceDAO events raised by the Licence contract.
type LicenceUpdatedLicenceDAOIterator struct {
	Event *LicenceUpdatedLicenceDAO // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedLicenceDAOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedLicenceDAO)
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
		it.Event = new(LicenceUpdatedLicenceDAO)
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
func (it *LicenceUpdatedLicenceDAOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedLicenceDAOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedLicenceDAO represents a UpdatedLicenceDAO event raised by the Licence contract.
type LicenceUpdatedLicenceDAO struct {
	NewDAO common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedLicenceDAO is a free log retrieval operation binding the contract event 0xd32c17b277c7e87842861153d758814a267634f4308ec2461f88756df7dd7068.
//
// Solidity: event UpdatedLicenceDAO(address _newDAO)
func (_Licence *LicenceFilterer) FilterUpdatedLicenceDAO(opts *bind.FilterOpts) (*LicenceUpdatedLicenceDAOIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedLicenceDAO")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedLicenceDAOIterator{contract: _Licence.contract, event: "UpdatedLicenceDAO", logs: logs, sub: sub}, nil
}

// WatchUpdatedLicenceDAO is a free log subscription operation binding the contract event 0xd32c17b277c7e87842861153d758814a267634f4308ec2461f88756df7dd7068.
//
// Solidity: event UpdatedLicenceDAO(address _newDAO)
func (_Licence *LicenceFilterer) WatchUpdatedLicenceDAO(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedLicenceDAO) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedLicenceDAO")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedLicenceDAO)
				if err := _Licence.contract.UnpackLog(event, "UpdatedLicenceDAO", log); err != nil {
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

// LicenceUpdatedTKNContractAddressIterator is returned from FilterUpdatedTKNContractAddress and is used to iterate over the raw logs and unpacked data for UpdatedTKNContractAddress events raised by the Licence contract.
type LicenceUpdatedTKNContractAddressIterator struct {
	Event *LicenceUpdatedTKNContractAddress // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedTKNContractAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedTKNContractAddress)
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
		it.Event = new(LicenceUpdatedTKNContractAddress)
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
func (it *LicenceUpdatedTKNContractAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedTKNContractAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedTKNContractAddress represents a UpdatedTKNContractAddress event raised by the Licence contract.
type LicenceUpdatedTKNContractAddress struct {
	NewTKN common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTKNContractAddress is a free log retrieval operation binding the contract event 0x2aeed92123e61fe64748a447c2ba122c4bfc0201d1ed5149e9ce9ede5adda545.
//
// Solidity: event UpdatedTKNContractAddress(address _newTKN)
func (_Licence *LicenceFilterer) FilterUpdatedTKNContractAddress(opts *bind.FilterOpts) (*LicenceUpdatedTKNContractAddressIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedTKNContractAddress")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedTKNContractAddressIterator{contract: _Licence.contract, event: "UpdatedTKNContractAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedTKNContractAddress is a free log subscription operation binding the contract event 0x2aeed92123e61fe64748a447c2ba122c4bfc0201d1ed5149e9ce9ede5adda545.
//
// Solidity: event UpdatedTKNContractAddress(address _newTKN)
func (_Licence *LicenceFilterer) WatchUpdatedTKNContractAddress(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedTKNContractAddress) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedTKNContractAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedTKNContractAddress)
				if err := _Licence.contract.UnpackLog(event, "UpdatedTKNContractAddress", log); err != nil {
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

// LicenceUpdatedTokenHolderIterator is returned from FilterUpdatedTokenHolder and is used to iterate over the raw logs and unpacked data for UpdatedTokenHolder events raised by the Licence contract.
type LicenceUpdatedTokenHolderIterator struct {
	Event *LicenceUpdatedTokenHolder // Event containing the contract specifics and raw log

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
func (it *LicenceUpdatedTokenHolderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenceUpdatedTokenHolder)
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
		it.Event = new(LicenceUpdatedTokenHolder)
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
func (it *LicenceUpdatedTokenHolderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenceUpdatedTokenHolderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenceUpdatedTokenHolder represents a UpdatedTokenHolder event raised by the Licence contract.
type LicenceUpdatedTokenHolder struct {
	NewHolder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenHolder is a free log retrieval operation binding the contract event 0xfa6bae0f250db86534a013b1c7a6c4076aa8f8d1ac248771a1c73f4ba366922a.
//
// Solidity: event UpdatedTokenHolder(address _newHolder)
func (_Licence *LicenceFilterer) FilterUpdatedTokenHolder(opts *bind.FilterOpts) (*LicenceUpdatedTokenHolderIterator, error) {

	logs, sub, err := _Licence.contract.FilterLogs(opts, "UpdatedTokenHolder")
	if err != nil {
		return nil, err
	}
	return &LicenceUpdatedTokenHolderIterator{contract: _Licence.contract, event: "UpdatedTokenHolder", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenHolder is a free log subscription operation binding the contract event 0xfa6bae0f250db86534a013b1c7a6c4076aa8f8d1ac248771a1c73f4ba366922a.
//
// Solidity: event UpdatedTokenHolder(address _newHolder)
func (_Licence *LicenceFilterer) WatchUpdatedTokenHolder(opts *bind.WatchOpts, sink chan<- *LicenceUpdatedTokenHolder) (event.Subscription, error) {

	logs, sub, err := _Licence.contract.WatchLogs(opts, "UpdatedTokenHolder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenceUpdatedTokenHolder)
				if err := _Licence.contract.UnpackLog(event, "UpdatedTokenHolder", log); err != nil {
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
