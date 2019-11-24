package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/marcelocantos/primal/pkg/primal"
)

var prefixRE = regexp.MustCompile(`^0[box]?`)

func usage(err error) {
	log.Fatalf("Usage: random_primes n [seed] (format of n determines output format)\n%v", err)
}

func main() {
	s := os.Args[1]
	seed := time.Now().UnixNano()
	if len(os.Args) > 2 {
		var err error
		seed, err = strconv.ParseInt(os.Args[2], 0, 64)
		if err != nil {
			usage(err)
		}
	}
	n, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		usage(err)
	}
	prefix := prefixRE.FindString(s)
	base, hasBase := map[string]int{"0b": 1, "0": 8, "0o": 8, "0x": 16}[strings.ToLower(prefix)]
	if !hasBase {
		base = 10
	}

	rand.Seed(seed)
	seen := map[uint64]struct{}{}
	for i := 0; i < int(n); i++ {
		for {
			u := rand.Uint64() | 1
			if _, has := seen[u]; !has {
				if primal.IsPrime(u) {
					seen[u] = struct{}{}
					fmt.Printf("%s%s\n", prefix, strconv.FormatUint(u, base))
					break
				}
			}
		}
	}
}
