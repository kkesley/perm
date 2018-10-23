package perm

import "github.com/kkesley/iteacloud-jwt/jwtidentity"

//ParseActionsRequest holds the request for parsing actions
type ParseActionsRequest struct {
	PolicyStr string `json:"policy_str"`
	Token     jwtidentity.TokenRequest
}
