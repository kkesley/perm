package perm

//Action holds the action specific permission
type Action struct {
	AllowedAction map[string][]string `json:"allow"`
	DeniedAction  map[string][]string `json:"deny"`
}
