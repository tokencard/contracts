# Instructions

In order to hide output for checks that are known to be unnecessary run the helper script:

    ./slither/slither.sh

This will show a prompt for each issue and ask you to select which output can be ignored.

# Development

Run trail of bits security toolbox and mount the contracts directory inside docker (make sure that the mount path is an absolute path):

    docker run -v $GOPATH/src/github.com/tokencard/contracts:/contracts -it trailofbits/eth-security-toolbox@sha256:<hash>

Set the solidity compiler to the same version as used in the contracts:

    solc-select x.x.xx

Navigate to the slither directory inside the container:

    cd /contracts/slither

Run slither in triage mode in order to mark issues that should not be displayed on the next run:

    slither --triage-mode ../contracts/wallet.sol

Exit the container and commit the `slither.db.json` changes to git.
