package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	cache map[int]int
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
	cache = make(map[int]int, capa)
	fmt.Println("OK")
}

func setVal(words []string) {
	if len(cache) >= capa {
		fmt.Println("OK")
		return
	}
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	cache[key], err = strconv.Atoi(words[2])
	if err != nil {
		log.Fatal(err)
	}
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
}
