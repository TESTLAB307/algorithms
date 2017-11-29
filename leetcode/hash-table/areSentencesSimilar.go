/* https://leetcode.com/problems/sentence-similarity/description/
Given two sentences words1, words2 (each represented as an array of strings), and a list of similar word pairs pairs, determine if two sentences are similar.

For example, "great acting skills" and "fine drama talent" are similar, if the similar word pairs are pairs = [["great", "fine"], ["acting","drama"], ["skills","talent"]].

Note that the similarity relation is not transitive. For example, if "great" and "fine" are similar, and "fine" and "good" are similar, "great" and "good" are not necessarily similar.

However, similarity is symmetric. For example, "great" and "fine" being similar is the same as "fine" and "great" being similar.

Also, a word is always similar with itself. For example, the sentences words1 = ["great"], words2 = ["great"], pairs = [] are similar, even though there are no specified similar word pairs.

Finally, sentences can only be similar if they have the same number of words. So a sentence like words1 = ["great"] can never be similar to words2 = ["doubleplus","good"].

Note:

	The length of words1 and words2 will not exceed 1000.
	The length of pairs will not exceed 2000.
	The length of each pairs[i] will be 2.
	The length of each words[i] and pairs[i][j] will be in the range [1, 20].
*/

package leetcode

func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

	if len(pairs) == 0 {
		for i, word := range words1 {
			if word != words2[i] {
				return false
			}
		}
		return true
	}

	pairsDict := make(map[string]map[string]int)
	for _, pair := range pairs {
		p1, p2 := pair[0], pair[1]
		if _, ok := pairsDict[p1]; !ok {
			pairsDict[p1] = make(map[string]int, 0)
		}
		pairsDict[p1][p2] = 1
		if _, ok := pairsDict[p2]; !ok {
			pairsDict[p2] = make(map[string]int, 0)
		}
		pairsDict[p2][p1] = 1
	}

	for i, word := range words1 {
		if word != words2[i] {
			if _, ok := pairsDict[word]; !ok {
				return false
			}
			if _, ok := pairsDict[word][words2[i]]; !ok {
				return false
			}
		}
	}
	return true
}
