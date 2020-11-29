// https://leetcode.com/problems/jump-game-iii/

package leetcode

func canReach(arr []int, start int) bool {
    visited := make([]bool, len(arr))
    
    var dfs func(int) bool
    dfs = func(i int) bool {
        switch {
        case i < 0 || i >= len(arr) || visited[i]:
            return false // OOB or already visited
        case arr[i] == 0:
            return true
        default:
            visited[i] = true
            return dfs(i + arr[i]) || dfs(i - arr[i])    
        }       
    }
    
    return dfs(start)
}
