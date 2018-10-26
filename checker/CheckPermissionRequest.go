package checker

import "github.com/kkesley/iteacloud-jwt/jwtidentity"

//CheckPermissionRequest holds the request for checking eligible permissions
type CheckPermissionRequest struct {
	Bucket    string   `json:"bucket"`
	Region    string   `json:"region"`
	Role      string   `json:"role"`
	PolicyStr *string  `json:"policy_str"`
	ClientID  string   `json:"client_id"`
	Path      string   `json:"path"`
	Actions   []string `json:"actions"`
	Token     jwtidentity.TokenRequest
}
