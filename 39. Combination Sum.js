// https://leetcode.com/problems/combination-sum/

/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
    const [result, wc] = [[], []];

    const backtrack = (i, remainder) => {
        if(i >= candidates.length || remainder < 0) return;
        if(remainder === 0) return result.push(wc.slice());

        const curr = candidates[i];
        wc.push(curr);
        backtrack(i, remainder - curr); // try using this item again
        wc.pop();
        backtrack(i+1, remainder); // try not using this item
    };

    backtrack(0, target);
    return result;
};