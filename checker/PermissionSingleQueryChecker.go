package checker

import (
	funk "github.com/thoas/go-funk"
)

//PermissionSingleQueryChecker build logic of permission checker for single operations (show, update, delete)
func (permission CheckResponse) PermissionSingleQueryChecker(targetID string, filterOutput FilterIDsOutput) bool {
	denyOwned := make([]string, 0)
	allowOwned := make([]string, 0)
	if permission.Allow.Owned && !permission.Deny.Owned {
		allowOwned = filterOutput.Owned
	} else if permission.Deny.Owned && !permission.Allow.Owned {
		denyOwned = filterOutput.Owned
	}
	if permission.Allow.All { //if permission has allow all
		if funk.ContainsString(append(permission.Deny.Resources, denyOwned...), targetID) {
			return false //if deny resources contains the target, deny the permission
		}
		requireCondition := false
		for _, condition := range permission.Allow.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}
		if requireCondition && !funk.ContainsString(append(append(filterOutput.Conditions, permission.Allow.Resources...), allowOwned...), targetID) {
			return false //if require condition & the condition and allowed resources does not contain the target, deny
		}
		return true //else allow the request
	} else if permission.Deny.All { //if permission has deny all
		if funk.ContainsString(append(permission.Allow.Resources, allowOwned...), targetID) {
			return true // if allow resources contain the target, allow the permission
		}
		requireCondition := false
		for _, condition := range permission.Deny.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition && !funk.ContainsString(append(append(filterOutput.Conditions, permission.Deny.Resources...), denyOwned...), targetID) {
			return true //if require condition & condition and denied resources does not contain the target, allow
		}
		return false //else deny the request
	} else { //if has individual allow/deny permission
		if (len(permission.Deny.Resources) > 0 || len(denyOwned) > 0) && funk.ContainsString(append(denyOwned, permission.Deny.Resources...), targetID) {
			return false //if deny resources contain the target, deny
		}
		if (len(permission.Allow.Resources) > 0 || len(allowOwned) > 0) && funk.ContainsString(append(allowOwned, permission.Allow.Resources...), targetID) {
			return true //if allow resources contain the target, allow
		}
	}
	return false //if empty permission, deny
}
