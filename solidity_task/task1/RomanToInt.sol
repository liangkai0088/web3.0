//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RomanToInt{
    function getNumByByte(bytes1 value) private pure returns (uint256){
        if(value == "I") return 1;
        if(value == "V") return 5;
        if(value == "X") return 10;
        if(value == "L") return 50;
        if(value == "C") return 100;
        if(value == "D") return 500;
        if(value == "M") return 1000;
        return 0;
    }

    function romanToInt(string memory num) public pure returns(uint256){
        bytes memory byteStr = bytes(num);
        uint256 len = byteStr.length;
        uint256 res = 0;
        for(uint256 i = 0;i < len;i ++){
            uint256 pre = getNumByByte(byteStr[i]);
            if(i + 1 < len){
                uint256 next = getNumByByte(byteStr[i + 1]);
                if(pre < next){
                    res += (next - pre);
                    i ++;
                    continue;
                }
            }
            res += pre;
        }
        return res;
    }
}