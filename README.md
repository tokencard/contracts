# TokenCard Smart Contracts

This repository contains the Smart Contract code needed to power the TokenCard, written in Solidity for execution on the EVM.

## Overview

The functionality in this repository has been designed to help users protect their tokens, within their own instance of a Consumer Smart Contract Wallet which they can configure to their liking. The functionality within the TokenWallet has been designed to limit a user's exposure to loss in the event that a user has had their Private Key compromised.

Each user deploys their own instance of the TokenWallet ([wallet.sol](/contracts/wallet.sol)) which interacts with an exchange-rate oracle ([oracle.sol](/contracts/oracle.sol)). The individual TokenWallet contracts use the [Ethereum Name Service (ENS)](https://ens.domains/) to resolve our exchange rate oracle. 

In order to help users protect their tokens in the event that their Private Key gets compromised, we present the TokenWallet. 


## Contract communication

```
                              ┌──────────────────┐ 
  ┌────────────────┐          │                  │ 
┌─┤                │          │                  │ 
│ │                │    ┌─────▶    oracle.sol    │ 
│ │                │    │     │                  │ 
│ │                │    │     │                  │ 
│ │   wallet.sol   ├────┘     └──────────────────┘ 
│ │                ├─────┐                         
│ │                │     │     ┌───────────────┐   
│ │                │     │     │               │   
│ │                │     │     │   ERC20       │   
│ └──────────────┬─┘     └─────▶   token.sol   │   
└────────────────┘             │               │   
                               │               │   
                               └───────────────┘   
                                                                          
```


## The solidity code for Smart Contracts can be found in the `/contracts/` folder.

wallet.sol is the primary TokenWallet contract which holds user's funds and offers additional security features (e.g. whitelist, daily spend/gas limit). The wallet contract is composed of multiple contracts, one of which is the Vault, which offers the same security features as the Wallet, but without the gas top up functionality. 

oracle.sol is an oracle contract which stores token/wei exchange rates for each of the supported tokens. Exchange rates have to be updated periodically by calling the Crypto Compare API through the Oraclize contract when the exchange rate increases/decreases beyond a certain threshold, relative to what is stored in the contract.

internal/resolver.sol is a resolver contract that provides dynamic address resolution for the oracle and oraclize smart contracts. 

internal/ownable.sol is an inheritable contract that provides owner authentication functionality to the child contract. 

internal/controller.sol is a storage contract that stores a list of controller addresses which can be used for together with the 'controllable.sol' contract.

internal/controllable.sol is an inheritable contract that integrates with the list of controllers and provides control functionality to the child contract.

mocks/token.sol is a partial implementation of the ERC20 token standard used for development purposes.






## Running contract tests

Get dependencies:
```sh
go mod vendor
```

Run tests:

```sh
go test -v ./test/...
```

## Building contracts

To build all contracts and generate corresponding Go bindings:

```sh
./build.sh
```
