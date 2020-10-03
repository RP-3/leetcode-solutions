// https://leetcode.com/problems/k-diff-pairs-in-an-array/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
// O(n.log(n)) time.
// O(1) space
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

// O(n) time
// O(n) space
var findPairs = function(nums, k) {
    const seen = new Set();
    const paired = new Set();
    let result = 0;

    for(let i=0; i<nums.length; i++){
        const w = [nums[i] - k, nums[i] + k];

        w.forEach((wanted) => {
            if(!seen.has(wanted)) return;
            const pair = wanted < nums[i] ? `${wanted}:${nums[i]}` : `${nums[i]}:${wanted}`;
            if(paired.has(pair)) return;
            result++;
            paired.add(pair)
        });

        seen.add(nums[i]);
    }

    return result;
}