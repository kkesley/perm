package perm

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kkesley/iteacloud-jwt/jwtidentity"
)

func TestParseActions(test *testing.T) {
	dat, err := ioutil.ReadFile("./sample_action.json")
	if err != nil {
		test.Error(err)
	}
	if resBank, err := ParseActions(ParseActionsRequest{
		PolicyStr: string(dat),
		Token: jwtidentity.TokenRequest{
			IsRoot: true,
		},
	}); err != nil {
		test.Error(err)
	} else {
		fmt.Println(resBank)
	}

}
