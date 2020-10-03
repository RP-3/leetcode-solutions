// https://leetcode.com/problems/k-diff-pairs-in-an-array/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findPairs = function(nums, k) {
    nums.sort((a, b) => a - b);
    let [result, l, r] = [0, 0, 1];

    while(r < nums.length){
        const diff = nums[r] - nums[l];

        if(l === r) r++;
        else if(diff < k) r++;
        else if(diff > k) l++;
        else if(diff === k){
            result++;
            r++;
            while(nums[r] === nums[r-1]) r++;
        }
    }

    return result
};