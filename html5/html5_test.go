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

func TestDocumentCopy(t *testing.T) {

	doc := Document("Test Document", "en")

	copy := doc.Copy()

	t.Logf("%#v", doc)
	t.Logf("%#v", copy)

	els.Expect(t, "Test Document", doc.html.head.tstr)
}

func TestHTMLCopy(t *testing.T) {

	html := HTML("Test Document", "en")

	copy := html.Copy()

	t.Logf("%#v", html)
	t.Logf("%#v", copy)

	els.Expect(t, "Test Document", html.head.tstr)
}

func TestHTML5ElementCopy(t *testing.T) {

	section := Section().
		AddChild(P().
			AddChild(els.Content("Hello World!")))

	copy := section.Copy()

	t.Logf("%#v", section)
	t.Logf("%#v", copy)

	els.Expect(t, "Hello World!", section.ChildN(0).(*els.Element).ChildN(0).String())
}

func TestElementBuilder(t *testing.T) {

	section := Section().AddClass("top").
		AddChild(P().AddClass("content")).
		AddChild(HR())

	t.Log(section)

	els.Expect(t, true, section.HasClass("top"))
	els.Expect(t, "hr", section.ChildN(1).(*HTML5Element).Tag())
}

func TestHTML5Elements(t *testing.T) {

	section := Section(
		H1(els.Content("Title")),
		P(els.Content("Hello World!")),
		HR(),
		Footer(P(els.Content("Goodbye"))),
	)

	t.Log(section)

	els.Expect(t, 4, len(section.Children()))
}
