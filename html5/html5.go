/* Copyright 2019 Kilobit Labs Inc. */

// HTML5 specific documents, tags and formatters.
//
package html5

import els "kilobit.ca/go/elements"

const html5Doctype string = "<!DOCTYPE html>"

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

func (doc *HTML5Document) Doctype() string {
	return html5Doctype
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

type HeadElement struct {
	*els.Element
	title *els.Element // Title of the document.
	tstr  string       // Title string.
}

func Head(tstr string) *HeadElement {

	title := els.NewElement("title").
		AddChild(Meta().SetAttr("charset", "utf-8")).
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

type elementBuilder func() *els.Element

func makeElementBuilder(tag string) elementBuilder {
	return func() *els.Element {
		return els.NewElement(tag)
	}
}

// Meta Tags
var Base elementBuilder = makeElementBuilder("base")
var Link elementBuilder = makeElementBuilder("link")
var Meta elementBuilder = makeElementBuilder("meta")
var Style elementBuilder = makeElementBuilder("style")
var Title elementBuilder = makeElementBuilder("title")

// Structural Tags
var Article elementBuilder = makeElementBuilder("article")
var Aside elementBuilder = makeElementBuilder("aside")
var BR elementBuilder = makeElementBuilder("br")
var Details elementBuilder = makeElementBuilder("details")
var Div elementBuilder = makeElementBuilder("div")
var H1 elementBuilder = makeElementBuilder("h1")
var H2 elementBuilder = makeElementBuilder("h2")
var H3 elementBuilder = makeElementBuilder("h3")
var H4 elementBuilder = makeElementBuilder("h4")
var H5 elementBuilder = makeElementBuilder("h5")
var H6 elementBuilder = makeElementBuilder("h6")
var Header elementBuilder = makeElementBuilder("header")
var HGroup elementBuilder = makeElementBuilder("hgroup")
var HR elementBuilder = makeElementBuilder("hr")
var Footer elementBuilder = makeElementBuilder("footer")
var Nav elementBuilder = makeElementBuilder("nav")
var P elementBuilder = makeElementBuilder("p")
var Section elementBuilder = makeElementBuilder("section")
var Span elementBuilder = makeElementBuilder("span")
var Summary elementBuilder = makeElementBuilder("summary")

// Form Tags
var Button elementBuilder = makeElementBuilder("button")
var DataList elementBuilder = makeElementBuilder("datalist")
var FieldSet elementBuilder = makeElementBuilder("fieldset")
var Form elementBuilder = makeElementBuilder("form")
var Input elementBuilder = makeElementBuilder("input")
var Keygen elementBuilder = makeElementBuilder("keygen")
var Label elementBuilder = makeElementBuilder("label")
var Legend elementBuilder = makeElementBuilder("legend")
var Meter elementBuilder = makeElementBuilder("meter")
var OptGroup elementBuilder = makeElementBuilder("optgroup")
var Option elementBuilder = makeElementBuilder("option")
var Select elementBuilder = makeElementBuilder("select")
var TextArea elementBuilder = makeElementBuilder("textarea")

// Formatting Tags
var Abbr elementBuilder = makeElementBuilder("abbr")
var Acronym elementBuilder = makeElementBuilder("acronym")
var Address elementBuilder = makeElementBuilder("address")
var B elementBuilder = makeElementBuilder("b")
var Bdi elementBuilder = makeElementBuilder("bdi")
var Bdo elementBuilder = makeElementBuilder("bdo")
var Big elementBuilder = makeElementBuilder("big")
var BlockQuote elementBuilder = makeElementBuilder("blockquote")
var Cite elementBuilder = makeElementBuilder("cite")
var Code elementBuilder = makeElementBuilder("code")
var Del elementBuilder = makeElementBuilder("del")
var Dfn elementBuilder = makeElementBuilder("dfn")
var Em elementBuilder = makeElementBuilder("em")
var I elementBuilder = makeElementBuilder("i")
var Ins elementBuilder = makeElementBuilder("ins")
var Kbd elementBuilder = makeElementBuilder("kbd")
var Mark elementBuilder = makeElementBuilder("mark")
var Output elementBuilder = makeElementBuilder("output")
var Pre elementBuilder = makeElementBuilder("pre")
var Progress elementBuilder = makeElementBuilder("progress")
var Q elementBuilder = makeElementBuilder("q")
var RP elementBuilder = makeElementBuilder("rp")
var RT elementBuilder = makeElementBuilder("rt")
var Ruby elementBuilder = makeElementBuilder("ruby")
var Samp elementBuilder = makeElementBuilder("samp")
var Small elementBuilder = makeElementBuilder("small")
var Strong elementBuilder = makeElementBuilder("strong")
var Sub elementBuilder = makeElementBuilder("sub")
var Sup elementBuilder = makeElementBuilder("sup")
var TT elementBuilder = makeElementBuilder("tt")
var Var elementBuilder = makeElementBuilder("var")
var WBR elementBuilder = makeElementBuilder("wbr")

// List Tags
var DD elementBuilder = makeElementBuilder("dd")
var Dir elementBuilder = makeElementBuilder("dir")
var DL elementBuilder = makeElementBuilder("dl")
var DT elementBuilder = makeElementBuilder("dt")
var LI elementBuilder = makeElementBuilder("li")
var OL elementBuilder = makeElementBuilder("ol")
var Menu elementBuilder = makeElementBuilder("menu")
var UL elementBuilder = makeElementBuilder("ul")

// Table Tags
var Caption elementBuilder = makeElementBuilder("caption")
var Col elementBuilder = makeElementBuilder("col")
var Colgroup elementBuilder = makeElementBuilder("colgroup")
var Table elementBuilder = makeElementBuilder("table")
var TBody elementBuilder = makeElementBuilder("tbody")
var TD elementBuilder = makeElementBuilder("td")
var TFoot elementBuilder = makeElementBuilder("tfoot")
var THead elementBuilder = makeElementBuilder("thead")
var TH elementBuilder = makeElementBuilder("th")
var TR elementBuilder = makeElementBuilder("tr")

// Scripting Tags
var NoScript elementBuilder = makeElementBuilder("noscript")
var Script elementBuilder = makeElementBuilder("script")

// Embedded Content Tags
var Area elementBuilder = makeElementBuilder("area")
var Audio elementBuilder = makeElementBuilder("audio")
var Canvas elementBuilder = makeElementBuilder("canvas")
var Embed elementBuilder = makeElementBuilder("embed")
var FigCaption elementBuilder = makeElementBuilder("figcaption")
var Figure elementBuilder = makeElementBuilder("figure")
var Frame elementBuilder = makeElementBuilder("frame")
var FrameSet elementBuilder = makeElementBuilder("frameset")
var IFrame elementBuilder = makeElementBuilder("iframe")
var Img elementBuilder = makeElementBuilder("img")
var Map elementBuilder = makeElementBuilder("map")
var NoFrames elementBuilder = makeElementBuilder("noframes")
var Object elementBuilder = makeElementBuilder("object")
var Param elementBuilder = makeElementBuilder("param")
var Source elementBuilder = makeElementBuilder("source")
var Time elementBuilder = makeElementBuilder("time")
var Video elementBuilder = makeElementBuilder("video")
