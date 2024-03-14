package terminal

import (
	"sort"
)

// Pair represents a string-frequency pair.
type Pair struct {
	Str  string
	Freq int
}

// Pairs is a slice of Pair.
type Pairs []Pair

// Len returns the length of the Pairs slice.
func (p Pairs) Len() int { return len(p) }

// Swap swaps the elements at indices i and j in the Pairs slice.
func (p Pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Less returns true if the frequency of the element at index i is greater than the frequency of the element at index j in the Pairs slice.
func (p Pairs) Less(i, j int) bool { return p[i].Freq > p[j].Freq }

// SortByFrequency sorts the strings by their frequency in descending order.
func SortByFrequency(strings []string) []Pair {
	freqMap := make(map[string]int)
	for _, s := range strings {
		freqMap[s]++
	}

	pairs := make(Pairs, 0, len(freqMap))
	for str, freq := range freqMap {
		pairs = append(pairs, Pair{Str: str, Freq: freq})
	}

	sort.Sort(pairs)
	return pairs
}
