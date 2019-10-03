/* Copyright 2019 Kilobit Labs Inc. */

package elements // import "kilobit.ca/go/elements"

// Definition for simple functions that format the element.
//
type Formatter func(s string) string

// Format for the element.
//
type Format struct {
	Close       bool   // Self closing
	Sep         string // Separator
	Prefix      string
	CPrefix     string // Child Prefix
	TagFmt      Formatter
	TagOpen     string
	TagClose    string
	EndTagOpen  string
	EndTagClose string
	AttrSep     string
	AttrKFmt    Formatter
	AttrVFmt    Formatter
}

func IDString(s string) string {
	return s
}

func QuoteString(s string) string {
	return "\"" + s + "\""
}

var defaultFormat Format = Format{
	Close:       false,
	Sep:         "",
	Prefix:      "",
	CPrefix:     "",
	TagFmt:      IDString,
	TagOpen:     "<",
	TagClose:    ">",
	EndTagOpen:  "</",
	EndTagClose: ">",
	AttrSep:     " ",
	AttrKFmt:    IDString,
	AttrVFmt:    QuoteString,
}
