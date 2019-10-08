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

	attr := el.Attr("tagged")
	Expect(t, "", attr)

	attr = el.Attr("foo")
	Expect(t, "bar", attr)
}

func TestElementCopy(t *testing.T) {

	el := NewElement("tag").
		AddAttr("tagged").
		SetAttr("foo", "bar").
		AddChild(NewElement("nav")).
		AddChild(NewElement("main")).
		AddChild(Content("Hello World!"))

	cel := el.Copy()

	t.Logf("%#v", cel)
	t.Log(cel)

	attr := cel.Attr("tagged")
	Expect(t, "", attr)

	attr = cel.Attr("foo")
	Expect(t, "bar", attr)

	Expect(t, 3, len(el.Children()))
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

func TestElementAttrs(t *testing.T) {

	el := NewElement("TAG").
		SetAttrs(
			map[string]string{
				"foo":    "bar",
				"bing":   "bang!",
				"cheery": "oh",
			})

	t.Log(el)

	Expect(t, "bar", el.Attr("foo"))
	Expect(t, "bang!", el.Attr("bing"))
	Expect(t, "oh", el.Attr("cheery"))
}
