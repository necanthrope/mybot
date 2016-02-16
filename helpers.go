package main

import (
	"strings"
	"unicode"
)

func filter(r rune) rune {
	if unicode.IsUpper(r) {
		r = unicode.ToLower(r)
	}
	if r >= 'a' && r <= 'z' {
		return r
	}
	return -1
}

// break down an input string into words, removing punctuation and lowercasing everything
func normalize(input string) []string {
	rval := make([]string, 0)
	for _, v := range strings.Fields(input) {
		v = strings.Map(filter, v)
		if ( v != "" ) {
		    rval = append(rval, v)
		}
	}
	return rval
}

func filterCandidates(normalized []string, candidates [][]string) [][]string {
	results := make([][]string, 0)
	skip := 0
	for i := range normalized {
		if skip > 0 {
			skip--
			continue
		}

		best := 0
		bestIdx := -1
	CandidateLoop:
		for idx, k := range candidates {
			if k[0] == normalized[i] {
				// first word matches, see if this is better
				rest := strings.Fields(k[1])
				if len(rest) >= best && i+len(rest) < len(normalized) {
					// might be better
					for l, m := range rest {
						if m != normalized[i+l+1] {
							continue CandidateLoop
						}
					}
					best = 1 + len(rest)
					bestIdx = idx
				}
			}
		}
		if bestIdx > -1 {

			results = append(results, candidates[bestIdx])
		}
		skip = best - 1
	}
	return results
}
