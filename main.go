package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(countBalancedPermutations("123"))
}

func rotateString(s string, goal string) bool {
	for i := 1; i < len(s); i++ {
		if strings.Contains(goal, s[:i]) {
			continue
		} else {
			idx := strings.Index(goal, s[:i-1])
			last := idx + len(s[:i-1])
			return goal[last:]+goal[:idx] == s[i-1:]
		}
	}
	return true
}

/*
Q1. Check Balanced String
*/
func isBalanced(num string) bool {
	even := 0
	odd := 0
	for i, elem := range num {
		digit := int(elem - '0') // Convert character to integer
		if i%2 == 0 {
			even += digit
		} else {
			odd += digit
		}
	}
	return even == odd
}

/*
Q2. Find Minimum Time to Reach Last Room I
*/

type PointTime struct {
	x, y, time int
}

type MinHeapTime []PointTime

func (h MinHeapTime) Len() int            { return len(h) }
func (h MinHeapTime) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeapTime) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeapTime) Push(x interface{}) { *h = append(*h, x.(PointTime)) }
func (h *MinHeapTime) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minTimeToReach1(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	minHeap := &MinHeapTime{{0, 0, 0}}
	heap.Init(minHeap)

	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(Point)
		i, j, curTime := cur.x, cur.y, cur.time
		for _, dir := range directions {
			nextI, nextJ := i+dir[0], j+dir[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m {
				nextTime := curTime + 1
				if nextTime <= moveTime[nextI][nextJ] {
					nextTime = moveTime[nextI][nextJ] + 1
				}

				if nextTime < dp[nextI][nextJ] {
					dp[nextI][nextJ] = nextTime
					heap.Push(minHeap, PointTime{nextI, nextJ, nextTime})
				}
			}
		}
	}

	return dp[n-1][m-1]
}

/*
Q3. Find Minimum Time to Reach Last Room II
*/
type Point struct {
	x, y, time, step int
}

type MinHeap []Point

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	minHeap := &MinHeap{{0, 0, 0, 0}}
	heap.Init(minHeap)

	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(Point)
		i, j, curTime, curStep := cur.x, cur.y, cur.time, cur.step

		timeIncrement := 1
		if curStep%2 == 1 {
			timeIncrement = 2
		}

		for _, dir := range directions {
			nextI, nextJ := i+dir[0], j+dir[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m {
				nextTime := max(curTime+timeIncrement, moveTime[nextI][nextJ]+timeIncrement)

				if nextTime < dp[nextI][nextJ] {
					dp[nextI][nextJ] = nextTime
					heap.Push(minHeap, Point{nextI, nextJ, nextTime, curStep + 1})
				}
			}
		}
	}

	return dp[n-1][m-1]
}

/*
Q4. Count Number of Balanced Permutations
*/
func countBalancedPermutations(num string) int {
	digitCount := make(map[rune]int)
	for _, ch := range num {
		digitCount[ch]++
	}

	var digits []rune
	for digit, count := range digitCount {
		for i := 0; i < count; i++ {
			digits = append(digits, digit)
		}
	}

	ans := 0
	n := len(digits)

	var backtrack func(start int, evenSum int, oddSum int)
	backtrack = func(start int, evenSum int, oddSum int) {
		if start == n {
			if evenSum == oddSum {
				ans++
			}
			return
		}

		memo := make(map[rune]bool)
		for i := start; i < n; i++ {
			if memo[digits[i]] {
				continue
			}
			memo[digits[i]] = true

			digits[start], digits[i] = digits[i], digits[start]

			if start%2 == 0 {
				backtrack(start+1, evenSum+int(digits[start]-'0'), oddSum)
			} else {
				backtrack(start+1, evenSum, oddSum+int(digits[start]-'0'))
			}

			digits[start], digits[i] = digits[i], digits[start]
		}
	}

	backtrack(0, 0, 0)
	return ans
}
