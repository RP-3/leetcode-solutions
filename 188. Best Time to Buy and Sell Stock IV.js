// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv

/**
 * @param {number} k
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(trx, prices) {

   if(trx > (prices.length / 2) ){
        let profit = 0;
        for(let i = 1; i < prices.length; i++){
            if(prices[i] > prices[i - 1]) profit += prices[i] - prices[i - 1];
        }
        return profit;
    }

    let prevK = new Array(prices.length).fill(0);
    let currentK = new Array(prices.length).fill(0);

    for(let k=0; k<trx; k++){
        [prevK, currentK] = [currentK, prevK];
        for(let i=0; i<prices.length; i++){

            let maxProfAti = 0;
            for(let j=0; j<i; j++){
                const profit = prices[i] - prices[j];
                const remainder = (j > 0) ? prevK[j-1] : 0;
                maxProfAti = Math.max(maxProfAti, profit + remainder);
            }
            const maxProfBeforei = i > 0 ? currentK[i-1] : 0;
            currentK[i] = Math.max(maxProfAti, maxProfBeforei);
        }
    }

    return currentK[prices.length-1];
};