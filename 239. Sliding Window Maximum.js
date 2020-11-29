// https://leetcode.com/problems/sliding-window-maximum/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var maxSlidingWindow = function(nums, k) {
    const q = [];
    const last = () => nums[q[q.length-1]];
    
    // initialize q
    for(let i=0; i<k; i++){
        while(q.length && nums[i] >= last()) q.pop();
        q.push(i);
    }
    
    const result = [nums[q[0]]];
    for(let i=k; i<nums.length; i++){
        while(q.length && nums[i] >= last()) q.pop();
        q.push(i); // add the newest member of our window
        while(i - q[0] >= k) q.shift(); // remove things not in our window
        result.push(nums[q[0]]);
    }
    
    return result;
};
