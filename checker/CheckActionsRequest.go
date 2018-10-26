package checker

import "github.com/kkesley/iteacloud-jwt/jwtidentity"

//CheckActionsRequest holds the request to check valid actions
type CheckActionsRequest struct {
	Bucket  string   `json:"bucket"`
	Region  string   `json:"region"`
	Role    string   `json:"role"`
	Actions []string `json:"actions"`
	Token   jwtidentity.TokenRequest
}
