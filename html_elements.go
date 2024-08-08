package elemx

func Body(c string, attrs ...SetAttr) string   { return Render(HTML_BODY_TAG, c, attrs) }
func Div(c string, attrs ...SetAttr) string    { return Render(HTML_DIV_TAG, c, attrs) }
func Head(c string, attrs ...SetAttr) string   { return Render(HTML_HEAD_TAG, c, attrs) }
func Html(c string, attrs ...SetAttr) string   { return Render(HTML_HTML_TAG, c, attrs) }
func Link(attrs ...SetAttr) string             { return Render(HTML_LINK_TAG, NO_CONTENT, attrs) }
func P(c string, attrs ...SetAttr) string      { return Render(HTML_P_TAG, c, attrs) }
func Script(c string, attrs ...SetAttr) string { return Render(HTML_SCRIPT_TAG, c, attrs) }
func Title(c string, attrs ...SetAttr) string  { return Render(HTML_TITLE_TAG, c, attrs) }
