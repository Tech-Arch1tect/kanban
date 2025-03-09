package taskquery

import (
	"regexp"
	"strings"
)

type QueryContext struct {
	BoardID  uint
	Statuses []string
	Assignee string
	Creator  string
	Title    string
	FreeText []string
}

type TokenHandler func(tokenValue string, ctx *QueryContext)

type QueryParser struct {
	regex    *regexp.Regexp
	registry map[string]TokenHandler
}

func NewQueryParser() *QueryParser {
	re := regexp.MustCompile(`\S+:"[^"]+"|\S+`)
	qp := &QueryParser{
		regex:    re,
		registry: make(map[string]TokenHandler),
	}
	qp.RegisterTokenHandler("status", func(tokenValue string, ctx *QueryContext) {
		parts := strings.Split(tokenValue, "|")
		for _, part := range parts {
			ctx.Statuses = append(ctx.Statuses, strings.ToLower(strings.TrimSpace(part)))
		}
	})
	qp.RegisterTokenHandler("assignee", func(tokenValue string, ctx *QueryContext) {
		ctx.Assignee = strings.Trim(tokenValue, "\"")
	})
	qp.RegisterTokenHandler("creator", func(tokenValue string, ctx *QueryContext) {
		ctx.Creator = strings.Trim(tokenValue, "\"")
	})
	qp.RegisterTokenHandler("title", func(tokenValue string, ctx *QueryContext) {
		ctx.Title = strings.Trim(tokenValue, "\"")
	})
	return qp
}

func (qp *QueryParser) RegisterTokenHandler(token string, handler TokenHandler) {
	qp.registry[strings.ToLower(token)] = handler
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
			if handler, ok := qp.registry[key]; ok {
				handler(value, ctx)
				continue
			}
		}
		ctx.FreeText = append(ctx.FreeText, token)
	}
	return ctx
}
