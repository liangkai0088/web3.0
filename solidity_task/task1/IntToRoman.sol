//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract IntToRoman{
    uint256[13] numArr = [1000,900,500,400,100,90,50,40,10,9,5,4,1];
    string[13] roman = [
        "M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"
    ];
    function intToRoman (uint256 num) public view returns (string memory){
        string memory res = "";
        uint256 len = numArr.length;
        for(uint256 i = 0;i < len;i ++){
            while(num >= numArr[i]){
                num -= numArr[i];
                res = string.concat(res,roman[i]);
                if(num == 0){
                    break;
                }
            }
        }
        return res;
    }
}