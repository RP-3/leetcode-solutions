// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/

/**
 * @param {number[][]} points
 * @return {number}
 */
var findMinArrowShots = function(points) {
    points.sort((a, b) => a[1] - b[1]); // asc order of end

    let [c, arrows] = [0, 0];
    while(c < points.length){
        arrows++; // shoot an arrow...
        let rightEdge = points[c++][1]; // at the right end of this balloon
        // Then skip all the balloons that were hit by that arrow
        while(c < points.length && points[c][0] <= rightEdge) c++;
    }

    return arrows;
};
