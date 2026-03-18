package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
	"syscall/js"
)

var DEBUG = false

// Map to act as our O(1) exclusion set
var exclude = map[string]bool{
	"": true, "am": true, "what": true, "like": true, "me": true, "through": true, "when": true,
	"where": true, "my": true, "a": true, "i": true, "an": true, "and": true, "are": true,
	"as": true, "at": true, "be": true, "but": true, "by": true, "for": true, "if": true,
	"in": true, "into": true, "is": true, "it": true, "no": true, "not": true, "of": true,
	"on": true, "or": true, "such": true, "that": true, "the": true, "their": true, "then": true,
	"there": true, "these": true, "they": true, "this": true, "to": true, "was": true, "will": true,
	"with": true, "from": true, "you": true, "your": true, "we": true, "he": true, "she": true,
	"his": true, "her": true, "them": true, "our": true, "us": true, "so": true, "too": true,
	"have": true, "has": true, "had": true, "it’s": true, "its": true, "some":true, "can":true, 
	"those":true, "about":true,
}

// getEligibleWords is the function we expose to JavaScript
func getEligibleWords(this js.Value, args []js.Value) any {
	poemsContent := args[0].String()
	// convert input to lowercase and split into words (strings.Fields splits by whitespace)
	words := strings.Fields(strings.ToLower(poemsContent))

	// filter excludable words and count frequencies of everything else
	counter := make(map[string]uint8)
	for _, w := range words {
		if !exclude[w] {
			counter[w]++
		}
	}

	if DEBUG {
		// print sorted dictionary equivalent (fmt.Println routes to console.log in WASM)
		type kv struct {
			Key   string
			Value uint8
		}
		var sortedKV []kv
		for k, v := range counter {
			sortedKV = append(sortedKV, kv{k, v})
		}
		sort.Slice(sortedKV, func(i, j int) bool {
			return sortedKV[i].Value > sortedKV[j].Value
		})
		fmt.Printf("Word counts: %v\n", sortedKV)
	}
	
	// extract counts > 1 and sort descending
	countFreqs := make(map[uint8]bool)
	for _, v := range counter {
		if v > 1 {
			countFreqs[v] = true
		}
	}
	counts := make([]uint8, 0, len(countFreqs))
	for k := range countFreqs {
		counts = append(counts, uint8(k))
	}
	//To reduce binary size, I'm avoiding using the sort lib
	//If I ever need to return to it, here it is:
	//sort.Slice(counts, func(i, j int) bool {
	//	return counts[i] > counts[j]
	//})
	for i := 0; i < len(counts); i++ {
		for j := 0; j < len(counts)-i-1; j++ {
			if counts[j] < counts[j+1] {
				counts[j], counts[j+1] = counts[j+1], counts[j]
			}
		}
	}
	// We only accept the top 4 frequencies:
	var minimumFrequency uint8
	if 4 > len(counts) {
		minimumFrequency = counts[len(counts) - 1]
	} else {
		minimumFrequency = counts[3]
	}
	if DEBUG {
		fmt.Print("Words of this frequency or higher are eligible: ")
		fmt.Println(minimumFrequency)
	}
	
	// collect eligible words
	var eligibleWords []string
	for word, count := range counter {
		if count >= minimumFrequency {
			eligibleWords = append(eligibleWords, word)
		}
	}
	if DEBUG {fmt.Printf("Eligible words: %v\n", eligibleWords)}

	// Handle case where no words matched to avoid panic
	if len(eligibleWords) == 0 {
		return []any{}
	}

	// shuffle the eligible words
	rand.Shuffle(len(eligibleWords), func(i, j int) {
		eligibleWords[i], eligibleWords[j] = eligibleWords[j], eligibleWords[i]
	})

	// pick a random number between 1 and 3 to display
	displayNWords := rand.IntN(3) + 2
	if displayNWords > len(eligibleWords) {
		displayNWords = len(eligibleWords)
	}
	
	selectedWords := eligibleWords[:displayNWords]
	
	// convert the Go slice of strings into a format JS can accept ([]any)
	jsArray := make([]any, len(selectedWords))
	for i, v := range selectedWords {
		jsArray[i] = v
	}
	
	return jsArray
}

func main() {
	// Bind the Go function to the JavaScript global scope so it can be called
	js.Global().Set("getEligibleWords", js.FuncOf(getEligibleWords))
	
	// Create an empty channel and block on it. 
	// This prevents the Go WASM program from exiting immediately.
	c := make(chan struct{})
	<-c
}