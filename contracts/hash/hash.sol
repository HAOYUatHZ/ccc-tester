// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Hash {
    function sha256(bytes memory input) public view
        returns (bytes memory out)
    {        
        (bool ok, bytes memory out) = address(2).staticcall(input);
        require(ok);
    }

    function sha256s(uint256 n) public
    {
        bytes memory input = abi.encode(999);
        for (uint i = 0; i < n; i++) {
            sha256(input);
        }
    }
}
