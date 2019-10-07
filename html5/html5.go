/* Copyright 2019 Kilobit Labs Inc. */

// HTML5 specific documents, tags and formatters.
//
package html5

import els "kilobit.ca/go/elements"
import "html"

const html5Doctype string = "<!DOCTYPE html>"

var html5Format els.Format = els.Format{
	Close:       false,
	Sep:         "",
	Prefix:      "",
	CPrefix:     "",
	TagFmt:      html.EscapeString,
	TagOpen:     "<",
	TagClose:    ">",
	EndTagOpen:  "</",
	EndTagClose: ">",
	AttrSep:     " ",
	AttrKFmt:    html.EscapeString,
	AttrVFmt:    func(s string) string { return els.QuoteString(html.EscapeString(s)) },
}

var closedFormat els.Format = els.Format{
	Close:       true,
	Sep:         "",
	Prefix:      "",
	CPrefix:     "",
	TagFmt:      html.EscapeString,
	TagOpen:     "<",
	TagClose:    ">",
	EndTagOpen:  "</",
	EndTagClose: ">",
	AttrSep:     " ",
	AttrKFmt:    html.EscapeString,
	AttrVFmt:    func(s string) string { return els.QuoteString(html.EscapeString(s)) },
}

type HTML5Document struct {
	html *HTMLElement
}

func Document(title string, lang string) *HTML5Document {

	html := HTML(title, lang)

	doc := &HTML5Document{
		html,
	}

	return doc
}

func (doc *HTML5Document) Copy() *HTML5Document {

	return &HTML5Document{
		doc.html.Copy(),
	}
}

func (doc *HTML5Document) Doctype() string {
	return html5Doctype
}

func (doc *HTML5Document) HTML() *HTMLElement {
	return doc.html
}

func (doc *HTML5Document) String() string {
	s := doc.Doctype() + doc.html.String()
	return s
}

type HTMLElement struct {
	*els.Element
	head *HeadElement
	body *BodyElement
}

func HTML(title string, lang string) *HTMLElement {

	head := Head(title)
	body := Body()

	html := &HTMLElement{
		els.NewElement("html").
			SetAttr("lang", lang).
			AddChild(head).
			AddChild(body),
		head,
		body,
	}

	return html
}

func (html *HTMLElement) Copy() *HTMLElement {

	el := html.Element.Copy()
	head := el.ChildN(0).(*HeadElement)
	body := el.ChildN(1).(*BodyElement)

	return &HTMLElement{el, head, body}
}

func (html *HTMLElement) Head() *HeadElement {
	return html.head
}

func (html *HTMLElement) Body() *BodyElement {
	return html.body
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
			AddChild(Meta().SetAttr("charset", "utf-8")).
			AddChild(title),
		title,
		tstr,
	}

	return head
}

func (head *HeadElement) Title() string {
	return head.tstr
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

type AnchorElement struct {
	*els.Element
	href string
}

func A(href string) *AnchorElement {
	return &AnchorElement{
		els.NewElement("A").
			SetAttr("href", href),
		href,
	}
}

func (el *AnchorElement) SetHref(href string) *AnchorElement {
	el.SetAttr("href", href)
	el.href = href

	return el
}

type HTML5Element struct {
	*els.Element
}

func NewHTML5Element(tag string, opts ...els.ElementOption) *HTML5Element {
	return &HTML5Element{
		els.NewElement(tag, opts...),
	}
}

type elementBuilder func() *HTML5Element

func makeElementBuilder(tag string, opts ...els.ElementOption) elementBuilder {
	return func() *HTML5Element {
		return NewHTML5Element(tag, opts...)
	}
}

var html5 = els.OptFormat(&html5Format)
var closed = els.OptFormat(&closedFormat)

// Meta Tags
var Base elementBuilder = makeElementBuilder("base", closed)
var Link elementBuilder = makeElementBuilder("link", closed)
var Meta elementBuilder = makeElementBuilder("meta", closed)
var Style elementBuilder = makeElementBuilder("style", html5)
var Title elementBuilder = makeElementBuilder("title", html5)

// Structural Tags
var Article elementBuilder = makeElementBuilder("article", html5)
var Aside elementBuilder = makeElementBuilder("aside", html5)
var BR elementBuilder = makeElementBuilder("br", closed)
var Details elementBuilder = makeElementBuilder("details", html5)
var Div elementBuilder = makeElementBuilder("div", html5)
var H1 elementBuilder = makeElementBuilder("h1", html5)
var H2 elementBuilder = makeElementBuilder("h2", html5)
var H3 elementBuilder = makeElementBuilder("h3", html5)
var H4 elementBuilder = makeElementBuilder("h4", html5)
var H5 elementBuilder = makeElementBuilder("h5", html5)
var H6 elementBuilder = makeElementBuilder("h6", html5)
var Header elementBuilder = makeElementBuilder("header", html5)
var HGroup elementBuilder = makeElementBuilder("hgroup", html5)
var HR elementBuilder = makeElementBuilder("hr", closed)
var Footer elementBuilder = makeElementBuilder("footer", html5)
var Nav elementBuilder = makeElementBuilder("nav", html5)
var P elementBuilder = makeElementBuilder("p", html5)
var Section elementBuilder = makeElementBuilder("section", html5)
var Span elementBuilder = makeElementBuilder("span", html5)
var Summary elementBuilder = makeElementBuilder("summary", html5)

// Form Tags
var Button elementBuilder = makeElementBuilder("button", html5)
var DataList elementBuilder = makeElementBuilder("datalist", html5)
var FieldSet elementBuilder = makeElementBuilder("fieldset", html5)
var Form elementBuilder = makeElementBuilder("form", html5)
var Input elementBuilder = makeElementBuilder("input", closed)
var Keygen elementBuilder = makeElementBuilder("keygen", html5)
var Label elementBuilder = makeElementBuilder("label", html5)
var Legend elementBuilder = makeElementBuilder("legend", html5)
var Meter elementBuilder = makeElementBuilder("meter", html5)
var OptGroup elementBuilder = makeElementBuilder("optgroup", html5)
var Option elementBuilder = makeElementBuilder("option", html5)
var Select elementBuilder = makeElementBuilder("select", html5)
var TextArea elementBuilder = makeElementBuilder("textarea", html5)

// Formatting Tags
var Abbr elementBuilder = makeElementBuilder("abbr", html5)
var Acronym elementBuilder = makeElementBuilder("acronym", html5)
var Address elementBuilder = makeElementBuilder("address", html5)
var B elementBuilder = makeElementBuilder("b", html5)
var Bdi elementBuilder = makeElementBuilder("bdi", html5)
var Bdo elementBuilder = makeElementBuilder("bdo", html5)
var Big elementBuilder = makeElementBuilder("big", html5)
var BlockQuote elementBuilder = makeElementBuilder("blockquote", html5)
var Cite elementBuilder = makeElementBuilder("cite", html5)
var Code elementBuilder = makeElementBuilder("code", html5)
var Del elementBuilder = makeElementBuilder("del", html5)
var Dfn elementBuilder = makeElementBuilder("dfn", html5)
var Em elementBuilder = makeElementBuilder("em", html5)
var I elementBuilder = makeElementBuilder("i", html5)
var Ins elementBuilder = makeElementBuilder("ins", html5)
var Kbd elementBuilder = makeElementBuilder("kbd", html5)
var Mark elementBuilder = makeElementBuilder("mark", html5)
var Output elementBuilder = makeElementBuilder("output", html5)
var Pre elementBuilder = makeElementBuilder("pre", html5)
var Progress elementBuilder = makeElementBuilder("progress", html5)
var Q elementBuilder = makeElementBuilder("q", html5)
var RP elementBuilder = makeElementBuilder("rp", html5)
var RT elementBuilder = makeElementBuilder("rt", html5)
var Ruby elementBuilder = makeElementBuilder("ruby", html5)
var Samp elementBuilder = makeElementBuilder("samp", html5)
var Small elementBuilder = makeElementBuilder("small", html5)
var Strong elementBuilder = makeElementBuilder("strong", html5)
var Sub elementBuilder = makeElementBuilder("sub", html5)
var Sup elementBuilder = makeElementBuilder("sup", html5)
var TT elementBuilder = makeElementBuilder("tt", html5)
var Var elementBuilder = makeElementBuilder("var", html5)
var WBR elementBuilder = makeElementBuilder("wbr", html5)

// List Tags
var DD elementBuilder = makeElementBuilder("dd", html5)
var Dir elementBuilder = makeElementBuilder("dir", html5)
var DL elementBuilder = makeElementBuilder("dl", html5)
var DT elementBuilder = makeElementBuilder("dt", html5)
var LI elementBuilder = makeElementBuilder("li", html5)
var OL elementBuilder = makeElementBuilder("ol", html5)
var Menu elementBuilder = makeElementBuilder("menu", html5)
var UL elementBuilder = makeElementBuilder("ul", html5)

// Table Tags
var Caption elementBuilder = makeElementBuilder("caption", html5)
var Col elementBuilder = makeElementBuilder("col", html5)
var Colgroup elementBuilder = makeElementBuilder("colgroup", html5)
var Table elementBuilder = makeElementBuilder("table", html5)
var TBody elementBuilder = makeElementBuilder("tbody", html5)
var TD elementBuilder = makeElementBuilder("td", html5)
var TFoot elementBuilder = makeElementBuilder("tfoot", html5)
var THead elementBuilder = makeElementBuilder("thead", html5)
var TH elementBuilder = makeElementBuilder("th", html5)
var TR elementBuilder = makeElementBuilder("tr", html5)

// Scripting Tags
var NoScript elementBuilder = makeElementBuilder("noscript", html5)
var Script elementBuilder = makeElementBuilder("script", html5)

// Embedded Content Tags
var Area elementBuilder = makeElementBuilder("area", html5)
var Audio elementBuilder = makeElementBuilder("audio", html5)
var Canvas elementBuilder = makeElementBuilder("canvas", html5)
var Embed elementBuilder = makeElementBuilder("embed", html5)
var FigCaption elementBuilder = makeElementBuilder("figcaption", html5)
var Figure elementBuilder = makeElementBuilder("figure", html5)
var Frame elementBuilder = makeElementBuilder("frame", html5)
var FrameSet elementBuilder = makeElementBuilder("frameset", html5)
var IFrame elementBuilder = makeElementBuilder("iframe", html5)
var Img elementBuilder = makeElementBuilder("img", closed)
var Map elementBuilder = makeElementBuilder("map", html5)
var NoFrames elementBuilder = makeElementBuilder("noframes", html5)
var Object elementBuilder = makeElementBuilder("object", html5)
var Param elementBuilder = makeElementBuilder("param", html5)
var Source elementBuilder = makeElementBuilder("source", html5)
var Time elementBuilder = makeElementBuilder("time", html5)
var Video elementBuilder = makeElementBuilder("video", html5)
