package main

import (
	"fmt"
	"strings"
)

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

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%d<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func (builder *HtmlBuilder) String() string {
	return builder.String()
}

func (builder *HtmlBuilder) AddChild(
	childName string, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	builder.root.elements = append(
		builder.root.elements, e)
}

func (builder *HtmlBuilder) AddChildFluent(
	childName string, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	builder.root.elements = append(
		builder.root.elements, e)
	return builder
}

func (builder *HtmlBuilder) Clear() {
	builder.root = HtmlElement{builder.rootName, "", []HtmlElement{}}
}

func main() {
	// if you want to build a simple HTML paragraph
	// using a strings.Builder
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Printf("%s\n", sb.String())
}
