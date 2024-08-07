package elemx

import (
    "testing"
)

func TestHtml(t *testing.T) {
    expected := "<!DOCTYPE html><html></html>"
    result := Render(Html{Content: ""})

    if result != expected {
        t.Errorf("Html(\"\") = %s; want %s", result, expected)
    }
}