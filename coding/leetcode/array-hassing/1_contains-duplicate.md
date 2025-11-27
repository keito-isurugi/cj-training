# Contains Duplicate

## Question
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

## Solution
- マップを使用して存在確認を行う
```.go
func hasDuplicate(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return true
		}
		
		m[n] = true
	}
	
	return false
}
```
