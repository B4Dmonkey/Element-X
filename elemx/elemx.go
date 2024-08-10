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

type Element func(attrs SetAttr, c ...string) string

func renderElementFunc(tag string, hasContent bool) Element {
	if hasContent {
		return func(attrs SetAttr, c ...string) string {
			return Render(tag, attrs, c)
		}
	} else {
		return func(attrs SetAttr, c ...string) string {
			return Render(tag, attrs, nil)
		}
	}
}

func Render(tag string, attrs SetAttr, content []string) string {
	element := HtmlElement{tag: tag, content: content[0]}
	if len(attrs) == 0 {
		return element.render(nil)
	}
	return element.render(attrs)
}

func (e *HtmlElement) render(attrs SetAttr) string {
	var includeHtmx bool

	if e.tag == HTML_HEAD_TAG && attrs == nil {
		includeHtmx = true
	}

	if e.tag == HTML_HEAD_TAG && attrs != nil {
		if excludeHtmx, err := strconv.ParseBool(attrs["excludeHtmx"]); err == nil {
			includeHtmx = !excludeHtmx
		} else {
			includeHtmx = true
		}
		delete(attrs, "excludeHtmx")
	}

	if includeHtmx {
		e.content += Script(SetAttr{SRC: HTMX_CDN_SOURCE}, NO_CONTENT)
	}

	attributes := ""
	if attrs != nil {
		// * This done to stop test from being flaky. I dont care about the order
		attributes_to_set := attrs
		keys := make([]string, 0, len(attributes_to_set))
		for key := range attributes_to_set {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			attributes += fmt.Sprintf(` %s="%s"`, key, attributes_to_set[key])
		}
	}

	if e.tag == HTML_HTML_TAG {
		return fmt.Sprintf("<!DOCTYPE html><%s%s>%s</%s>", e.tag, attributes, e.content, e.tag)
	}

	if e.isSelfClosing() {
		return fmt.Sprintf("<%s%s />", e.tag, attributes)
	}

	return fmt.Sprintf("<%s%s>%s</%s>", e.tag, attributes, e.content, e.tag)
}

func (e *HtmlElement) isSelfClosing() bool {
	for _, tag := range []string{HTML_LINK_TAG} {
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

// ? There may be an opportunity to do type safe attribute setting
// var lookUpSetter map[string]func(string) SetAttributes = map[string]func(string) SetAttributes{
// 	LANG: Lang,
// 	SRC:  Src,
// 	REL:  Rel,
// 	HREF: Href,
// }

// func Lang(lang string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(LANG, lang) }
// }

// func Src(source string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(SRC, source) }
// }

// func Rel(rel string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(REL, rel) }
// }

// func Href(href string) SetAttributes {
// 	return func(e *HtmlElement) { e.SetAttribute(HREF, href) }
// }
