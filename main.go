package elemx

import "fmt"

type HtmlElement struct {
	Tag        string
	Content    string
	Attributes map[string]string
}

type SetAttributes func(*HtmlElement)

func Render(tag string, content string, attrs []SetAttributes) string {
	element := HtmlElement{Tag: tag, Content: content}
	for _, attr := range attrs {
		attr(&element)
	}

	if tag == HTML {
		return fmt.Sprintf("<!DOCTYPE html>%s", element.render())
	}

	return element.render()
}

func (e *HtmlElement) render() string {
	attributes := ""
	for key, value := range e.Attributes {
		attributes += fmt.Sprintf(` %s="%s"`, key, value)
	}
	return fmt.Sprintf("<%s%s>%s</%s>", e.Tag, attributes, e.Content, e.Tag)
}

func (e *HtmlElement) SetAttribute(key, value string) {
	if e.Attributes == nil {
		e.Attributes = make(map[string]string)
	}
	e.Attributes[key] = value
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
	return Script("", ApplyHtmxCDNSource())
}

func Html(c string, attrs ...SetAttributes) string   { return Render(HTML, c, attrs) }
func Body(c string, attrs ...SetAttributes) string   { return Render(BODY, c, attrs) }
func Div(c string, attrs ...SetAttributes) string    { return Render(DIV, c, attrs) }
func Script(c string, attrs ...SetAttributes) string { return Render(SCRIPT, c, attrs) }
func Title(c string, attrs ...SetAttributes) string  { return Render(TITLE, c, attrs) }
func Head(c string, attrs ...SetAttributes) string {
	if len(attrs) == 0 {
		c = c + IncludeHtmx()
	}
	return Render(HEAD, c, attrs)
}
