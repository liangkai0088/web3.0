//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BinarySearch{
    function binarySearch(uint256[] memory arr,uint256 num) public pure returns (bool){
        uint256 left = 0;
        uint256 right = arr.length - 1;
        while(left <= right){
            uint256 mid = left + (right - left) / 2;
            if(arr[mid] == num){
                return true;
            }else if(arr[mid] > num){
                right = mid - 1;
            }else{
                left = mid + 1;
            }
        }
        return false;
    }
}