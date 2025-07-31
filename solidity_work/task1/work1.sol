// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

//创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数
contract Voting {
    mapping (string => uint) public votes;

    string[] public candidates;

    function addCandidate(string memory candidate) public {
        require(votes[candidate] == 0, "Candidate already exists");
        candidates.push(candidate);
    }

    function getVotes(string memory candidate) public view returns (uint) {
        return votes[candidate];
    }

    function vote(string memory candidate) public  {
        require(votes[candidate] != 0 || isCandidate(candidate), "Invalid Candidate");
        votes[candidate] += 1;
    }

    function isCandidate(string memory candidate) public view returns (bool) {
        for (uint i = 0; i < candidates.length; i++) {
            if(keccak256(bytes(candidates[i])) == keccak256(bytes(candidate))) {
                return true;
            }
        }
        return false;
    }

    function resetVotes() public {
        for (uint i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }
}
           
           