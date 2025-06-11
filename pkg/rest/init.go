package rest

import "strings"

var (
	responseJSONPretty       = true
	responseJSONPrettyIndent = "    "
)

func SetResponseJSONPretty(value bool) {
	responseJSONPretty = value
}

func SetResponseJSONPrettyIndent(value int) {
	responseJSONPrettyIndent = strings.Repeat(" ", value)
}
