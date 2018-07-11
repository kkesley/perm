package perm

//CheckRequest holds the request for checking eligible permissions
type CheckRequest struct {
	PolicyStr string   `json:"policy_str"`
	ClientID  string   `json:"client_id"`
	Path      string   `json:"path"`
	Actions   []string `json:"actions"`
}
