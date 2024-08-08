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
	return func(e *HtmlElement) {
		e.SetAttribute("lang", "en")
	}
}

func Html(c string, attrs ...SetAttributes) string { return Render("html", c, attrs) }
