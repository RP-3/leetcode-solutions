// https://leetcode.com/problems/flipping-an-image/

/**
 * @param {number[][]} A
 * @return {number[][]}
 */
var flipAndInvertImage = function(A) {
    return new Array(A.length).fill(0)
    .map((_, r) => A[r].map((_, c) => ~A[r][A[0].length - c - 1] & 1));
};