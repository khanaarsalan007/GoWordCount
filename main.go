package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Response struct {
	Result string `json:"result"`
}

type WordCountPair struct {
	Word  string
	Count int32
}

type PairList []WordCountPair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Count < p[j].Count }

func main() {
	words := "Create a service that accepts input as text and provides Json Output as Top ten most used words and times of occurrence in the text"
	fmt.Println(WordsCount(words))
}

func WordsCount(text string) string {
	//creating map to count words
	wordsCount := make(map[string]int32)
	words := strings.Split(text, " ")
	for _, value := range words {
		wordsCount[value]++
	}
	// converting the map[string]int to list in order to sort the values
	wordCountList := make(PairList, len(wordsCount))
	i := 0
	for k, v := range wordsCount {
		wordCountList[i] = WordCountPair{k, v}
		i++
	}
	//sorting desc to get top 10 words
	sort.Sort(sort.Reverse(wordCountList))
	wordCountList = wordCountList[:10]

	//marshalling the object into json
	bytes, _ := json.Marshal(wordCountList)

	return string(bytes)
}
