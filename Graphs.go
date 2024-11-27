package main

import (
	"container/heap"
	"math"
)

/*
1905. Count Sub Islands
You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water) and 1's (representing land).
An island is a group of 1's connected 4-directionally (horizontal or vertical). Any cells outside of the grid are
considered water cells.

An island in grid2 is considered a sub-island if there is an island in grid1 that contains all the cells that make up this island in grid2.
Return the number of islands in grid2 that are considered sub-islands.
*/
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	return floodfillSub(grid2, grid1)
}

func floodfillSub(grid [][]int, grid2 [][]int) int {
	var visited [][2]int
	m := len(grid)
	n := len(grid[0])
	islands := 0
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if notContain(visited, i, j) && grid[i][j] == 1 {
				visited = helperFloodFill(grid, i, j, visited, m, n)
				subIsland := 0
				groundTrue := len(visited) - index
				for i := index; i < len(visited); i++ {
					if grid2[visited[i][0]][visited[i][1]] == 1 {
						subIsland++
					}
					index++
				}
				if subIsland == groundTrue {
					islands++
				}
			}
		}
	}
	return islands
}

func helperFloodFill(grid [][]int, x int, y int, visited [][2]int, m int, n int) [][2]int {
	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	stack := Stack[[2]int]()
	stack.Push([2]int{x, y})
	newVisited := visited
	for stack.Length() > 0 {
		position := stack.Pop()
		newVisited = append(newVisited, position)
		for _, dir := range directions {
			m1, n1 := dir[0]+position[0], dir[1]+position[1]
			if m1 > -1 && m1 < m && n1 > -1 && n1 < n {
				if notContain(newVisited, m1, n1) && grid[m1][n1] == 1 {
					stack.Push([2]int{m1, n1})
				}
			}
		}
	}
	return newVisited
}

func notContain(arr [][2]int, x, y int) bool {
	for _, pair := range arr {
		if pair[0] == x && pair[1] == y {
			return false
		}
	}
	return true
}

/*
*
874. Walking Robot Simulation
A robot on an infinite XY-plane starts at point (0, 0) facing north. The robot can receive a sequence of these three
possible types of commands:

	-2: Turn left 90 degrees.
	-1: Turn right 90 degrees.
	1 <= k <= 9: Move forward k units, one unit at a time.

Some of the grid squares are obstacles. The ith obstacle is at grid point obstacles[i] = (xi, yi).
If the robot runs into an obstacle, then it will instead stay in its current location and move on to the next command.
Return the maximum Euclidean distance that the robot ever gets from the origin squared.
*/
func robotSim(commands []int, obstacles [][]int) int {
	direction := [2]int{0, 1}
	position := [2]int{0, 0}
	maxEuclidean := 0

	obstacleMap := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		obstacleMap[[2]int{obstacle[0], obstacle[1]}] = true
	}

	for _, command := range commands {
		if command == -2 {
			// Turn left: [0,1] -> [-1,0] -> [0,-1] -> [1,0] -> [0,1]
			direction = [2]int{-direction[1], direction[0]}
		} else if command == -1 {
			// Turn right: [0,1] -> [1,0] -> [0,-1] -> [-1,0] -> [0,1]
			direction = [2]int{direction[1], -direction[0]}
		} else {
			for step := 0; step < command; step++ {
				nextPosition := [2]int{position[0] + direction[0], position[1] + direction[1]}
				if obstacleMap[nextPosition] {
					break
				}
				position = nextPosition // update position
				// Update the maximum Euclidean distance squared
				distanceSquared := position[0]*position[0] + position[1]*position[1]
				if distanceSquared > maxEuclidean {
					maxEuclidean = distanceSquared
				}
			}
		}
	}
	return maxEuclidean
}

/*
3217. Delete Nodes From Linked List Present in Array
You are given an array of integers nums and the head of a linked list. Return the head of the modified linked list after
removing all nodes from the linked list that have a value that exists in nums.
*/
func modifiedList(nums []int, head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	var helper func(*ListNode) *ListNode
	helper = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}
		if numSet[node.Val] {
			return helper(node.Next)
		}
		node.Next = helper(node.Next)
		return node
	}
	return helper(head)
}

/*
3243. Shortest Distance After Road Addition Queries I
https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/description/
*/
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	// Adjacency list to store the graph
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}
	// Add default edges
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

	// Dijkstra's algorithm implementation
	dijkstra := func() int {
		pq := &PriorityQueue{}
		heap.Init(pq)
		heap.Push(pq, &Item{node: 0, distance: 0})

		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0

		for pq.Len() > 0 {
			current := heap.Pop(pq).(*Item)
			currentNode, currentDist := current.node, current.distance

			if currentNode == n-1 {
				return currentDist
			}

			for _, neighbor := range graph[currentNode] {
				newDist := currentDist + 1
				if newDist < dist[neighbor] {
					dist[neighbor] = newDist
					heap.Push(pq, &Item{node: neighbor, distance: newDist})
				}
			}
		}
		return -1 // Return -1 if the target is unreachable
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		graph[query[0]] = append(graph[query[0]], query[1])
		ans[i] = dijkstra()
	}
	return ans
}

type Item struct {
	node     int
	distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
