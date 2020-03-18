# Slither

Steps to hide output for checks that are known to be unnecessary.

Run trail of bits security toolbox and mount the contracts directory:

    docker run -v ./contracts:/contracts -it trailofbits/eth-security-toolbox:latest

Set the solidity compiler to the same version as used in the contracts (e.g.):

    solc-select 0.5.15

Navigate to the slither directory inside the container:

    cd /contracts/slither

Run slither in triage mode in order to mark issues that should not be displayed on the next run:

    slither --triage-mode ../contracts/wallet.sol

Exit the container and commit the `slither.db.json` changes to git.
