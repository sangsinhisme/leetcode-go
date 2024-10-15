package main

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
650. 2 Keys Keyboard
There is only one character 'A' on the screen of a notepad. You can perform one of two operations on this notepad for each step:

Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
Paste: You can paste the characters which are copied last time.
Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.
. . .
*/

func minSteps(n int) int {
	var factors = 0
	for n%2 == 0 {
		n = n / 2
		factors += 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			n = n / i
			factors += i
		}
	}

	if n > 2 {
		factors += n
	}
	return factors
}

func getSneakyNumbers(nums []int) []int {
	freqMap := frequency(nums)
	var result []int

	// Loop through the frequency map and select numbers that have a frequency of 2
	for num, count := range freqMap {
		if count == 2 {
			result = append(result, num)
		}
	}

	return result
}

func frequency(nums []int) map[int]int {
	freqMap := make(map[int]int)

	// Loop through each number in the slice and update its frequency count
	for _, num := range nums {
		freqMap[num]++
	}

	return freqMap
}

func findMinDifference(timePoints []string) int {
	var helper func(time string) int
	helper = func(time string) int {
		parser := strings.Split(time, ":")
		hour, _ := strconv.Atoi(parser[0])
		minute, _ := strconv.Atoi(parser[1])
		return hour*60 + minute
	}
	n := len(timePoints)

	times := make([]int, n)
	for i, time := range timePoints {
		times[i] = helper(time)
	}
	sort.Ints(times)
	minDiff := 1440 + times[0] - times[n-1]
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, times[i]-times[i-1])
	}
	return minDiff
}

/*
386. Lexicographical Numbers
Given an integer n, return all the numbers in the range [1, n] sorted in lexicographical order.
You must write an algorithm that runs in O(n) time and uses O(1) extra space.
*/
func lexicalOrder(n int) []int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	sort.Slice(nums, func(i, j int) bool {
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])
		if num1[0] == num2[0] {
			return num1 < num2
		}
		concat1, _ := strconv.Atoi(num1 + num2)
		concat2, _ := strconv.Atoi(num2 + num1)
		return concat1 < concat2
	})
	return nums
}

func findKthNumber(n int, k int) int {
	arr := lexicalOrder(n)
	return arr[k]
}

/*
1497. Check If Array Pairs Are Divisible by k
Given an array of integers arr of even length n and an integer k.
We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.
Return true If you can find a way to do that or false otherwise
*/
func canArrange(arr []int, k int) bool {
	frequency := make(map[int]int)
	for _, num := range arr {
		divisor := ((num % k) + k) % k
		frequency[divisor]++
	}
	for key := range frequency {
		left := k - key
		if key == 0 {
			if frequency[key]%2 != 0 {
				return false
			}
		} else if frequency[key] != frequency[left] {
			return false
		}
	}
	return true
}

/*
1331. Rank Transform of an Array
https://leetcode.com/problems/rank-transform-of-an-array/description/
*/
func arrayRankTransform(arr []int) []int {
	temp := make([]int, len(arr)) // Create a new slice with the same length as arr
	copy(temp, arr)
	indexes := map[int]int{}
	sort.Ints(temp)
	idx := 1
	for _, num := range temp {
		if indexes[num] == 0 {
			indexes[num] = idx
			idx++
		}
	}
	for i, num := range arr {
		arr[i] = indexes[num]
	}
	return arr
}

/*
2491. Divide Players Into Teams of Equal Skill
https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/description/
*/
func dividePlayers(skill []int) int64 {
	total := sum(skill)
	n := len(skill)
	pair := total / (n / 2)
	freq := make(map[int]int)
	for _, i := range skill {
		freq[i]++
	}
	var ans int64 = 0
	for key, value := range freq {
		left := pair - key
		if freq[left] != value {
			return -1
		}
		if left == key {
			if value%2 != 0 {
				return -1
			}
			ans += int64(key * (value / 2) * left)
		} else {
			ans += int64(key * value * left)
		}
		delete(freq, key)
		delete(freq, left)
	}
	return ans
}

/*
962. Maximum Width Ramp
https://leetcode.com/problems/maximum-width-ramp/description/
*/
func maxWidthRamp(nums []int) int {
	n := len(nums)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		if nums[indices[i]] == nums[indices[j]] {
			return indices[i] < indices[j]
		}
		return nums[indices[i]] < nums[indices[j]]
	})

	maxWidth := 0
	minIndex := indices[0]

	for _, i := range indices {
		if i > minIndex {
			maxWidth = max(maxWidth, i-minIndex)
		}
		minIndex = min(minIndex, i)
	}
	return maxWidth
}

/*
2406. Divide Intervals Into Minimum Number of Groups
https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/description/
*/
func minGroups(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for _, interval := range intervals {
		if minHeap.Len() > 0 && (*minHeap)[0] < interval[0] {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, interval[1])
	}

	return minHeap.Len()
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
632. Smallest Range Covering Elements from K Lists
https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/description
*/
func smallestRange(nums [][]int) []int {
	k := len(nums)
	indices := make([]int, k)
	for i := range indices {
		indices[i] = 0
	}
	rangeList := make([]int, 2)
	rangeList[0], rangeList[1] = 0, math.MaxInt

	for {
		curMin, curMax := math.MaxInt, math.MinInt
		minListIdx := 0
		for i := range k {
			curElem := nums[i][indices[i]]
			if curElem < curMin {
				curMin = curElem
				minListIdx = i
			}
			if curElem > curMax {
				curMax = curElem
			}
		}
		if curMax-curMin < rangeList[1]-rangeList[0] {
			rangeList[1] = curMax
			rangeList[0] = curMin
		}
		indices[minListIdx] += 1
		if indices[minListIdx] == len(nums[minListIdx]) {
			break
		}
	}
	return rangeList
}

/*
2530. Maximal Score After Applying K Operations
https://leetcode.com/problems/maximal-score-after-applying-k-operations/description
*/
func maxKelements(nums []int, k int) int64 {
	pq := IntHeap{}
	heap.Init(&pq)

	for _, i := range nums {
		heap.Push(&pq, i)
	}

	total := int64(0)
	for i := 0; i < k; i++ {
		if pq.Len() > 0 {
			elem := heap.Pop(&pq).(int)
			total += int64(elem)
			newElem := int(math.Ceil(float64(elem) / 3))
			heap.Push(&pq, newElem)
		}
	}

	return total
}

/*
2938. Separate Black and White Balls
https://leetcode.com/problems/separate-black-and-white-balls/description/
*/
func minimumSteps(s string) int64 {
	n := len(s)
	l, r := 0, n-1
	var swap int64
	for l < r {
		if s[l] == '0' {
			l++
			continue
		}
		if s[r] == '1' {
			r--
			continue
		}
		swap += int64(r - l)
		r--
		l++
	}
	return swap
}
