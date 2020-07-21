// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

import "../wallet.sol";


interface IWallet {
    function isValidSignature(bytes calldata, bytes calldata) external view returns (bytes4);
}


contract IsValidSignatureExporter {
    address walletAddress;

    constructor(address _wallet) public {
        walletAddress = _wallet;
    }

    /// @dev exports isValidSignature(bytes,bytes) aka EIP-1271, so it can tested (no overloading in Go)
    function isValidSignature(bytes calldata _data, bytes calldata _signature) external view returns (bytes4) {
        return IWallet(walletAddress).isValidSignature(_data, _signature);
    }
}
