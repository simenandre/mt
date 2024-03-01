package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/simenandre/mt/task"
	"github.com/spf13/pflag"
)

func main() {
	// Set the directory to scan; using current directory as default.
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	tasks, err := getTasks(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	jsonOut := pflag.BoolP("json", "j", false, "output as json")
	all := pflag.BoolP("all", "a", false, "output all tasks")
	pflag.Parse()

	if !*all {
		tasks = task.FilterAndSortTasks(tasks)
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(tasks)
	} else {
		for _, t := range tasks {
			fmt.Println(t.Title)
		}
	}
}

func getFileList(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getTasks(dir string) ([]task.Task, error) {
	files, err := getFileList(dir)
	if err != nil {
		return nil, err
	}

	var tasks []task.Task
	for _, file := range files {
		t, err := extractAndParseTasks(file)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *t...)
	}
	return tasks, nil
}

func extractAndParseTasks(filePath string) (*[]task.Task, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// make an array of tasks
	var tasks []task.Task
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "- [ ]") {

			t, err := task.ParseTaskLine(line)
			if err != nil {
				return nil, err
			}

			tasks = append(tasks, t)
		}
	}
	return &tasks, nil
}
