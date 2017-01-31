package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"golang.org/x/text/width"
	"math/rand"
	"os"
)

// Range of full-width kana characters
const (
	KanaLo = 0x30A0
	KanaHi = 0x30FF
)

func main() {
	kana := flag.Bool("kana", false, "append random full-width kana")
	flag.Parse()

	seeded := false

	ｒｕｎｅｓｃａｎｎｅｒ := bufio.NewScanner(os.Stdin)
	ｒｕｎｅｓｃａｎｎｅｒ.Split(bufio.ScanRunes)
	for ｒｕｎｅｓｃａｎｎｅｒ.Scan() {
		r := ｒｕｎｅｓｃａｎｎｅｒ.Bytes()
		os.Stdout.Write(width.Widen.Bytes(r))

		if !seeded { // first rune is the seed, deterministic
			seed, _ := binary.Varint(r)
			rand.Seed(seed)
			seeded = true
		}
	}

	if kana != nil && *kana {
		for i := rand.Intn(6); i >= 0; i-- {
			os.Stdout.Write([]byte(string(rune(KanaLo + rand.Intn(KanaHi-KanaLo)))))
		}
	}
}
