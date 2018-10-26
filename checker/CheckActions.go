package checker

import (
	"strings"

	"github.com/kkesley/perm"
)

//CheckActions checks if a role is allowed for a certain action
func CheckActions(request CheckActionsRequest) bool {
	if request.Token.IsRoot {
		return true
	}
	permissions := perm.GetActions(request.Bucket, request.Region, request.Role)
	allowAll := false
	denyAll := false
	allow := false
	deny := false
	for _, action := range request.Actions {
		for _, permission := range permissions {
			mode := "DENY"
			if string(permission[0]) == "+" {
				mode = "ALLOW"
			}
			sections := strings.Split(string(permission[1:]), "::")
			actionSections := strings.Split(action, "::")
			for idx, section := range sections {
				if len(actionSections)-1 < idx {
					break //if required action is shorter than the permission, it mustn't be a match
				}
				if section != actionSections[idx] && section != "*" {
					break //if current section is not the same as the required action and section is not *, mustn't be a match
				}
				if idx == len(sections)-1 {
					if section == "*" {
						if mode == "ALLOW" {
							allowAll = true //if last and section is * and mode is ALLOW, it means allow all
						} else {
							denyAll = true //if last and section is * and mode is DENY, it means deny all
						}
					} else if section == actionSections[len(actionSections)-1] {
						if mode == "ALLOW" {
							allow = true //if last and section is * and mode is ALLOW, it means allow all
						} else {
							deny = true //if last and section is * and mode is DENY, it means deny all
						}
					}
				}
			}
		}
		if allowAll && deny {
			continue
		} else if denyAll && allow {
			return true
		} else if allow || allowAll {
			return true
		}
	}
	return false
}
