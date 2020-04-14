# Slither

In order to hide output for checks that are known to be unnecessary run the helper script:

    ./tools/slither/slither.sh

This will show a prompt for each issue and ask you to select which output can be ignored.

## Configuration

Unnecessary detectors can be disabled permanently inside the `slither.config.json` file.

For a full list of configuration options go to [GitHub](https://github.com/crytic/slither/wiki/Usage#configuration-file).

## Development

Run trail of bits security toolbox and mount the contracts directory inside docker (make sure that the mount path is an absolute path):

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox@sha256:<hash>

Set the solidity compiler to the same version as used in the contracts:

    solc-select <version>

Navigate to the slither directory inside the container:

    cd /contracts/tools/slither

Run slither in triage mode in order to mark issues that should not be displayed on the next run:

    slither --triage-mode ../../contracts/<contract>.sol

Exit the container and commit the `slither.db.json` changes to git.
