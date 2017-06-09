/* https://leetcode.com/problems/longest-uncommon-subsequence-i/#/description
Given a group of two strings, you need to find the longest uncommon subsequence of this group of two strings.
The longest uncommon subsequence is defined as the longest subsequence of one of these strings and this subsequence should not be any subsequence of the other strings.

A subsequence is a sequence that can be derived from one sequence by deleting some characters without changing the order of the remaining elements.
Trivially, any string is a subsequence of itself and an empty string is a subsequence of any string.

The input will be two strings, and the output needs to be the length of the longest uncommon subsequence. If the longest uncommon subsequence doesn't exist, return -1.

Example 1:
    Input: "aba", "cdc"
    Output: 3
    Explanation: The longest uncommon subsequence is "aba" (or "cdc"),
    because "aba" is a subsequence of "aba",
    but not a subsequence of any other strings in the group of two strings.

Note:
    Both strings' lengths will not exceed 100.
    Only letters from a ~ z will appear in input strings.
*/

package leetcode

func findLUSlength(a string, b string) int {
	// 两个字符串相等，一定没有非共同子序列
	// 反之，不等，较长的那个字符串就是最长非共同子序列
	if a == b {
		return -1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(len(a), len(b))
}