// https://leetcode.com/problems/numbers-with-same-consecutive-differences/

/**
 * @param {number} N
 * @param {number} K
 * @return {number[]}
 */
var numsSameConsecDiff = function(N, K) {

    if(N === 1) return [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

    const result = [];
    const wc = [];

    const build = () => {

        const last = wc[wc.length-1];
        if(last < 0 || last > 9) return;
        if(wc.length === N) return result.push(wc.join(''));

        wc.push(last + K); // try last + K
        build();
        wc.pop(); // backtrack

        if(!K) return;

        wc.push(last - K); // try last - K
        build();
        wc.pop(); // backtrack
    };

    for(let i=1; i<=9; i++){
        wc.push(i);
        build();
        wc.pop();
    }

    return result.map((digits) => parseInt(digits, 10));
};