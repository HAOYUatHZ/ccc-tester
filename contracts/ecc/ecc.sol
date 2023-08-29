// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Ecc {
    /* ECC Functions */
    // https://etherscan.io/address/0x41bf00f080ed41fa86201eac56b8afb170d9e36d#code
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

    // https://etherscan.io/address/0x41bf00f080ed41fa86201eac56b8afb170d9e36d#code
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

    // https://github.com/tornadocash/tornado-core/blob/master/contracts/Verifier.sol#L79
    function ecPairing(
        uint256[2] memory a1,
        uint256[2][2] memory a2,
        uint256[2] memory b1,
        uint256[2][2] memory b2
    ) public view
        returns (bool)
    {
        uint256[2][2] memory p1 = [a1, b1];
        uint256[2][2][2] memory p2 = [a2, b2];

        uint256 inputSize = 2*6;
        uint256[] memory input = new uint256[](inputSize);
        for (uint256 i = 0; i < 2; i++) {
            uint256 j = i * 6;
            input[j + 0] = p1[i][0];
            input[j + 1] = p1[i][1];
            input[j + 2] = p2[i][0][0];
            input[j + 3] = p2[i][0][1];
            input[j + 4] = p2[i][1][0];
            input[j + 5] = p2[i][1][1];
        }

        uint256[1] memory out;
        assembly {
            if iszero(staticcall(not(0), 0x08, add(input, 0x20), mul(inputSize, 0x20), out, 0x20)) {
                revert(0, 0)
            }
        }
        return out[0] != 0;
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

    function ecPairings(uint256 n) public
    {
        uint256[2] memory a1;
        uint256[2][2] memory a2;
        uint256[2] memory b1;
        uint256[2][2] memory b2;

        a1[0] = 0x0000000000000000000000000000000000000000000000000000000000000001;
        a1[1] = 0x0000000000000000000000000000000000000000000000000000000000000002;
        a2[0][0] = 0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed;
        a2[0][1] = 0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2;
        a2[1][0] = 0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa;
        a2[1][1] = 0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b;

        b1[0] = 0x1aa125a22bd902874034e67868aed40267e5575d5919677987e3bc6dd42a32fe;
        b1[1] = 0x1bacc186725464068956d9a191455c2d6f6db282d83645c610510d8d4efbaee0;
        b2[0][0] = 0x1b7734c80605f71f1e2de61e998ce5854ff2abebb76537c3d67e50d71422a852;
        b2[0][1] = 0x10d5a1e34b2388a5ebe266033a5e0e63c89084203784da0c6bd9b052a78a2cac;
        b2[1][0] = 0x275739c5c2cdbc72e37c689e2ab441ea76c1d284b9c46ae8f5c42ead937819e1;
        b2[1][1] = 0x018de34c5b7c3d3d75428bbe050f1449ea3d9961d563291f307a1874f7332e65;

        for (uint i = 0; i < n; i++) {
            ecPairing(a1, a2, b1, b2);
        }
    }
}
