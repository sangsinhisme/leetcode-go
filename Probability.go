package main

/*
*
2028. Find Missing Observations
You have observations of n + m 6-sided dice rolls with each face numbered from 1 to 6. n of the observations went missing,
and you only have the observations of m rolls. Fortunately, you have also calculated the average value of the n + m rolls.

You are given an integer array rolls of length m where rolls[i] is the value of the ith observation. You are also given
the two integers mean and n.

Return an array of length n containing the missing observations such that the average value of the n + m rolls is exactly mean.
If there are multiple valid answers, return any of them. If no such array exists, return an empty array.
*/
func missingRolls(rolls []int, mean int, n int) []int {
	total := mean*(n+len(rolls)) - sum(rolls)
	if total > n*6 || total < n {
		return nil
	}
	arr := make([]int, n)
	currTotal := n
	for i := range arr {
		leftTotal := total - currTotal
		if leftTotal > 5 {
			arr[i] = 6
			currTotal += 5
		} else if leftTotal > 0 {
			arr[i] = 1 + leftTotal
			currTotal = total
		} else {
			arr[i] = 1
		}
	}
	return arr
}

func sum(arr []int) int {
	total := 0
	for _, i := range arr {
		total += i
	}
	return total
}
