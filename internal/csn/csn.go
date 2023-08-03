package csn

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const envPages = `CSN_PAGES`

const UsageMessage = `
  cheatsheet navigator (csn) is a CLI to search and navigate cheatsheets.

  Usage:
    csn [flags]
    csn <page> [term]

  Args:
    page                Cheatsheet page name.
    term                Terms to search inside the page. Terms with multiple words should use "".
    

  Flags:
    -ls                 List all the cheatsheet pages available.
    -h                  Display this message.

  Examples:
    csn                 It's the same as csn -ls.
    csn nvim            It will display the full nvim cheatsheet page.
    csn nvim save       It will show only the sections of nvim page with the word "save".
    csn nvim "save as"  Same as previous example but with a term that contains two words.
`

type Page struct {
	fullPath string
	name     string
}

// Commands function will handle the parse of flags and commands and
// call the proper function to be executed

// getRepository will read the Env Var to get the directory to scan for pages.
// If the variable is not set, not valid or not a directly it will return an error.
func getRepository() (string, error) {
	repoPath := os.Getenv(envPages)
	if repoPath == "" {
		return "", fmt.Errorf("environment variable %q not found", envPages)
	}

	repoPath = expandHomeDir(repoPath)

	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return "", fmt.Errorf("repository path %q not valid", repoPath)
	}

	stat, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("repository path %q doesn't exists", absPath)
		}
		return "", fmt.Errorf("error checking repository path %w", err)
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("repository path %q is not a directory", absPath)
	}

	return absPath, nil
}

// expandHomeDir will get a directory path and if it starts with ~/ or $HOME/
// it will replace that with the value from os.UserHomeDir.
func expandHomeDir(dir string) string {
	if strings.HasPrefix(dir, "~/") {
		dirname, _ := os.UserHomeDir()
		return filepath.Join(dirname, dir[2:])
	}

	if strings.HasPrefix(dir, "$HOME/") {
		dirname, _ := os.UserHomeDir()
		return filepath.Join(dirname, dir[6:])
	}
	return dir
}

// Set Bat function will check if bat is installed and if it is returns a string
// with the bat command to print markdown with colors
func genBatCommand() (string, error) {
	path, err := exec.LookPath("bat")
	if err != nil {
		return "", fmt.Errorf("bat cli not found: %w", err)
	}

	return path + " -l markdown --color always", nil
}

// Show Result function will call bat with the result to print it

// List Pages function will search the repository folder for all the markdown
// files and return a slice with all the files
func getPages(repo string) ([]Page, error) {
	pages := make([]Page, 0, 20)

	err := filepath.WalkDir(repo, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		base := filepath.Base(path)

		if strings.HasPrefix(base, ".") {
			return nil
		}

		if !d.IsDir() && filepath.Ext(path) == ".md" {
			pages = append(pages, Page{
				fullPath: path,
				name:     strings.ReplaceAll(strings.TrimPrefix(path, repo+string(os.PathSeparator)), " ", "_"),
			})
		}

		return nil
	})

	return pages, err
}

// showPages will get the repository path, scan for pages and print the result
// one page per line.
func ShowPages() error {
	// get repo path
	repo, err := getRepository()
	if err != nil {
		return err
	}
	// get pages from repo
	pages, err := getPages(repo)
	if err != nil {
		return err
	}
	// print pages as list
	for i := range pages {
		fmt.Printf("%s\n", pages[i].name)
	}
	println("")
	return nil
}

// Read Page will open a file by name and read the content and return an io.Reader
func ReadPage(page Page, filter string) (string, error) {
	_, err := os.ReadFile(page.fullPath)
  if err != nil {
    return "", err
  }
}

// Filter Page will get an io.Reader and parse the content to find the sections





















