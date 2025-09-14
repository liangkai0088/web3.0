// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Count{
    int256 public  i = 0;
    function plusOne() public {
        i = i + 1;
    }

}