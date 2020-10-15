// https://leetcode.com/problems/rotate-array/

var rotate = function(nums, k) { // 10 mins
    k = k%nums.length;
    reverse(0, nums.length-1);
    reverse(0, k-1);
    reverse(k, nums.length-1);


    function reverse(i, j){
        while(i<j){
            [nums[i], nums[j]] = [nums[j], nums[i]];
            i++; j--;
        }
    }
};