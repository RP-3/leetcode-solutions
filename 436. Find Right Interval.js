// https://leetcode.com/problems/find-right-interval/

/**
 * @param {number[][]} intervals
 * @return {number[]}
 */
var findRightInterval = function(intervals) {
    const events = [];
    intervals.forEach(([start, end], i) => {
        events.push([true, start, i]);
        events.push([false, end, i]);
    });
    events.sort((a, b) => {
        if(a[1] !== b[1]) return a[1] - b[1]; // asc order of pos
        if(a[0]) return 1; // if pos is eq, ends come first
        return -1;
    });

    const result = new Array(intervals.length).fill(-1);
    const ended = new Set();

    events.forEach(([isStart, pos, i]) => {
        if(!isStart) return ended.add(i);
        for(let endedIntervalIndex of ended){
            ended.delete(endedIntervalIndex);
            result[endedIntervalIndex] = i;
        }
    });

    return result;
};