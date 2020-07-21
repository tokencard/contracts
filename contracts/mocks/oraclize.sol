// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.11;

contract OraclizeConnector {
    address public cbAddress;
    uint256 gasPrice;
    bytes1 public proofType;

    constructor(address _cbAddress) public {
        cbAddress = _cbAddress;
    }

    function query(string memory _arg) private pure returns (bytes32) {
        return keccak256(bytes(_arg));
    }

    function query(uint256 _timestamp, string calldata _datasource, string calldata _arg) external payable returns (bytes32) {
        if (_timestamp == 0 && bytes(_datasource).length == 0) {
            return query(_arg);
        }
        return keccak256(bytes(_arg));
    }

    function query_withGasLimit(uint256 _timestamp, string calldata _datasource, string calldata _arg, uint256 _gaslimit) external payable returns (bytes32) {
        if (_timestamp == 0 && bytes(_datasource).length == 0 && _gaslimit == 0) {
            return query(_arg);
        }
        return keccak256(bytes(_arg));
    }

    function getPrice(string memory _datasource) public pure returns (uint256) {
        if (bytes(_datasource).length == 0) {
            return 1000000;
        }
        return 1000000; //10^6
    }

    function getPrice(string memory _datasource, uint256 gaslimit) public pure returns (uint256) {
        if (gaslimit == 0) {
            return getPrice(_datasource);
        }
        return 1000000; //10^6
    }

    function setProofType(bytes1 _proofType) external {
        proofType = _proofType;
    }

    function setCustomGasPrice(uint256 _gasPrice) external {
        gasPrice = _gasPrice;
    }
}


contract OraclizeAddrResolver {
    address private oraclizedAddress;

    constructor(address _oraclizedAddress) public {
        oraclizedAddress = _oraclizedAddress;
    }

    function getAddress() public view returns (address) {
        return oraclizedAddress;
    }
}
