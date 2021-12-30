package requests

import (
	"bytes"
	"fmt"
	"strings"
)

type AQL struct {
	query     string
	bindVars  map[string]interface{}
	cache     *bool
	batchSize int
}

func NewAQL(query string, params ...interface{}) *AQL {
	return &AQL{query: processAQL(fmt.Sprintf(query, params...))}
}

func processAQL(query string) string {
	buf := bytes.NewBuffer(nil)
	space := false
	for _, c := range query {
		switch c {
		case ' ', '\n', '\t':
			if !space {
				buf.WriteRune(' ')
				space = true
			}
		case '"':
			buf.WriteRune('\'')
		default:
			buf.WriteRune(c)
			if space {
				space = false
			}
		}
	}
	return strings.TrimSpace(buf.String())
}
