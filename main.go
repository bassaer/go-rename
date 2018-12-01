package main

import (
	"fmt"
	"io"
	"os"
)

func read(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, f)
}

func cp(src, dst string) {
	srcfile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcfile.Close()
	dstfile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer dstfile.Close()
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		panic(err)
	}
}

func create(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "ccc")
}

func main() {
	a := "./a.txt"
	b := "./b.txt"
	c := "./c.txt"
	read(a)
	cp(a, b)
	read(a)
	read(b)
	create(c)
	read(c)
	os.Rename(c, a)
	read(a)
	os.Rename(b, a)
	read(a)
}
