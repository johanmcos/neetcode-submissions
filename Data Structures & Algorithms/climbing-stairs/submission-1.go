func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    prev2 := 1  // ways to climb 1 step
    prev1 := 2  // ways to climb 2 steps
    
    for i := 3; i <= n; i++ {
        current := prev1 + prev2
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}