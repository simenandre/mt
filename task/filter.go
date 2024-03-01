package task

import (
	"sort"
	"time"
)

// filterAndSortTasks filters and sorts tasks based on specified criteria.
func FilterAndSortTasks(tasks []Task) []Task {
	today := time.Now().Truncate(24 * time.Hour) // Consider today's date without the time part

	// Filter tasks
	filtered := make([]Task, 0)
	for _, task := range tasks {
		if (task.Starts != nil && !task.Starts.After(today)) ||
			(task.Scheduled != nil && !task.Scheduled.After(today)) ||
			(task.Starts == nil && task.Scheduled == nil) {
			filtered = append(filtered, task)
		}
	}

	// Sort tasks by priority (considering nil as lowest priority), then by due date (ascending)
	sort.Slice(filtered, func(i, j int) bool {
		// Handle nil priorities by treating them as lower than any other priority
		iPriority := int(^uint(0) >> 1) // Set to max int if nil
		if filtered[i].Priority != nil {
			iPriority = *filtered[i].Priority
		}

		jPriority := int(^uint(0) >> 1) // Set to max int if nil
		if filtered[j].Priority != nil {
			jPriority = *filtered[j].Priority
		}

		if iPriority == jPriority {
			if filtered[i].Due == nil {
				return false // Tasks without a due date come last
			}
			if filtered[j].Due == nil {
				return true // Ensure tasks with a due date are sorted first
			}
			return filtered[i].Due.Before(*filtered[j].Due)
		}
		return iPriority < jPriority
	})

	return filtered
}
