// https://leetcode.com/problems/subarray-product-less-than-k/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var numSubarrayProductLessThanK = function(nums, k) {
    let [result, l, prod] = [0, 0, 1];

    for(let r = 0; r < nums.length; r++){
        prod*=nums[r];
        while(l < r && prod >= k) prod /= nums[l++];
        if(prod < k) result += (r - l + 1);
    }

    return result;
};