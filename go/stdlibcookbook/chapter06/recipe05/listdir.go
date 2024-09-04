// List a directory
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("List by ReadDir")
	listDirByReadDir(".")
	fmt.Println()
	fmt.Println("List by Walk")
	listDirByWalk(".")
}

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		// Walk given dir without printing it
		if wPath == path {
			return nil
		}

		// if given path is folder then stop list recursively and print as folder
		if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}

		if wPath != path {
			fmt.Println(wPath)
		}
		return nil
	})
}

func listDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}
