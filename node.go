/* Copyright 2019 Kilobit Labs Inc. */

package elements

import "html"
import "fmt"

type Node interface {
	fmt.Stringer
}

type ContentNode string

func (s ContentNode) String() string {
	return html.EscapeString(string(s))
}

func Content(s string) ContentNode {
	return ContentNode(s)
}
