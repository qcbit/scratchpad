// Write a Go problem that reads files concurrently and counts the lines in each file.
// Use goroutines and channels.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

// FileLineCounter defines the interface for counting lines in a file.
type FileLineCounter interface {
	CountLines(filePath string) (int, error)
}

// FileLineCounterFunc is a function type that implements the FileLineCounter interface.
// This allows using ordinary functions as line counters.
type FileLineCounterFunc func(filePath string) (int, error)

// CountLines calls the underlying function.
func (f FileLineCounterFunc) CountLines(filePath string) (int, error) {
	return f(filePath)
}

// defaultFileLineCounter is a concrete implementation of FileLineCounter.
type defaultFileLineCounter struct{}

// CountLines opens a file and counts the number of lines.
func (c *defaultFileLineCounter) CountLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}

// CountLinesInFiles concurrently counts lines in a list of files.
// It returns a map of file paths to their line counts for successfully processed files,
// and a slice of errors for any files that failed.
func CountLinesInFiles(filePaths []string, counter FileLineCounter) (map[string]int, []error) {
	// A struct to hold the result of a single file processing.
	type result struct {
		filePath  string
		lineCount int
		err       error
	}

	// Use a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup
	// Channel to receive results from goroutines.
	resultCh := make(chan result, len(filePaths))

	for _, filePath := range filePaths {
		wg.Add(1)
		go func(fp string) {
			defer wg.Done()
			lineCount, err := counter.CountLines(fp)
			resultCh <- result{filePath: fp, lineCount: lineCount, err: err}
		}(filePath)
	}

	// Wait for all goroutines to complete, then close the channel.
	wg.Wait()
	close(resultCh)

	// Collect results and errors from the channel.
	results := make(map[string]int)
	var errs []error
	for res := range resultCh {
		if res.err != nil {
			errs = append(errs, fmt.Errorf("error processing %s: %w", res.filePath, res.err))
		} else {
			results[res.filePath] = res.lineCount
		}
	}

	if len(errs) > 0 {
		return results, errs
	}

	return results, nil
}

func main() {
	// Create some temporary files for demonstration.
	files := []struct {
		name    string
		content string
	}{
		{"file1.txt", "one\ntwo\nthree"},
		{"file2.txt", "four\nfive"},
		{"file3.txt", "six\nseven\neight\nnine"},
	}

	var filePaths []string
	tempDir, err := os.MkdirTemp("", "countinglines-example")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // clean up the whole directory

	for _, f := range files {
		tmpfile, err := os.CreateTemp(tempDir, "testfile-*.txt")
		if err != nil {
			log.Fatal(err)
		}

		if _, err := tmpfile.WriteString(f.content); err != nil {
			log.Fatal(err)
		}
		filePaths = append(filePaths, tmpfile.Name())
		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	}

	// Add a non-existent file to demonstrate error handling.
	filePaths = append(filePaths, "non-existent-file.txt")

	// Create a counter and run the line counting.
	counter := &defaultFileLineCounter{}
	results, errs := CountLinesInFiles(filePaths, counter)

	fmt.Println("--- Results ---")
	for path, count := range results {
		fmt.Printf("%s: %d lines\n", path, count)
	}

	if len(errs) > 0 {
		fmt.Println("\n--- Errors ---")
		for _, err := range errs {
			fmt.Println(err)
		}
	}
}