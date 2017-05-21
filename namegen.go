package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	f int = 2

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

	vowel []string = []string{
		"a",
		"e",
		"i",
		"o",
		"u",
	}
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	
	var length int = rand.Intn(8) + 5
	var vow = false
	var name string
	
	for ; length > 0; length-- {
		if vow {
			name = name + vowel[int(rand.Float64()*float64(len(vowel)))]
		} else {
			name = name + consonant[int(rand.Float64()*float64(len(consonant)))]
		}
		vow = !vow
	}
	fmt.Println(name)
}
