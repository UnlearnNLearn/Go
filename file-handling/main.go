package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// File Handling
func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println("file contains ", str)

	bs, err = ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	fmt.Println(string(bs))

	file1, err := os.Create("go.txt")
	if err != nil {
		return
	}

	defer file1.Close()

	file1.WriteString("this is writing by a process")

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // instruct it to read all of the files
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
