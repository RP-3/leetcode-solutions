https://leetcode.com/problems/sort-array-by-parity/

/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortArrayByParity = function(A) {
    let [l,r] = [0,A.length-1]
    while(l<r){
      if(A[l] %2){
        [A[l],A[r]] = [A[r],A[l]]
        r--
      } else {
        l++
      }
    }
    return A;
};