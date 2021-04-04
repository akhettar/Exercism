package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency compute the freequency of letters in the given text
func ConcurrentFrequency(texts []string) FreqMap {
	freq := make(FreqMap)
	freqch := make(chan FreqMap, 10)

	for _, text := range texts {
		go func(s string) {
			freqch <- Frequency(s)
		}(text)
	}

	// process frequencies as they come through
	for range texts {
		for k, v := range <-freqch {
			freq[k] += v
		}
	}
	return freq
}
