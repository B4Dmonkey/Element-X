package elemx

import (
	"fmt"
)

type GlobalAttributes struct {
	Lang string //The lang attribute specifies the language of the element's content.
}

type HtmlAttributes struct {
	GlobalAttributes
	Xmls string
}
type HtmlElement interface {
	Render() string
}

func Render(el HtmlElement) string {
	return el.Render()
}
type Html struct {
	Content string
	Attrs *HtmlAttributes
}

func (h Html) Render() string {
	template := "<!DOCTYPE html><html>%s</html>"
	return fmt.Sprintf(template, h.Content)
}


