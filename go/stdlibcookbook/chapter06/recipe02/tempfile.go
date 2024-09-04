// Creating temporary files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	tFile, err := ioutil.TempFile("", "gostdcookbook")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tFile.Name())

	fmt.Println(tFile.Name())
	tDir, err := ioutil.TempDir("", "gostdcookbookdir")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tDir)
	fmt.Println(tDir)
}
