package util

import "mapreduce_service/vojo"

func Heapsort(nums []vojo.TopServiceKV) []vojo.TopServiceKV {
	lenght := len(nums)
	if lenght == 0 {
		return nil
	}
	for i := 0; i < lenght; i++ {
		MiniHeap(nums[i:])
	}
	return nums
}

func MiniHeap(nums []vojo.TopServiceKV) {
	length := len(nums)
	floor := length/2 - 1
	for i := floor; i >= 0; i-- {

		baba := i
		left := 2*i + 1
		right := 2*i + 2

		if right < length && nums[baba].Times > nums[right].Times {
			nums[baba], nums[right] = nums[right], nums[baba]
		}

		if left < length && nums[baba].Times > nums[left].Times {
			nums[baba], nums[left] = nums[left], nums[baba]
		}

	}
}
