package perm

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
)

//PermissionSingleQueryCheckerAsync build logic of permission checker for single operations (show, update, delete) in async mode
func (permission CheckResponse) PermissionSingleQueryCheckerAsync(targetID string, conditionIDs []string, wg *sync.WaitGroup, handleData chan<- *string) {
	defer wg.Done()
	if permission.PermissionSingleQueryChecker(targetID, conditionIDs) {
		handleData <- aws.String(targetID)
	} else {
		handleData <- nil
	}
}
