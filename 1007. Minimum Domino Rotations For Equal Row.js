// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/

/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number}
 */
var minDominoRotations = function(A, B) {
    if(!A.length) return 0;

    const count = (v) => {
        let [aCount, bCount] = [0, 0];
        // convert A[x] to v
        for(let i=0; i<A.length; i++){
            if(A[i] !== v){
                if(B[i] === v) aCount++;
                else aCount = Infinity;
            }
        }

        // convert B[x] to v
        for(let i=0; i<B.length; i++){
            if(B[i] !== v){
                if(A[i] === v) bCount++;
                else bCount = Infinity;
            }
        }

        return [aCount, bCount];
    };

    const result = Math.min(...count(A[0]), ...count(B[0]));

    return result === Infinity ? -1 : result;
}