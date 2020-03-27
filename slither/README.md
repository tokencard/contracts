# Slither

In order to hide output for checks that are known to be unnecessary run the `slither.sh` script from root directory:

    ./slither/slither.sh

This will show a prompt for each issue and ask you to select which output can be ignored.

# Manual

Run trail of bits security toolbox and mount the contracts directory, you need to make sure that the path is an absolute path, e.g.

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox@sha256:<hash>

Set the solidity compiler to the same version as used in the contracts (e.g.):

    solc-select 0.5.15

Navigate to the slither directory inside the container:

    cd /contracts/slither

Run slither in triage mode in order to mark issues that should not be displayed on the next run:

    slither --triage-mode ../contracts/wallet.sol

Exit the container and commit the `slither.db.json` changes to git.
