package task

import (
	"time"

	taskspec "github.com/taskspec/taskspec-go"
)

type Task struct {
	Title     string     `json:"title"`
	Starts    *time.Time `json:"starts,omitempty"`
	Scheduled *time.Time `json:"scheduled,omitempty"`
	Due       *time.Time `json:"due,omitempty"`
	Priority  *int       `json:"priority,omitempty"`
}

func ParseTaskLine(line string) (Task, error) {
	parser := taskspec.NewParser()
	taskspecTask, err := parser.Parse(line)
	if err != nil {
		return Task{}, err
	}

	// If parser returned nil (not a valid task), return empty task
	if taskspecTask == nil {
		return Task{}, nil
	}

	var t Task
	t.Title = taskspecTask.Description

	if taskspecTask.StartDate != nil {
		t.Starts = taskspecTask.StartDate
	}

	if taskspecTask.ScheduledDate != nil {
		t.Scheduled = taskspecTask.ScheduledDate
	}

	if taskspecTask.DueDate != nil {
		t.Due = taskspecTask.DueDate
	}

	// Convert taskspec Priority to int pointer if set
	if taskspecTask.Priority != taskspec.PriorityUnknown {
		priority := int(taskspecTask.Priority)
		t.Priority = &priority
	}

	return t, nil
}
