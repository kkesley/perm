package perm

import (
	"fmt"
	"strings"

	parser "github.com/kkesley/s3-parser"
)

//GetActions get actions of a user in s3
//Requires ROLE_OUTPUT_REGION, ROLE_OUTPUT_BUCKET
func GetActions(bucket string, region string, role string) []string {
	roleStr := ""
	if roleByte, err := parser.GetS3DocumentDefault(region, bucket, strings.Replace(role, "::", "_", -1)+"/action__role.json"); err == nil {
		roleStr = string(roleByte)
	} else {
		fmt.Println(err)
	}
	actions, err := ParseActions(ParseActionsRequest{
		PolicyStr: roleStr,
	})
	if err != nil {
		return make([]string, 0)
	}
	return actions
}
