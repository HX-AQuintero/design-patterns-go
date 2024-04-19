package main

import (
	"fmt"
	"strings"
)

// a better approach
const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

// string method to print as an HTML
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

// here comes the builder
type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName, HtmlElement{rootName, "", []HtmlElement{}}}
}

func (builder *HtmlBuilder) String() string {
	return builder.root.String()
}

func (builder *HtmlBuilder) AddChild(childName string, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	builder.root.elements = append(builder.root.elements, e)
}

// using the pointer to continue adding children
func (builder *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	builder.root.elements = append(builder.root.elements, e)

	return builder
}

func main() {
	// let's convert string into HTML elements
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	// that was easy, wasn't?
	// what if there are many strings?
	words := []string{"hello", "world"}
	sb.Reset()
	//<ul><li>...</li><li>...</li></ul>
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// well, not so bad. But, what if it has a lot of different HTML tags 🤔❓
	// let's use structs
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	// using the childFluent method
	b2 := NewHtmlBuilder("ul")
	b2.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b2.String())

}
