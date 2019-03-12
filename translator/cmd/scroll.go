package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

type scroll struct {
	words []string
}

var alphabetMap = map[string]int{"s": 0, "x": 1, "o": 2, "c": 3, "q": 4, "n": 5, "m": 6, "w": 7, "p": 8, "f": 9,
	"y": 10, "h": 11, "e": 12, "l": 13, "j": 14, "r": 15, "d": 16, "g": 17, "u": 18, "i": 19}
var fooLetters = "udxsmpf"
var barLetters = "ocqnwyheljrgi"
var powDict = map[float64]float64{}

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
	var us []string

	for _, w := range ws {
		if !distinct[w] {
			us = append(us, w)
			distinct[w] = true
		}
	}
	return scroll{words: us}
}

func (s scroll) prettyNumbers() []int {
	var pn []int

	for _, w := range s.words {
		n := wordToNumber(w)
		if isPrettyNumber(n) {
			pn = append(pn, n)
		}
	}

	return pn
}

func isPrettyNumber(n int) bool {
	if n >= 81827 && n%3 == 0 {
		return true
	}
	return false
}

func wordToNumber(s string) int {
	n := float64(0)
	for j, l := range s {
		n += float64(alphabetMap[string(l)]) * significance(j)
	}
	return int(n)

}

func significance(p int) float64 {
	p64 := float64(p)

	v, ok := powDict[p64]
	if ok {
		return v
	} else {
		v = math.Pow(20, p64)
		powDict[p64] = v
	}
	return v

}

func (s scroll) printScrollSummary() {

	prep, verb, verbSubj, number := s.classify()
	orderVocabulary := distinctWords(prep, verb, verbSubj, number).sortVocabulary()
	prettyN := distinctWords(prep, verb, verbSubj, number).prettyNumbers()

	fmt.Printf("\n         +-----------------------------------+\n")
	fmt.Printf("         |       The Googlon Language        |\n")
	fmt.Printf("         |       by  {  Martin Ruiz  }       |\n")
	fmt.Printf("         +-----------------------------------+\n\n")
	fmt.Printf("+---------------+---------+--------------+-------------------+\n")
	fmt.Printf("|  Preposition  |  Verbs  |  Subj verbs  |  Pretty numbers   |\n")
	fmt.Printf("+---------------+---------+--------------+-------------------+\n")
	fmt.Printf("| %13d | %7d | %12d | %16d  |\n", len(prep.words), len(verb.words), len(verbSubj.words), len(prettyN))
	fmt.Printf("+------------------------------------------------------------+\n")
	fmt.Printf("|                         Vocabulary                         |\n")
	fmt.Printf("|                                                            |\n")

	lines := int(math.Ceil(float64(len(orderVocabulary.words)) / float64(5)))

	for i := 0; i < lines; i++ {
		if i+1 == lines {
			fmt.Printf("|  %10s  |\n", orderVocabulary.words[i*5:])
			continue
		}
		fmt.Printf("|  %10s  |\n", orderVocabulary.words[i*5:i*5+5])
	}
	fmt.Printf("+-----------------------------------------------------------+\n")
}
