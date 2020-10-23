// https://leetcode.com/problems/132-pattern/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var find132pattern = function(nums) {
    if(nums.length < 3) return false;

    const mins = new Array(nums.length).fill(nums[0]);
    for(let i=1; i<nums.length; i++){
        mins[i] = Math.min(mins[i-1], nums[i]);
    }

    const rg = [nums[nums.length-1]];
    for(let i=nums.length-2; i>0; i--){
        while(rg.length && rg[rg.length-1] <= mins[i-1]) rg.pop();
        if(!rg.length || nums[i] < rg[rg.length-1]) rg.push(nums[i]);
        if(rg.length && nums[i] > rg[rg.length-1]) return true;
    }

    return false;
};