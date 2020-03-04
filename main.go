package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ParseRepository returns a new repository
func ParseRepository(input string) (string, error) {
	if strings.HasPrefix(input, "http") {
		return parseHTTP(input), nil
	}
	if strings.HasSuffix(input, ".git") {
		return parseGit(input), nil
	}
	return "", fmt.Errorf("Invalid Path")
}

var re = regexp.MustCompile(`^[a-z]+@([a-z\.]+):([a-z\./]+)`)

func parseGit(input string) string {
	input = strings.TrimRight(input, ".git")
	matches := re.FindStringSubmatch(input)
	host := matches[1]
	return filepath.Join(host, matches[2])
}
func parseHTTP(input string) string {
	input = strings.TrimRight(input, ".git")
	return input[8:]
}

var srcPath = flag.String("src", filepath.Join(os.Getenv("HOME"), "src"), "source directory")

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		os.Exit(0)
	}

	repoPath, err := ParseRepository(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(filepath.Join(*srcPath, repoPath))
}
