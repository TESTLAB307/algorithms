/* https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/#/description
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

You may assume that each input would have exactly one solution and you may not use the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2
*/

package leetcode

func twoSum2(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		} else if sum > target {
			end--
		} else {
			start++
		}
	}
	return []int{}
}