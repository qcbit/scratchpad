// Creating files and directories
package main

import "os"

func main() {
	f, err := os.Create("created.file")
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.OpenFile("created.byopen", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f.Close()

	if err = os.Mkdir("createdDir", 0777); err != nil {
		panic(err)
	}

	if err = os.MkdirAll("sampleDir/path1/path2", 0777); err != nil {
		panic(err)
	}
}
