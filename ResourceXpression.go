package perm

type ResourceXpression struct {
	All        bool                `json:"all"`
	Self       bool                `json:"self"`
	Owned      bool                `json:"owned"`
	Resources  []string            `json:"resources"`
	Conditions []map[string]string `json:"conditions"`
}
