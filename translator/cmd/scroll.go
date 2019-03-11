package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type scroll struct {
	words []string
}

var fooLettersArr = []string{"u", "d", "x", "s", "m", "p", "f"}
var alphabetArr = []string{"s", "x", "o", "c", "q", "n", "m", "w", "p", "f", "y", "h", "e", "l", "j", "r", "d", "g", "u", "i"}
var alphabetMap = map[string]int{"s": 0, "x": 1, "o": 2, "c": 3, "q": 4, "n": 5, "m": 6, "w": 7, "p": 8, "f": 9, "y": 10, "h": 11, "e": 12,
	"l": 13, "j": 14, "r": 15, "d": 16, "g": 17, "u": 18, "i": 19}
var barLettersArr = []string{"o", "c", "q", "n", "w", "y", "h", "e", "l", "j", "r", "g", "i"}
var fooLetters = "udxsmpf"
var barLetters = "ocqnwyheljrgi"

func newScrollFromFile(filename string) scroll {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Fatal error: ", err)
		os.Exit(1)
	}

	cleanFile := strings.Replace(string(file), "\n", " ", -1)

	data := strings.Split(cleanFile, " ")
	return scroll{words: data}
}

func (s scroll) classify() (scroll, scroll, scroll, scroll) {
	var prep scroll
	var verb scroll
	var verbSubj scroll
	var number scroll

	for _, v := range s.words {
		if isPrep(v) {
			prep.words = append(prep.words, v)
		} else if isVerb(v) {
			if isVerbSubj(v) {
				verbSubj.words = append(verbSubj.words, v)
			}
			verb.words = append(verb.words, v)
		} else {
			number.words = append(number.words, v)
		}
	}
	return prep, verb, verbSubj, number
}

func isPrep(s string) bool {
	if len(s) == 6 {
		if strings.ContainsAny(s[len(s)-1:], fooLetters) {
			if !strings.ContainsAny(s, "u") {
				return true
			}
		}
	}
	return false
}

func isVerb(s string) bool {
	if len(s) >= 6 {
		if strings.ContainsAny(s[len(s)-1:], barLetters) {
			return true
		}
	}
	return false
}

func isVerbSubj(s string) bool {
	if strings.ContainsAny(s[0:1], barLetters) {
		return true
	}
	return false
}

func (s scroll) sortVocabulary() scroll {

	sort.Slice(s.words, func(i, j int) bool {
		wi := s.words[i]
		wj := s.words[j]
		var minLen int

		if len(wi) < len(wj) {
			minLen = len(wi)
		} else {
			minLen = len(wj)
		}

		for i := 0; i < minLen; i++ {
			l1 := string(wi[i])
			l2 := string(wj[i])
			if alphabetMap[l1] == alphabetMap[l2] {
			} else if alphabetMap[l1] < alphabetMap[l2] {
				return true
			} else {
				return false
			}
		}
		return false
	})
	return s
}

func distinctWords(scrolls ...scroll) scroll {
	var ws []string

	for _, s := range scrolls {
		for _, w := range s.words {
			ws = append(ws, w)
		}
	}

	distinct := make(map[string]bool, len(ws))
	us := []string{}

	for _, w := range ws {
		if !distinct[w] {
			us = append(us, w)
			distinct[w] = true
		}
	}

	return scroll{words: us}

}
