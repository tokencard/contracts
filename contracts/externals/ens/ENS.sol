/**
 * BSD 2-Clause License
 *
 * Copyright (c) 2018, True Names Limited
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * * Redistributions of source code must retain the above copyright notice, this
 *   list of conditions and the following disclaimer.
 *
 * * Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
pragma solidity ^0.5.0;

interface ENS {

    // Logged when the owner of a node assigns a new owner to a subnode.
    event NewOwner(bytes32 indexed _node, bytes32 indexed _label, address _owner);

    // Logged when the owner of a node transfers ownership to a new account.
    event Transfer(bytes32 indexed _node, address _owner);

    // Logged when the resolver for a node changes.
    event NewResolver(bytes32 indexed _node, address _resolver);

    // Logged when the TTL of a node changes
    event NewTTL(bytes32 indexed _node, uint64 _ttl);

    // Logged when an operator is added or removed.
    event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);

    function setRecord(bytes32 _node, address _owner, address _resolver, uint64 _ttl) external;
    function setSubnodeRecord(bytes32 _node, bytes32 _label, address _owner, address _resolver, uint64 _ttl) external;
    function setSubnodeOwner(bytes32 _node, bytes32 _label, address _owner) external returns(bytes32);
    function setResolver(bytes32 _node, address _resolver) external;
    function setOwner(bytes32 _node, address _owner) external;
    function setTTL(bytes32 _node, uint64 _ttl) external;
    function setApprovalForAll(address _operator, bool _approved) external;
    function owner(bytes32 _node) external view returns (address);
    function resolver(bytes32 _node) external view returns (address);
    function ttl(bytes32 _node) external view returns (uint64);
    function recordExists(bytes32 _node) external view returns (bool);
    function isApprovedForAll(address _owner, address _operator) external view returns (bool);
}
