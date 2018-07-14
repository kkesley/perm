package perm

import (
	"fmt"
	"testing"
)

func TestDownloadRole(test *testing.T) {
	policies, err := DownloadRole("iteacloud-platform-api-role-dev-ap-southeast-1-input", "ap-southeast-1", "test")
	if err != nil {
		fmt.Println(err)
		test.Error(err)
	}
	fmt.Println(policies)
}
