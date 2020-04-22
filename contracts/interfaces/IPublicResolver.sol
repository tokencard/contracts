pragma solidity ^0.6.0;


interface IPublicResolver {

    function addr(bytes32) external view returns (address);

}
