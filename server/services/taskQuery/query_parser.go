package taskquery

import (
	"regexp"
	"strings"
)

type QueryParser struct {
	regex *regexp.Regexp
}

func NewQueryParser() *QueryParser {
	re := regexp.MustCompile(`\S+:"[^"]+"|\S+`)
	return &QueryParser{regex: re}
}

func (qp *QueryParser) Parse(q string, boardID uint) *QueryContext {
	ctx := &QueryContext{BoardID: boardID}
	tokens := qp.regex.FindAllString(q, -1)
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		if colon := strings.Index(token, ":"); colon > 0 {
			key := strings.ToLower(token[:colon])
			value := token[colon+1:]
			value = strings.Trim(value, "\"")
			if handler, ok := registry[key]; ok {
				handler(value, ctx)
				continue
			}
		}
		ctx.FreeText = append(ctx.FreeText, token)
	}
	return ctx
}
