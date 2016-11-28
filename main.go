package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Get config directory
func configDir() (string, error) {
	path := fmt.Sprintf("%s/.config/repos", os.Getenv("HOME"))
	if _, err := os.Stat(path); err != nil {
		return path, os.MkdirAll(path, 0755)
	}
	return path, nil
}

// Repo is a Git repository
type Repo struct {
	Path string // Path to the repo
}

// Pending returns all files in the repo with uncommitted
// or unstaged files
func (r Repo) Pending() []string {
	return lines("git", "-C", r.Path, "status", "--porcelain")
}

// Stashed returns all files in the repo with stashed changes
func (r Repo) Stashed() []string {
	return lines("git", "-C", r.Path, "stash", "list", "--porcelain")
}

// Collections are named groups of repositories
type Collection struct {
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}

// Config is loaded from ~/.config/repos/config.json
type Config struct {
	configDir   string
	BasePath    string        `json:"basePath"`
	Collections []*Collection `json:"collections"`
}

// Load a new Config object from ~/.config/repos/config.json
func NewConfig() (*Config, error) {
	path, err := configDir()
	if err != nil {
		return nil, err
	}
	confPath := fmt.Sprintf("%s/config.json", path)
	if _, err := os.Stat(confPath); err != nil {
		return nil, fmt.Errorf("Need to create %s", confPath)
	}
	raw, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	config := &Config{
		configDir: path,
	}
	if err := json.Unmarshal(raw, config); err != nil {
		return nil, err
	}
	return config, nil
}

// lines collects lines of output from a command
func lines(name string, args ...string) []string {
	lines := []string{}
	raw, err := exec.Command(name, args...).Output()
	if err != nil {
		return lines
	}
	for _, line := range strings.SplitAfter(string(raw), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// Flat walks directories in each collection returning []Repo
func flat(collections []*Collection) []*Repo {
	flat := []*Repo{}
	for _, collection := range collections {
		matches, err := filepath.Glob(collection.Pattern)
		failOnErr(err)
		for _, match := range matches {
			repo := &Repo{
				Path: match,
			}
			flat = append(flat, repo)
		}
	}
	return flat
}

func failOnErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}

func main() {
	cfg, err := NewConfig()
	failOnErr(err)
	failOnErr(os.Chdir(cfg.BasePath))
	pending := flag.Bool("pending", false, "Show repositories with pending changes")
	stashed := flag.Bool("stashed", false, "Show repositories with stashed changes")
	flag.Parse()
	switch {
	case *pending:
		for _, repo := range flat(cfg.Collections) {
			if len(repo.Pending()) > 0 {
				fmt.Println(repo.Path)
			}
		}
	case *stashed:
		for _, repo := range flat(cfg.Collections) {
			if len(repo.Stashed()) > 0 {
				fmt.Println(repo.Path)
			}
		}
	default:
		for _, repo := range flat(cfg.Collections) {
			pending := repo.Pending()
			stashed := repo.Stashed()
			if len(pending) > 0 || len(stashed) > 0 {
				fmt.Println(repo.Path)
			}
		}
	}
}
