package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Closer()
}

type ReadCloser interface {
	Reader
	Closer
}
type File struct {
	Name string
}

func (f File) Read() {
	fmt.Println("Reading file:", f.Name)
}

func (f File) Closer() {
	fmt.Println("Closing file:", f.Name)
}

func main() {
	var d1 ReadCloser = File{Name: "google.com"}
	d1.Read()
	d1.Closer()
}
