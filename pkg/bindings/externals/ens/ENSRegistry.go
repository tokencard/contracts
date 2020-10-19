// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ens

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

// ENSRegistryABI is the input ABI used to generate the binding from.
const ENSRegistryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ENSRegistryBin is the compiled bytecode used for deploying new contracts.
<<<<<<< HEAD
var ENSRegistryBin = "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055610939806100596000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80635b0fc9c3116100715780635b0fc9c3146101e85780635ef2c7f014610214578063a22cb46514610261578063cf4088231461028f578063e985e9c5146102d6578063f79fe53814610318576100b4565b80630178b8bf146100b957806302571be3146100f257806306ab59231461010f57806314ab90381461015357806316a25cbd146101825780631896f70a146101bc575b600080fd5b6100d6600480360360208110156100cf57600080fd5b5035610335565b604080516001600160a01b039092168252519081900360200190f35b6100d66004803603602081101561010857600080fd5b5035610356565b6101416004803603606081101561012557600080fd5b50803590602081013590604001356001600160a01b0316610386565b60408051918252519081900360200190f35b6101806004803603604081101561016957600080fd5b508035906020013567ffffffffffffffff16610456565b005b61019f6004803603602081101561019857600080fd5b503561052a565b6040805167ffffffffffffffff9092168252519081900360200190f35b610180600480360360408110156101d257600080fd5b50803590602001356001600160a01b0316610550565b610180600480360360408110156101fe57600080fd5b50803590602001356001600160a01b0316610616565b610180600480360360a081101561022a57600080fd5b5080359060208101359060408101356001600160a01b03908116916060810135909116906080013567ffffffffffffffff166106b9565b6101806004803603604081101561027757600080fd5b506001600160a01b03813516906020013515156106db565b610180600480360360808110156102a557600080fd5b5080359060208101356001600160a01b03908116916040810135909116906060013567ffffffffffffffff16610749565b610304600480360360408110156102ec57600080fd5b506001600160a01b0381358116916020013516610764565b604080519115158252519081900360200190f35b6103046004803603602081101561032e57600080fd5b5035610792565b6000818152602081905260409020600101546001600160a01b03165b919050565b6000818152602081905260408120546001600160a01b031630811415610380576000915050610351565b92915050565b60008381526020819052604081205484906001600160a01b0316338114806103d157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6103da57600080fd5b6040805160208082018990528183018890528251808303840181526060909201909252805191012061040c81866107af565b604080516001600160a01b03871681529051879189917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360200190a39695505050505050565b60008281526020819052604090205482906001600160a01b0316338114806104a157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6104aa57600080fd5b6040805167ffffffffffffffff85168152905185917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a25050600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b03163381148061059b57506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6105a457600080fd5b604080516001600160a01b0385168152905185917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a2505060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b03163381148061066157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61066a57600080fd5b61067484846107af565b604080516001600160a01b0385168152905185917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250505050565b60006106c6868686610386565b90506106d38184846107dd565b505050505050565b3360008181526001602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6107538484610616565b61075e8483836107dd565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6000908152602081905260409020546001600160a01b0316151590565b60009182526020829052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6000838152602081905260409020600101546001600160a01b03838116911614610863576000838152602081815260409182902060010180546001600160a01b0319166001600160a01b0386169081179091558251908152915185927f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a092908290030190a25b60008381526020819052604090206001015467ffffffffffffffff828116600160a01b90920416146108ff5760008381526020818152604091829020600101805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff8616908102919091179091558251908152915185927f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa6892908290030190a25b50505056fea265627a7a723158207f3c9ab060926571865045bde8d1bcab6d1fe5e070d4d31e19b5c37def6112e864736f6c63430005110032"
||||||| constructed merge base
var ENSRegistryBin = "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055610939806100596000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80635b0fc9c3116100715780635b0fc9c3146101e85780635ef2c7f014610214578063a22cb46514610261578063cf4088231461028f578063e985e9c5146102d6578063f79fe53814610318576100b4565b80630178b8bf146100b957806302571be3146100f257806306ab59231461010f57806314ab90381461015357806316a25cbd146101825780631896f70a146101bc575b600080fd5b6100d6600480360360208110156100cf57600080fd5b5035610335565b604080516001600160a01b039092168252519081900360200190f35b6100d66004803603602081101561010857600080fd5b5035610356565b6101416004803603606081101561012557600080fd5b50803590602081013590604001356001600160a01b0316610386565b60408051918252519081900360200190f35b6101806004803603604081101561016957600080fd5b508035906020013567ffffffffffffffff16610456565b005b61019f6004803603602081101561019857600080fd5b503561052a565b6040805167ffffffffffffffff9092168252519081900360200190f35b610180600480360360408110156101d257600080fd5b50803590602001356001600160a01b0316610550565b610180600480360360408110156101fe57600080fd5b50803590602001356001600160a01b0316610616565b610180600480360360a081101561022a57600080fd5b5080359060208101359060408101356001600160a01b03908116916060810135909116906080013567ffffffffffffffff166106b9565b6101806004803603604081101561027757600080fd5b506001600160a01b03813516906020013515156106db565b610180600480360360808110156102a557600080fd5b5080359060208101356001600160a01b03908116916040810135909116906060013567ffffffffffffffff16610749565b610304600480360360408110156102ec57600080fd5b506001600160a01b0381358116916020013516610764565b604080519115158252519081900360200190f35b6103046004803603602081101561032e57600080fd5b5035610792565b6000818152602081905260409020600101546001600160a01b03165b919050565b6000818152602081905260408120546001600160a01b031630811415610380576000915050610351565b92915050565b60008381526020819052604081205484906001600160a01b0316338114806103d157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6103da57600080fd5b6040805160208082018990528183018890528251808303840181526060909201909252805191012061040c81866107af565b604080516001600160a01b03871681529051879189917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360200190a39695505050505050565b60008281526020819052604090205482906001600160a01b0316338114806104a157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6104aa57600080fd5b6040805167ffffffffffffffff85168152905185917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a25050600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b03163381148061059b57506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6105a457600080fd5b604080516001600160a01b0385168152905185917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a2505060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b03163381148061066157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61066a57600080fd5b61067484846107af565b604080516001600160a01b0385168152905185917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250505050565b60006106c6868686610386565b90506106d38184846107dd565b505050505050565b3360008181526001602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6107538484610616565b61075e8483836107dd565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6000908152602081905260409020546001600160a01b0316151590565b60009182526020829052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6000838152602081905260409020600101546001600160a01b03838116911614610863576000838152602081815260409182902060010180546001600160a01b0319166001600160a01b0386169081179091558251908152915185927f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a092908290030190a25b60008381526020819052604090206001015467ffffffffffffffff828116600160a01b90920416146108ff5760008381526020818152604091829020600101805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff8616908102919091179091558251908152915185927f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa6892908290030190a25b50505056fea265627a7a723158205acb335f362434847a26a419f5452838532c6709e8e789e9f4852c20ac4f95cf64736f6c634300050f0032"
=======
<<<<<<< HEAD
var ENSRegistryBin = "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055610939806100596000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80635b0fc9c3116100715780635b0fc9c3146101e85780635ef2c7f014610214578063a22cb46514610261578063cf4088231461028f578063e985e9c5146102d6578063f79fe53814610318576100b4565b80630178b8bf146100b957806302571be3146100f257806306ab59231461010f57806314ab90381461015357806316a25cbd146101825780631896f70a146101bc575b600080fd5b6100d6600480360360208110156100cf57600080fd5b5035610335565b604080516001600160a01b039092168252519081900360200190f35b6100d66004803603602081101561010857600080fd5b5035610356565b6101416004803603606081101561012557600080fd5b50803590602081013590604001356001600160a01b0316610386565b60408051918252519081900360200190f35b6101806004803603604081101561016957600080fd5b508035906020013567ffffffffffffffff16610456565b005b61019f6004803603602081101561019857600080fd5b503561052a565b6040805167ffffffffffffffff9092168252519081900360200190f35b610180600480360360408110156101d257600080fd5b50803590602001356001600160a01b0316610550565b610180600480360360408110156101fe57600080fd5b50803590602001356001600160a01b0316610616565b610180600480360360a081101561022a57600080fd5b5080359060208101359060408101356001600160a01b03908116916060810135909116906080013567ffffffffffffffff166106b9565b6101806004803603604081101561027757600080fd5b506001600160a01b03813516906020013515156106db565b610180600480360360808110156102a557600080fd5b5080359060208101356001600160a01b03908116916040810135909116906060013567ffffffffffffffff16610749565b610304600480360360408110156102ec57600080fd5b506001600160a01b0381358116916020013516610764565b604080519115158252519081900360200190f35b6103046004803603602081101561032e57600080fd5b5035610792565b6000818152602081905260409020600101546001600160a01b03165b919050565b6000818152602081905260408120546001600160a01b031630811415610380576000915050610351565b92915050565b60008381526020819052604081205484906001600160a01b0316338114806103d157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6103da57600080fd5b6040805160208082018990528183018890528251808303840181526060909201909252805191012061040c81866107af565b604080516001600160a01b03871681529051879189917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360200190a39695505050505050565b60008281526020819052604090205482906001600160a01b0316338114806104a157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6104aa57600080fd5b6040805167ffffffffffffffff85168152905185917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a25050600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b03163381148061059b57506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6105a457600080fd5b604080516001600160a01b0385168152905185917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a2505060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b03163381148061066157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61066a57600080fd5b61067484846107af565b604080516001600160a01b0385168152905185917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250505050565b60006106c6868686610386565b90506106d38184846107dd565b505050505050565b3360008181526001602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6107538484610616565b61075e8483836107dd565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6000908152602081905260409020546001600160a01b0316151590565b60009182526020829052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6000838152602081905260409020600101546001600160a01b03838116911614610863576000838152602081815260409182902060010180546001600160a01b0319166001600160a01b0386169081179091558251908152915185927f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a092908290030190a25b60008381526020819052604090206001015467ffffffffffffffff828116600160a01b90920416146108ff5760008381526020818152604091829020600101805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff8616908102919091179091558251908152915185927f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa6892908290030190a25b50505056fea265627a7a723158205acb335f362434847a26a419f5452838532c6709e8e789e9f4852c20ac4f95cf64736f6c634300050f0032"
||||||| constructed merge base
var ENSRegistryBin = "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b031916331790556104e9806100596000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806314ab90381161005b57806314ab90381461010c57806316a25cbd146101395780631896f70a146101735780635b0fc9c31461019f5761007d565b80630178b8bf1461008257806302571be3146100bb57806306ab5923146100d8575b600080fd5b61009f6004803603602081101561009857600080fd5b50356101cb565b604080516001600160a01b039092168252519081900360200190f35b61009f600480360360208110156100d157600080fd5b50356101e9565b61010a600480360360608110156100ee57600080fd5b50803590602081013590604001356001600160a01b0316610204565b005b61010a6004803603604081101561012257600080fd5b508035906020013567ffffffffffffffff166102c1565b6101566004803603602081101561014f57600080fd5b5035610365565b6040805167ffffffffffffffff9092168252519081900360200190f35b61010a6004803603604081101561018957600080fd5b50803590602001356001600160a01b031661038b565b61010a600480360360408110156101b557600080fd5b50803590602001356001600160a01b0316610421565b6000908152602081905260409020600101546001600160a01b031690565b6000908152602081905260409020546001600160a01b031690565b60008381526020819052604090205483906001600160a01b0316331461022957600080fd5b60408051602080820187905281830186905282518083038401815260608301808552815191909201206001600160a01b0386169091529151859187917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360800190a3600090815260208190526040902080546001600160a01b0319166001600160a01b039390931692909217909155505050565b60008281526020819052604090205482906001600160a01b031633146102e657600080fd5b6040805167ffffffffffffffff84168152905184917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a250600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b031633146103b057600080fd5b604080516001600160a01b0384168152905184917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a25060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b0316331461044657600080fd5b604080516001600160a01b0384168152905184917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a25060009182526020829052604090912080546001600160a01b0319166001600160a01b0390921691909117905556fea265627a7a72315820824ee377f3d2b0c86755b6945656550383349adf36b6244d4ebfd08bc542950664736f6c634300050f0032"
=======
var ENSRegistryBin = "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b031916331790556104e9806100596000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806314ab90381161005b57806314ab90381461010c57806316a25cbd146101395780631896f70a146101735780635b0fc9c31461019f5761007d565b80630178b8bf1461008257806302571be3146100bb57806306ab5923146100d8575b600080fd5b61009f6004803603602081101561009857600080fd5b50356101cb565b604080516001600160a01b039092168252519081900360200190f35b61009f600480360360208110156100d157600080fd5b50356101e9565b61010a600480360360608110156100ee57600080fd5b50803590602081013590604001356001600160a01b0316610204565b005b61010a6004803603604081101561012257600080fd5b508035906020013567ffffffffffffffff166102c1565b6101566004803603602081101561014f57600080fd5b5035610365565b6040805167ffffffffffffffff9092168252519081900360200190f35b61010a6004803603604081101561018957600080fd5b50803590602001356001600160a01b031661038b565b61010a600480360360408110156101b557600080fd5b50803590602001356001600160a01b0316610421565b6000908152602081905260409020600101546001600160a01b031690565b6000908152602081905260409020546001600160a01b031690565b60008381526020819052604090205483906001600160a01b0316331461022957600080fd5b60408051602080820187905281830186905282518083038401815260608301808552815191909201206001600160a01b0386169091529151859187917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360800190a3600090815260208190526040902080546001600160a01b0319166001600160a01b039390931692909217909155505050565b60008281526020819052604090205482906001600160a01b031633146102e657600080fd5b6040805167ffffffffffffffff84168152905184917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a250600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b031633146103b057600080fd5b604080516001600160a01b0384168152905184917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a25060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b0316331461044657600080fd5b604080516001600160a01b0384168152905184917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a25060009182526020829052604090912080546001600160a01b0319166001600160a01b0390921691909117905556fea265627a7a72315820ac7b3495cd4219baba322aaf208f04d49951a18fb0984dd27d8e414dc2924f1164736f6c63430005110032"
>>>>>>> Upgrade to solc 0.6.4 for Wallet
>>>>>>> Upgrade to solc 0.6.4 for Wallet

// DeployENSRegistry deploys a new Ethereum contract, binding an instance of ENSRegistry to it.
func DeployENSRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ENSRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ENSRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSRegistry{ENSRegistryCaller: ENSRegistryCaller{contract: contract}, ENSRegistryTransactor: ENSRegistryTransactor{contract: contract}, ENSRegistryFilterer: ENSRegistryFilterer{contract: contract}}, nil
}

// ENSRegistry is an auto generated Go binding around an Ethereum contract.
type ENSRegistry struct {
	ENSRegistryCaller     // Read-only binding to the contract
	ENSRegistryTransactor // Write-only binding to the contract
	ENSRegistryFilterer   // Log filterer for contract events
}

// ENSRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSRegistrySession struct {
	Contract     *ENSRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSRegistryCallerSession struct {
	Contract *ENSRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ENSRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSRegistryTransactorSession struct {
	Contract     *ENSRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ENSRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRegistryRaw struct {
	Contract *ENSRegistry // Generic contract binding to access the raw methods on
}

// ENSRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSRegistryCallerRaw struct {
	Contract *ENSRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ENSRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSRegistryTransactorRaw struct {
	Contract *ENSRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSRegistry creates a new instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistry(address common.Address, backend bind.ContractBackend) (*ENSRegistry, error) {
	contract, err := bindENSRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSRegistry{ENSRegistryCaller: ENSRegistryCaller{contract: contract}, ENSRegistryTransactor: ENSRegistryTransactor{contract: contract}, ENSRegistryFilterer: ENSRegistryFilterer{contract: contract}}, nil
}

// NewENSRegistryCaller creates a new read-only instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryCaller(address common.Address, caller bind.ContractCaller) (*ENSRegistryCaller, error) {
	contract, err := bindENSRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryCaller{contract: contract}, nil
}

// NewENSRegistryTransactor creates a new write-only instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSRegistryTransactor, error) {
	contract, err := bindENSRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryTransactor{contract: contract}, nil
}

// NewENSRegistryFilterer creates a new log filterer instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSRegistryFilterer, error) {
	contract, err := bindENSRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryFilterer{contract: contract}, nil
}

// bindENSRegistry binds a generic wrapper to an already deployed contract.
func bindENSRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistry *ENSRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSRegistry.Contract.ENSRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistry *ENSRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistry.Contract.ENSRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistry *ENSRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistry.Contract.ENSRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistry *ENSRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistry *ENSRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistry *ENSRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ENSRegistry *ENSRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ENSRegistry *ENSRegistrySession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ENSRegistry.Contract.IsApprovedForAll(&_ENSRegistry.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ENSRegistry *ENSRegistryCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ENSRegistry.Contract.IsApprovedForAll(&_ENSRegistry.CallOpts, _owner, _operator)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 _node) constant returns(address)
func (_ENSRegistry *ENSRegistryCaller) Owner(opts *bind.CallOpts, _node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "owner", _node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 _node) constant returns(address)
func (_ENSRegistry *ENSRegistrySession) Owner(_node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Owner(&_ENSRegistry.CallOpts, _node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 _node) constant returns(address)
func (_ENSRegistry *ENSRegistryCallerSession) Owner(_node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Owner(&_ENSRegistry.CallOpts, _node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 _node) constant returns(bool)
func (_ENSRegistry *ENSRegistryCaller) RecordExists(opts *bind.CallOpts, _node [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "recordExists", _node)
	return *ret0, err
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 _node) constant returns(bool)
func (_ENSRegistry *ENSRegistrySession) RecordExists(_node [32]byte) (bool, error) {
	return _ENSRegistry.Contract.RecordExists(&_ENSRegistry.CallOpts, _node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 _node) constant returns(bool)
func (_ENSRegistry *ENSRegistryCallerSession) RecordExists(_node [32]byte) (bool, error) {
	return _ENSRegistry.Contract.RecordExists(&_ENSRegistry.CallOpts, _node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 _node) constant returns(address)
func (_ENSRegistry *ENSRegistryCaller) Resolver(opts *bind.CallOpts, _node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "resolver", _node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 _node) constant returns(address)
func (_ENSRegistry *ENSRegistrySession) Resolver(_node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Resolver(&_ENSRegistry.CallOpts, _node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 _node) constant returns(address)
func (_ENSRegistry *ENSRegistryCallerSession) Resolver(_node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Resolver(&_ENSRegistry.CallOpts, _node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 _node) constant returns(uint64)
func (_ENSRegistry *ENSRegistryCaller) Ttl(opts *bind.CallOpts, _node [32]byte) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ENSRegistry.contract.Call(opts, out, "ttl", _node)
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 _node) constant returns(uint64)
func (_ENSRegistry *ENSRegistrySession) Ttl(_node [32]byte) (uint64, error) {
	return _ENSRegistry.Contract.Ttl(&_ENSRegistry.CallOpts, _node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 _node) constant returns(uint64)
func (_ENSRegistry *ENSRegistryCallerSession) Ttl(_node [32]byte) (uint64, error) {
	return _ENSRegistry.Contract.Ttl(&_ENSRegistry.CallOpts, _node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ENSRegistry *ENSRegistrySession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetApprovalForAll(&_ENSRegistry.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetApprovalForAll(&_ENSRegistry.TransactOpts, _operator, _approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 _node, address _owner) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetOwner(opts *bind.TransactOpts, _node [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setOwner", _node, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 _node, address _owner) returns()
func (_ENSRegistry *ENSRegistrySession) SetOwner(_node [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetOwner(&_ENSRegistry.TransactOpts, _node, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 _node, address _owner) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetOwner(_node [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetOwner(&_ENSRegistry.TransactOpts, _node, _owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 _node, address _owner, address _resolver, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetRecord(opts *bind.TransactOpts, _node [32]byte, _owner common.Address, _resolver common.Address, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setRecord", _node, _owner, _resolver, _ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 _node, address _owner, address _resolver, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetRecord(_node [32]byte, _owner common.Address, _resolver common.Address, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetRecord(&_ENSRegistry.TransactOpts, _node, _owner, _resolver, _ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 _node, address _owner, address _resolver, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetRecord(_node [32]byte, _owner common.Address, _resolver common.Address, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetRecord(&_ENSRegistry.TransactOpts, _node, _owner, _resolver, _ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 _node, address _resolver) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetResolver(opts *bind.TransactOpts, _node [32]byte, _resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setResolver", _node, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 _node, address _resolver) returns()
func (_ENSRegistry *ENSRegistrySession) SetResolver(_node [32]byte, _resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetResolver(&_ENSRegistry.TransactOpts, _node, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 _node, address _resolver) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetResolver(_node [32]byte, _resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetResolver(&_ENSRegistry.TransactOpts, _node, _resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 _node, bytes32 _label, address _owner) returns(bytes32)
func (_ENSRegistry *ENSRegistryTransactor) SetSubnodeOwner(opts *bind.TransactOpts, _node [32]byte, _label [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setSubnodeOwner", _node, _label, _owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 _node, bytes32 _label, address _owner) returns(bytes32)
func (_ENSRegistry *ENSRegistrySession) SetSubnodeOwner(_node [32]byte, _label [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeOwner(&_ENSRegistry.TransactOpts, _node, _label, _owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 _node, bytes32 _label, address _owner) returns(bytes32)
func (_ENSRegistry *ENSRegistryTransactorSession) SetSubnodeOwner(_node [32]byte, _label [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeOwner(&_ENSRegistry.TransactOpts, _node, _label, _owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 _node, bytes32 _label, address _owner, address _resolver, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetSubnodeRecord(opts *bind.TransactOpts, _node [32]byte, _label [32]byte, _owner common.Address, _resolver common.Address, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setSubnodeRecord", _node, _label, _owner, _resolver, _ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 _node, bytes32 _label, address _owner, address _resolver, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetSubnodeRecord(_node [32]byte, _label [32]byte, _owner common.Address, _resolver common.Address, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeRecord(&_ENSRegistry.TransactOpts, _node, _label, _owner, _resolver, _ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 _node, bytes32 _label, address _owner, address _resolver, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetSubnodeRecord(_node [32]byte, _label [32]byte, _owner common.Address, _resolver common.Address, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeRecord(&_ENSRegistry.TransactOpts, _node, _label, _owner, _resolver, _ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 _node, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetTTL(opts *bind.TransactOpts, _node [32]byte, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setTTL", _node, _ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 _node, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetTTL(_node [32]byte, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetTTL(&_ENSRegistry.TransactOpts, _node, _ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 _node, uint64 _ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetTTL(_node [32]byte, _ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetTTL(&_ENSRegistry.TransactOpts, _node, _ttl)
}

// ENSRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ENSRegistry contract.
type ENSRegistryApprovalForAllIterator struct {
	Event *ENSRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ENSRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryApprovalForAll)
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
		it.Event = new(ENSRegistryApprovalForAll)
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
func (it *ENSRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryApprovalForAll represents a ApprovalForAll event raised by the ENSRegistry contract.
type ENSRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ENSRegistry *ENSRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ENSRegistryApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryApprovalForAllIterator{contract: _ENSRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ENSRegistry *ENSRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ENSRegistryApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryApprovalForAll)
				if err := _ENSRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ENSRegistry *ENSRegistryFilterer) ParseApprovalForAll(log types.Log) (*ENSRegistryApprovalForAll, error) {
	event := new(ENSRegistryApprovalForAll)
	if err := _ENSRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ENSRegistryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENSRegistry contract.
type ENSRegistryNewOwnerIterator struct {
	Event *ENSRegistryNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewOwner)
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
		it.Event = new(ENSRegistryNewOwner)
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
func (it *ENSRegistryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewOwner represents a NewOwner event raised by the ENSRegistry contract.
type ENSRegistryNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed _node, bytes32 indexed _label, address _owner)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewOwner(opts *bind.FilterOpts, _node [][32]byte, _label [][32]byte) (*ENSRegistryNewOwnerIterator, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}
	var _labelRule []interface{}
	for _, _labelItem := range _label {
		_labelRule = append(_labelRule, _labelItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewOwner", _nodeRule, _labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewOwnerIterator{contract: _ENSRegistry.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed _node, bytes32 indexed _label, address _owner)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewOwner, _node [][32]byte, _label [][32]byte) (event.Subscription, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}
	var _labelRule []interface{}
	for _, _labelItem := range _label {
		_labelRule = append(_labelRule, _labelItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewOwner", _nodeRule, _labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewOwner)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed _node, bytes32 indexed _label, address _owner)
func (_ENSRegistry *ENSRegistryFilterer) ParseNewOwner(log types.Log) (*ENSRegistryNewOwner, error) {
	event := new(ENSRegistryNewOwner)
	if err := _ENSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ENSRegistryNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENSRegistry contract.
type ENSRegistryNewResolverIterator struct {
	Event *ENSRegistryNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewResolver)
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
		it.Event = new(ENSRegistryNewResolver)
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
func (it *ENSRegistryNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewResolver represents a NewResolver event raised by the ENSRegistry contract.
type ENSRegistryNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed _node, address _resolver)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewResolver(opts *bind.FilterOpts, _node [][32]byte) (*ENSRegistryNewResolverIterator, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewResolver", _nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewResolverIterator{contract: _ENSRegistry.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed _node, address _resolver)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewResolver, _node [][32]byte) (event.Subscription, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewResolver", _nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewResolver)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed _node, address _resolver)
func (_ENSRegistry *ENSRegistryFilterer) ParseNewResolver(log types.Log) (*ENSRegistryNewResolver, error) {
	event := new(ENSRegistryNewResolver)
	if err := _ENSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ENSRegistryNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the ENSRegistry contract.
type ENSRegistryNewTTLIterator struct {
	Event *ENSRegistryNewTTL // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewTTL)
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
		it.Event = new(ENSRegistryNewTTL)
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
func (it *ENSRegistryNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewTTL represents a NewTTL event raised by the ENSRegistry contract.
type ENSRegistryNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed _node, uint64 _ttl)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewTTL(opts *bind.FilterOpts, _node [][32]byte) (*ENSRegistryNewTTLIterator, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewTTL", _nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewTTLIterator{contract: _ENSRegistry.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed _node, uint64 _ttl)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewTTL, _node [][32]byte) (event.Subscription, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewTTL", _nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewTTL)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed _node, uint64 _ttl)
func (_ENSRegistry *ENSRegistryFilterer) ParseNewTTL(log types.Log) (*ENSRegistryNewTTL, error) {
	event := new(ENSRegistryNewTTL)
	if err := _ENSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ENSRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENSRegistry contract.
type ENSRegistryTransferIterator struct {
	Event *ENSRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *ENSRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryTransfer)
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
		it.Event = new(ENSRegistryTransfer)
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
func (it *ENSRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryTransfer represents a Transfer event raised by the ENSRegistry contract.
type ENSRegistryTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed _node, address _owner)
func (_ENSRegistry *ENSRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, _node [][32]byte) (*ENSRegistryTransferIterator, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "Transfer", _nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryTransferIterator{contract: _ENSRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed _node, address _owner)
func (_ENSRegistry *ENSRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSRegistryTransfer, _node [][32]byte) (event.Subscription, error) {

	var _nodeRule []interface{}
	for _, _nodeItem := range _node {
		_nodeRule = append(_nodeRule, _nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "Transfer", _nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryTransfer)
				if err := _ENSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed _node, address _owner)
func (_ENSRegistry *ENSRegistryFilterer) ParseTransfer(log types.Log) (*ENSRegistryTransfer, error) {
	event := new(ENSRegistryTransfer)
	if err := _ENSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
