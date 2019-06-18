package main

import "github.com/arxdsilva/POCs/playground/funcaspar"

func main() {
	var fn = func(s string) string { return s }
	funcaspar.RandomFunc("aaa", fn)
	funcaspar.RandomFuncFF("aaa", fn)
}
