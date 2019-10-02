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

	doc := Document("Test Document")

	t.Log(doc.html)

	els.Expect(t, "Test Document", doc.html.head.tstr)
}
