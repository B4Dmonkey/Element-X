package elemx

import (
	"fmt"
)
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

func Head(el string) string {
	return fmt.Sprintf("<head>%s</head>", el)
}

func Body(el string) string {
	return fmt.Sprintf("<body>%s</body>", el)
}

func H1(el string) string {
	return fmt.Sprintf("<h1>%s</h1>", el)
}
