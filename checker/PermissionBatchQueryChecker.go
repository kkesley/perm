package checker

import (
	"sync"
)

//PermissionBatchQueryChecker build logic of permission checker for batch operations (batch delete / batch update)
func (permission CheckResponse) PermissionBatchQueryChecker(targetIDs []string, filterOutput FilterIDsOutput) []string {
	idChannel := make(chan *string, len(targetIDs))
	var wg sync.WaitGroup
	wg.Add(len(targetIDs))
	for _, targetID := range targetIDs {
		go permission.PermissionSingleQueryCheckerAsync(targetID, filterOutput, &wg, idChannel)
	}
	wg.Wait()
	close(idChannel)

	ids := make([]string, 0)
	for id := range idChannel {
		if id != nil && len(*id) > 0 {
			ids = append(ids, *id)
		}
	}
	return ids
}
