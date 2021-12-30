package requests

import "encoding/json"

type SetCachceProperties struct {
	Mode       string `json:"mode,omitempty"`
	MaxResults int    `json:"maxResults,omitempty"`
}

func (r *SetCachceProperties) Path() string {
	return "/_api/query-cache/properties"
}

func (r *SetCachceProperties) Method() string {
	return "PUT"
}

func (r *SetCachceProperties) Generate() []byte {
	m, _ := json.Marshal(r)
	return m
}

type GetCacheProperties struct{}
