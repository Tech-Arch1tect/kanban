package taskquery

import (
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

var registry = map[string]TokenHandler{}

func RegisterTokenHandler(token string, handler TokenHandler) {
	registry[strings.ToLower(token)] = handler
}

func init() {
	RegisterTokenHandler("status", func(tokenValue string, ctx *QueryContext) {
		parts := strings.Split(tokenValue, "|")
		for _, part := range parts {
			ctx.Statuses = append(ctx.Statuses, strings.ToLower(strings.TrimSpace(part)))
		}
	})
	RegisterTokenHandler("assignee", func(tokenValue string, ctx *QueryContext) {
		ctx.Assignee = strings.Trim(tokenValue, "\"")
	})
	RegisterTokenHandler("creator", func(tokenValue string, ctx *QueryContext) {
		ctx.Creator = strings.Trim(tokenValue, "\"")
	})
	RegisterTokenHandler("title", func(tokenValue string, ctx *QueryContext) {
		ctx.Title = strings.Trim(tokenValue, "\"")
	})
}
