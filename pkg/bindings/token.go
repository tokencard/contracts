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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"Launch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockedTokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeRemainders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockTokenHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnerSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"neverPauseAgain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"launched\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"setOwnerFreeDay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenholder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingDone\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pausingMechanismLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remaindersSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"uint256[]\"}],\"name\":\"multiMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"completeMinting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerTokensFreeDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimAuctionableTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_remainingOwner\",\"type\":\"uint256\"},{\"name\":\"_remainingAuctionable\",\"type\":\"uint256\"}],\"name\":\"setRemainders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingAuctionable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_th\",\"type\":\"address\"}],\"name\":\"setTokenHolder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logTokenTransfer\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `608060405260086002556007805462ffffff19169055600d805460a860020a61ffff02191690553480156200003357600080fd5b5060008054600160a060020a0319163317905560408051808201909152600c8082527f4d6f6e6f6c69746820544b4e000000000000000000000000000000000000000060209092019182526200008c91600891620000f0565b506002546009805460ff191660ff9092169190911790556040805180820190915260038082527f544b4e00000000000000000000000000000000000000000000000000000000006020909201918252620000e991600a91620000f0565b5062000195565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013357805160ff191683800117855562000163565b8280016001018555821562000163579182015b828111156200016357825182559160200191906001019062000146565b506200017192915062000175565b5090565b6200019291905b808211156200017157600081556001016200017c565b90565b6119ed80620001a56000396000f3006080604052600436106102195763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302ac8168811461021e57806306fdde0314610235578063095ea7b3146102bf57806317e139a7146102f757806318160ddd1461030c5780632276774c1461033357806323b872dd146103485780632bbeac9114610372578063313ce567146103875780633da1eff5146103b25780633f4ba83a146103c757806340c10f19146103dc57806342966c68146104005780635c975abb146104185780636348eae61461042d578063661884631461044257806370a0823114610466578063771282f61461048757806379ba50971461049c5780638091f3bf146104b157806381e529cd146104c65780638456cb59146104de57806384eba00c146104f35780638da5cb5b146105245780638e2ae5641461053957806390912d091461054e57806392eefe9b1461056357806395d89b411461058457806398a9ae44146105995780639a0e4ebb146105ae578063a4f91a2e14610603578063a6f9dae114610618578063a844545d14610639578063a9059cbb1461064e578063ac4d2e9514610672578063b357a55214610687578063c27549d91461069f578063c6e81b07146106ba578063cae9ca51146106cf578063d73dd62314610738578063dd62ed3e1461075c578063df8de3e714610783578063f29d2f28146107a4578063f77c4791146107c5575b600080fd5b34801561022a57600080fd5b506102336107da565b005b34801561024157600080fd5b5061024a610800565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561028457818101518382015260200161026c565b50505050905090810190601f1680156102b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102cb57600080fd5b506102e3600160a060020a036004351660243561088e565b604080519115158252519081900360200190f35b34801561030357600080fd5b506102e3610947565b34801561031857600080fd5b50610321610968565b60408051918252519081900360200190f35b34801561033f57600080fd5b50610233610972565b34801561035457600080fd5b506102e3600160a060020a036004358116906024351660443561099a565b34801561037e57600080fd5b50610233610b23565b34801561039357600080fd5b5061039c610b71565b6040805160ff9092168252519081900360200190f35b3480156103be57600080fd5b50610233610b7a565b3480156103d357600080fd5b50610233610c0b565b3480156103e857600080fd5b50610233600160a060020a0360043516602435610c6d565b34801561040c57600080fd5b506102e3600435610d28565b34801561042457600080fd5b506102e3610e6e565b34801561043957600080fd5b50610233610e7e565b34801561044e57600080fd5b506102e3600160a060020a0360043516602435610ece565b34801561047257600080fd5b50610321600160a060020a0360043516610fcc565b34801561049357600080fd5b50610321610fde565b3480156104a857600080fd5b50610233610fe4565b3480156104bd57600080fd5b506102e3611029565b3480156104d257600080fd5b50610233600435611032565b3480156104ea57600080fd5b5061023361105b565b3480156104ff57600080fd5b506105086110c3565b60408051600160a060020a039092168252519081900360200190f35b34801561053057600080fd5b506105086110d2565b34801561054557600080fd5b506102e36110e1565b34801561055a57600080fd5b506102e36110f0565b34801561056f57600080fd5b50610233600160a060020a0360043516611112565b34801561059057600080fd5b5061024a611162565b3480156105a557600080fd5b506102e36111bd565b3480156105ba57600080fd5b5060408051602060048035808201358381028086018501909652808552610233953695939460249493850192918291850190849080828437509497506111cb9650505050505050565b34801561060f57600080fd5b506103216112d3565b34801561062457600080fd5b50610233600160a060020a03600435166112d9565b34801561064557600080fd5b5061023361131f565b34801561065a57600080fd5b506102e3600160a060020a0360043516602435611349565b34801561067e57600080fd5b50610321611451565b34801561069357600080fd5b50610233600435611457565b3480156106ab57600080fd5b50610233600435602435611532565b3480156106c657600080fd5b50610321611569565b3480156106db57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102e3948235600160a060020a031694602480359536959460649492019190819084018382808284375094975061156f9650505050505050565b34801561074457600080fd5b506102e3600160a060020a036004351660243561168a565b34801561076857600080fd5b50610321600160a060020a03600435811690602435166116cb565b34801561078f57600080fd5b50610233600160a060020a03600435166116e8565b3480156107b057600080fd5b50610233600160a060020a03600435166118d6565b3480156107d157600080fd5b50610508611944565b600054600160a060020a031633146107f157600080fd5b6007805460ff19166001179055565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108865780601f1061085b57610100808354040283529160200191610886565b820191906000526020600020905b81548152906001019060200180831161086957829003601f168201915b505050505081565b6000600261089e3660441461195a565b82158015906108cf5750336000908152600c60209081526040808320600160a060020a038816845290915290205415155b156108dd5760009150610940565b336000818152600c60209081526040808320600160a060020a03891680855290835292819020879055805187815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a3600191505b5092915050565b600d5474010000000000000000000000000000000000000000900460ff1681565b6004546003540190565b600054600160a060020a0316331461098957600080fd5b6007805461ff001916610100179055565b60075460009081906109b39060ff16151560011461195a565b600d5460b060020a900460ff16156109ca57600080fd5b60036109d83660641461195a565b600160a060020a03851615156109f15760009250610b1a565b600160a060020a0386166000908152600b6020526040902054841115610a1a5760009250610b1a565b600160a060020a0386166000908152600c60209081526040808320338452909152902054915083821015610a515760009250610b1a565b600160a060020a0385166000908152600b6020526040902054610a749085611969565b600160a060020a038087166000908152600b60205260408082209390935590881681522054610aa3908561198d565b600160a060020a0387166000908152600b6020526040902055610ac6828561198d565b600160a060020a038088166000818152600c6020908152604080832033845282529182902094909455805188815290519289169391926000805160206119a2833981519152929181900390910190a3600192505b50509392505050565b600054600160a060020a03163314610b3a57600080fd5b600d805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b60095460ff1681565b600054600160a060020a03163314610b9157600080fd5b600654421015610ba057600080fd5b6004541515610bae57600080fd5b600754610100900460ff161515610bc457600080fd5b60008054600160a060020a03168152600b6020526040902054600454610bea9190611969565b60008054600160a060020a03168152600b6020526040812091909155600455565b600054600160a060020a03163314610c2257600080fd5b600d547501000000000000000000000000000000000000000000900460ff1615610c4b57600080fd5b600d805476ff0000000000000000000000000000000000000000000019169055565b600054600160a060020a03163314610c8457600080fd5b6002610c923660441461195a565b60075462010000900460ff1615610ca857600080fd5b600160a060020a0383166000908152600b6020526040902054610ccb9083611969565b600160a060020a0384166000908152600b6020526040902055600354610cf19083611969565b600355604080518381529051600160a060020a038516916000916000805160206119a28339815191529181900360200190a3505050565b600d5460009060b060020a900460ff1615610d4257600080fd5b336000908152600b6020526040902054821115610d6157506000610e69565b336000908152600b6020526040902054610d7b908361198d565b336000908152600b6020526040902055600354610d98908361198d565b600355600d54604080517f9dc29fac000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a0390921691639dc29fac916044808201926020929091908290030181600087803b158015610e0857600080fd5b505af1158015610e1c573d6000803e3d6000fd5b505050506040513d6020811015610e3257600080fd5b50519050801515610e4257600080fd5b60408051838152905160009133916000805160206119a28339815191529181900360200190a35b919050565b600d5460b060020a900460ff1681565b600054600160a060020a03163314610e9557600080fd5b600d805475ff00000000000000000000000000000000000000000019167501000000000000000000000000000000000000000000179055565b6000806002610edf3660441461195a565b336000908152600c60209081526040808320600160a060020a0389168452909152902054915081841115610f3657336000908152600c60209081526040808320600160a060020a0389168452909152812055610f65565b610f40828561198d565b336000908152600c60209081526040808320600160a060020a038a1684529091529020555b336000818152600c60209081526040808320600160a060020a038a168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a3506001949350505050565b600b6020526000908152604090205481565b60035481565b600154600160a060020a0316331415611027576001546000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790555b565b60075460ff1681565b600054600160a060020a0316331461104957600080fd5b6006541561105657600080fd5b600655565b600054600160a060020a0316331461107257600080fd5b600d547501000000000000000000000000000000000000000000900460ff161561109b57600080fd5b600d805476ff00000000000000000000000000000000000000000000191660b060020a179055565b600d54600160a060020a031681565b600054600160a060020a031681565b60075462010000900460ff1681565b600d547501000000000000000000000000000000000000000000900460ff1681565b600054600160a060020a0316331461112957600080fd5b60078054600160a060020a0390921663010000000276ffffffffffffffffffffffffffffffffffffffff00000019909216919091179055565b600a805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108865780601f1061085b57610100808354040283529160200191610886565b600754610100900460ff1681565b60008054819081908190600160a060020a031633146111e957600080fd5b60075462010000900460ff16156111ff57600080fd5b600092505b84518310156112c3578451600160a060020a039086908590811061122457fe5b9060200190602002015116915074010000000000000000000000000000000000000000858481518110151561125557fe5b9060200190602002015181151561126857fe5b600160a060020a0384166000818152600b602090815260408083208054969095049586019094558351858152935198850198949550919390926000805160206119a283398151915292908290030190a3600190920191611204565b5050600380549092019091555050565b60045481565b600054600160a060020a031633146112f057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461133657600080fd5b6007805462ff0000191662010000179055565b6007546000906113609060ff16151560011461195a565b600d5460b060020a900460ff161561137757600080fd5b60026113853660441461195a565b336000908152600b60205260409020548311156113a55760009150610940565b600160a060020a03841615156113be5760009150610940565b336000908152600b60205260409020546113d8908461198d565b336000908152600b602052604080822092909255600160a060020a038616815220546114049084611969565b600160a060020a0385166000818152600b60209081526040918290209390935580518681529051919233926000805160206119a28339815191529281900390910190a35060019392505050565b60065481565b60075463010000009004600160a060020a0316331461147557600080fd5b60055481111561148457600080fd5b60075463010000009004600160a060020a03166000908152600b60205260409020546114b09082611969565b60075463010000009004600160a060020a03166000908152600b60205260409020556003546114df9082611969565b6003556005546114ef908261198d565b600555600754604080518381529051600160a060020a03630100000090930492909216916000916000805160206119a2833981519152919081900360200190a350565b600054600160a060020a0316331461154957600080fd5b600754610100900460ff161561155e57600080fd5b600491909155600555565b60055481565b600061157b848461088e565b151561158657600080fd5b6040517f8f4ffcb10000000000000000000000000000000000000000000000000000000081523360048201818152602483018690523060448401819052608060648501908152865160848601528651600160a060020a038a1695638f4ffcb195948a94938a939192909160a490910190602085019080838360005b83811015611619578181015183820152602001611601565b50505050905090810190601f1680156116465780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561166857600080fd5b505af115801561167c573d6000803e3d6000fd5b506001979650505050505050565b600080600261169b3660441461195a565b336000908152600c60209081526040808320600160a060020a03891684529091529020549150610f408285611969565b600c60209081526000928352604080842090915290825290205481565b600080548190600160a060020a0316331461170257600080fd5b600160a060020a03831615156117535760008054604051600160a060020a0390911691303180156108fc02929091818181858888f1935050505015801561174d573d6000803e3d6000fd5b506118d1565b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051849350600160a060020a038416916370a082319160248083019260209291908290030181600087803b1580156117b757600080fd5b505af11580156117cb573d6000803e3d6000fd5b505050506040513d60208110156117e157600080fd5b505160008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810185905290519394509085169263a9059cbb92604480840193602093929083900390910190829087803b15801561185757600080fd5b505af115801561186b573d6000803e3d6000fd5b505050506040513d602081101561188157600080fd5b505060005460408051600160a060020a0380871682529092166020830152818101839052517f977a8f1bdcf5f444d404662ea2c090d707ebcef1be61b37fe6ce74d0c6288fb89181900360600190a15b505050565b600054600160a060020a031633146118ed57600080fd5b600d5474010000000000000000000000000000000000000000900460ff161561191557600080fd5b600d805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60075463010000009004600160a060020a031681565b80151561196657600080fd5b50565b60008282016119868482108015906119815750838210155b61195a565b9392505050565b600061199b8383111561195a565b509003905600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a723058206b9865ffdd16655efc6c8fc757dccb556a2a1214cf28f0209b90d0fef90aa9440029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Token *TokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Token *TokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Token *TokenSession) Controller() (common.Address, error) {
	return _Token.Contract.Controller(&_Token.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_Token *TokenCallerSession) Controller() (common.Address, error) {
	return _Token.Contract.Controller(&_Token.CallOpts)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_Token *TokenCaller) CurrentSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currentSupply")
	return *ret0, err
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_Token *TokenSession) CurrentSupply() (*big.Int, error) {
	return _Token.Contract.CurrentSupply(&_Token.CallOpts)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_Token *TokenCallerSession) CurrentSupply() (*big.Int, error) {
	return _Token.Contract.CurrentSupply(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Launched is a free data retrieval call binding the contract method 0x8091f3bf.
//
// Solidity: function launched() constant returns(bool)
func (_Token *TokenCaller) Launched(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "launched")
	return *ret0, err
}

// Launched is a free data retrieval call binding the contract method 0x8091f3bf.
//
// Solidity: function launched() constant returns(bool)
func (_Token *TokenSession) Launched() (bool, error) {
	return _Token.Contract.Launched(&_Token.CallOpts)
}

// Launched is a free data retrieval call binding the contract method 0x8091f3bf.
//
// Solidity: function launched() constant returns(bool)
func (_Token *TokenCallerSession) Launched() (bool, error) {
	return _Token.Contract.Launched(&_Token.CallOpts)
}

// LockedTokenHolder is a free data retrieval call binding the contract method 0x17e139a7.
//
// Solidity: function lockedTokenHolder() constant returns(bool)
func (_Token *TokenCaller) LockedTokenHolder(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "lockedTokenHolder")
	return *ret0, err
}

// LockedTokenHolder is a free data retrieval call binding the contract method 0x17e139a7.
//
// Solidity: function lockedTokenHolder() constant returns(bool)
func (_Token *TokenSession) LockedTokenHolder() (bool, error) {
	return _Token.Contract.LockedTokenHolder(&_Token.CallOpts)
}

// LockedTokenHolder is a free data retrieval call binding the contract method 0x17e139a7.
//
// Solidity: function lockedTokenHolder() constant returns(bool)
func (_Token *TokenCallerSession) LockedTokenHolder() (bool, error) {
	return _Token.Contract.LockedTokenHolder(&_Token.CallOpts)
}

// MintingDone is a free data retrieval call binding the contract method 0x8e2ae564.
//
// Solidity: function mintingDone() constant returns(bool)
func (_Token *TokenCaller) MintingDone(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "mintingDone")
	return *ret0, err
}

// MintingDone is a free data retrieval call binding the contract method 0x8e2ae564.
//
// Solidity: function mintingDone() constant returns(bool)
func (_Token *TokenSession) MintingDone() (bool, error) {
	return _Token.Contract.MintingDone(&_Token.CallOpts)
}

// MintingDone is a free data retrieval call binding the contract method 0x8e2ae564.
//
// Solidity: function mintingDone() constant returns(bool)
func (_Token *TokenCallerSession) MintingDone() (bool, error) {
	return _Token.Contract.MintingDone(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// OwnerTokensFreeDay is a free data retrieval call binding the contract method 0xac4d2e95.
//
// Solidity: function ownerTokensFreeDay() constant returns(uint256)
func (_Token *TokenCaller) OwnerTokensFreeDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ownerTokensFreeDay")
	return *ret0, err
}

// OwnerTokensFreeDay is a free data retrieval call binding the contract method 0xac4d2e95.
//
// Solidity: function ownerTokensFreeDay() constant returns(uint256)
func (_Token *TokenSession) OwnerTokensFreeDay() (*big.Int, error) {
	return _Token.Contract.OwnerTokensFreeDay(&_Token.CallOpts)
}

// OwnerTokensFreeDay is a free data retrieval call binding the contract method 0xac4d2e95.
//
// Solidity: function ownerTokensFreeDay() constant returns(uint256)
func (_Token *TokenCallerSession) OwnerTokensFreeDay() (*big.Int, error) {
	return _Token.Contract.OwnerTokensFreeDay(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenCallerSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// PausingMechanismLocked is a free data retrieval call binding the contract method 0x90912d09.
//
// Solidity: function pausingMechanismLocked() constant returns(bool)
func (_Token *TokenCaller) PausingMechanismLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pausingMechanismLocked")
	return *ret0, err
}

// PausingMechanismLocked is a free data retrieval call binding the contract method 0x90912d09.
//
// Solidity: function pausingMechanismLocked() constant returns(bool)
func (_Token *TokenSession) PausingMechanismLocked() (bool, error) {
	return _Token.Contract.PausingMechanismLocked(&_Token.CallOpts)
}

// PausingMechanismLocked is a free data retrieval call binding the contract method 0x90912d09.
//
// Solidity: function pausingMechanismLocked() constant returns(bool)
func (_Token *TokenCallerSession) PausingMechanismLocked() (bool, error) {
	return _Token.Contract.PausingMechanismLocked(&_Token.CallOpts)
}

// RemaindersSet is a free data retrieval call binding the contract method 0x98a9ae44.
//
// Solidity: function remaindersSet() constant returns(bool)
func (_Token *TokenCaller) RemaindersSet(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "remaindersSet")
	return *ret0, err
}

// RemaindersSet is a free data retrieval call binding the contract method 0x98a9ae44.
//
// Solidity: function remaindersSet() constant returns(bool)
func (_Token *TokenSession) RemaindersSet() (bool, error) {
	return _Token.Contract.RemaindersSet(&_Token.CallOpts)
}

// RemaindersSet is a free data retrieval call binding the contract method 0x98a9ae44.
//
// Solidity: function remaindersSet() constant returns(bool)
func (_Token *TokenCallerSession) RemaindersSet() (bool, error) {
	return _Token.Contract.RemaindersSet(&_Token.CallOpts)
}

// RemainingAuctionable is a free data retrieval call binding the contract method 0xc6e81b07.
//
// Solidity: function remainingAuctionable() constant returns(uint256)
func (_Token *TokenCaller) RemainingAuctionable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "remainingAuctionable")
	return *ret0, err
}

// RemainingAuctionable is a free data retrieval call binding the contract method 0xc6e81b07.
//
// Solidity: function remainingAuctionable() constant returns(uint256)
func (_Token *TokenSession) RemainingAuctionable() (*big.Int, error) {
	return _Token.Contract.RemainingAuctionable(&_Token.CallOpts)
}

// RemainingAuctionable is a free data retrieval call binding the contract method 0xc6e81b07.
//
// Solidity: function remainingAuctionable() constant returns(uint256)
func (_Token *TokenCallerSession) RemainingAuctionable() (*big.Int, error) {
	return _Token.Contract.RemainingAuctionable(&_Token.CallOpts)
}

// RemainingOwner is a free data retrieval call binding the contract method 0xa4f91a2e.
//
// Solidity: function remainingOwner() constant returns(uint256)
func (_Token *TokenCaller) RemainingOwner(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "remainingOwner")
	return *ret0, err
}

// RemainingOwner is a free data retrieval call binding the contract method 0xa4f91a2e.
//
// Solidity: function remainingOwner() constant returns(uint256)
func (_Token *TokenSession) RemainingOwner() (*big.Int, error) {
	return _Token.Contract.RemainingOwner(&_Token.CallOpts)
}

// RemainingOwner is a free data retrieval call binding the contract method 0xa4f91a2e.
//
// Solidity: function remainingOwner() constant returns(uint256)
func (_Token *TokenCallerSession) RemainingOwner() (*big.Int, error) {
	return _Token.Contract.RemainingOwner(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Tokenholder is a free data retrieval call binding the contract method 0x84eba00c.
//
// Solidity: function tokenholder() constant returns(address)
func (_Token *TokenCaller) Tokenholder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "tokenholder")
	return *ret0, err
}

// Tokenholder is a free data retrieval call binding the contract method 0x84eba00c.
//
// Solidity: function tokenholder() constant returns(address)
func (_Token *TokenSession) Tokenholder() (common.Address, error) {
	return _Token.Contract.Tokenholder(&_Token.CallOpts)
}

// Tokenholder is a free data retrieval call binding the contract method 0x84eba00c.
//
// Solidity: function tokenholder() constant returns(address)
func (_Token *TokenCallerSession) Tokenholder() (common.Address, error) {
	return _Token.Contract.Tokenholder(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Launch is a paid mutator transaction binding the contract method 0x02ac8168.
//
// Solidity: function Launch() returns()
func (_Token *TokenTransactor) Launch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "Launch")
}

// Launch is a paid mutator transaction binding the contract method 0x02ac8168.
//
// Solidity: function Launch() returns()
func (_Token *TokenSession) Launch() (*types.Transaction, error) {
	return _Token.Contract.Launch(&_Token.TransactOpts)
}

// Launch is a paid mutator transaction binding the contract method 0x02ac8168.
//
// Solidity: function Launch() returns()
func (_Token *TokenTransactorSession) Launch() (*types.Transaction, error) {
	return _Token.Contract.Launch(&_Token.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _amount, bytes _extraData) returns(bool success)
func (_Token *TokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approveAndCall", _spender, _amount, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _amount, bytes _extraData) returns(bool success)
func (_Token *TokenSession) ApproveAndCall(_spender common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _amount, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _amount, bytes _extraData) returns(bool success)
func (_Token *TokenTransactorSession) ApproveAndCall(_spender common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _amount, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool result)
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool result)
func (_Token *TokenSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool result)
func (_Token *TokenTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Token *TokenTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Token *TokenSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeOwner(&_Token.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Token *TokenTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeOwner(&_Token.TransactOpts, _newOwner)
}

// ClaimAuctionableTokens is a paid mutator transaction binding the contract method 0xb357a552.
//
// Solidity: function claimAuctionableTokens(uint256 amount) returns()
func (_Token *TokenTransactor) ClaimAuctionableTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimAuctionableTokens", amount)
}

// ClaimAuctionableTokens is a paid mutator transaction binding the contract method 0xb357a552.
//
// Solidity: function claimAuctionableTokens(uint256 amount) returns()
func (_Token *TokenSession) ClaimAuctionableTokens(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAuctionableTokens(&_Token.TransactOpts, amount)
}

// ClaimAuctionableTokens is a paid mutator transaction binding the contract method 0xb357a552.
//
// Solidity: function claimAuctionableTokens(uint256 amount) returns()
func (_Token *TokenTransactorSession) ClaimAuctionableTokens(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAuctionableTokens(&_Token.TransactOpts, amount)
}

// ClaimOwnerSupply is a paid mutator transaction binding the contract method 0x3da1eff5.
//
// Solidity: function claimOwnerSupply() returns()
func (_Token *TokenTransactor) ClaimOwnerSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimOwnerSupply")
}

// ClaimOwnerSupply is a paid mutator transaction binding the contract method 0x3da1eff5.
//
// Solidity: function claimOwnerSupply() returns()
func (_Token *TokenSession) ClaimOwnerSupply() (*types.Transaction, error) {
	return _Token.Contract.ClaimOwnerSupply(&_Token.TransactOpts)
}

// ClaimOwnerSupply is a paid mutator transaction binding the contract method 0x3da1eff5.
//
// Solidity: function claimOwnerSupply() returns()
func (_Token *TokenTransactorSession) ClaimOwnerSupply() (*types.Transaction, error) {
	return _Token.Contract.ClaimOwnerSupply(&_Token.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Token *TokenTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Token *TokenSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Token.Contract.ClaimTokens(&_Token.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Token *TokenTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Token.Contract.ClaimTokens(&_Token.TransactOpts, _token)
}

// CompleteMinting is a paid mutator transaction binding the contract method 0xa844545d.
//
// Solidity: function completeMinting() returns()
func (_Token *TokenTransactor) CompleteMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "completeMinting")
}

// CompleteMinting is a paid mutator transaction binding the contract method 0xa844545d.
//
// Solidity: function completeMinting() returns()
func (_Token *TokenSession) CompleteMinting() (*types.Transaction, error) {
	return _Token.Contract.CompleteMinting(&_Token.TransactOpts)
}

// CompleteMinting is a paid mutator transaction binding the contract method 0xa844545d.
//
// Solidity: function completeMinting() returns()
func (_Token *TokenTransactorSession) CompleteMinting() (*types.Transaction, error) {
	return _Token.Contract.CompleteMinting(&_Token.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_Token *TokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_Token *TokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseApproval(&_Token.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_Token *TokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseApproval(&_Token.TransactOpts, _spender, _subtractedValue)
}

// FinalizeRemainders is a paid mutator transaction binding the contract method 0x2276774c.
//
// Solidity: function finalizeRemainders() returns()
func (_Token *TokenTransactor) FinalizeRemainders(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "finalizeRemainders")
}

// FinalizeRemainders is a paid mutator transaction binding the contract method 0x2276774c.
//
// Solidity: function finalizeRemainders() returns()
func (_Token *TokenSession) FinalizeRemainders() (*types.Transaction, error) {
	return _Token.Contract.FinalizeRemainders(&_Token.TransactOpts)
}

// FinalizeRemainders is a paid mutator transaction binding the contract method 0x2276774c.
//
// Solidity: function finalizeRemainders() returns()
func (_Token *TokenTransactorSession) FinalizeRemainders() (*types.Transaction, error) {
	return _Token.Contract.FinalizeRemainders(&_Token.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_Token *TokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_Token *TokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseApproval(&_Token.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_Token *TokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseApproval(&_Token.TransactOpts, _spender, _addedValue)
}

// LockTokenHolder is a paid mutator transaction binding the contract method 0x2bbeac91.
//
// Solidity: function lockTokenHolder() returns()
func (_Token *TokenTransactor) LockTokenHolder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "lockTokenHolder")
}

// LockTokenHolder is a paid mutator transaction binding the contract method 0x2bbeac91.
//
// Solidity: function lockTokenHolder() returns()
func (_Token *TokenSession) LockTokenHolder() (*types.Transaction, error) {
	return _Token.Contract.LockTokenHolder(&_Token.TransactOpts)
}

// LockTokenHolder is a paid mutator transaction binding the contract method 0x2bbeac91.
//
// Solidity: function lockTokenHolder() returns()
func (_Token *TokenTransactorSession) LockTokenHolder() (*types.Transaction, error) {
	return _Token.Contract.LockTokenHolder(&_Token.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_Token *TokenSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_Token *TokenTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, addr, amount)
}

// MultiMint is a paid mutator transaction binding the contract method 0x9a0e4ebb.
//
// Solidity: function multiMint(uint256[] data) returns()
func (_Token *TokenTransactor) MultiMint(opts *bind.TransactOpts, data []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "multiMint", data)
}

// MultiMint is a paid mutator transaction binding the contract method 0x9a0e4ebb.
//
// Solidity: function multiMint(uint256[] data) returns()
func (_Token *TokenSession) MultiMint(data []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MultiMint(&_Token.TransactOpts, data)
}

// MultiMint is a paid mutator transaction binding the contract method 0x9a0e4ebb.
//
// Solidity: function multiMint(uint256[] data) returns()
func (_Token *TokenTransactorSession) MultiMint(data []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MultiMint(&_Token.TransactOpts, data)
}

// NeverPauseAgain is a paid mutator transaction binding the contract method 0x6348eae6.
//
// Solidity: function neverPauseAgain() returns()
func (_Token *TokenTransactor) NeverPauseAgain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "neverPauseAgain")
}

// NeverPauseAgain is a paid mutator transaction binding the contract method 0x6348eae6.
//
// Solidity: function neverPauseAgain() returns()
func (_Token *TokenSession) NeverPauseAgain() (*types.Transaction, error) {
	return _Token.Contract.NeverPauseAgain(&_Token.TransactOpts)
}

// NeverPauseAgain is a paid mutator transaction binding the contract method 0x6348eae6.
//
// Solidity: function neverPauseAgain() returns()
func (_Token *TokenTransactorSession) NeverPauseAgain() (*types.Transaction, error) {
	return _Token.Contract.NeverPauseAgain(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Token *TokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Token *TokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetController(&_Token.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Token *TokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetController(&_Token.TransactOpts, _controller)
}

// SetOwnerFreeDay is a paid mutator transaction binding the contract method 0x81e529cd.
//
// Solidity: function setOwnerFreeDay(uint256 day) returns()
func (_Token *TokenTransactor) SetOwnerFreeDay(opts *bind.TransactOpts, day *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setOwnerFreeDay", day)
}

// SetOwnerFreeDay is a paid mutator transaction binding the contract method 0x81e529cd.
//
// Solidity: function setOwnerFreeDay(uint256 day) returns()
func (_Token *TokenSession) SetOwnerFreeDay(day *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetOwnerFreeDay(&_Token.TransactOpts, day)
}

// SetOwnerFreeDay is a paid mutator transaction binding the contract method 0x81e529cd.
//
// Solidity: function setOwnerFreeDay(uint256 day) returns()
func (_Token *TokenTransactorSession) SetOwnerFreeDay(day *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetOwnerFreeDay(&_Token.TransactOpts, day)
}

// SetRemainders is a paid mutator transaction binding the contract method 0xc27549d9.
//
// Solidity: function setRemainders(uint256 _remainingOwner, uint256 _remainingAuctionable) returns()
func (_Token *TokenTransactor) SetRemainders(opts *bind.TransactOpts, _remainingOwner *big.Int, _remainingAuctionable *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setRemainders", _remainingOwner, _remainingAuctionable)
}

// SetRemainders is a paid mutator transaction binding the contract method 0xc27549d9.
//
// Solidity: function setRemainders(uint256 _remainingOwner, uint256 _remainingAuctionable) returns()
func (_Token *TokenSession) SetRemainders(_remainingOwner *big.Int, _remainingAuctionable *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetRemainders(&_Token.TransactOpts, _remainingOwner, _remainingAuctionable)
}

// SetRemainders is a paid mutator transaction binding the contract method 0xc27549d9.
//
// Solidity: function setRemainders(uint256 _remainingOwner, uint256 _remainingAuctionable) returns()
func (_Token *TokenTransactorSession) SetRemainders(_remainingOwner *big.Int, _remainingAuctionable *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetRemainders(&_Token.TransactOpts, _remainingOwner, _remainingAuctionable)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(address _th) returns()
func (_Token *TokenTransactor) SetTokenHolder(opts *bind.TransactOpts, _th common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTokenHolder", _th)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(address _th) returns()
func (_Token *TokenSession) SetTokenHolder(_th common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetTokenHolder(&_Token.TransactOpts, _th)
}

// SetTokenHolder is a paid mutator transaction binding the contract method 0xf29d2f28.
//
// Solidity: function setTokenHolder(address _th) returns()
func (_Token *TokenTransactorSession) SetTokenHolder(_th common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetTokenHolder(&_Token.TransactOpts, _th)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TokenLogTokenTransferIterator is returned from FilterLogTokenTransfer and is used to iterate over the raw logs and unpacked data for LogTokenTransfer events raised by the Token contract.
type TokenLogTokenTransferIterator struct {
	Event *TokenLogTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenLogTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogTokenTransfer)
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
		it.Event = new(TokenLogTokenTransfer)
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
func (it *TokenLogTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogTokenTransfer represents a LogTokenTransfer event raised by the Token contract.
type TokenLogTokenTransfer struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTokenTransfer is a free log retrieval operation binding the contract event 0x977a8f1bdcf5f444d404662ea2c090d707ebcef1be61b37fe6ce74d0c6288fb8.
//
// Solidity: event logTokenTransfer(address token, address to, uint256 amount)
func (_Token *TokenFilterer) FilterLogTokenTransfer(opts *bind.FilterOpts) (*TokenLogTokenTransferIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "logTokenTransfer")
	if err != nil {
		return nil, err
	}
	return &TokenLogTokenTransferIterator{contract: _Token.contract, event: "logTokenTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTokenTransfer is a free log subscription operation binding the contract event 0x977a8f1bdcf5f444d404662ea2c090d707ebcef1be61b37fe6ce74d0c6288fb8.
//
// Solidity: event logTokenTransfer(address token, address to, uint256 amount)
func (_Token *TokenFilterer) WatchLogTokenTransfer(opts *bind.WatchOpts, sink chan<- *TokenLogTokenTransfer) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "logTokenTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogTokenTransfer)
				if err := _Token.contract.UnpackLog(event, "logTokenTransfer", log); err != nil {
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
