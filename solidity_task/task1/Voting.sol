//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting{
    //mapping来存储首选人得票数
    string[] private names;
    mapping(string => uint256) data;

    //vote 用户投票给某个人
    function vote(string calldata candidate) public{
        uint256 votes = data[candidate];
        data[candidate] = votes + 1;
        names.push(candidate);
    }


    //getVotes返回某个首选人的得票数
    function getVotes(string calldata candidate) public view returns(uint256){
        return data[candidate];
    }

    //resetVotes重置所有候选人的得票数
    function resetVotes() public {
        uint256 len = names.length;
        for(uint256 i = 0;i < len;i ++){
            delete data[names[i]];
        }
        delete names;
    }

}