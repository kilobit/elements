/* Copyright 2019 Kilobit Labs Inc. */

package elements

import "html"

type value *string

type Attributes map[string]value

type Element struct {
	tag   string     // Tag name
	attrs Attributes // Attributes
	cs    []Node     // Children
}

func NewElement(tag string) *Element {
	return &Element{tag,
		make(Attributes),
		[]Node{},
	}
}

func (el *Element) Tag() string {
	return el.tag
}

func (el *Element) AddAttr(k string) *Element {
	el.attrs[k] = nil
	return el
}

func (el *Element) SetAttr(k string, v string) *Element {
	el.attrs[k] = value(&v)
	return el
}

func (el *Element) Attr(k string) (string, bool) {
	v, ok := el.attrs[k]
	if v == nil {
		return "", ok
	}

	return *v, ok
}

func (el *Element) AddChild(c Node) *Element {
	el.cs = append(el.cs, c)
	return el
}

func (el *Element) Children() []Node {
	return el.cs
}

func (el *Element) ChildN(n int) Node {
	if len(el.cs) <= n {
		return nil
	}

	return el.cs[n]
}

func (el *Element) String() string {

	sattrs := ""
	for k, v := range el.attrs {
		if v == nil {
			sattrs += " \"" + k + "\""
		} else {
			sattrs += " \"" + k + "\"" + "=" + "\"" + *v + "\""
		}
	}

	s := "<" + html.EscapeString(el.tag) + sattrs + ">"

	// Single tag without children
	if len(el.cs) == 0 {
		return s
	}

	for _, c := range el.cs {
		s += c.String()
	}

	s += "</" + html.EscapeString(el.tag) + ">"

	return s
}
