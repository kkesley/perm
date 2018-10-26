package checker

import (
	"strconv"

	"github.com/kkesley/arn"
	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

//CheckPrivatePermission check if a user can access a private resource.
//They can access the resource if the resource is themselves
//They can access the resource if the resource is their client only if they are the root
//They can access the resource if the resource is their peers only if they are the root
func CheckPrivatePermission(ARN string, token jwtidentity.TokenRequest, allowPeer bool, allowMember bool) bool {
	if token.ClientARN == ARN && !token.IsRoot && !allowMember {
		return false
	} else if token.UserARN != ARN && token.ClientARN != ARN && !token.IsRoot && !allowPeer {
		return false
	} else if arn.GetPartResourceID(ARN, "itea") != strconv.FormatUint(token.ClientID, 10) {
		return false
	}
	return true
}
