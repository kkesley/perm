package checker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

func TestCheckPermission(test *testing.T) {
	// dat, err := ioutil.ReadFile("./sample_input.json")
	// if err != nil {
	// 	test.Error(err)
	// }
	if resBank, err := CheckPermission(CheckPermissionRequest{
		Bucket:   "iteacloud-platform-api-dev-ap-southeast-1-role",
		Region:   "ap-southeast-1",
		Role:     "arn::itea::1::platform::role::3",
		ClientID: "1",
		Path:     "itea::platform::user",
		Actions:  []string{"read"},
		Token: jwtidentity.TokenRequest{
			IsRoot: false,
		},
	}); err != nil {
		test.Error(err)
	} else {
		marByte, err := json.Marshal(resBank)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(marByte))
	}
}
