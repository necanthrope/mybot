package main

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	res := normalize("A b, c'+^")
	if len(res) != 3 {
		t.Errorf("Expected 3, got %d\n", len(res))
	}
	if res[0] != "a" {
		t.Errorf("Expected a, got %s\n", res[0])
	}
	if res[1] != "b" {
		t.Errorf("Expected b, got %s\n", res[0])
	}
	if res[2] != "c" {
		t.Errorf("Expected b, got %s\n", res[0])
	}
}
func TestFilter(t *testing.T) {
	res := filterCandidates([]string{"a", "b", "c"}, [][]string{{"a", "b", "def"},{"a","","ghi"}})
	if len(res) != 1 {
		t.Errorf("Expected 1, got %d\n", len(res))
	}
	if (res[0][2] != "def") {
		t.Errorf("Expected def, got %s\n", res[0][2])
	}
}
func TestBest(t *testing.T) {
	res := filterCandidates([]string{"a", "b", "c"}, [][]string{{"b", "", "def 1"},{"b","c","def 2"}})
	if len(res) != 1 {
		t.Errorf("Expected 1, got %d: %q\n", len(res), res)
	}
	if (res[0][2] != "def 2") {
		t.Errorf("Expected def 2, got %s\n", res[0][2])
	}
}
func TestBoth(t *testing.T) {
	res := filterCandidates(normalize("A b, c"), [][]string{{"a", "b c", "def 1"},{"a","","def 2"},{"b","","def 3"}})
	if len(res) != 1 {
		t.Errorf("Expected 1, got %d %q\n", len(res), res)
	}
}
func TestEdge(t *testing.T) {
	res := filterCandidates(normalize("A b, c"), [][]string{{"b", "c d", "def 1"},{"c","d","def 2"}})
	if len(res) != 0 {
		t.Errorf("Expected 0, got %d %q\n", len(res), res)
	}
	res = filterCandidates(normalize("A b, c"), [][]string{{"b", "c d", "def 1"},{"c","d","def 2"},{"c","","def 3"}})
	if len(res) != 1 {
		t.Errorf("Expected 1, got %d %q\n", len(res), res)
	}
}
