package checker

import (
	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

//CheckActionsFromTokenRequest holds the request to CheckActionsFromToken
type CheckActionsFromTokenRequest struct {
	Token   jwtidentity.TokenRequest
	Actions []string
}
