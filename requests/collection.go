package requests

type CreateCollection struct {
	JournalSize    int                    `json:"journalSize,omitempty"`
	KeyOptions     map[string]interface{} `json:"keyOptions,omitempty"`
	Name           string                 `json:"name"`
	WaitForSync    bool                   `json:"waitForSync,omitempty"`
	DoCompact      bool                   `json:"doCompact,omitempty"`
	IsVolatile     bool                   `json:"isVolatile,omitempty"`
	ShardKeys      []string               `json:"shardKeys,omitempty"`
	NumberOfShards int                    `json:"numberOfShards,omitempty"`
	IsSystem       bool                   `json:"isSystem,omitempty"`
	Type           int                    `json:"type,omitempty"`
	IndexBuckets   int                    `json:"indexBuckets,omitempty"`
}

func (r *CreateCollection) Path() string {
	return "/_api/collection"
}

func (r *CreateCollection) Method() string {
	return "POST"
}
