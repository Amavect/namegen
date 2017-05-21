package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// A list of consonants.
	consonant []string = []string{
		"b",
		"c",
		"d",
		"f",
		"g",
		"h",
		"j",
		"k",
		"l",
		"m",
		"n",
		"p",
		"q",
		"r",
		"s",
		"t",
		"v",
		"w",
		"x",
		"y",
		"z",
	}
	// A list of vowels.
	vowel []string = []string{
		"a",
		"e",
		"i",
		"o",
		"u",
	}
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond())) // Seed randomness based on nanoseconds.

	var length int = rand.Intn(7) + 5 // Generate length of word.
	var isVowel bool
	if rand.Float64() > 0.7 { // Determines if the first letter is a vowel.
		isVowel = true
	} else {
		isVowel = false
	}

	var name string

	for ; length > 0; length-- { // Generates the word.
		if isVowel {
			name = name + vowel[int(rand.Float64()*float64(len(vowel)))]
		} else {
			name = name + consonant[int(rand.Float64()*float64(len(consonant)))]
		}
		isVowel = !isVowel
	}

	fmt.Println(name) // Prints the word.
}
