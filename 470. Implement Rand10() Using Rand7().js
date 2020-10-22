// https://leetcode.com/problems/implement-rand10-using-rand7/

/**
 * The rand7() API is already defined for you.
 * var rand7 = function() {}
 * @return {number} a random integer in the range 1 to 7
 */
/*
IDEA: Make two calls to rand7 to get a value from this table:
  1. 2  3. 4. 5. 6. 7
1 2  3  4  5  6  7  8
2 3  4  5  6  7  8  9
3 4  5  6  7  8  9  10
4 5  6  7  8  9  10 11
5 6  7  8  9  10 11 12
6 7  8  9  10 11 12 13
7 8  9  10 11 12 13 14

So the values we can get from 2 calls range from 2 - 14.
What is the probability that we get each of those vals?

 2 1  4               x
 3 2  3               x
 4 3  4               x
 5 4  1               x
 6 5  6 and c1 not 5  x
 7 6  8 and c1 <= 4   x
 8 7  9 and c1 <= 4   x
 9 6  10 and <=5
10 5  7 and c1 <= 6   x
11 4  2               x
12 3  5               x
13 2  3               x
14 1  5               x
*/
var rand10 = function() {
    while(true){
        let [c1, c2] = [rand7(), rand7()];
        let score = c1 + (c2-1)*7;
        if(score <= 40) return 1 + ((score - 1) % 10)
    }
};