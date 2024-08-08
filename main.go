package elemx

import (
	"fmt"
	"sort"
	"strconv"
)

type HtmlElement struct {
	tag        string
	content    string
	attributes SetAttr
}

type SetAttributes func(*HtmlElement) string
type SetAttr map[string]string

// var lookUpSetter map[string]func(string) SetAttributes = map[string]func(string) SetAttributes{
// 	LANG: Lang,
// 	SRC:  Src,
// 	REL:  Rel,
// 	HREF: Href,
// }

func Render(tag string, content string, attrs []SetAttr) string {
	element := HtmlElement{tag: tag, content: content}
	if len(attrs) == 0 {
		return element.render(nil)
	}
	return element.render(attrs)
}

func (e *HtmlElement) render(attrs []SetAttr) string {
	var hasAttributesToSet SetAttr
	if attrs != nil {
		hasAttributesToSet = attrs[0]
	} else {
		hasAttributesToSet = nil
	}

	var includeHtmx bool

	if e.tag == HEAD && hasAttributesToSet == nil {
		includeHtmx = true
	}

	if e.tag == HEAD && hasAttributesToSet != nil {
		if excludeHtmx, err := strconv.ParseBool(hasAttributesToSet["excludeHtmx"]); err == nil {
			includeHtmx = !excludeHtmx
		} else {
			includeHtmx = true
		}
		delete(attrs[0], "excludeHtmx")
	}

	if includeHtmx {
		e.content += IncludeHtmx()
	}

	attributes := ""
	if attrs != nil {
		// * This done to stop test from being flaky. I dont care about the order
		attributes_to_set := attrs[0]
		keys := make([]string, 0, len(attributes_to_set))
		for key := range attributes_to_set {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			attributes += fmt.Sprintf(` %s="%s"`, key, attributes_to_set[key])
		}
	}

	if e.tag == HTML {
		return fmt.Sprintf("<!DOCTYPE html><%s%s>%s</%s>", e.tag, attributes, e.content, e.tag)
	}

	if e.isSelfClosing() {
		return fmt.Sprintf("<%s%s />", e.tag, attributes)
	}

	return fmt.Sprintf("<%s%s>%s</%s>", e.tag, attributes, e.content, e.tag)
}

func (e *HtmlElement) isSelfClosing() bool {
	for _, tag := range []string{LINK} {
		if e.tag == tag {
			return true
		}
	}
	return false
}

func (e *HtmlElement) SetAttribute(key, value string) {
	if e.attributes == nil {
		e.attributes = make(map[string]string)
	}
	e.attributes[key] = value
}

// func Lang(lang string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(LANG, lang) }
// }

// func Src(source string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(SRC, source) }
// }

func IncludeHtmx() string {
	return Script(NO_CONTENT, SetAttr{SRC: HTMX_CDN_SOURCE})
}

// func Rel(rel string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(REL, rel) }
// }

// func Href(href string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(HREF, href) }
// }

