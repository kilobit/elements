/* Copyright 2019 Kilobit Labs Inc. */

package elements // import "kilobit.ca/go/elements"

import "strings"

type value *string

type ElementOption func(el *Element)

type Attributes map[string]value

type Element struct {
	tag   string     // Tag name
	attrs Attributes // Attributes
	cs    []Node     // Children
	fmt   *Format    // Format
}

func NewElement(tag string, opts ...ElementOption) *Element {
	el := &Element{
		tag,
		make(Attributes),
		[]Node{},
		&defaultFormat,
	}

	for _, opt := range opts {
		opt(el)
	}

	return el
}

func (el *Element) Copy() *Element {

	attrs := make(Attributes, len(el.attrs))
	for k, v := range el.attrs {
		attrs[k] = v
	}

	cs := make([]Node, len(el.cs))
	for i, c := range el.cs {
		switch v := c.(type) {
		case *Element:
			cs[i] = v.Copy()
		default:
			cs[i] = c
		}
	}

	fmt := *el.fmt

	return &Element{
		el.tag,
		attrs,
		cs,
		&fmt,
	}
}

func (el *Element) Tag() string {
	return el.tag
}

func (el *Element) AddAttr(ks ...string) *Element {
	for _, k := range ks {
		el.attrs[k] = nil
	}

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

func (el *Element) SetID(id string) *Element {
	el.SetAttr("id", id)
	return el
}

func (el *Element) ID() string {
	id, ok := el.Attr("id")
	if !ok {
		return ""
	}

	return id
}

func (el *Element) AddClass(class string) *Element {

	cur, ok := el.Attr("class")
	if !ok {
		el.SetAttr("class", class)
		return el
	}

	cur += " " + class

	el.SetAttr("class", cur)
	return el
}

func (el *Element) HasClass(class string) bool {

	classes, ok := el.Attr("class")
	if !ok {
		return false
	}

	for _, c := range strings.Split(classes, " ") {
		if strings.Index(c, class) != -1 {
			return true
		}
	}

	return false
}

func (el *Element) AddChild(c ...Node) *Element {
	el.cs = append(el.cs, c...)
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

func (el *Element) SetFormat(fmt *Format) {
	el.fmt = fmt
}

func (el *Element) Format() *Format {
	return el.fmt
}

func (el *Element) String() string {

	sattrs := ""
	for k, v := range el.attrs {
		if v == nil {
			sattrs += el.fmt.AttrSep + el.fmt.AttrKFmt(k)
		} else {
			sattrs += el.fmt.AttrSep + el.fmt.AttrKFmt(k) + "=" +
				el.fmt.AttrVFmt(*v)
		}
	}

	s := el.fmt.Prefix + el.fmt.TagOpen + el.fmt.TagFmt(el.tag) + sattrs +
		el.fmt.TagClose

	// Single tag without children
	if len(el.cs) == 0 && el.fmt.Close {
		return s
	}

	s += el.fmt.AttrSep

	for _, c := range el.cs {
		s += el.fmt.CPrefix + c.String() + el.fmt.Sep
	}

	s += el.fmt.Prefix + el.fmt.EndTagOpen + el.fmt.TagFmt(el.tag) +
		el.fmt.EndTagClose + el.fmt.Sep

	return s
}

func OptFormat(fmt *Format) ElementOption {
	return func(el *Element) {
		el.fmt = fmt
	}
}
