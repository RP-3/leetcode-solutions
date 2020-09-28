// https://leetcode.com/problems/pancake-sorting/

/**
 * @param {number[]} A
 * @return {number[]}
 */
var pancakeSort = function(A) {

    const wc = [];
    let end = A.length-1;

    while(end > 0) {
        const i = indexOfLargest(end);
        if(i !== end){
            wc.push(i+1);
            flip(i);
            wc.push(end+1);
            flip(end);
        }
        end--;
    };

    return wc;

    // private helpers
    function indexOfLargest(end){
        let [largest, index] = [A[0], 0];
        for(let i=0; i<=end; i++){
            if(A[i] > largest) [largest, index] = [A[i], i];
        }
        return index;
    }

    function flip(index){
        let [l, r] = [0, index];
        while(l < r){
            [A[l], A[r]] = [A[r], A[l]];
            l++; r--;
        }
    }
};