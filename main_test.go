package elemx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description     string
		expected        string
		renderedElement func() string
	}{
		{
			description:     "Render Empty Element",
			expected:        "<html></html>",
			renderedElement: func() string { return Html("") },
		},
		{
			description:     "Render Element with Content",
			expected:        "<html>Hey young world</html>",
			renderedElement: func() string { return Html("Hey young world") },
		},
		{
			description:     "Render Element with Attribute",
			expected:        "<html lang=\"en\"></html>",
			renderedElement: func() string { return Html("", Lang("en")) },
		},
		{
			description:     "Render Nested Element",
			expected:        "<html><body></body></html>",
			renderedElement: func() string { return Html(Body("")) },
		},
	}

	for _, test := range tests {
		assert.Equal(test.renderedElement(), test.expected, test.description)
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
