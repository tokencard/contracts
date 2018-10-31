pragma solidity ^0.4.25;

 /* solium-disable-next-line */ 
import "../externals/base64.sol";


contract Base64Exporter is Base64 {

    /// @dev export _base64decode() as a public function.
    function base64decode(bytes _encoded) public pure returns (bytes) {
        return _base64decode(_encoded);
    }

}
