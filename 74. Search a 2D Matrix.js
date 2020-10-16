// https://leetcode.com/problems/search-a-2d-matrix/

/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(m, target) {
    if(!m.length || !m[0].length) return false;

    const [rows, cols] = [m.length, m[0].length];
    const totalCells = rows * cols;
    let [l, r] = [0, totalCells-1];

    while(l <= r){
        const mid = Math.floor((l+r)/2);
        const row = Math.floor(mid / cols);
        const col = mid % cols;
        const val = m[row][col];

        if(val < target) l = mid + 1;
        else if(val > target) r = mid - 1;
        else return true;
    }

    return false;
};
