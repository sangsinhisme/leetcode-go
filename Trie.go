package main

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
	Children [26]*NodeTrie
	EndWord  bool
}

func NewNode(char string) *NodeTrie {
	node := &NodeTrie{Char: char}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}
	return node
}

func Constructor() Trie {
	root := NewNode("\000")
	return Trie{RootNode: root}
}

func (t *Trie) Insert(word string) {
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
		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return true
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
