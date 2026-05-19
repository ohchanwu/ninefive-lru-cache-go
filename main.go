package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type item struct {
	value      int
	lastUsedAt time.Time
}

var (
	cache map[int]*item
	capa  = 0
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
	cache = make(map[int]*item, capa)
	fmt.Println("OK")
}

func setVal(words []string) {
	if len(cache) >= capa {
		(*cache[findLRU()]).value = -1
	}
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	val, err := strconv.Atoi(words[2])
	if err != nil {
		log.Fatal(err)
	}
	cache[key] = &item{
		value:      val,
		lastUsedAt: time.Now(),
	}
	fmt.Println("OK")
}

func getVal(words []string) {
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	item, ok := cache[key]
	if !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println((*item).value)
	(*item).lastUsedAt = time.Now()
}

func findLRU() int {
	var lruKey, count int
	for k := range cache {
		if count == 0 {
			lruKey = k
		}
		if time.Since((*cache[k]).lastUsedAt) > time.Since((*cache[lruKey]).lastUsedAt) {
			lruKey = k
		}
		count++
	}
	return lruKey
}
