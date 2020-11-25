// https://leetcode.com/problems/non-overlapping-intervals/

/**
 * @param {number[][]} intervals
 * @return {number}
 */
var eraseOverlapIntervals = function(intervals) {

    intervals.sort((a, b) => a[0] - b[0]); // asc by start time

    let result = 0;
    let slow = 0; let fast = 1;

    while(fast < intervals.length){

        let [s, f] = [intervals[slow], intervals[fast]];

        if(s[1] > f[0] && f[1] > s[0]){ // overlaps

            result++;

            if(intervals[slow][1] <= intervals[fast][1]){
                // delete the fast one
                fast++;
            }else{
                // delete the slow one
                slow = fast;
                fast++;
            }
        }else{ // no overlap

            slow = fast;
            fast++;
        }
    }

    return result;
};