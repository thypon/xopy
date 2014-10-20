package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

var (
	x, y, output string

	extend             bool
	operation, pattern string
	threshold          int
)

func init() {
	flag.StringVar(&x, "x", "STDIN", "The X-file")
	flag.StringVar(&y, "y", "", "The Y-file")
	flag.StringVar(&output, "output",
		"STDOUT", "The Output File, STDOUT if not specified")

	flag.BoolVar(&extend, "extend", false, "Extend the file read to the longer one")
	flag.StringVar(&operation, "operation", "eq", "Select the operation: eq, and, or, xor, cut")
	flag.StringVar(&pattern, "pattern", "\x00", "Select the replacing pattern")
	flag.IntVar(&threshold, "threshold", 8, "selection threshold in cut mode")
}

type Op func(x, y byte) byte

func Eq(x, y byte) byte {
	if x == y {
		return x
	} else {
		return pattern[0]
	}
}

func And(x, y byte) byte {
	return x & y
}

func Or(x, y byte) byte {
	return x | y
}

func Xor(x, y byte) byte {
	return x ^ y
}

func CutCreator(threshold int) Op {
	i := 0
	return func(x, y byte) byte {
		if i >= threshold {
			i += 1
			return y
		} else {
			i += 1
			return x
		}
	}
}

func main() {
	flag.Parse()

	var op Op

	switch operation {
	case "xor":
		op = Xor
	case "and":
		op = And
	case "or":
		op = Or
	case "cut":
		op = CutCreator(threshold)
	default:
		op = Eq
	}

	if y == "" {
		flag.Usage()
		os.Exit(1)
	}

	var xbuff io.ByteReader
	if x == "STDIN" {
		xbuff = bufio.NewReader(os.Stdin)
	} else {
		xfile, err := os.Open(x)
		if err != nil {
			log.Fatal(err)
		}
		xbuff = bufio.NewReader(xfile)
	}

	yfile, err := os.Open(y)
	if err != nil {
		log.Fatal(err)
	}
	var ybuff io.ByteReader = bufio.NewReader(yfile)

	var outbuff *bufio.Writer
	if output == "STDOUT" {
		outbuff = bufio.NewWriter(os.Stdout)
	} else {
		outfile, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		outbuff = bufio.NewWriter(outfile)
	}
	defer outbuff.Flush()

	for {
		xbyte, xerr := xbuff.ReadByte()
		if xerr != nil && !extend {
			break
		}
		ybyte, yerr := ybuff.ReadByte()
		if yerr != nil && !extend {
			break
		}
		if yerr != nil && xerr != nil {
			break
		}

		outbuff.WriteByte(op(xbyte, ybyte))
	}
}
