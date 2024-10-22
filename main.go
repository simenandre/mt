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
	"github.com/spf13/viper"
)

func initConfig() error {
	// Get OS-specific config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("error getting config directory: %w", err)
	}

	// Create mt config directory if it doesn't exist
	mtConfigDir := filepath.Join(configDir, "mt")
	if err := os.MkdirAll(mtConfigDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Set up Viper defaults
	viper.SetDefault("directory", ".")

	// Set up config file details
	viper.SetConfigName("mt-config") // name of config file (without extension)
	viper.SetConfigType("yaml")      // YAML format
	viper.AddConfigPath(mtConfigDir) // add mt config directory as search path

	// Read config file - ignore error if config doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error occurred
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Bind CLI flags
	pflag.StringP("dir", "d", "", "directory to scan (overrides config)")
	pflag.BoolP("json", "j", false, "output as json")
	pflag.BoolP("all", "a", false, "output all tasks")
	pflag.Parse()

	// Bind flags to viper
	viper.BindPFlag("directory", pflag.Lookup("dir"))

	// Handle positional argument for backward compatibility
	if args := pflag.Args(); len(args) > 0 {
		viper.Set("directory", args[0])
	}

	return nil
}

func main() {
	if err := initConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing config: %v\n", err)
		os.Exit(1)
	}

	tasks, err := getTasks(viper.GetString("directory"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if !viper.GetBool("all") {
		tasks = task.FilterAndSortTasks(tasks)
	}

	if viper.GetBool("json") {
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(tasks)
	} else {
		for _, t := range tasks {
			fmt.Printf("- %s\n", t.Title)
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
