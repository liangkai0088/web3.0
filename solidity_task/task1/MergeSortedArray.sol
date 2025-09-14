//SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract MergeSortedArray{
    function mergeSortedArray(uint256[] memory arr1,uint256[] memory arr2) public pure returns (uint256[] memory){
        uint256 len1 = arr1.length;
        uint256 len2 = arr2.length;
        uint256[] memory newArr = new uint256[](len1 + len2);
        uint256 idx = 0;
        uint256 a = 0;
        uint256 b = 0;
        while(a < len1 && b < len2){
            if(arr1[a] <= arr2[b]){
                newArr[idx ++] = arr1[a];
                a ++;
            }else{
                newArr[idx ++] = arr1[b];
                b ++;
            }
        }
        while(a < len1){
            newArr[idx ++] = arr1[a];
            a ++;
        }
        while(b < len2){
            newArr[idx ++] = arr2[b];
            b ++;
        }   
        return newArr;
    }
}
