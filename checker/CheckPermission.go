package checker

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/kkesley/perm"
	parser "github.com/kkesley/s3-parser"
)

//CheckPermission returns allowed objects. If empty, it means the permission is denied
func CheckPermission(request CheckPermissionRequest) (*CheckResponse, error) {
	if request.Token.IsRoot {
		return &CheckResponse{
			Allow: perm.ResourceXpression{
				All:        true,
				Self:       true,
				Owned:      true,
				Conditions: make([]map[string]string, 0),
				Resources:  make([]string, 0),
			},
		}, nil
	} else if request.PolicyStr == nil || len(*request.PolicyStr) <= 0 {
		if len(request.Role) <= 0 {
			return &CheckResponse{
				Allow: perm.ResourceXpression{
					All:        false,
					Self:       false,
					Owned:      false,
					Conditions: make([]map[string]string, 0),
					Resources:  make([]string, 0),
				},
			}, nil
		}
		if roleByte, err := parser.GetS3DocumentDefault(request.Region, request.Bucket, strings.Replace(request.Role, "::", "_", -1)+"/final__role.json"); err == nil {
			request.PolicyStr = aws.String(string(roleByte))
		} else {
			return &CheckResponse{
				Allow: perm.ResourceXpression{
					All:        false,
					Self:       false,
					Owned:      false,
					Conditions: make([]map[string]string, 0),
					Resources:  make([]string, 0),
				},
			}, nil
		}
	}
	var role perm.Role
	//unmarshal policy string to a role
	if err := json.Unmarshal([]byte(*request.PolicyStr), &role); err != nil {
		fmt.Println(err)
		return nil, err
	}
	checkResponse := CheckResponse{
		Allow: getResource(role.AllowPolicy.Groups, request), //get allowed policies
		Deny:  getResource(role.DenyPolicy.Groups, request),  //get denied policies
	}
	return &checkResponse, nil
}

//getResource get the resource of eligible policies
func getResource(groups map[string]map[string]map[string]map[string][]*perm.Permission, request CheckPermissionRequest) perm.ResourceXpression {
	sections := strings.Split(request.Path, "::")
	resourcexpression := perm.ResourceXpression{
		Resources:  make([]string, 0),
		Conditions: make([]map[string]string, 0),
	}
	//sections must have a length at least 3 e.g. itea::platform::user
	if len(sections) < 3 {
		return resourcexpression
	}
	//check if 3 basic section is available in the permission
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

	//grab the resources
	if len(sections) == 3 {
		//if the length of section is exactly 3,
		//fill the resource in current permission
		for _, permission := range groups[request.ClientID][sections[0]][sections[1]][sections[2]] {
			fillResourceExpression(permission, request.Actions, &resourcexpression)
		}

	} else {
		//if the length of section is longer than 3, find all the children
		for _, permission := range groups[request.ClientID][sections[0]][sections[1]][sections[2]] {
			mapThroughChildren(permission, sections[3:], request.Actions, &resourcexpression)
		}

	}
	return resourcexpression
}

//mapThroughChildren traverse children to get permission
func mapThroughChildren(permission *perm.Permission, remainingSections []string, actions []string, resourcexpression *perm.ResourceXpression) {
	if len(remainingSections) <= 0 {
		//if remaining sections is empty, it should break
		return
	}
	currentSection := remainingSections[0] //current section to check
	for _, resource := range permission.Resources {
		if _, ok := resource.Children[currentSection]; !ok {
			//if no children align with the current section, check the next
			continue
		}
		if len(remainingSections) <= 1 {
			//this is the target. Fill the eligible resources
			fillResourceExpression(resource.Children[currentSection], actions, resourcexpression)
		} else {
			//traverse next level of children. remove the first remainingSections
			mapThroughChildren(resource.Children[currentSection], remainingSections[1:], actions, resourcexpression)
		}
	}
}

//fillResourceExpression fills the expression with eligible resources
func fillResourceExpression(permission *perm.Permission, actions []string, resourcexpression *perm.ResourceXpression) {
	//check if the actions are eligible to get these resources
	if checkEligibility(permission.Actions, actions) {
		//loop through the resources
		for _, resource := range permission.Resources {
			if resource.Target == "*" {
				//if target is *, the resource can be all targets with a certain conditions
				resourcexpression.All = true
				resourcexpression.Conditions = append(resourcexpression.Conditions, resource.Conditions)
			} else if resource.Target == "self" {
				//if the target is self, the resource is the one who call the api
				resourcexpression.Self = true
			} else if resource.Target == "owned" {
				//if the target is owned, the resource is all owned resources by the caller
				resourcexpression.Owned = true
			} else {
				//specific resources by id
				found := false
				//append all resources without having duplicates
				for _, existingResource := range resourcexpression.Resources {
					if existingResource == resource.Target {
						found = true
					}
				}
				if !found {
					//if no resource exist, append the resource
					resourcexpression.Resources = append(resourcexpression.Resources, resource.Target)
				}
			}
		}
	}
}

//checkEligibility checks the intersection of 2 action arrays and returns whether there's an intersection
func checkEligibility(actions []string, requiredActions []string) bool {
	//loop through required actions and actions to find intersection
	for _, requiredAction := range requiredActions {
		for _, action := range actions {
			if action == requiredAction || action == "*" {
				//if 1 intersection found or action is *, the actions are eligible
				return true
			}
		}
	}
	//if no intersection found, the actions are not eligible
	return false
}
