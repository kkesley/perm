package perm

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

func TestCheckActions(test *testing.T) {
	dat, err := ioutil.ReadFile("./sample_action.json")
	if err != nil {
		test.Error(err)
	}
	if resBank, err := CheckActions(CheckRequest{
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
		fmt.Println(resBank)
	}

}
