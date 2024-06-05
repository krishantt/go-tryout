package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	sourceFileName := filepath.Base("./hello.go")

	// Open the source file
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	tempDir := os.TempDir()

	destinationFilePath := filepath.Join(tempDir, sourceFileName )

	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file contents:", err)
		return
	}

	fmt.Println("File copied to temporary directory:", destinationFilePath)
}
