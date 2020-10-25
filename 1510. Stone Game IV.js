// https://leetcode.com/problems/stone-game-iv/

/**
 * @param {number} n
 * @return {boolean}
 */
var winnerSquareGame = function(n) {

    const memo = new Array(n+1).fill(0).map(() => []);

    const aliceAlwaysWins = (n, aliceTurn) => {
        const root = Math.sqrt(n); // if the root is zero
        if(Math.floor(root) === root) return aliceTurn; // the current player wins

        if(memo[n][aliceTurn] !== undefined) return memo[n][aliceTurn]; //

        if(aliceTurn){ // if it's alice's turn, alice always wins if
            for(let i=1; i<=root; i++){ // there's a move she can make that causes her to win
                if(aliceAlwaysWins(n - i*i, false)) return memo[n][aliceTurn] = true; // with n-i*i stones
            }
            return memo[n][aliceTurn] = false;
        }

        // And if its bob's turn
        for(let i=1; i<=root; i++){ // if there's a single move bob can make that makes alice not win
            if(!aliceAlwaysWins(n - i*i, true)) return memo[n][aliceTurn] = false; // then alice doesn't always win
        }
        return memo[n][aliceTurn] = true // but if there's no such move, bob wins
    };

    return aliceAlwaysWins(n, true);
};