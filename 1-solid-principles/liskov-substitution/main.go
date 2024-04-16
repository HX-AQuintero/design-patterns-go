package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of ", expectedArea, ", but got ", actualArea)
}

// this is working!
// let's define a new struct that "inherits" from Rectangle
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// and two new methods
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// this is breaking the LSP ⚠️ why?
// Square redefines "SetWidth" and "SetHeight" methods,
// creating an unexpected behavior incoherent to Sized expectations.

// LSP --> if a function takes an interface and works with a struct A that
// implements that interface, then any other struct B that embeds A should
// also work within that function.

// derived types must behave consistently with the base interface
// and not introduce unexpected behaviors.

// a pretty much approach
type Square2 struct {
	size int
}

// rectangle still works correctly
func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq) // Expected an area of  50 , but got  100 ❌

}
