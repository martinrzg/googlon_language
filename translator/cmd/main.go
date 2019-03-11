package cmd

import "fmt"

func main() {

	//TODO uncomment line
	//argFilename := os.Args[1]
	text := newScrollFromFile("resource/input.txt")
	fmt.Println("Scroll word count: ", len(text.words))

	prep, verb, verbSubj, number := text.classify()

	orderVocabulary := distinctWords(prep, verb, verbSubj, number).sortVocabulary()

	fmt.Println(len(prep.words), len(verb.words), len(verbSubj.words), len(number.words))

	fmt.Println(orderVocabulary, len(orderVocabulary.words))

}
