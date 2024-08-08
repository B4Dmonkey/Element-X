package elemx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description string
		expected    string
		result      string
	}{
		{
			description: "Render Empty Element",
			expected:    "<!DOCTYPE html><html></html>",
			result:      Html(""),
		},
		{
			description: "Render Element with Content",
			expected:    "<div>Hey young world</div>",
			result:      Div("Hey young world"),
		},
		{
			description: "Render Element with Attribute",
			expected:    "<div lang=\"en\"></div>",
			result:      Div("", SetAttr{"lang": "en"}),
		},
		{
			description: "Render Nested Element",
			expected:    "<body><div></div></body>",
			result:      Body(Div("")),
		},
		{
			description: "Render Multiple Elements on the Same Level",
			expected:    "<div>Hey young world</div><div>The world is yours</div>",
			result:      Div("Hey young world") + Div("The world is yours"),
		},
		{
			description: "Render Multiple Elements Nested",
			expected:    "<body><div></div><div></div></body>",
			result:      Body(Div("") + Div("")),
		},
		{
			description: "It renders self closing tags",
			expected:    "<link rel=\"stylesheet\" href=\"styles.css\" />",
			result:      Link(SetAttr{"rel": "stylesheet", "href": "styles.css"}),
		},
	}

	for _, test := range tests {
		assert.Equal(test.expected, test.result, test.description)
	}
}
func TestHtmxCompatibility(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description string
		expected    string
		rendered    string
	}{
		// {
		// 	description: "Include htmx",
		// 	expected:    "<script src=\"https://unpkg.com/htmx.org@2.0.1\"></script>",
		// 	rendered:    Script("", ApplyHtmxCDNSource()),
		// },
		{
			description: "Empty Head Tag should include htmx by default", // ? not sure if this is correct at the moment
			expected:    "<head><script src=\"https://unpkg.com/htmx.org@2.0.1\"></script></head>",
			rendered:    Head(""),
		},
		{
			description: "Head Tag should include htmx by default",
			expected:    "<head><title>ElementX</title><script src=\"https://unpkg.com/htmx.org@2.0.1\"></script></head>",
			rendered:    Head(Title("ElementX")),
		},
		{
			description: "Optionally exclude htmx from Head Tag",
			expected:    "<head><title>ElementX</title><link rel=\"stylesheet\" href=\"styles.css\" /></head>",
			rendered: Head(
				Title("ElementX")+
					Link(SetAttr{"rel": "stylesheet", "href": "styles.css"}),
				SetAttr{"excludeHtmx": "true"},
			),
		},
	}

	for _, test := range tests {
		assert.Equal(test.expected, test.rendered, test.description)
	}
}
func TestSetAttribute(t *testing.T) {
	assert := assert.New(t)

	expected := "<div lang=\"en\"></div>"
	element := HtmlElement{tag: "div"}
	element.SetAttribute("lang", "en")
	result := element.render()

	assert.Equal(expected, result, "Test SetAttribute")

	tests := []struct {
		description string
		expected    string
		attribute   SetAttributes
	}{
		{
			description: "SetAttribute",
			expected:    "<div lang=\"en\"></div>",
			attribute:   Lang("en"),
		},
	}

	for _, test := range tests {
		element := HtmlElement{tag: "div"}
		test.attribute(&element)
		assert.Equal(test.expected, element.render(), test.description)
	}
}
