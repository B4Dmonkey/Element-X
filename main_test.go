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
			result:      Div("", Lang("en")),
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
		{
			description: "Include htmx",
			expected:    "<script src=\"https://unpkg.com/htmx.org@2.0.1\"></script>",
			rendered:    Script("", ApplyHtmxCDNSource()),
		},
		{
			description: "Empty Head Tag default include htmx", // ? not sure if this is correct at the moment
			expected:    "<head><script src=\"https://unpkg.com/htmx.org@2.0.1\"></script></head>",
			rendered:    Head(""),
		},
		{
			description: "Empty Head Tag without htmx",
			expected:    "<head><title>ElementX</title><script src=\"https://unpkg.com/htmx.org@2.0.1\"></script></head>",
			rendered:    Head(Title("ElementX")),
		},
	}

	for _, test := range tests {
		assert.Equal(test.expected, test.rendered, test.description)
	}
}
func TestSetAttribute(t *testing.T) {
	assert := assert.New(t)

	expected := "<html lang=\"en\"></html>"
	element := HtmlElement{Tag: "html"}
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
			expected:    "<html lang=\"en\"></html>",
			attribute:   Lang("en"),
		},
	}

	for _, test := range tests {
		element := HtmlElement{Tag: "html"}
		test.attribute(&element)
		assert.Equal(test.expected, element.render(), test.description)
	}
}
