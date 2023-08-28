// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Ecc {
    /* ECC Functions */
    function ecAdd(uint256[2] memory p0, uint256[2] memory p1) public view
        returns (uint256[2] memory retP)
    {
        uint256[4] memory i = [p0[0], p0[1], p1[0], p1[1]];
        
        assembly {
            // call ecadd precompile
            // inputs are: x1, y1, x2, y2
            if iszero(staticcall(not(0), 0x06, i, 0x80, retP, 0x40)) {
                revert(0, 0)
            }
        }
    }

    function ecMul(uint256[2] memory p, uint256 s) public view
        returns (uint256[2] memory retP)
    {
        // With a public key (x, y), this computes p = scalar * (x, y).
        uint256[3] memory i = [p[0], p[1], s];
        
        assembly {
            // call ecmul precompile
            // inputs are: x, y, scalar
            if iszero(staticcall(not(0), 0x07, i, 0x60, retP, 0x40)) {
                revert(0, 0)
            }
        }
    }


    /* Bench */
    function ecAdds(uint256 n) public
    {
        uint256[2] memory p0;
        p0[0] = 1;
        p0[1] = 2;
        uint256[2] memory p1;
        p1[0] = 1;
        p1[1] = 2;

        for (uint i = 0; i < n; i++) {
            ecAdd(p0, p1);
        }
    }

    function ecMuls(uint256 n) public
    {
        uint256[2] memory p0;
        p0[0] = 1;
        p0[1] = 2;

        for (uint i = 0; i < n; i++) {
            ecMul(p0, 3);
        }
    }
}