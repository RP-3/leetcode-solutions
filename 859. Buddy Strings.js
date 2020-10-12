// https://leetcode.com/problems/buddy-strings/

/**
 * @param {string} A
 * @param {string} B
 * @return {boolean}
 */
var buddyStrings = function(A, B) {
    if(A.length !== B.length || A.length < 2) return false;

    const d = []; // diffs
    for(let i=0; i<A.length; i++){
        if(A[i] !== B[i]) d.push([A[i], B[i]]);
    }

    if(d.length > 2 || d.length === 1) return false;
    if(d.length === 2) return d[0][0] === d[1][1] && d[0][1] === d[1][0];

    // no diffs. If a has two of any char we can swap them
    const counts = new Array(26).fill(0);
    for(let i=0; i<A.length; i++){
        const c = A.charCodeAt(i) - 97;
        if(++counts[c] > 1) return true;
    }

    return false;
};