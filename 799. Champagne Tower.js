// https://leetcode.com/problems/champagne-tower/

/**
 * @param {number} poured
 * @param {number} query_row
 * @param {number} query_glass
 * @return {number}
 */
var champagneTower = function(poured, row, col) {
    if(!row) return Math.min(poured, 1);

    const memo = new Array(row+1).fill(0).map((_, i) => new Array(i+1).fill(undefined));

    const leftOverFrom = (row, col) => {
        if(row === 0) return (col !== 0 || poured <= 1) ? 0 : ((poured - 1) / 2);
        if(col < 0 || col >= row+1) return 0; // OOB

        if(memo[row][col] !== undefined) return memo[row][col]; // memoized

        const totalPouredHere = leftOverFrom(row-1, col-1) + leftOverFrom(row-1, col);
        if(totalPouredHere <= 1) return memo[row][col] = 0;
        return memo[row][col] = (totalPouredHere - 1) / 2;
    };

    return Math.min(leftOverFrom(row-1, col-1) + leftOverFrom(row-1, col), 1);
};
