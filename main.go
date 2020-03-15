package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func resolveBasePath() string {
	basePath := os.Getenv("REPOS_PATH")
	if basePath == "" {
		return "."
	}
	return basePath
}

func main() {
	var (
		basePath string
	)
	flag.StringVar(&basePath, "path", resolveBasePath(), "base path to traverse for git repos")
	flag.Parse()
	maybe(filepath.Walk(basePath, func(pathName string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			stat, err := os.Stat(path.Join(pathName, "/.git"))
			if err == nil && stat.IsDir() {
				// this contains a git directory
				fmt.Println(pathName)
				return filepath.SkipDir
			}
		}

		return nil
	}))
}
