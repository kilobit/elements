/* Copyright 2019 Kilobit Labs Inc. */

package elements

const HTML5Doctype string = "<!DOCTYPE html>"

// A fully encapsulated markup document.
//
type Document interface {
	Doctype() string
}

type HTML5Document struct {
}

func (doc *HTML5Document) Doctype() string {
	return HTML5Doctype
}
