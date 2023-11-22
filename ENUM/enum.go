package main

// sample of an basic enum based on int
const (
	A = iota // 0
	B        // 1
	C        // 2
)

// sample of an enum based on string
const (
	D = "a" // "a"
	E = "b" // "b"
	F = "c" // "c"
)

// sample of an enum based on own type
// this dont allow to compare two enums with towo different types

type MyType int

// to prevent having no value, we could asign it a value of 0 and set an undifined name => Default enum value

const (
	undifined MyType = iota // 0
	G                       // 1
	H                       // 2
)

// a good pratice to use with typed enum is to use an receiver func to print the name or the value

func (m MyType) String() string {
	switch m {
	case G:
		return "G name"
	case H:
		return "H name"
	default:
		return "undefined"
	}
}

func (m MyType) Value() int {
	switch m {
	case G:
		return int(G)
	case H:
		return int(H)
	default:
		return int(undifined)
	}
}

func (m MyType) Name() string {
	switch m {
	case G:
		return "G name"
	case H:
		return "H name"
	default:
		return "undefined"
	}
}
