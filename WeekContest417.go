package main

/*
Q1. Find the K-th Character in String Game I
*/
func kthCharacter(k int) byte {
	word := "a"
	for len(word) < k {
		newWord := ""
		for _, char := range word {
			nextChart := char + 1
			if nextChart == 122 {
				nextChart = 97
			}
			newWord = newWord + string(nextChart)
		}
		word += newWord
	}
	return word[k-1]
}

/*
Q2. Count of Substrings Containing Every Vowel and K Consonants I
*/
func countOfSubstringsQ2(word string, k int) int64 {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	output := 0
	n := len(word)

	for i := 0; i < n; i++ {
		initVowels := make(map[rune]bool)
		consonants := 0

		for j := i; j < n; j++ {
			char := rune(word[j])

			// Check if the character is a vowel
			if vowels[char] {
				initVowels[char] = true
			} else {
				consonants++
			}

			// If consonants exceed k, break out of the loop
			if consonants > k {
				break
			}

			// If we have all vowels and exactly k consonants, count the substring
			if len(initVowels) == len(vowels) && consonants == k {
				output++
			}
		}
	}

	return int64(output)
}

/*
Q3. Count of Substrings Containing Every Vowel and K Consonants II
*/
func countOfSubstrings(word string, k int) int64 {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}
	output := 0
	n := len(word)
	limit := len(vowels) + k

	for i := 0; i <= n-limit; i++ {
		initVowels := make(map[rune]bool)
		consonants := 0

		// Check the substring from i to i + limit
		for j := i; j < i+limit && j < n; j++ {
			char := rune(word[j])
			if vowels[char] {
				initVowels[char] = true
			} else {
				consonants++
			}
		}

		// Check if we have all vowels and exactly k consonants
		if len(initVowels) == len(vowels) && consonants == k {
			output++

			// Now, check for additional valid substrings
			for j := i + limit; j < n; j++ {
				if vowels[rune(word[j])] {
					output++
				} else {
					break
				}
			}
		}
	}

	return int64(output)
}

/*
Q4. Find the K-th Character in String Game II
*/
func kthCharacterQ4(k int64, operations []int) byte {
	return 0
}
