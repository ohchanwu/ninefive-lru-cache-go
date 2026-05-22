# LRU Cache

A least-recently-used (LRU) cache in Go with **O(1)** `GET` and `SET`. Hand-coded for [ninefive](https://www.95ninefive.dev/), a gamified app and CLI for practicing coding by hand.

## What it is

An LRU cache holds a fixed number of entries. When it is full and a new entry arrives, it removes the least-recently-used entry to make room. This program reads commands from standard input:

| Command         | Behavior                                                                  |
| --------------- | ------------------------------------------------------------------------- |
| `CAPACITY n`    | Set the cache capacity to `n`.                                            |
| `SET key value` | Store a value. Evict the least-recently-used entry if the cache is full.  |
| `GET key`       | Return the value, or `-1` if absent, and mark the key most recently used. |

## How it works

`GET`, `SET`, and eviction all run in O(1) time. The cache uses two data structures:

- A hash map (`map[int]*Node`) maps each key to its node, which gives constant-time lookup.
- A doubly linked list keeps the nodes ordered from least to most recently used.

The map returns a node directly, so that node can be unlinked and moved to the most-recently-used end of the list with a fixed number of pointer updates, without scanning. Eviction removes the node at the least-recently-used end.

Two sentinel nodes sit at the head and tail of the list and are never removed. Because they are always present, insertion and removal never need a special case for an empty list or for the ends of the list.

Standard library only. No external dependencies, one file.

## Run it

```bash
go build -o ./app && ./app
```

The program reads one command per line from standard input. It replies `OK` to each `CAPACITY` and `SET`, and prints the stored value — or `-1` if the key is absent — for each `GET`. Typing the commands on the left produces the output on the right:

| You type     | Program prints                               |
| ------------ | -------------------------------------------- |
| `CAPACITY 2` | `OK`                                         |
| `SET 1 100`  | `OK`                                         |
| `SET 2 200`  | `OK`                                         |
| `GET 1`      | `100`                                        |
| `SET 3 300`  | `OK` — evicts key 2, the least recently used |
| `GET 2`      | `-1` — key 2 was evicted                     |

Press `Ctrl-D` to send EOF and exit.

## About this quest

This is a side quest, "Build Your Own LRU Cache", and my first completed quest on ninefive. I built it in steps: the commit history goes from a naive version, through capacity-aware eviction and recency promotion, to a doubly linked list that makes every operation O(1).

The exercise included this reference reading, which I read while building it:

- [Cache replacement policies](https://en.wikipedia.org/wiki/Cache_replacement_policies)
- [Doubly linked list](https://en.wikipedia.org/wiki/Doubly_linked_list)
- [Hash table](https://en.wikipedia.org/wiki/Hash_table)

## Lessons learned

**Splicing a node takes four pointer updates, not three.** Inserting a node `X` between neighbors `A` and `B` needs all four of `X.prev = A`, `X.next = B`, `A.next = X`, and `B.prev = X`. I wrote three and missed `A.next = X`, the predecessor's forward pointer. The backward chain was still correct, so the bug did not appear until a forward traversal reached the gap. The fix was one line. The lesson: count the pointer updates instead of assuming a splice is correct because it looks correct. (See the comments in `insertNewNode`.)

**Passing the tests is not the same as meeting the spec.** My first version tracked recency order with a slice. It passed every functional test, but `slices.Index()` and `slices.Delete()` are O(n), and the exercise required O(1). The output was correct, which hid the wrong complexity class. I now treat a performance requirement as an acceptance criterion and check the cost of the standard-library calls I depend on.

## AI usage

I hand-coded this entire project. When I got stuck, I asked an AI assistant for conceptual hints only, not the solution and not the code. The data structure, the bug fix, and the complexity analysis are my own work.
