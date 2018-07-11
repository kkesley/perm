package perm

import (
	"encoding/json"
	"fmt"
	"strings"
)

func CheckPermission(request CheckRequest) (*ResourceBank, error) {
	var role Role
	if err := json.Unmarshal([]byte(request.PolicyStr), &role); err != nil {
		fmt.Println(err)
		return nil, err
	}
	resourceBank := ResourceBank{
		Allow: GetResource(role.AllowPolicy.Groups, request),
		Deny:  GetResource(role.DenyPolicy.Groups, request),
	}
	return &resourceBank, nil
}

func GetResource(groups map[string]map[string]map[string]map[string][]*Permission, request CheckRequest) ResourceXpression {
	sections := strings.Split(request.Path, "::")
	resourcexpression := ResourceXpression{
		Resources:  make([]string, 0),
		Conditions: make([]map[string]string, 0),
	}
	if len(sections) < 3 {
		return resourcexpression
	}
	if _, ok := groups[request.ClientID]; !ok {
		return resourcexpression
	}
	if _, ok := groups[request.ClientID][sections[0]]; !ok {
		return resourcexpression
	}
	if _, ok := groups[request.ClientID][sections[0]][sections[1]]; !ok {
		return resourcexpression
	}
	if _, ok := groups[request.ClientID][sections[0]][sections[1]][sections[2]]; !ok {
		return resourcexpression
	}

	if len(sections) <= 3 {
		for _, permission := range groups[request.ClientID][sections[0]][sections[1]][sections[2]] {
			fillResourceExpression(permission, request.Actions, &resourcexpression)
		}

	} else {
		for _, permission := range groups[request.ClientID][sections[0]][sections[1]][sections[2]] {
			mapThroughChildren(permission, sections[3:], request.Actions, &resourcexpression)
		}

	}
	return resourcexpression
}

func mapThroughChildren(permission *Permission, remainingSections []string, actions []string, resourcexpression *ResourceXpression) {
	if len(remainingSections) <= 0 {
		return
	}
	currentSection := remainingSections[0]
	for _, resource := range permission.Resources {
		if _, ok := resource.Children[currentSection]; !ok {
			return
		}
		if len(remainingSections) <= 1 {
			//this is the target
			fillResourceExpression(resource.Children[currentSection], actions, resourcexpression)
		} else {
			mapThroughChildren(resource.Children[currentSection], remainingSections[1:], actions, resourcexpression)
		}
	}
}

func fillResourceExpression(permission *Permission, actions []string, resourcexpression *ResourceXpression) {
	if checkEligibility(permission.Actions, actions) {
		for _, resource := range permission.Resources {
			if resource.Target == "*" {
				resourcexpression.All = true
				resourcexpression.Conditions = append(resourcexpression.Conditions, resource.Conditions)
			} else if resource.Target == "self" {
				resourcexpression.Self = true
			} else if resource.Target == "owned" {
				resourcexpression.Owned = true
			} else {
				found := false
				for _, existingResource := range resourcexpression.Resources {
					if existingResource == resource.Target {
						found = true
					}
				}
				if !found {
					resourcexpression.Resources = append(resourcexpression.Resources, resource.Target)
				}
			}
		}
	}
}

func checkEligibility(permissions []string, requiredPermissions []string) bool {
	for _, requiredPermission := range requiredPermissions {
		for _, permission := range permissions {
			if permission == requiredPermission {
				return true
			}
		}
	}
	return false
}
