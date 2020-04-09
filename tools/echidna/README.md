# Echidna

Normally the echidna tests run in CircleCI but they can also be invoked locally for development purposes:

    ./tools/echidna/echidna.sh

New test contracts can be added by appending the `echidna-test` command to the `echidna.sh` script.

## Configuration

Individual tests can be configured using yaml configuration files.

For a full list of configuration options go to [GitHub](https://github.com/crytic/echidna/blob/master/examples/solidity/basic/default.yaml).

## Development

Sometimes it's useful to test different `echidna-test` command line options when running the fuzz tests, for this it's possible to run the security toolbox container.

To run trail of bits security toolbox and mount the contracts directory inside docker (make sure that the mount path is an absolute path):

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox@sha256:<hash>

Set the solidity compiler to the same version as used in the contracts:

    solc-select <version>

We expect some functions to be `public` instead of `external` so that they can be called from within the test contract, to achieve this we have to flatten the contracts using `slither-flat`:

    find contracts -type f -name '*.sol' -exec slither-flat --convert-external {} \;

This can also be acheved using the provided `./tools/slither/flatten.sh` script.

Run `echidna-test` command on each test contract with a corresponding configuration file:

    echidna-test <contract>.sol --config=<contract>.yaml --contract=TEST
