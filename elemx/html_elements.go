package elemx

// ? This creates a cyclic dependency when following the pattern below
func Script(c string, attrs ...SetAttr) string { return Render(HTML_SCRIPT_TAG, c, attrs) }

var (
	A      = renderElementFunc(HTML_A_TAG, true)
	Body   = renderElementFunc(HTML_BODY_TAG, true)
	Button = renderElementFunc(HTML_BUTTON_TAG, true)
	Div    = renderElementFunc(HTML_DIV_TAG, true)
	Form   = renderElementFunc(HTML_FORM_TAG, true)
	H1     = renderElementFunc(HTML_H1_TAG, true)
	H2     = renderElementFunc(HTML_H2_TAG, true)
	H3     = renderElementFunc(HTML_H3_TAG, true)
	H4     = renderElementFunc(HTML_H4_TAG, true)
	H5     = renderElementFunc(HTML_H5_TAG, true)
	H6     = renderElementFunc(HTML_H6_TAG, true)
	Head   = renderElementFunc(HTML_HEAD_TAG, true)
	Html   = renderElementFunc(HTML_HTML_TAG, true)
	Img    = renderElementFunc(HTML_IMG_TAG, false)
	Input  = renderElementFunc(HTML_INPUT_TAG, false)
	Label  = renderElementFunc(HTML_LABEL_TAG, true)
	Li     = renderElementFunc(HTML_LI_TAG, true)
	// Link   = renderElementFunc(HTML_LINK_TAG, false)
	Ol     = renderElementFunc(HTML_OL_TAG, true)
	P      = renderElementFunc(HTML_P_TAG, true)
	Span   = renderElementFunc(HTML_SPAN_TAG, true)
	Table  = renderElementFunc(HTML_TABLE_TAG, true)
	Td     = renderElementFunc(HTML_TD_TAG, true)
	Title  = renderElementFunc(HTML_TITLE_TAG, true)
	Tr     = renderElementFunc(HTML_TR_TAG, true)
	Ul     = renderElementFunc(HTML_UL_TAG, true)
)
