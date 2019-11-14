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

func main() {
	s := os.Args[1]
	n, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		log.Fatalf("Usage: random_primes n (format of n determines output format)\n%v", err)
	}
	prefix := prefixRE.FindString(s)
	base := computeBase(s)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(n); i++ {
		for {
			u := rand.Uint64()
			if primal.IsPrime(u) {
				fmt.Printf("%s%s\n", prefix, strconv.FormatUint(u, base))
				break
			}
		}
	}
}
