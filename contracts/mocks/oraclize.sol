pragma solidity ^0.5.3;


contract OraclizeConnector {

    address public cbAddress;
    uint gasPrice;
    byte public proofType;

    constructor(address _cbAddress) public {
        cbAddress = _cbAddress;
    }

    function query(string memory _arg) private pure returns (bytes32){
        return keccak256(bytes(_arg));
    }

    function query(uint _timestamp, string calldata _datasource, string calldata _arg) external payable returns (bytes32) {
        if (_timestamp == 0 && bytes(_datasource).length == 0){
            return query(_arg);
        }
        return keccak256(bytes(_arg));
    }

    function query_withGasLimit(uint _timestamp, string calldata _datasource, string calldata _arg, uint _gaslimit) external payable returns (bytes32) {
        if (_timestamp == 0 && bytes(_datasource).length == 0 && _gaslimit == 0){
            return query(_arg);
        }
        return keccak256(bytes(_arg));
    }

    function getPrice(string memory _datasource) public pure returns (uint) {
        if (bytes(_datasource).length == 0){
            return 1000000;
        }
        return 1000000; //10^6
    }

    function getPrice(string memory _datasource, uint gaslimit) public pure returns (uint) {
        if (gaslimit == 0){
            return getPrice(_datasource);
        }
        return 1000000; //10^6
    }

    function setProofType(byte _proofType) external {
        proofType = _proofType;
    }

    function setCustomGasPrice(uint _gasPrice) external {
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
