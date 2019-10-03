/* Copyright 2019 Kilobit Labs Inc. */

package html5

import els "kilobit.ca/go/elements"
import "testing"

func TestDocumentTest(t *testing.T) {
	els.Expect(t, true, true)
}

func TestHTML5Doctype(t *testing.T) {

	doc := HTML5Document{}
	els.Expect(t, html5Doctype, doc.Doctype())
}

func TestEmptyDocument(t *testing.T) {

	doc := Document("Test Document", "en")

	t.Log(doc)

	els.Expect(t, "Test Document", doc.html.head.tstr)
}

func TestElementBuilder(t *testing.T) {

	section := Section().AddClass("top").
		AddChild(P().AddClass("content")).
		AddChild(HR())

	t.Log(section)

	els.Expect(t, true, section.HasClass("top"))
	els.Expect(t, "hr", section.ChildN(1).(*HTML5Element).Tag())
}
