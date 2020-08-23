package parser

/**
  Parse jekyll style markdown post into Page struct
*/

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Page represents a page in a website
type Page struct {
	Title     string
	Date      string
	Content   string
	Permalink string
}

func parseHeader(s string) map[string]string {
	var m = make(map[string]string)
	slices := strings.Split(s, "\n")
	for i := 0; i < len(slices); i++ {
		values := strings.Split(slices[i], ":")
		if len(values) >= 2 {
			m[values[0]] = strings.Trim(strings.Join(values[1:], ":"), " \"")
		}
	}
	return m
}

// GetPagesFromDir returns an array of Pages created from a directory
func GetPagesFromDir(dirname string) []Page {
	mds, err := filepath.Glob(dirname + "/*")
	if err != nil {
		panic(err)
	}
	var pages []Page
	for _, md := range mds {
		data, err := ioutil.ReadFile(md)
		if err != nil {
			fmt.Printf("Error processing %s\n", md)
			continue
		}
		page := GetPageFromString(string(data))
		pages = append(pages, page)
	}
	return pages
}

// GetPagesMapFromDir generate Pages from a directory name
// Keyed by permalink
func GetPagesMapFromDir(dirname string) map[string]Page {
	mds, err := filepath.Glob(dirname + "/*")
	if err != nil {
		panic(err)
	}
	var pages = make(map[string]Page)
	for _, md := range mds {
		data, err := ioutil.ReadFile(md)
		if err != nil {
			fmt.Printf("Error processing %s\n", md)
			continue
		}
		page := GetPageFromString(string(data))
		pages[page.Permalink] = page
	}
	return pages
}

// GetPageFromString returns Page struct from a string
func GetPageFromString(s string) Page {
	slices := strings.Split(s, "---")
	if len(slices) != 3 {
		return Page{Title: "", Date: "", Content: ""}
	}
	header := slices[1]
	m := parseHeader(header)
	var title = ""
	if t, found := m["title"]; found {
		title = t
	}
	var date = ""
	if d, found := m["date"]; found {
		date = d
	}
	var permalink = ""
	if p, found := m["permalink"]; found {
		permalink = p
	}
	body := strings.Trim(slices[2], " \n\r")
	return Page{Title: title, Date: date, Content: body, Permalink: permalink}
}
