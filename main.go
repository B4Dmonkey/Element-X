package elemx

import (
	"fmt"
	// "reflect"
)

type HtmlElement struct {
	tag        string
	content    string
	attributes SetAttr
}

type SetAttributes func(*HtmlElement)
type SetAttr map[string]string

var lookUpSetter map[string]func(string) SetAttributes = map[string]func(string) SetAttributes{
	LANG: Lang,
	SRC:  Src,
	REL:  Rel,
	HREF: Href,
}

func Render(tag string, content string, attrs []SetAttr) string {
	element := HtmlElement{tag: tag, content: content}
	if len(attrs) == 0 {
		return element.render()
	}

	attributes := attrs[0]
	for attr, value := range attributes {
		setter := lookUpSetter[attr]
		if setter != nil {
			setter(value)(&element)
		} else {
			element.SetAttribute(attr, value)
		}
	}
	return element.render()
}

func (e *HtmlElement) render() string {
	attributes := ""
	for key, value := range e.attributes {
		attributes += fmt.Sprintf(` %s="%s"`, key, value)
	}

	if e.tag == HTML {
		return fmt.Sprintf("<!DOCTYPE html><%s%s>%s</%s>", e.tag, attributes, e.content, e.tag)
	}

	if e.isSelfClosing() {
		return fmt.Sprintf("<%s%s />", e.tag, attributes)
	}

	return fmt.Sprintf("<%s%s>%s</%s>", e.tag, attributes, e.content, e.tag)
}

func (e *HtmlElement) isSelfClosing() bool {
	for _, tag := range []string{LINK} {
		if e.tag == tag {
			return true
		}
	}
	return false
}

func (e *HtmlElement) SetAttribute(key, value string) {
	if e.attributes == nil {
		e.attributes = make(map[string]string)
	}
	e.attributes[key] = value
}

func Lang(lang string) SetAttributes {
	return func(e *HtmlElement) { e.SetAttribute(LANG, lang) }
}

func Src(source string) SetAttributes {
	return func(e *HtmlElement) { e.SetAttribute(SRC, source) }
}

func ApplyHtmxCDNSource() SetAttributes {
	return func(e *HtmlElement) { e.SetAttribute(SRC, HTMX_CDN_SOURCE) }
}

func IncludeHtmx() string {
	return Script(NO_CONTENT, SetAttr{SRC: HTMX_CDN_SOURCE})
}

func ExcludeHtmx() SetAttributes {
	return func(e *HtmlElement) { delete(e.attributes, SRC) }
}

func Rel(rel string) SetAttributes {
	return func(e *HtmlElement) { e.SetAttribute(REL, rel) }
}

func Href(href string) SetAttributes {
	return func(e *HtmlElement) { e.SetAttribute(HREF, href) }
}

func P(c string, attrs ...SetAttr) string      { return Render("p", c, attrs) }
func Html(c string, attrs ...SetAttr) string   { return Render(HTML, c, attrs) }
func Body(c string, attrs ...SetAttr) string   { return Render(BODY, c, attrs) }
func Div(c string, attrs ...SetAttr) string    { return Render(DIV, c, attrs) }
func Script(c string, attrs ...SetAttr) string { return Render(SCRIPT, c, attrs) }
func Title(c string, attrs ...SetAttr) string  { return Render(TITLE, c, attrs) }
func Link(attrs ...SetAttr) string             { return Render(LINK, NO_CONTENT, attrs) }
func Head(c string, attrs ...SetAttr) string {
	if len(attrs) == 0 {
		c = c + IncludeHtmx()
	}

	if len(attrs) > 0 {
		if _, includeHtmx := attrs[0]["excludeHtmx"]; !includeHtmx {
			c = c + IncludeHtmx()
		}
		delete(attrs[0], "excludeHtmx")
	}
	return Render(HEAD, c, attrs)
}
