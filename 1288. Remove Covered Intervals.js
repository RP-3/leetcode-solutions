// https://leetcode.com/problems/remove-covered-intervals/

/**
 * @param {number[][]} intervals
 * @return {number}
 */
var removeCoveredIntervals = function(intervals) {
    intervals.sort(([aStart, aEnd], [bStart, bEnd]) => {
        if(aStart !== bStart) return aStart - bStart; // asc order of start
        return bEnd - aEnd; // desc order of end. E.g., [[1,3], [1,2] ...]
    });

    let [prev, deleted] = [intervals[0], 0];
    for(let i=1; i<intervals.length; i++){
        if(intervals[i][1] <= prev[1]) deleted++;
        else prev = intervals[i];
    }

    return intervals.length - deleted;
};