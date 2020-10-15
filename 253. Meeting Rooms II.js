// https://leetcode.com/problems/meeting-rooms-ii/

/**
 * @param {number[][]} intervals
 * @return {number}
 */
var minMeetingRooms = function(intervals) {
    const events = [];
    let [result, meetingsInProgress] = [0, 0];

    for(let i=0; i<intervals.length; i++){
        events.push([intervals[i][0], 1]); // start
        events.push([intervals[i][1], -1]); // end
    }
    // asc order of time. If tied, ends come first
    events.sort((a, b) => (a[0] !== b[0]) ? (a[0] - b[0]) : (a[1] - b[1]));

    for(let i=0; i<events.length; i++){
        meetingsInProgress += events[i][1];
        result = Math.max(result, meetingsInProgress);
    }

    return result;
};