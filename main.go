package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fabiodcorreia/cheatsheets-navigator/internal/csn"
)

func main() {
	listFlag := flag.Bool("ls", false, "List all teh cheatsheet pages available.")
	helpFlag := flag.Bool("h", false, "Display this message.")

	flag.Parse()

	flag.Usage = func() {
		fmt.Print(csn.UsageMessage)
	}

	if *helpFlag {
		flag.Usage()
		return
	}

	args := flag.Args()

	if *listFlag || len(args) == 0 {
		// fmt.Println(listPages())
		csn.ShowPages()
		return
	}

	if len(args) == 1 {
		_, err := readPage(args[0])




		if err != nil {
			fmt.Println(err)
			return
		}
		// TODO: Improve this checking if bat is installed
		// writeLess(page, "bat -l markdown --color always")
	}

	if len(args) > 1 {
		_, err := readAndFilterMarkdownFile(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		// writeLess(strings.Join(content, "\n"), "bat")
	}
}

func readPage(pageName string) (string, error) {
	pagePath := ""
	pagePath = pagePath + "/" + pageName + ".md"

	content, err := os.ReadFile(pagePath)
	if err != nil {
		return "", fmt.Errorf("fail to read the repository page: %w", err)
	}

	return string(content), nil
}

func readAndFilterMarkdownFile(pageName, filterWord string) ([]string, error) {
	var filteredContent []string
	var section string

	pagePath := ""

	pagePath = pagePath + "/" + pageName + ".md"

	content, err := os.ReadFile(pagePath)
	if err != nil {
		return filteredContent, fmt.Errorf("fail to read the repository page: %w", err)
	}
	// BUG: For some reason it not getting all the matches on NvChad Keys and Commands if I search by C- for example
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "##") || strings.HasPrefix(line, "###") {
			if strings.Contains(section, filterWord) {
				sectionHl := strings.ReplaceAll(section, filterWord, "\033[1;31m"+filterWord+"\033[0m")
				filteredContent = append(filteredContent, sectionHl)
			}
			section = ""
		}

		section += line + "\n"
	}

	return filteredContent, nil
}
