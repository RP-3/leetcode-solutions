// https://leetcode.com/problems/length-of-last-word/

/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    let [result, started] = [0, false];

    for(let i = s.length-1; i>=0; i--){
        if(s[i] !== ' '){
            result++;
            started = true;
        }
        else if(started) return result;
    }

    return result;
};