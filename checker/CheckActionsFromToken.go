package checker

import (
	"os"
)

//CheckActionsFromToken retrieve boolean checking from jwt token. (convenient but no customization)
func CheckActionsFromToken(request CheckActionsFromTokenRequest) bool {
	role := ""
	if len(request.Token.RoleARN) > 0 {
		role = request.Token.RoleARN[0]
	}
	if !CheckActions(CheckActionsRequest{
		Bucket:  os.Getenv("ROLE_OUTPUT_BUCKET"),
		Region:  os.Getenv("ROLE_OUTPUT_REGION"),
		Actions: request.Actions,
		Role:    role,
		Token:   request.Token,
	}) {
		return false
	}
	return true
}
