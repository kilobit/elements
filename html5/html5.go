/* Copyright 2019 Kilobit Labs Inc. */

// HTML5 specific documents, tags and formatters.
//
package html5

import els "kilobit.ca/go/elements"

const html5Doctype string = "<!DOCTYPE html>"

type HTML5Document struct {
	html *HTMLElement
}

func Document(title string) *HTML5Document {

	html := HTML(title)

	doc := &HTML5Document{
		html,
	}

	return doc
}

func (doc *HTML5Document) Doctype() string {
	return html5Doctype
}

type HTMLElement struct {
	*els.Element
	head *HeadElement
	body *BodyElement
}

func HTML(title string) *HTMLElement {

	head := Head(title)
	body := Body()

	html := &HTMLElement{
		els.NewElement("html").
			AddChild(head).
			AddChild(body),
		head,
		body,
	}

	return html
}

type HeadElement struct {
	*els.Element
	title *els.Element // Title of the document.
	tstr  string       // Title string.
}

func Head(tstr string) *HeadElement {

	title := els.NewElement("title").
		AddChild(els.Content(tstr))

	head := &HeadElement{
		els.NewElement("head").
			AddChild(title),
		title,
		tstr,
	}

	return head
}

type BodyElement struct {
	*els.Element
}

func Body() *BodyElement {

	body := &BodyElement{
		els.NewElement("body"),
	}

	return body
}
