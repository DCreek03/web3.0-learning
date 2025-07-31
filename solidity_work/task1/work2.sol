// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract leetcode {
    // 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
    function revertString(string memory str) public pure returns (string memory) {
        bytes memory inputBytes = bytes(str);
        uint len = inputBytes.length;
        for (uint i = 0; i < len / 2; i++) {
            uint j = len - 1 - i ;
            bytes1 temp = inputBytes[i];
            inputBytes[i] = inputBytes[j];
            inputBytes[j] = temp;
        }
        return string(inputBytes);
        
    }

    // 用 solidity 实现整数转罗马数字 1<= num <= 3999
    function int2RomanNum(uint256 num) public pure returns (string memory res) {
        require(num > 0 && num < 4000, "Number out of range");

        bytes memory roman;

        while (num > 1000) {
            roman = abi.encodePacked(roman, "M");
            num -= 1000;
        }

        roman = _appendDigit(roman, num / 100, "C", "D", "M");
        num %= 100;

        roman = _appendDigit(roman, num / 10, "X", "L", "C");
        num %= 10;

        roman = _appendDigit(roman, num / 1, "I", "V", "X");
        
        return string(roman);
    }

    function _appendDigit(
        bytes memory roman, 
        uint256 num, 
        string memory one, 
        string memory five, 
        string memory ten) 
        private pure returns (bytes memory) {
            if (num == 0) return roman;
            if (num <= 3) {
                for (uint256 i = 0; i < num; i++) {
                    roman = abi.encodePacked(roman, one);
                }
            } else if (num == 4) {
                roman = abi.encodePacked(roman, one, five);
            } else if (num <= 8) {
                roman = abi.encodePacked(roman, five);
                for (uint i = 5; i < num; i++ ) {
                    roman = abi.encodePacked(roman, one);                    
                }
            } else  {
                roman = abi.encodePacked(roman, one, ten);
            }
            return roman;
        }

    // 用 solidity 实现罗马数字转数整数
    function roman2Num (string memory str) public pure returns (uint256) {
        bytes memory roman = bytes(str);
        uint len = roman.length;
        int256 num = 0;
        for (uint i = 0; i < len; i++) {
            if (i < len - 1 && _getValue(roman[i]) < _getValue(roman[i + 1])) {
                num -= int(_getValue(roman[i]));
            } else {
                num += int(_getValue(roman[i]));
            }
        }
        return uint256(num);
    }

    function _getValue (bytes1 num) private pure returns (uint256) {
        if (num == "I") return 1;
        if (num == "V") return 5;
        if (num == "X") return 10;
        if (num == "L") return 50;
        if (num == "C") return 100;
        if (num == "D") return 500;
        if (num == "M") return 1000;
        return 0;
    }

    // 将两个有序数组合并为一个有序数组
    function mergeTwoArray(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        uint[] memory merged = new uint[](len1 + len2);
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i < len1 && j < len2) {
            if (arr1[i] < arr2[j]) {
                merged[k] = arr1[i];
                i++;
            } else {
                merged[k] = arr2[j];
                j++;
            }
            k++;
        }
        while (i < len1) {
            merged[k] = arr1[i];
            i++;
            k++;
        }
        while (j < len2) {
            merged[k] = arr2[j];
            j++;
            k++;
        }
        return merged;
    }

    // 二分查找 在一个有序数组中查找目标值
    function binarySearch(uint[] memory arr, uint target) public pure returns (int) {
        if (arr.length == 0) return -1;
        
        uint left = 0;
        uint right = arr.length - 1;
        
        while (left <= right) {
            uint mid = left + (right - left) / 2;
            
        
            if (arr[mid] == target) {              
                return int(mid);
            } else if (arr[mid] < target) {           
                left = mid + 1;
            } else {              
                if (mid == 0) break;
                right = mid - 1;
            }
        }
        
       
        return -1;
    }
        
}

