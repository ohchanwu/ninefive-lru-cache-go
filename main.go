package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

var (
	cache     map[int]*Node
	capa      = 0
	lastNode  = &Node{value: -2}
	firstNode = &Node{value: -1}
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
	if capa < 1 {
		log.Fatal("Capacity must be greater than 0")
	}
	cache = make(map[int]*Node, capa)
	fmt.Println("OK")
}

func setVal(words []string) {
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	val, err := strconv.Atoi(words[2])
	if err != nil {
		log.Fatal(err)
	}
	node, ok := cache[key]
	if ok {
		node.value = val
		moveNodeToMRU(node)
	} else {
		node = &Node{
			key:   key,
			value: val,
		}
		insertNewNode(node)
		cache[key] = node
	}
	fmt.Println("OK")
}

func getVal(words []string) {
	key, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	node, ok := cache[key]
	if !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println(node.value)
	moveNodeToMRU(node)
}

func insertNewNode(node *Node) {
	switch len(cache) {
	case 0:
		node.prev = firstNode
		node.next = lastNode
		firstNode.next = node
		lastNode.prev = node
	case capa:
		// delete from cache
		delete(cache, firstNode.next.key)

		// delete LRU node
		firstNode.next.next.prev = firstNode
		firstNode.next = firstNode.next.next

		// insert new node
		node.prev = lastNode.prev
		node.next = lastNode
		// FIX: Forgot this one too (see comment below).
		lastNode.prev.next = node
		lastNode.prev = node
	default:
		node.prev = lastNode.prev
		node.next = lastNode
		// FIX: I forgot this operation, which caused a bug where firstNode's next node stayed as 1 even after 2 was inserted.
		// Remember, to splice a new node X into a doubly linked list, 4 pointers must be updated:
		// X.prev = A			 X.next = B			A.next = X			B.prev = X
		lastNode.prev.next = node
		lastNode.prev = node
	}
}

func moveNodeToMRU(node *Node) {
	// remove node
	node.next.prev = node.prev
	node.prev.next = node.next

	// insert node before lastNode
	node.next = lastNode
	node.prev = lastNode.prev
	lastNode.prev.next = node
	lastNode.prev = node
}
