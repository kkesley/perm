package checker

import (
	"os"
	"strconv"
)

//CheckPermissionFromToken returns permission expression from token (convenient but no customization)
func CheckPermissionFromToken(request CheckPermissionFromTokenRequest) (*CheckResponse, error) {
	role := ""
	if len(request.Token.RoleARN) > 0 {
		role = request.Token.RoleARN[0]
	}
	return CheckPermission(CheckPermissionRequest{
		Bucket:   os.Getenv("ROLE_OUTPUT_BUCKET"),
		Region:   os.Getenv("ROLE_OUTPUT_REGION"),
		Role:     role,
		ClientID: strconv.FormatUint(request.Token.ClientID, 10),
		Path:     request.Path,
		Actions:  request.BasicActions,
		Token:    request.Token,
	})
}
