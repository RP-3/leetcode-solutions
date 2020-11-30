// https://leetcode.com/problems/maximum-average-subarray-ii/

// MARK: repeat

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {
    let l = Math.min(...nums), r = Math.max(...nums);
    const window = Math.pow(10, -5);
    
    let result = l;
    while(r - l >= window){
        const mid = (l + r) / 2;
        if(answerExists(mid)){
            result = Math.max(result, mid);
            l = mid;
        }else{
            r = mid;
        }
    }
    
    return result;
    
    // private helpers
    function answerExists(target){
        let head = 0, tail = 0, min = 0;
        
        for(let i=0; i<k; i++) head += (nums[i] - target);
        if(head >= 0) return true;
        
        for(let i=k; i<nums.length; i++){
            head += (nums[i] - target);
            tail += (nums[i-k] - target);
            min = Math.min(min, tail);
            if(head >= min) return true;
        }
        
        return false;
    }
};
