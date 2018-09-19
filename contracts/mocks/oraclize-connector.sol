pragma solidity 0.4.24;


contract Oraclize {
    address public cbAddress;

    constructor(address _cbAddress) public {
      cbAddress = _cbAddress;
    }

    function query(uint _timestamp, string _datasource, string _arg) external payable returns (bytes32 _id) {
      return 0;
    }

    function query_withGasLimit(uint _timestamp, string _datasource, string _arg, uint _gaslimit) external payable returns (bytes32 _id) {
      return 0;
    }

    function query2(uint _timestamp, string _datasource, string _arg1, string _arg2) public payable returns (bytes32 _id) {
      return 0;
    }

    function query2_withGasLimit(uint _timestamp, string _datasource, string _arg1, string _arg2, uint _gaslimit) external payable returns (bytes32 _id) {
      return 0;
    }

    function queryN(uint _timestamp, string _datasource, bytes _argN) public payable returns (bytes32 _id) {
      return 0;
    }

    function queryN_withGasLimit(uint _timestamp, string _datasource, bytes _argN, uint _gaslimit) external payable returns (bytes32 _id) {
      return 0;
    }

    function getPrice(string _datasource) public returns (uint _dsprice) {
      return 0;
    }

    function getPrice(string _datasource, uint gaslimit) public returns (uint _dsprice) {
      return 0;
    }

    function setProofType(byte _proofType) external {

    }

    function setCustomGasPrice(uint _gasPrice) external {

    }

    function randomDS_getSessionPubKeyHash() external constant returns(bytes32) {
      return 0;
    }
}

