// https://leetcode.com/problems/climbing-stairs/

/**
 * @param {number} n
 * @return {number}
 */

// memoized
var climbStairs = function(n) {

    const memo = new Array(n).fill(-1);

    const r = (step) => {
        if(step > n) return 0;
        if(step === n) return 1;
        if(memo[step] !== -1) return memo[step];
        return memo[step] = r(step + 1) + r(step + 2);
    };

    return r(0);
};

// brute force backtracking
// var climbStairs = function(n) {

//     const r = (step) => {
//         if(step > n) return 0;
//         if(step === n) return 1;
//         return r(step + 1) + r(step + 2);
//     };

//     return r(0);
// };


// var climbStairs = function(n) {

//     if(n === 1) return 1;
//     if(n === 2) return 2;

//     let step = 2;
//     let twoAgo = 1;
//     let oneAgo = 1;

//     let result = 0; // placeholder

//     while(step++ <= n){
//         result = twoAgo + oneAgo;
//         twoAgo = oneAgo;
//         oneAgo = result;
//     }

//     return result;
// };

//  0. 1. 2. 3. 4. 5
// [1, 1, 2, 0, 0, 0]
