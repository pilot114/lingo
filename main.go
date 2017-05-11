package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "sort"
)

// A data structure to hold key/value pairs
type Pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func ValidLines(filename string) (c chan string) {
    c = make(chan string)
    buff := ""
    go func() {
        file, err := os.Open(filename)
        if err != nil {
            close(c)
            return
        }
        reg, err := regexp.Compile("[^A-Za-z0-9а-яА-Я ]+")
        if err != nil {
            log.Fatal(err)
        }

        reader := bufio.NewReader(file)
        for {
            line, err := reader.ReadString('\n')
            if err != nil {
                close(c)
                return
            }
            // очистка от мусорных символов
            line = reg.ReplaceAllString(line, "")
            if line == "" {
                continue
            }
            buff += line
            c <- buff
            buff = ""
        }
    }()
    return c
}

func Sort() {
	
}

func main() {
	words := make(map[string]int)
    for line := range ValidLines("lordring.txt") {
    	for _, w := range strings.Fields(strings.ToLower(line)) {
			words[w]++
		}
    }
    sortWords := make(PairList, len(words))
	i := 0
	for k, v := range words {
		sortWords[i] = Pair{k, v}
		i++
	}
	sort.Sort(sortWords)
	
    fmt.Println(len(sortWords))
}