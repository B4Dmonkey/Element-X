package elemx

type GlobalAttributes struct {
	Lang string //The lang attribute specifies the language of the element's content.
}

type HtmlAttributes struct {
	GlobalAttributes
	Xmls string
}
