// https://leetcode.com/problems/house-robber-ii/

/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {

    const memo = new Array(nums.length).fill(0).map(() => new Array(2));

    const r = (i, canRobLast) => {
        if(i >= nums.length) return 0;
        if(i === nums.length-1) return canRobLast ? nums[i] : 0;

        if(memo[i][canRobLast] !== undefined) return memo[i][canRobLast];

        if(!i){
            return memo[i][canRobLast] = Math.max(
                nums[i] + r(i+2, false), // rob this house
                0 + r(i+1, true) // do not rob this house
            );
        }

        return memo[i][canRobLast] = Math.max(
            nums[i] + r(i+2, canRobLast), // rob this house
            0 + r(i+1, canRobLast) // do not rob this house
        );
    };

    return r(0, true);
};
