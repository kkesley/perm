package perm

import (
	"fmt"
	"testing"
)

func TestUploadRole(test *testing.T) {
	err := UploadRole("iteacloud-platform-api-role-dev-ap-southeast-1-input", "ap-southeast-1", "test", []RawPolicy{
		RawPolicy{
			Effect: "Allow",
			Resources: []string{
				"test",
			},
			Actions: []string{
				"test_action",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		test.Error(err)
	}
}
