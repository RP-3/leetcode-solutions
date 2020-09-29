// https://leetcode.com/problems/subsets-ii/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsetsWithDup = function(nums) {

    const result = [];
    const wc = [];
    nums.sort((a, b) => a - b); // asc

    // decide which elements exist at position i
    const backtrack = (i) => {

        result.push(wc.slice());

        for(let j=i; j<nums.length; j++){
            if(j > i && nums[j] === nums[j-1]) continue;
            wc.push(nums[j]);
            backtrack(j+1);
            wc.pop();
        }
    };

    backtrack(0);
    return result;
};
