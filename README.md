# Tokencard Solidity Contracts

## Running contract tests

```sh
npm install
npm run test
```


## Code coverage reports

Running tests will create coverage reports that can be viewed in a browser.

```sh
open coverage/index.html
```

## Building contracts

You need abigen:
```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum
make all
cd cmd/abigen
go build
mv abigen $GOPATH/bin
```

To build all contracts and generate corresponding Go bindings:

```sh
./build.sh
```
