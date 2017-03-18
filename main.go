package main

import (
	"bufio"
	"bytes"
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

	outbuffer := bytes.NewBuffer(nil)

	var seed int64
	ｒｕｎｅｓｃａｎｎｅｒ := bufio.NewScanner(os.Stdin)
	ｒｕｎｅｓｃａｎｎｅｒ.Split(bufio.ScanRunes)
	var r []byte
	for ｒｕｎｅｓｃａｎｎｅｒ.Scan() {
		r = ｒｕｎｅｓｃａｎｎｅｒ.Bytes()
		outbuffer.Write(width.Widen.Bytes(r))

		s, _ := binary.Varint(r)
		seed ^= s
	}
	rand.Seed(seed)

	if string(r) == "\n" {
		outbuffer.Truncate(outbuffer.Len() - 1)
	}

	if kana != nil && *kana {
		for i := rand.Intn(6); i >= 0; i-- {
			outbuffer.Write([]byte(string(rune(KanaLo + rand.Intn(KanaHi-KanaLo)))))
		}
	}

	if string(r) == "\n" {
		outbuffer.Write([]byte("\n"))
	}

	outbuffer.WriteTo(os.Stdout)
}
