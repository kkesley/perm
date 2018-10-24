package perm

import (
	funk "github.com/thoas/go-funk"
)

//PermissionSingleQueryChecker build logic of permission checker for single operations (show, update, delete)
func (permission CheckResponse) PermissionSingleQueryChecker(targetID string, conditionIDs []string) bool {
	if permission.Allow.All { //if permission has allow all
		if funk.ContainsString(permission.Deny.Resources, targetID) {
			return false //if deny resources contains the target, deny the permission
		}
		requireCondition := false
		for _, condition := range permission.Allow.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}
		if requireCondition && !funk.ContainsString(funk.UniqString(append(conditionIDs, permission.Allow.Resources...)), targetID) {
			return false //if require condition & the condition and allowed resources does not contain the target, deny
		}
		return true //else allow the request
	} else if permission.Deny.All { //if permission has deny all
		if funk.ContainsString(permission.Allow.Resources, targetID) {
			return true // if allow resources contain the target, allow the permission
		}
		requireCondition := false
		for _, condition := range permission.Deny.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition && !funk.ContainsString(funk.UniqString(append(conditionIDs, permission.Deny.Resources...)), targetID) {
			return true //if require condition & condition and denied resources does not contain the target, allow
		}
		return false //else deny the request
	} else { //if has individual allow/deny permission
		if len(permission.Deny.Resources) > 0 && funk.ContainsString(permission.Deny.Resources, targetID) {
			return false //if deny resources contain the target, deny
		}
		if len(permission.Allow.Resources) > 0 && funk.ContainsString(permission.Allow.Resources, targetID) {
			return true //if allow resources contain the target, allow
		}
	}
	return false //if empty permission, deny
}
