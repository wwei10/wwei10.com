package parser_test

import (
	"testing"

	. "github.com/wwei10/wwei10.com/parser"
)

func TestGetPageFromString(t *testing.T) {
	s := `
---
layout: page
title: Title
permalink: "/about/"
date: 2020-08-22 12:00:00
---

Content`
	page := GetPageFromString(s)
	if page.Title != "Title" || page.Date != "2020-08-22 12:00:00" || page.Content != "Content" || page.Permalink != "/about/" {
		t.Errorf("Parse fail: title: %s date: %s content: %s", page.Title, page.Date, page.Content)
	}
}
