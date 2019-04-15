# The Consumer Contract Wallet 

This repository contains the Smart Contracts needed to power the TokenCard, written in Solidity, for execution in the EVM.

The TokenCard is the world's first non-custodial VISA card, that allows people to hold their own assets, whilst being able to seemlessly move funds to a VISA debit card. 

## Overview

The functionality encoded in the Smart Contracts found in this repository have been designed to help users protect their tokens, by holding them within their own instance of a *Consumer Contract Wallet* which they can configure to their liking. The functionality within the *Consumer Contract Wallet* has been designed to limit a user's exposure to loss of tokens in the event that a user has had their Private Key compromised.

We believe there are two major problems facing the adoptiona and use of cryptocurrecy, protection from: 
 - Users having their private key compromised and losing all of their assets
 - Users losing their private key

This, first version, of the *Consumer Contract Wallet* protects users by limiting their exposure to theft if their private key gets compromised. We are working on things that will address the loss of one's private key in the future... 

Each user deploys their own instance of the *Consumer Contract Wallet* ([wallet.sol](/contracts/wallet.sol)) to the Ethereum Network which interacts with an exchange rate *oracle* ([oracle.sol](/contracts/oracle.sol)) that exists to provide exchange rates needed to secure a user's tokens. The individual *Consumer Contract Wallet* contracts use the [Ethereum Name Service (ENS)](https://ens.domains/) to resolve the location of the exchange rate oracle, as well as to resolve the location of the *controller* ([controller.sol](contracts/internals/controller.sol)) contract, and the *tokenwhitelist* ([tokenWhitelist.sol](contracts/internals/tokenWhitelist.sol)) contract. The *tokenWhitelist* is a whitelist of tokens and their exchange rates that are used to secure a user's tokens within their wallet, it also determines which tokens can be used to load the TokenCard and which are burnable by the TKN Holder Contract. This *controller* contract is used for administrative purposes only, rest assured this has no access to user's tokens. The controllers are used to perform 2FA functionality, used to help a user topup there gas when stuck, and are used to perform administrative tasks on the oracle and on the (token). 

### High-level Architecture

```
                                                                                               +-------------------------+
                                                                                               |                         |
                                                                                               |                         |
                 +------------------------------------ENS-------------------------------------->     controller.sol      |
                 |                                                                             |                         |
                 |                                                                             |                         |
                 |                                                                             |                         |
                 |                                      +---------------------------+          +-------------------------+
                 |                                      |                           |
    +------------+---------------+                      |                           |
+---+                            |                      |                           |
|   |                            +--------ENS----------->        oracle.sol         +-----+
|   |                            |                      |                           |     |
|   |                            |                      |                           |     |
|   |                            |                      |                           |     |
|   |                            |                      |                           |     E
|   |                            |                      +---------------------------+     N
|   |                            |                                                        S
|   |                            |                      +---------------------------+     |
|   |                            |                      |                           |     |
|   |   Consumer Contract Wallet |                      |                           |     |
|   |                            |                      |                           |     |
|   |   wallet.sol               +--------ENS----------->      tokenWhitelist.sol   <-----+
|   |                            |                      |                           |               +------------------------+
|   |                            |                      |                           |               |                        |
|   |                            |                      |                           |               |                        |
|   |                            |                      |                           |      +--------+   Gnosis Multisig      |
|   |                            |                      +---------------------------+      |        |                        |
|   |                            |                                                         |        |   CryptoFloat          |
|   |                            |                      +---------------------------+      |        |                        |
|   |                            |                      |                           |      |        +------------------------+
|   |                            |                      |                           |      |
|   |                            |                      |                           +------+
|   |                            +--------ENS----------->      licence.sol          |
|   |                            |                      |                           |
|   |                            |                      |                           +------+
|   +-------------------------+--+                      |                           |      |
|                             |                         |                           |      |
+-----------------------------+                         +---------------------------+      |        +------------------------+
                                                                                           |        |                        |
                                                                                           |        |                        |
                                                                                           +--------+  TKN Holder Contract   |
                                                                                                    |                        |
                                                                                                    |  holder.sol            |
                                                                                                    |                        |
                                                                                                    +------------------------+

```

## Assumptions

- Every user will have their own Public and Private key pair, aka the `Owner Address`.
- Users SHOULD NEVER have to share the Private Key of their `Owner Address` with anyone.
- There are a number of different "pots of tokens" for a given user:
     - The user’s entire ETH and ERC20 token assets stored within the *Consumer Contract Wallet*.
     - An amount of ETH used to pay for the gas - *Gas Tank*. The *Gas Tank* is a representation of the ETH on the user's `Owner Address`. It should be noted that this ETH is NOT protected by the security features in the *Consumer Contract Wallet* as it resides outside of the Smart Contract.

## Requirements

- This `Owner Address` will own all of the user’s Smart Contracts and will be referred to as the *Owner*, this is sometime referred to as an *Externally Owned Address*
- The `Controller` - Is a set of Addresses, owned and operated by Token Group Ltd, used to provide services to the end user
- The *Consumer Contract Wallet* needs to allow its *Owner* to configure how they wish to secure their tokens in their wallet.
- There needs to be a convenient way to "top-up" the amount of ETH that lives on our user’s `Owner Address` aka *Gas Tank* via the Smart Contracts
- The wallet's design is intended to be as decentralised as possible. This will be achieved by eliminating access to user assets by third-parties and minimising reliance of third-party infrastructure in running the *Consumer Contract Wallet*.
- Must help user protect their funds by minimising the risk in the event of their `Owner Address`'s private key being compromised.

## Security Features

In order to help users protect their tokens in the event that their Private Key gets compromised, we present the following security features: 

- *A Whitelist of Addresses* - akin to a whitelist of payees in a banking application, this whitelist should be configured with a list of trusted addresses for each *Owner* of the *Consumer Contract Wallet*.
- *Daily Spending Limit* - denominated in ETH. This is used to define how much can a user can transfer in a given day if transferring assets to addresses outside their *Whitelist*.
- *Daily Gas Tank Top-up Limit* - (*Gas Tank*) top-up daily limit denominated in ETH. This is used to define the daily limit of ETH that can be sent from a user's *Consumer Contract Wallet* to their `Address`; this ETH is what is used to pay the network for gas.
- *Daily Card Load Limit* - (*Card Load Limit*) card loading daily limit denominated in USD (technically speaking in a stablecoin like USDC). This is used to define the daily limit of tokens or ETH that can be sent from a user's *Consumer Contract Wallet* for loading of the user's TokenCard. This is currently set to 10k USD.

### Configuration

There are three ways to configure a Consumer Contract Wallet: 

- via Constructor: Upon deployment of a new *Consumer Contract Wallet* the above security features can be configured by passing the desired values to the constructor of the *Consumer Contract Wallet* smart contract when deploying this to the Ethereum network. These are the values set when deploying a new instance of the *Consumer Contract Wallet*. 
- via a 1-time write pattern: Aside from default values passed in via the Constructor the user may do a 1-time write to the aforementioned `Security Features`. These allow the `Address` to change the values that power the security features. It is advised that users of the *Consumer Contract Wallet* set their security settings so that they can not longer be tampered with in the event that a user's private key is compromised. Users should set these values once, otherwise an attacker would be able to configure their Smart Contract. 
- via a 2FA pattern: Where a user can *submitChange* a new value for one of the Security Features, then one of the `Controller` addresses needs to either OK the value change or not. It should be noted that due to the nature of the user *AddressWhitelist* where a user may add or remove items from their whitelist via the 2FA pattern, only one pending change to the user's address whitelist can be in flight at a given point in time.

## Naming convention

This section details the naming convention adopted in this codebase: 
 
 - *Contracts* - should be Nouns
 - *functions* - should be Verbs
 - *ending in able* - Smart Contracts that are meant to be inherited and not standalone, i.e. they are Mixins, Snippets, Decorators ...
 - *private contract scoped variables* - all start with an underscore `_`
 - *private / internal functions* - all start with an underscore `_`
 - *constructor parameters* - all start with an underscore `_` and end in one too, e.g. `_ens_` this is to avoid shadowing
 - *function parameters* - all start with an underscor `_`
 - *local variables* - to functions should start without an underscore
 - *public contract scoped variables* - should start without an underscore
 - *public functions* - should start without an underscore
 - *crud functions* - when there exists multiple actions on the same variable we will use the suffix to illustrate the action, for example : `dailySpendLimitSet`, and `dailySpendLimitUpdate`


## Solidity code in the `/contracts/` folder

It should be noted that this codebase makes heavy use of inheritance. 

[wallet.sol](/contracts/wallet.sol) is the primary *Consumer Contract Wallet* contract that helps user's secure their funds. The Wallet communicates with the *oracle*, the controllers, and other ERC20 contracts; see([wallet inhertiance digram](/docs/wallet.inheritance.png)).

[oracle.sol](/contracts/oracle.sol) is an exchange rate *oracle* contract that stores exchange rates for a set of supported ERC20 tokens. Exchange rates are updated periodically by calling the Crypto Compare API through the Oraclize contract. It should noted that the *Consumer Contract Wallet* only protects the ERC20 tokens supported by the *oracle* in its Security Features, tokens not listed as supported by the *oracle* will not be secured by the *Consumer Contract Wallet*; see ([oracle inhertiance diagram](/docs/oracle.inheritance.png)).

[licence.sol](/contracts/licence.sol) is the *TKN licence* contract, and it is used to take a 1% fee of all loads of the user's TokenCard so that it can be sent to the TKN Holder contract. This contract is aware of the CryptoFloat where the remain tokens are to be sent for Token Group Ltd for loading to the TokenCards and it is aware of the address of the TKN Holder contract (that is currently still in development). The licence contract has been created in a way to allow for a DAO to change some of its configured features, this way placed to future proof the implementation; see ([licence inhertiance diagram](/docs/licence.inheritance.png)).


### Solidity code in the `/contracts/internals/` folder

[claimable.sol](/contracts/internals/claimable.sol) is an inheritable contract that allows for tokens sent to a smart contract to be claimed by the owner. This is inherited by the oracle, the tokenWhitelist, and the licence contract.

[controllable.sol](/contracts/internals/controllable.sol) is an inheritable contract that integrates with the list of controllers and provides control functionality to the child contract.

[controller.sol](/contracts/internals/controller.sol) is a storage contract that stores a list of controller addresses which power the 'controllable.sol' contract.

[date.sol](/contracts/internals/date.sol) is a simple date parsing contract with a single method, used to parse out a comparable number from the date in an HTTP header

[ensResolvable.sol](/contracts/internals/ensResolvable.sol) implements a inheritable contract that allows contracts to looked up others via ENS.

[ownable.sol](/contracts/internals/ownable.sol) is an inheritable contract that provides owner authentication functionality to the owned contract. 

[parseIntScientific.sol](/contracts/internals/parseIntScientific.sol) provides floating point in scientific notation (e.g. e-5) parsing functionality. This has been built to support floating point scientific notation returned in JSON. 

[tokenWhitelistable.sol](/contracts/internals/tokenWhitelistable.sol) is an inheritable contract that interfaces with the tokenWhitelist. 

[tokenWhitelist.sol](/contracts/internals/tokenWhitelist.sol) is a storage contract that stores a whitelist of tokens for use within the platform. This whitelist is used to determine which tokens are secured by the security settings, along with which tokens are loadable to the TokenCard and which are burnable by the TKN holder contract.

### Solidity code in the `/contracts/mocks/` folder

[base64-exporter.sol](/contracts/mocks/base64-exporter.sol]) is a mocked out version of a contract that pulls in the base64 encoder for unit testing purposes.

[oraclize-connector.sol](/contracts/mocks/oraclize-connector.sol) is a mocked out version of the oraclize-connector, this is for testing purposes only.

[oraclize-resolver.sol](/contracts/mocks/oraclize-resolver.sol]) is a mocked out version of the oraclize-resolver, which is used to locate the address of oraclize connector.

[parseIntScientific-exporter.sol](/contracts/mocks/parseIntScientific-exporter.sol]) is a mocked out version of a contract that pulls in the parseIntScientific contract used to parse floating points that include scientific notation out of JSON. 

[token.sol](/contracts/mocks/token.sol) is a partial implementation of the ERC20 token standard used for testing and development purposes.


### Solidity code in the `/contracts/externals/` folder

All of the third-party code we rely on can be found in this folder. The below table details the third-party code used and their licenses.


| File            | License       | 
| --------------- | ------------- |
| SafeMath.sol    | [MIT](https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/LICENSE) |
| base64.sol      | [GPLv3](https://github.com/vcealicu/melonport-price-feed/blob/master/LICENSE) |
| ENS             | [BSD2](https://github.com/ensdomains/ens/blob/master/LICENSE) |
| strings.sol     | [Apache v2](https://github.com/Arachnid/solidity-stringutils/blob/master/LICENSE) |
| oraclizeAPI     | [MIT](https://github.com/oraclize/ethereum-api/blob/master/LICENSE) |
| cbor            | [MIT](https://github.com/smartcontractkit/solidity-cborutils/blob/master/LICENSE) |
| gnosis MulitSig | [GPLv3](https://github.com/gnosis/MultiSigWallet/blob/master/LICENSE) |


## Building contracts

To build all contracts and generate corresponding Go bindings:

```sh
./build.sh
```

## Running contract tests

### Dependencies

- go version >1.11 is required. 
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

