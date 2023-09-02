package main

import (
	"fmt"
	"os"
)

func createFiles(basePath string, files ...string) {
	for _, name := range files {
		filePath := fmt.Sprintf("%s/%s.%s", basePath, name, "txt")
		f, err := os.Create(filePath)
		//call f.Close() before the function ends to avoid memory leeks
		defer f.Close()
		if err != nil {
			fmt.Printf("Error while creating file %s:%v\n", filePath, err)
			os.Exit(1)
		}
		fmt.Printf("file %s sucessfuly created", f.Name())
	}
	fmt.Println("Files successfuly created!")
}

func main() {
	tmp := os.TempDir()
	createFiles(tmp, "arquivo1", "novo", "teste")
}