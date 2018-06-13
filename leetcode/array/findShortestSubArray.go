/* https://leetcode.com/problems/count-binary-substrings/description/
Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example 1:
    Input: [1, 2, 2, 3, 1]
    Output: 2
    Explanation:
    The input array has a degree of 2 because both elements 1 and 2 appear twice.
    Of the subarrays that have the same degree:
    [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
    The shortest length is 2. So return 2.
Example 2:
    Input: [1,2,2,3,1,4,2]
    Output: 6

Note:
    nums.length will be between 1 and 50,000.
    nums[i] will be an integer between 0 and 49,999.
*/

package larray

func findShortestSubArray(nums []int) int {
	type Info struct {
		first, end int
		count      int
	}
	maps := make(map[int]*Info)
	for i, num := range nums {
		if info, ok := maps[num]; !ok {
			info = &Info{i, i, 1}
			maps[num] = info
		} else {
			info.end = i
			info.count++
		}
	}

	tmp, res := 0, 0
	for _, v := range maps {
		switch {
		case v.count > tmp:
			tmp, res = v.count, v.end-v.first+1
		case v.count == tmp:
			if length := v.end - v.first + 1; length <= res {
				res = length
			}
		}
	}
	return res
}
