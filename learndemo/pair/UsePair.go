package main

import "fmt"

type Write interface {
	WriteBook()
}
type Reader interface {
	ReadBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("****read book")
}
func (this *Book) WriteBook() {
	fmt.Println("****write book")
}

func main() {
	//pair<type:Book,value:book{}地址>
	book := &Book{}

	//pair<type:,value:>
	var re Reader
	//pair<type:Book,value:book{}地址>
	re = book

	re.ReadBook()

	var wr Write = book
	wr.WriteBook()
}
