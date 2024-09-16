package main

type stack = []float64

// pops a value off the top of the stack, and returns the top value (This doesn't bounds-check the stack, so that must be done before calling the function)
func Pop(nums stack) (float64, stack) {
	val := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	return val, nums
}
