package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give all the required arguments!")
		os.Exit(1)
	}

	targetURL := os.Args[1]
	crawlAndScan(targetURL)
}

func crawlAndScan(url string) {
	visited := make(map[string]bool)
	queue := []string{url}

	for len(queue) > 0 {
		currentURL := queue[0]
		queue = queue[1:]

		if visited[currentURL] {
			continue
		}

		visited[currentURL] = true

		fmt.Println("Scanning:", currentURL)
		analyzeResponse(currentURL)

		links := discoverLinks(currentURL)
		for _, link := range links {
			if !visited[link] {
				queue = append(queue, link)
			}
		}
	}
}

func analyzeResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	body := resp.Body
	// Analyze the response body for SQL Injection vulnerabilities
	if containsString(body, "' OR '1'='1") {
		reportVulnerability(url, "Potential SQL Injection vulnerability")
	}
}

func discoverLinks(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL:", err)
		return nil
	}
	defer resp.Body.Close()

	// Implement link discovery logic here
	return nil
}

func reportVulnerability(url, vulnerability string) {
	fmt.Printf("Vulnerability detected at %s: %s\n", url, vulnerability)
	// You can log the vulnerability or save it to a report file
}

func containsString(data io.Reader, target string) bool {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, data)
	if err != nil {
		log.Println("Error reading response body:", err)
		return false
	}

	return strings.Contains(buf.String(), target)
}
