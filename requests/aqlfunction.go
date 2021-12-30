package requests

import "encoding/json"

type CreateAQLFunction struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	IsDeterministic bool
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

type AQLFunction struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetAQLFunctionsResult []AQLFunction
