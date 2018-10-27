# TokenWallet - Consumer Contract Wallet 

This repository contains the Smart Contracts needed to power the TokenCard, written in Solidity, for execution in the EVM.

## Overview

The functionality encoded in the Smart Contracts found in this repository have been designed to help users protect their tokens, by holding them within their own instance of a *Consumer Contract Wallet* aka the *TokenWallet* which they can configure to their liking. The functionality within the TokenWallet has been designed to limit a user's exposure to loss in the event that a user has had their Private Key compromised.

Each user deploys their own instance of the TokenWallet ([wallet.sol](/contracts/wallet.sol)) to the Ethereum Network which interacts with an exchange rate oracle ([oracle.sol](/contracts/oracle.sol)) that exists to provide exchange rates needed to secure a user's tokens. The individual TokenWallet contracts use the [Ethereum Name Service (ENS)](https://ens.domains/) to resolve the location of the exchange rate oracle.

### High-level Architecture

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


## Assumptions

- Every user will have their own Public and Private key pair, aka their `Address`
- TokenCard's systems will NEVER have access to the Private Key for the user’s `Address`
- There are a number of different "pots of tokens" for a given user:
     - The user’s entire ETH and ERC20 token assets - TokenWallet
     - An amount of ETH to pay for Gas - *Gas Level*. The *Gas Level* is an amount of ETH associate with the `Address`

## Requirements

- This `Address` will own all of the user’s Smart Contracts aka *Owner*
- TokenCard `Controller` - A set of Addresses, owned and operated by TokenCard, used to provide services
- ERC20 tokens to be stored in the TokenWallet so that user can configure best in class security
- We need to have a convenient way to "top-up" the amount of ETH that lives on our user’s `Address` aka *Gas Level* via our Smart Contracts
- We need to be as decentralised as possible, minimise TokenCard’s access to user funds as well as minimising the contracts dependency on TokenCard’s infrastructure. 
- Help user protect their funds by minimising the risk in the event of their `Address`'s private key being compromised.

## Security Features

In order to help users protect their tokens in the event that their Private Key gets compromised, we present the following security features: 

- A Whitelist of Addresses - akin to a whitelist of payees in a banking application, this whitelist is exempt from the configured daily spend limits.
- Daily Spending Limit - denominated in ETH. This is used to define how much can a user spend in a given day if transferring assets to addresses outside their Whitelist.
- Daily Top Up Limit - (*Gas Level*) denominated in ETH. This is used to define how much ETH can be spent from a user's TokenWallet to their `Address` to pay the network for gas.

### Configuration

There are three ways to configure a TokenWallet: 

- via Constructor: Upon creation of a new TokenWallet the above security features can be configured by passing the desired values to the contructor of the TokenWallet smart contract when deploying this to the Ethereum network. These are essentially the default values set when creating the TokenWallet. 
- via a 1-time write pattern: Aside from default values passed in via the Constructor the user's `Address` may do a 1-time write to the aforementioned Security Features. These allow the `Address` to change the values that power the security features. It is advised that users of the TokenWallet set their security settings so that they can not longer be tampered with in the event that a user's private key is compromised. 
- via a 2FA pattern: Where a user can *submitChange* a new value for one of the Security Features, then the `Controller` needs to either OK the value change, after the user performed some form of 2nd factor authentication the `Controller` may decide to respond with a *confirmChange* method or a *cancelChange* method.


## Solidity code in the `/contracts/` folder

[wallet.sol](/contracts/wallet.sol) is the primary TokenWallet contract that helps user's secure their funds. The Wallet is composed of multiple contracts, one of which is the Vault, which offers the same security features as the Wallet, but without the initial GasTopUp functionality. 

[oracle.sol](/contracts/oracle.sol) is an exchange rate oracle contract that stores exchange rates for a set of supported ERC20 tokens. Exchange rates are updated periodically by calling the Crypto Compare API through the Oraclize contract.


### Solidity code in the `/contracts/internals/` folder

[controllable.sol](/contracts/internals/controllable.sol) is an inheritable contract that integrates with the list of controllers and provides control functionality to the child contract.

[controller.sol](/contracts/internals/controller.sol) is a storage contract that stores a list of controller addresses which can be used for together with the 'controllable.sol' contract.

[ownable.sol](/contracts/internals/ownable.sol) is an inheritable contract that provides owner authentication functionality to the owned contract. 

[date.sol](/contracts/internals/date.sol) is a simple Date parsing contract with a single method, used to parse out a comparable number from the Date in an HTTP header

[json.sol](/contracts/internals/json.sol) a simple and bespoke JSON parser, this is not a generic JSON parser. 


### Solidity code in the `/contracts/mocks/` folder

[token.sol](/contracts/mocks/token.sol) is a partial implementation of the ERC20 token standard used for testing and development purposes.

[oraclize-connector.sol](/contracts/mocks/oraclize-connector.sol) is a mocked out version of the oraclize-connector, which the contract that does most of the oraclize work.

[oraclize-resolver.sol](/contracts/mocks/oraclize-resolver.sol]) is a mocked out version of the oraclize-resolver, which is used to locate the address of oraclize connector

### Solidity code in the `/contracts/externals/` folder

All of the third-party code we rely on can be found in this folder. The below table details the third-party code used and their licenses.

| File          | License       | 
| ------------- | ------------- |
| SafeMath.sol  | [MIT](https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/LICENSE) |
| base64.sol    | [GPLv3](https://github.com/vcealicu/melonport-price-feed/blob/master/LICENSE) |
| ENS           | [BSD2](https://github.com/ensdomains/ens/blob/master/LICENSE) |
| strings.sol   | [Apache v2](https://github.com/Arachnid/solidity-stringutils/blob/master/LICENSE) |
| oraclizeAPI   | [MIT](https://github.com/oraclize/ethereum-api/blob/master/LICENSE) |
| cbor          | [MIT](https://github.com/smartcontractkit/solidity-cborutils/blob/master/LICENSE) |


## Building contracts

To build all contracts and generate corresponding Go bindings:

```sh
./build.sh
```

## Running contract tests

### Dependencies

- go version 1.11 is required. 
- go modules (experimental in go 1.11) are needed. `export GO111MODULE=on`

### Running 

Get dependencies:

```sh
go mod vendor
```

Run tests:

```sh
go test -v ./test/...
```

