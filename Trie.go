package main

import (
	"math"
	"sort"
)

/*
Trie
https://leetcode.com/problems/implement-trie-prefix-tree/description
208. Implement Trie (Prefix Tree)
*/
type Trie struct {
	RootNode *NodeTrie
}

type NodeTrie struct {
	Char     string
	Children [27]*NodeTrie
	EndWord  bool
}

func NewNode(char string) *NodeTrie {
	node := &NodeTrie{Char: char}
	for i := 0; i < 27; i++ {
		node.Children[i] = nil
	}
	return node
}

func ConstructorTrie() Trie {
	root := NewNode("\000")
	return Trie{RootNode: root}
}

func (t *Trie) Insert(word string) {
	current := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if index == 206 {
			index = 26
		}
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(word[i]))
		}
		current = current.Children[index]
	}
	current.EndWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	if current.EndWord != true {
		return false
	}
	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.RootNode
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if index == 206 {
			index = 26
		}
		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return true
}

func (t *Trie) IsParentOfAny(subfolder string) bool {
	current := t.RootNode
	for i := 0; i < len(subfolder); i++ {
		index := subfolder[i] - 'a'
		if index == 206 {
			index = 26
		}
		if current != nil && current.EndWord == true && index == 26 {
			return true
		}
		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return false
}

/*
14. Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix/description
*/
func longestCommonPrefix(strs []string) string {
	idx := 0
	flag := true
	for idx < len(strs[0]) && flag {
		idx++
		currStr := strs[0][:idx]
		for i := 1; i < len(strs); i++ {
			if idx > len(strs[i]) {
				flag = false
				idx--
				break
			}
			if strs[i][:idx] != currStr {
				flag = false
				idx--
				break
			}
		}
	}
	return strs[0][:idx]
}

/*
139. Word Break
https://leetcode.com/problems/word-break/description/?envType=problem-list-v2&envId=trie
*/
func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	mapWord := make(map[string]bool, n)
	for _, word := range wordDict {
		mapWord[word] = true
	}

	dp := make([]int, n+1)
	dp[0] = 0
	var founded []int
	for i := 1; i <= n; i++ {
		if mapWord[s[:i]] {
			dp[i] = i
			founded = append(founded, i)
		} else {
			for _, idx := range founded {
				if mapWord[s[idx:i]] {
					dp[i] = i
					founded = append(founded, i)
					break
				}
			}
		}

	}
	return dp[n] == n
}

/*
140. Word Break II
https://leetcode.com/problems/word-break-ii/description/?envType=problem-list-v2&envId=trie
*/
func wordBreakII(s string, wordDict []string) []string {
	n := len(s)

	mapWord := make(map[string]bool, n)
	for _, word := range wordDict {
		mapWord[word] = true
	}
	founded := make(map[int][]string, n)
	for i := 1; i <= n; i++ {
		if mapWord[s[:i]] {
			founded[i] = []string{s[:i]}
		}
		for key, value := range founded {
			if mapWord[s[key:i]] {
				for j := range value {
					founded[i] = append(founded[i], value[j]+" "+s[key:i])
				}
			}
		}
	}
	return founded[n]
}

/*
WordDictionary
211. Design Add and Search Words Data Structure
https://leetcode.com/problems/design-add-and-search-words-data-structure/description
*/
type WordDictionary struct {
	RootNode *NodeTrie
}

func Constructor() WordDictionary {
	root := NewNode("\000")
	return WordDictionary{RootNode: root}
}

func (t *WordDictionary) AddWord(word string) {
	current := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(word[i]))
		}
		current = current.Children[index]
	}
	current.EndWord = true
}

func (t *WordDictionary) Search(word string) bool {

	var helper func(current *NodeTrie, word string) bool
	helper = func(current *NodeTrie, word string) bool {
		for i := 0; i < len(word); i++ {
			if word[i] == '.' {
				for _, nextWord := range current.Children {
					if nextWord != nil {
						if helper(nextWord, word[i+1:]) {
							return true
						}
					}
				}
				return false
			} else {
				index := word[i] - 'a'
				if current == nil || current.Children[index] == nil {
					return false
				}
				current = current.Children[index]
			}
		}
		if current.EndWord != true {
			return false
		}
		return true
	}

	return helper(t.RootNode, word)

}

/*
212. Word Search II
https://leetcode.com/problems/word-search-ii/?envType=problem-list-v2&envId=trie
*/

func (t *Trie) SearchNotNested(word string) bool {
	current := t.RootNode
	visited := make(map[*NodeTrie]bool)
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		visited[current] = true
		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
		if visited[current] {
			return false
		}
	}
	return true
}

func findWords(board [][]byte, words []string) []string {
	//trie
	trie := ConstructorTrie()
	mapTrie := make(map[[2]int]*NodeTrie)

	m := len(board[0])
	n := len(board)

	//directions
	directions := map[int][2]int{
		0: {0, 1},  // Right
		1: {1, 0},  // Down
		2: {0, -1}, // Left
		3: {-1, 0}, // Up
	}

	root := trie.RootNode
	//index := charIndex - 'a'

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			charIndex := board[i][j]
			index := charIndex - 'a'
			position := [2]int{i, j}
			mapTrie[position] = NewNode(string(charIndex))
			if root.Children[index] == nil {
				root.Children[index] = mapTrie[position]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			position := [2]int{i, j}
			currNode := root.Children[mapTrie[position].Char[0]-'a']

			for _, dir := range directions {
				n1, m1 := dir[0]+position[0], dir[1]+position[1]
				if m1 > -1 && m1 < m && n1 > -1 && n1 < n && mapTrie[[2]int{n1, m1}] != nil {
					linkNode := mapTrie[[2]int{n1, m1}]
					currNode.Children[linkNode.Char[0]-'a'] = linkNode
					mapTrie[position].Children[linkNode.Char[0]-'a'] = linkNode
				}
			}
		}
	}

	var output []string
	for _, word := range words {
		if trie.SearchNotNested(word) {
			output = append(output, word)
		}
	}

	return output
}

/*
1233. Remove Sub-Folders from the Filesystem
https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/description/
*/
func removeSubfolders(folder []string) []string {
	trie := ConstructorTrie()
	var output []string
	sort.Slice(folder, func(i, j int) bool {
		return len(folder[i]) < len(folder[j])
	})
	for _, subfolder := range folder {
		if !trie.IsParentOfAny(subfolder) {
			trie.Insert(subfolder)
			output = append(output, subfolder)
		}
	}
	return output
}

/*
3291. Minimum Number of Valid Strings to Form Target I
https://leetcode.com/problems/minimum-number-of-valid-strings-to-form-target-i/description
*/
func minValidStrings(words []string, target string) int {
	trie := ConstructorTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	dp := make([]int, len(target)+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= len(target); i++ {
		for j := 0; j < i; j++ {
			if dp[j] != math.MaxInt32 && trie.StartsWith(target[j:i]) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	if dp[len(target)] == math.MaxInt32 {
		return -1
	}
	return dp[len(target)]
}
