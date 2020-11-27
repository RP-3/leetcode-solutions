// https://leetcode.com/problems/partition-equal-subset-sum/

package leetcode

func canPartitionTopDown(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	/*
	memo strategy: use:
	0 = unset
	1 = true
	2 = false
	*/
	memo := make([][]uint8, target+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]uint8, len(nums))
	}

	var search func(i, r int) bool
	search = func(i, r int) bool {
		switch {
		case r == 0: // found valid answer
			return true
		case i >= len(nums) || r < 0: // out of bounds
			return false
		case memo[r][i] != 0: // answer memoized, stop early
			return memo[r][i] == 1
		default: // try using nums[i] || skipping nums[i]
			if search(i+1, r-nums[i]) || search(i+1, r) {
				memo[r][i] = 1
			} else {
				memo[r][i] = 2
			}
			return memo[r][i] == 1
		}
	}

	return search(0, target)
}

func canPartitionBottomUp(nums []int) bool {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    
    if sum % 2 != 0 {
        return false
    }
    target := sum / 2
    
	memo := make([][]bool, target+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]bool, len(nums))
	}
    
    for r := 0; r <= target; r++ { // can we get to this remainder
        for i:=0; i<len(nums); i++ { // using vals nums[0..i]
            switch {
            case r == 0: // we can always make 0 with any elements
                memo[r][i] = true 
            case r == nums[i]: // and if the remainder is exactly what we want, that works too
                memo[r][i] = true
            case i == 0: // an empty set can never yield a non-zero answer
                memo[r][i] = false
			case nums[i] <= r && memo[r - nums[i]][i-1]: // if using this gets us to a state
                memo[r][i] = true // that we previously found was valid, that works too
            case memo[r][i-1] == true: // or if the previous state was valid with no help from this one
                memo[r][i] = true
            default: // if none of the above worked, this is a dead end
                memo[r][i] = false
            }
        }
    }
    
    return memo[target][len(nums)-1]
}