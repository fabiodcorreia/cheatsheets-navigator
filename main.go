package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fabiodcorreia/cheatsheets-navigator/internal/csn"
)

// TODO: Add ENV description to the help page.
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

	if !*listFlag && len(args) == 0 {
		fmt.Print(csn.UsageMessage)
		return
	}

	repo, _ := os.LookupEnv("CSN_PAGES")

	if *listFlag {
		csn.ShowPages(repo)
		return
	}
	// If args > 0 csn.ScanForPages and store it
	// Find the page from the input, if not found return error
	// With that page found call ReadPage or FilterPage

	pages, err := csn.ScanForPages(repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	page, err := findPageByName(args[0], pages)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(args) == 1 {

		r, err := csn.ReadPage(page)
		if err != nil {
			fmt.Println(err)
			return
		}
		// TODO: Improve this checking if bat is installed
		// writeLess(page, "bat -l markdown --color always")

		if len(args) > 1 {
			_, err := csn.FilterPage(args[0], r)
			if err != nil {
				fmt.Println(err)
				return
			}
			// writeLess(strings.Join(content, "\n"), "bat")
		}
	}
}

func findPageByName(pageName string, pages []csn.Page) (csn.Page, error) {
	for i := range pages {
		if pageName == pages[i].Name {
			return pages[i], nil
		}
	}
	return csn.Page{}, fmt.Errorf("page not found: %q", pageName)
}
