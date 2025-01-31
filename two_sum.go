package main

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	myMap[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		a := target - nums[i]
		b := myMap[a]
		if b == 0 {
			myMap[nums[i]] = i + 1
		} else {
			return []int{i, b - 1}
		}
	}
	return []int{}
}
