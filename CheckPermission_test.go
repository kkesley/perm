package perm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCheckPermission(test *testing.T) {
	dat, err := ioutil.ReadFile("./sample_input.json")
	if err != nil {
		test.Error(err)
	}
	if resBank, err := CheckPermission(CheckRequest{
		PolicyStr: string(dat),
		ClientID:  "1",
		Path:      "itea::platform::user",
		Actions:   []string{"read"},
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
