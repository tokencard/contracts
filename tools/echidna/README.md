# Configuration

Individual tests can be configured using yaml configuration files.

A full list of test configuration options can be found here: [default.yaml](https://github.com/crytic/echidna/blob/master/examples/solidity/basic/default.yaml)

# Instructions

Normally the echidna tests run in CircleCI but they can also be invoked locally for development purposes:

    ./echidna/echidna.sh

New test contracts can be added by appending the `echidna-test` command to the `echidna.sh` script.

# Development

Sometimes it's useful to test different `echidna-test` command line options when running the fuzz tests, for this it's possible to run the security toolbox container.

To run trail of bits security toolbox and mount the contracts directory inside docker (make sure that the mount path is an absolute path):

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox:latest

Set the solidity compiler to the same version as used in the contracts:

    solc-select x.x.xx

We expect some functions to be `public` instead of `external` so that they can be called from within the test contract, to achieve this we have to flatten the contracts using `slither-flat`:

    find contracts -type f -name '*.sol' -exec slither-flat --convert-external {} \;

Run `echidna-test` command on each test contract:

    echidna-test echidna/spendLimit.sol --config=echidna/spendLimit.yaml --contract=TEST
