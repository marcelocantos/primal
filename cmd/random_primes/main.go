package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/marcelocantos/primal/pkg/primal"
)

var prefixRE = regexp.MustCompile(`^0[box]?`)

func computeBase(s string) int {
	base := 10
	if s[0] == '0' {
		switch {
		case len(s) >= 3 && lower(s[1]) == 'b':
			base = 2
			s = s[2:]
		case len(s) >= 3 && lower(s[1]) == 'o':
			base = 8
			s = s[2:]
		case len(s) >= 3 && lower(s[1]) == 'x':
			base = 16
			s = s[2:]
		default:
			base = 8
			s = s[1:]
		}
	}
	return base
}

func lower(c byte) byte {
	return c | ('x' - 'X')
}

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
	base := computeBase(s)

	rand.Seed(seed)
	seen := map[uint64]struct{}{}
	for i := 0; i < int(n); i++ {
		for {
			u := rand.Uint64()
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
