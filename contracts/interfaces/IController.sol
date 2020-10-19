<<<<<<< HEAD
// SPDX-License-Identifier: GPLv3

pragma solidity ^0.6.0;
=======
pragma solidity >0.5.0;
>>>>>>> f6d9eb58... Move more contracts(controller, holder, licence) to solc 0.6

/// @title The IController interface provides access to the isController and isAdmin checks.
interface IController {
    function isController(address) external view returns (bool);

    function isAdmin(address) external view returns (bool);
}