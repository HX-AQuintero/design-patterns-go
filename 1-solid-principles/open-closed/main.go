package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//...
}

// this is going to be working fine!
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// but, what if we need to filter by size? 🤔
// well, we could copy FilterbyColor and change it but ... what if we need another one or both?
// this is a violation of the OCP ⚠️

// using Specification Pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

// we define structs implementing the Specification interface
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

// better filter has a method which receives the spec and implements the interface
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// using OCP: Specification interface is open to extension (implement)
	// closed to modification (change its structure)
	bf := BetterFilter{}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	
	// new spec just for testing
	largeSpec := SizeSpecification{large}
	
	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}
	
	// using both specs
	fmt.Printf("Green and large products (new):\n")
	largeGreenSpec := AndSpecification{greenSpec, largeSpec}

	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}
