package main

import (
	"bufio"
	"fmt"
	"os"
)

/** Shaking bits a little */
func bit_permute(x, m, shift byte) byte {
	var t byte;
	t = ((x >> shift) ^ x) & m
	x = (x ^ t) ^ (t << shift)
	return x
}

/** Feistel crypto function (which is a result of lots of random tries :D) */
/** Consists of both s-block and p-block */
func f (key byte, block byte) byte {
	x := byte(key*key + block) // mod by 256 because it's byte
	x = bit_permute(x, 0x01, 8)
	x = bit_permute(x, 0x05, 4)
	x = bit_permute(x, 0x40, 1)
	x = bit_permute(x, 0x22, 2)
	return x
}

/** Simplistic hashing using Feistel network chaining */
func main() {
	reader := bufio.NewReader(os.Stdin)

	//var hash uint16 = 0
	a := byte(63)
	b := byte(206)

	// Hash data itself
	for {
		x, err := reader.ReadByte()
		if err != nil {
			break
		}
		res := f(x, a)
		a = b
		b = b ^ res
		//fmt.Printf("a: %08b, b: %08b\n", a, b)
	}

	// Appending some data to end, to get at least uniform-looking distribution :D
	for _, x := range([]byte{159, 45, 22, 112, 209}) {
		res := f(x, a)
		a = b
		b = b ^ res
		//fmt.Printf("a: %08b, b: %08b\n", a, b)
	}

	// version for 16-bit hash
	//hash = uint16(a) << 8 | uint16(b);
	hash := byte(a ^ b)
	fmt.Printf("%02x", hash)
}
