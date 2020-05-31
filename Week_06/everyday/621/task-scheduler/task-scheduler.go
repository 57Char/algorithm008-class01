package task_scheduler

import "sort"

// https://leetcode.com/problems/task-scheduler/

// Runtime: 4 ms, faster than 99.36% of Go online submissions for Task Scheduler.
// Memory Usage: 6 MB, less than 100.00% of Go online submissions for Task Scheduler.
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	task := make([]int, 26)
	for _, t := range tasks {
		task[t-'A']++
	}
	sort.Ints(task)

	max := task[25] - 1
	count := max * n

	for i := 24; i >= 0 && task[i] > 0; i-- {
		count -= min(task[i], max)
	}

	if count > 0 {
		return count + len(tasks)
	}

	return len(tasks)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}