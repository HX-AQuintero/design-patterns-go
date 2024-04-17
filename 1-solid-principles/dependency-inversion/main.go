package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	//...
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low-level module
type Relationships struct {
	relations []Info
}

// a better approach using DIP
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// break DIP ⚠️ why?
	// HLM doesn't depend on LLM
	// relationships Relationships
	browser RelationshipBrowser
	//...
}

// however, this actually works!
func (r *Research) Investigate() {
	// this is using the internals of the Relationships module (LLM). but...
	// what if relationships changes the way relations are been stored ❓🤔
	// it's going to break the HLM
	/* relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	} */

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	// high-level module relies on low-level module
	research := Research{&relationships}
	research.Investigate()
}
