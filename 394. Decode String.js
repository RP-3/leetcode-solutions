// https://leetcode.com/problems/decode-string/

/**
 * @param {string} s
 * @return {string}
 */
var decodeString = function(s) {
    const stack = [];
    // utilities to tell what's at the top of the stack without fiddly syntax
    const lastMatches = (re) => stack.length && stack[stack.length-1].match(re);
    const [sq, d] = [new RegExp('[^\[]'), new RegExp('\\d')]; // detect not '[' and any digit, respectively

    s.split('').forEach((c) => {
        if(c !== ']') return stack.push(c);
        const [reverseSeg, reverseNum] = [[], []];
        while(lastMatches(sq)) reverseSeg.push(stack.pop()); // fetch the letters
        stack.pop(); // delete the trailing '['
        while(lastMatches(d)) reverseNum.push(stack.pop());  // fetch the numbers
        const seg = reverseSeg.reverse().join('');           // parse out the segment
        const mul = parseInt(reverseNum.reverse().join('')); // parse out the multiplier
        stack.push(...seg.repeat(mul).split('')); // simplify and add back to stack
    });

    return stack.join('');
};

