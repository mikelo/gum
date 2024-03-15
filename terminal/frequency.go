package terminal

import (
	"sort"
)

// SortByFrequency sorts a slice of strings by their frequency and returns the sorted strings.
func SortByFrequency(strings []string) []string {
	// Count the frequency of each string
	freqMap := make(map[string]int)
	for _, s := range strings {
		freqMap[s]++
	}

	// Convert the frequency map to pairs
	type Pair struct {
		Str  string
		Freq int
	}
	pairs := make([]Pair, 0, len(freqMap))
	for str, freq := range freqMap {
		pairs = append(pairs, Pair{Str: str, Freq: freq})
	}

	// Define custom sort for pairs based on frequency
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Freq > pairs[j].Freq
	})

	// Extract the sorted strings from pairs
	sortedStrings := make([]string, len(pairs))
	for i, pair := range pairs {
		sortedStrings[i] = pair.Str
	}

	return sortedStrings
}
