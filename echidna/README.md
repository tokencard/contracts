# Properties

This directory contains some properties for testing Solidity contracts where multiple users are defined:

- `echidna_deployer`: Deployer address used to execute the contract constructor.

- `echidna_owner`: Address of the user that owns an ownable contract.

- `echidna_controller`: Controller address assigned to a controllable contract.

- `echidna_attacker`: An unprivileged user, used to execute the randomly generated transactions.

# Configuration

Individual tests can be configured using yaml configuration files.

A full list of test configuration options can be found here: [default.yaml](https://github.com/crytic/echidna/blob/master/examples/solidity/basic/default.yaml)

# Instructions

Normally the echidna tests run in CircleCI but they can also be invoked locally for development purposes.

Run trail of bits security toolbox and mount the contracts directory, you need to make sure that the path is an absolute path:

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox:latest

Run `echidna-test` command on each test contract:

    echidna-test echidna/spendLimit.sol --config=echidna/spendLimit.yaml --contract=TEST

    Analyzing contract: spendLimit.sol:TEST
    echidna_fixed_spendLimitValue: passed!
    echidna_bounded_spendLimitAvailable: passed!
    echidna_fixed_owner: passed!
    echidna_daily_spendLimitAvailable: passed!

    Unique instructions: 900
    Unique codehashes: 1
