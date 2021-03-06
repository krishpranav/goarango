package requests

import (
	"encoding/json"
)

type CreateAQLFunction struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	IsDeterministic bool   `json:"isDeterministic"`
}

func (r *CreateAQLFunction) Path() string {
	return "/_api/aqlfunction"
}

func (r *CreateAQLFunction) Method() string {
	return "POST"
}

func (r *CreateAQLFunction) Generate() []byte {
	m, _ := json.Marshal(r)
	return m
}

type DeleteAQLFunction struct {
	Name  string
	Group bool
}

func (r *DeleteAQLFunction) Path() string {
	path := "/_api/aqlfunction/" + r.Name

	if r.Group {
		path += "?group=true"
	}

	return path
}

func (r *DeleteAQLFunction) Method() string {
	return "DELETE"
}

func (r *DeleteAQLFunction) Generate() []byte {
	return nil
}

type GetAQLFunctions struct {
	Namespace string
}

func (r *GetAQLFunctions) Path() string {
	path := "/_api/aqlfunction"

	if r.Namespace != "" {
		path += "?namespace=" + r.Namespace
	}

	return path
}

func (r *GetAQLFunctions) Method() string {
	return "GET"
}

func (r *GetAQLFunctions) Generate() []byte {
	return nil
}

type AQLFunction struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetAQLFunctionsResult []AQLFunction
