// https://leetcode.com/problems/remove-duplicate-letters/

/**
 * @param {string} s
 * @return {string}
 */
var removeDuplicateLetters = function(s) {
    const result = [];
    const index = new Array(26).fill(-1); // last index of each char
    for(let i=0; i<s.length; i++) index[s.charCodeAt(i) - 97] = i;

    const lastChar = () => result[result.length-1];
    const lastCharCode = () => lastChar().charCodeAt(0) - 97;

    for(let i=0; i<s.length; i++){
        const c = s[i]; // for each char
        if(result.indexOf(c) > -1) continue; // if it already exists, skip it. Else:
        // while c is less than last char AND last char occurs later (i.e., result IS improvable)
        while(result.length && index[lastCharCode()] > i && lastChar() > c) result.pop(); // remove the last char
        result.push(c); // then improve the result, with confidence that we'll add the chars we just removed later
    }

    return result.join('');
};
