// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/

/**
 * @param {number[]} nums
 * @param {number} threshold
 * @return {number}
 */
var smallestDivisor = function(nums, threshold) {

    const sumDiv = (d) => nums.reduce((a, b) => a + Math.ceil(b/d), 0);
    const max = Math.max(...nums);

    let [l, r, result] = [1, max, max];
    while(l <= r){
        const mid = Math.floor((l + r) / 2);
        const sum = sumDiv(mid);
        if(sum <= threshold){
            result = Math.min(result, mid);
            r = mid-1;
        }else{
            l = mid + 1;
        }
    }

    return result;
};