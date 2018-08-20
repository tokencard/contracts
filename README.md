# Smart contracts

This repository contains TokenWallet and TokenCard smart contracts.

Wallet.sol is the primary TokenWallet contract which holds user's funds and offers additional security features (e.g. whitelist, daily spend/gas limit). The wallet contract is composed of multiple contracts, one of which is the Vault, which offers the same security features as the Wallet, but without the gas top up functionality. 

Oracle.sol is an oracle contract which stores token/wei exchange rates for each of the supported tokens. Exchange rates have to be updated periodically by Token when the exchange rate increases/decreases beyond a certain threshold, relative to what is stored in the contract.

Token.sol is a partial implementation of the ERC20 token standard used for development purposes.

Card.sol contract stores funds that can be spent when using the TokenCard (old contract).

Controller.sol allows Token to interact with the Card.sol contract. (old contract)

## Contract communication

```
                              ┌──────────────────┐ 
  ┌────────────────┐          │                  │ 
┌─┤                │          │                  │ 
│ │                │    ┌─────▶    Oracle.sol    │ 
│ │                │    │     │                  │ 
│ │                │    │     │                  │ 
│ │   Wallet.sol   ├────┘     └──────────────────┘ 
│ │                ├─────┐                         
│ │                │     │     ┌───────────────┐   
│ │                │     │     │               │   
│ │                │     │     │               │   
│ └──────────────┬─┘     └─────▶   Token.sol   │   
└────────────────┘             │               │   
                               │               │   
                               └───────────────┘   
                                                                          
                                                   
                                   ┌──────────────┐
  ┌──────────────────────┐         │              │
  │                      │         │              │
  │                      │         │              │
  │                      │         │              │
  │    Controller.sol    │─────────▶   Card.sol   │
  │                      │         │              │
  │                      │         │              │
  │                      │         │              │
  └──────────────────────┘         │              │
                                   └──────────────┘
```


## Running contract tests

Get dependencies:
```sh
dep ensure
```

Run tests:

```sh
go test -v ./test/...
```

## Code coverage reports

Running tests will create coverage reports that can be viewed in a browser.

```sh
open coverage/index.html
```

## Building contracts

To build all contracts and generate corresponding Go bindings:

```sh
./build.sh
```
