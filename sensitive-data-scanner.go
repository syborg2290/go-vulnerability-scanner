package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give the directory path as an argument!")
		os.Exit(1)
	}

	directoryPath := os.Args[1]
	scanForSensitiveData(directoryPath)
}

func scanForSensitiveData(directoryPath string) {
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error accessing path:", err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(path, ".js") {
			checkFileForSensitiveData(path)
		}
		return nil
	})

	if err != nil {
		log.Println("Error scanning directory:", err)
	}
}

func checkFileForSensitiveData(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	sensitiveKeywords := []string{"password", "api_key", "secret_key", "access_token"}
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(string(content), keyword) {
			reportVulnerability(filePath, "Sensitive data exposure")
			return
		}
	}
}

func reportVulnerability(filePath, vulnerability string) {
	fmt.Printf("Vulnerability detected in %s: %s\n", filePath, vulnerability)
	// You can log the vulnerability or save it to a report file
}
