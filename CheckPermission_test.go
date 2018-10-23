package perm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

func TestCheckPermission(test *testing.T) {
	dat, err := ioutil.ReadFile("./sample_input.json")
	if err != nil {
		test.Error(err)
	}
	if resBank, err := CheckPermission(CheckPermissionRequest{
		PolicyStr: string(dat),
		ClientID:  "1",
		Path:      "itea::platform::user",
		Actions:   []string{"read"},
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
