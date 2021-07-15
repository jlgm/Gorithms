// https://leetcode.com/problems/subsets/

func rec(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{nums}
    }
    
    head := nums[0]
    ans := make([][]int, 0)
    ans = append (ans, []int{head})
    
    set := rec(nums[1:len(nums)])
    for _, elem := range set {
        ans = append(ans, elem)
        tmp := make([]int, 0)
        for _, num := range elem {
            tmp = append(tmp, num)
        }
        tmp = append(tmp, head)
        ans = append(ans, tmp)
    }
    return ans
}

func subsets(nums []int) [][]int {
    ans := rec(nums)
    ans = append(ans, []int{})
    return ans
}
