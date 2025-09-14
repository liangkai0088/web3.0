//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ReserseString{
    function reverseString(string memory str) public pure returns (string memory){
        bytes memory arr = bytes(str);
        uint256 len = arr.length;
        bytes memory res = new bytes(len);
        for(uint256 i = 0;i < len;i ++){
            res[i] = arr[len - 1 - i];
        }
        return string(res);
    }
}