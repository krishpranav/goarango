package requests

import (
	"bytes"
	"encoding/json"
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

func (a *AQL) Cache(enable bool) *AQL {
	a.cache = &enable
	return a
}

func (a *AQL) BatchSize(size int) *AQL {
	a.batchSize = size
	return a
}

func (a *AQL) Bind(name string, value interface{}) *AQL {
	if a.bindVars == nil {
		a.bindVars = make(map[string]interface{})
	}
	a.bindVars[name] = value
	return a
}

func (a *AQL) Path() string {
	return "/_api/cursor"
}

func (a *AQL) Method() string {
	return "POST"
}

func (a *AQL) Generate() []byte {
	type AQLFmt struct {
		Query     string                 `json:"query"`
		BindVars  map[string]interface{} `json:"bindVars,omitempty"`
		Cache     *bool                  `json:"cache,omitempty"`
		BatchSize int                    `json:"batchSize,omitempty"`
	}

	jsonAQL, _ := json.Marshal(&AQLFmt{Query: a.query, BindVars: a.bindVars, Cache: a.cache, BatchSize: a.batchSize})

	return jsonAQL
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
