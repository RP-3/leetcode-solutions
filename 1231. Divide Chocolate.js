// https://leetcode.com/problems/divide-chocolate/

/**
 * @param {number[]} sweetness
 * @param {number} K
 * @return {number}
 */
var maximizeSweetness = function(a, K) {
    let l = Math.min(...a);
    let r = a.reduce((a, b) => a + b, 0);
    if(!K) return r;

    let result = l;

    while(l <= r) {
        const mid = Math.floor((l + r) / 2);
        const good = answerExists(mid);
        if(good){
            result = Math.max(result, mid);
            l = mid+1;
        }else{
            r = mid - 1;
        }
    }

    return result;

    function answerExists(mid) {
        let sum = 0;
        let cuts = K;
        for(let i=0; i<a.length; i++){
            sum += a[i];
            if(sum >= mid && i !== (a.length-1)){
                if(cuts > 0){
                    cuts--;
                    sum = 0;
                }
            }
        }
        return cuts <= 0 && sum >= mid;
    }
};
