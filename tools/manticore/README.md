# Manticore

Manticore symbolic execution can be invoked for all contracts using the provided script:

	./tools/manticore/manticore.sh

Additional contracts can be added by appending them to the command string inside `manticore.sh`.

## Development

To run trail of bits security toolbox and mount the contracts directory inside docker (make sure that the mount path is an absolute path):

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox@sha256:<hash>

Set the solidity compiler to the same version as used in the contracts:

    solc-select <version>

To perform the symbolic execution/analysis on a contract, use the manticore command:

	manticore 
