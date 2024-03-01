package task

import (
	"fmt"
	"regexp"
	"time"
)

type Task struct {
	Title     string     `json:"title"`
	Starts    *time.Time `json:"starts,omitempty"`
	Scheduled *time.Time `json:"scheduled,omitempty"`
	Due       *time.Time `json:"due,omitempty"`
	Priority  *int       `json:"priority,omitempty"`
}

func parseDate(emoji string, s string) (*time.Time, error) {
	re := regexp.MustCompile(emoji + `\s*(\d{4}-\d{2}-\d{2})`)

	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		// matches[1] contains the first captured group, i.e., the date string
		dateStr := matches[1]

		// Parse the extracted date string into a time.Time object
		const layout = "2006-01-02" // Go's reference date format
		parsedDate, err := time.Parse(layout, dateStr)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}

		return &parsedDate, nil

	} else {
		return nil, nil
	}
}

func parseTaskTitle(s string) string {
	datePattern := regexp.MustCompile(`[🛫⏳📅] \d{4}-\d{2}-\d{2}`)
	checklistPattern := regexp.MustCompile(`- \[ \] `)

	s = datePattern.ReplaceAllString(s, "")
	s = checklistPattern.ReplaceAllString(s, "")
	return markdownToText(s)
}

func ParseTaskLine(task string) (Task, error) {
	var t Task

	t.Title = parseTaskTitle(task)
	starts, err := parseDate("🛫", task)
	if err != nil {
		return t, err
	}
	if starts != nil {
		t.Starts = starts
	}

	scheduled, err := parseDate("⏳", task)
	if err != nil {
		return t, err
	}
	if scheduled != nil {
		t.Scheduled = scheduled
	}

	due, err := parseDate("📅", task)
	if err != nil {
		return t, err
	}
	if due != nil {
		t.Due = due
	}

	return t, nil
}
