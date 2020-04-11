/*
 * hash表检索法
 * 时间复杂度O(N)，空间复杂度O(N)
 */

func twoSum(nums []int, target int) []int {
    nums_map := make(map[int]int)
    
    for k, v := range nums {
        complement := target - v
        elt, ok := nums_map[complement]
        if ok {
            return []int{elt, k}
        }
        nums_map[v] = k
    }
    return nil
}