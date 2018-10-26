package checker

import (
	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

//CheckPermissionFromTokenRequest holds the request to CheckPermissionFromToken
type CheckPermissionFromTokenRequest struct {
	Token        jwtidentity.TokenRequest
	Path         string
	BasicActions []string
}
