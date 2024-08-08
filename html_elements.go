package elemx

func Body(c string, attrs ...SetAttr) string   { return Render(BODY, c, attrs) }
func Div(c string, attrs ...SetAttr) string    { return Render(DIV, c, attrs) }
func Head(c string, attrs ...SetAttr) string   { return Render(HEAD, c, attrs) }
func Html(c string, attrs ...SetAttr) string   { return Render(HTML, c, attrs) }
func Link(attrs ...SetAttr) string             { return Render(LINK, NO_CONTENT, attrs) }
func P(c string, attrs ...SetAttr) string      { return Render(P_TAG, c, attrs) }
func Script(c string, attrs ...SetAttr) string { return Render(SCRIPT, c, attrs) }
func Title(c string, attrs ...SetAttr) string  { return Render(TITLE, c, attrs) }
