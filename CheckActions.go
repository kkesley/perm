package perm

import (
	"encoding/json"
	"fmt"
)

//CheckActions check an action
func CheckActions(request CheckRequest) (bool, error) {
	if request.Token.IsRoot {
		return true, nil
	} else if len(request.PolicyStr) <= 0 {
		return false, nil
	}
	var action Action
	//unmarshal policy string to a role
	if err := json.Unmarshal([]byte(request.PolicyStr), &action); err != nil {
		fmt.Println(err)
		return false, err
	}
	allowed := false
	allowAll := false
	if permissions, ok := action.AllowedAction[request.Path]; ok {
		for _, permission := range permissions {
			for _, requestedPermission := range request.Actions {
				if requestedPermission == permission || permission == "*" {
					allowed = true
					if permission == "*" {
						allowAll = true
					}
					break
				}
			}
			if allowed {
				break
			}
		}
	}
	if !allowed {
		return false, nil
	}
	if permissions, ok := action.DeniedAction[request.Path]; ok {
		for _, permission := range permissions {
			for _, requestedPermission := range request.Actions {
				if requestedPermission == permission || (permission == "*" && !allowAll) {
					allowed = false
					break
				}
			}
			if !allowed {
				break
			}
		}
	}
	return allowed, nil
}
