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
	result := Render(&element)

	assert.Equal(expected, result, "Test SetAttribute")

	tests := []struct {
		description string
		expected    string
		attribute   SetElementAttributes
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
		assert.Equal(test.expected, Render(&element), test.description)
	}
}
