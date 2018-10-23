package perm

type CheckActionsRequest struct {
	Bucket  string   `json:"bucket"`
	Region  string   `json:"region"`
	Role    string   `json:"role"`
	Actions []string `json:"actions"`
}
