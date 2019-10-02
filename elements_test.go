/* Copyright 2019 Kilobit Labs Inc. */

// Tests for the elements unit.
//
package elements

import "testing"

func TestElementsTest(t *testing.T) {
	Expect(t, true, true)
}

func TestElement(t *testing.T) {

	el := NewElement("tag").
		AddAttr("tagged").
		SetAttr("foo", "bar")

	t.Log(el)
	
	Expect(t, "tag", el.Tag())

	attr, ok := el.Attr("tagged")
	Expect(t, true, ok)
	Expect(t, "", attr)

	attr, ok = el.Attr("foo")
	Expect(t, true, ok)
	Expect(t, "bar", attr)
}

func TestElementNested(t *testing.T) {

	el := NewElement("section").
		AddChild(NewElement("nav")).
		AddChild(NewElement("main"))

	t.Log(el)
	
	Expect(t, 2, len(el.Children()))
}

func TestElementContent(t *testing.T) {

	c := Content("Hello > World!")

	el := NewElement("p").
		AddChild(c)

	t.Log(el)
	
	Expect(t, c, el.ChildN(0))
	Expect(t, 1, len(el.Children()))
}
