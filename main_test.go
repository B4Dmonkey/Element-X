package elemx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	assert := assert.New(t)
	expected := "<html></html>"
	result := Html("")
	assert.Equal(expected, result, "Test Render")
}

func TestSetAttribute(t *testing.T) {
	assert := assert.New(t)

	expected := "<html lang=\"en\"></html>"
	element := HtmlElement{Tag: "html"}
	element.SetAttribute("lang", "en")
	result := Render(&element)

	assert.Equal(expected, result, "Test SetAttribute")

	tests:= []struct {
		description string
		expected string
		attribute SetElementAttributes
	}{
		{
			description: "SetAttribute",
			expected: "<html lang=\"en\"></html>",
			attribute:  Lang("en"),
		},
	}

	for _, test := range tests {
		element := HtmlElement{Tag: "html"}
		test.attribute(&element)
		assert.Equal(test.expected, Render(&element), test.description)
	}
}
