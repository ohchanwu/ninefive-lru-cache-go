package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	cache map[int]int
	// keys are ordered from LRU to MRU
	keys []int
	capa = 0
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		switch words[0] {
		case "CAPACITY":
			setCapa(words)
		case "SET":
			setVal(words)
		case "GET":
			getVal(words)
		}
	}
}

func setCapa(words []string) {
	var err error
	capa, err = strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	cache = make(map[int]int, capa)
	keys = make([]int, 0, capa)
	fmt.Println("OK")
}

func setVal(words []string) {
	if len(cache) >= capa {
		// remove LRU node from cache
		delete(cache, keys[0])
		// shift key from slice
		keys = slices.Delete(keys, 0, 1)
	}
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	val, err := strconv.Atoi(words[2])
	if err != nil {
		log.Fatal(err)
	}
	cache[key] = val
	keys = append(keys, key)
	fmt.Println("OK")
}

func getVal(words []string) {
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	val, ok := cache[key]
	if !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println(val)
	index := slices.Index(keys, key)
	keys = slices.Delete(keys, index, index+1)
	keys = append(keys, key)
}
