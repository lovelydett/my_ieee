package crawler_conference

import "testing"

func TestDoCrawler(t *testing.T) {
	// The conference number for RTSS 2022 is 9984704
	res := DoCrawler("9984704")
	names, links := res[0], res[1]
	if len(names) != len(links) {
		panic("len(names) != len(links)")
	}
	if len(names) == 0 {
		panic("No articles found")
	}
}
