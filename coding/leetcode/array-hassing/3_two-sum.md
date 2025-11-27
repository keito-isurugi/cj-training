# Two Sum

## Question
Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.

## Solution
```.go
func twoSum(nums []int, target int) []int {
    // value => indexを入れるmapを作る
    nm := make(map[int]int)
    for i, n := range nums {
        // keyにtargetとnの差分が存在するかを確認する
        if j, ok := nm[target - n]; ok {
            return []int{j, i}
        }
        // 存在しなければマップに追加する
        nm[n] = i
    }

    return nil
}
```
