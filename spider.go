package main

import (
	"fmt"
	"net/url"
)

func Crawl(url string) []string {
	var results []string

	results = append(results, fmt.Sprintf("%s/home", url))
	return results
}

func GetAllScripts(url string) []string {
	var results []string
	root := GetRootNode(url)

	scripts := GetNodesByTagName(root, "script")
	for _, s := range scripts {
		r := GetAttrsByName(&s, "src")
		if r != url && r != "" {
			r = FixLink(r, url)
			results = AppendUnique(results, r)
		}
	}
	return results
}
func AppendUnique(slice []string, item string) []string {
	for _, s := range slice {
		if s == item {
			return slice
		}
	}
	return append(slice, item)
}
func FixLink(link string, baseURL string) string {
	u, err := url.Parse(link)
	if err != nil {
		return link
	}

	if u.IsAbs() {
		return link
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return link
	}

	u = base.ResolveReference(u)

	return u.String()
}
