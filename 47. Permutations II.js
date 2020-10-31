https://leetcode.com/problems/permutations-ii/

var permuteUnique = function(nums) {

    if(nums.length <= 1) return [nums];

    const result = [];
    nums.sort((a, b) => a - b);

    const rightFoot = () => {
        for(let i=nums.length-1; i>0; i--){
            if(nums[i] > nums[i-1]) return i-1;
        }
        return -1;
    };

    const reverse = (l, r) => {
        while(l < r){
            [nums[l], nums[r]] = [nums[r], nums[l]];
            l++; r--;
        }
    }

    // find the next largest permutation
    const permute = () => {
        let foot = rightFoot();
        if(foot === -1) return false;

        let i = foot + 1;
        while(nums[i] > nums[foot] && i < nums.length) i++;
        i--;
        [nums[foot], nums[i]] = [nums[i], nums[foot]];

        reverse(foot + 1, nums.length-1);
        return true;
    };

    result.push(nums.slice());
    while(permute()){
        result.push(nums.slice());
    }

    return result;
};