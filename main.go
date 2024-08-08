package elemx

import "fmt"

func Render(element *HtmlElement) string {
	return element.toString()
}

type HtmlElement struct {
	Tag        string
	Content    string
	Attributes map[string]string
}

func (e *HtmlElement) toString() string {
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

type SetElementAttributes func(*HtmlElement)

func Lang(lang string) SetElementAttributes {
	return func(e *HtmlElement) {
		e.SetAttribute("lang", "en")
	}
}

func Html(content string) string { return Render(&HtmlElement{Tag: "html", Content: content} )}
