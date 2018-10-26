package checker

import (
	"fmt"
	"testing"
)

func TestCheckActions(test *testing.T) {
	res := CheckActions(CheckActionsRequest{
		Role:    "arn::itea::1::platform::role::3",
		Actions: []string{"itea::platform::user::read"},
		Bucket:  "iteacloud-platform-api-dev-ap-southeast-1-role",
		Region:  "ap-southeast-1",
	})
	fmt.Println(res)
}
