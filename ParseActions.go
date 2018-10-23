package perm

import (
	"encoding/json"
	"fmt"
)

//ParseActions parses actions based on policy in s3
func ParseActions(request ParseActionsRequest) ([]string, error) {
	if request.Token.IsRoot {
		return []string{"+*"}, nil
	} else if len(request.PolicyStr) <= 0 {
		return make([]string, 0), nil
	}
	var action Action
	//unmarshal policy string to a role
	if err := json.Unmarshal([]byte(request.PolicyStr), &action); err != nil {
		fmt.Println(err)
		return make([]string, 0), nil
	}

	finalPermissions := make([]string, 0)
	for path, permissions := range action.AllowedAction {
		if path == "*" {
			finalPermissions = append(finalPermissions, "+"+path)
		} else {
			for _, permission := range permissions {
				finalPermissions = append(finalPermissions, "+"+path+"::"+permission)
			}
		}
	}
	for path, permissions := range action.DeniedAction {
		if path == "*" {
			finalPermissions = append(finalPermissions, "-"+path)
		} else {
			for _, permission := range permissions {
				finalPermissions = append(finalPermissions, "-"+path+"::"+permission)
			}
		}
	}
	return finalPermissions, nil
}
