package main

// handler interface

type Department interface {
	execute(*Patient)
	setNext(Department)
}
