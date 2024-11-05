package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if pages[keys[i]] == pages[keys[j]] {
			return keys[i] < keys[j]
		}
		return pages[keys[i]] < pages[keys[j]]
	})

	for _, k := range keys {
		fmt.Printf("Found %v internal links to %s\n\n", pages[k], k)
	}

}
