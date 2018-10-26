package checker

import (
	"github.com/jinzhu/gorm"
	funk "github.com/thoas/go-funk"
)

//PermissionListQueryBuilder build logic of permission query for list operations
func (permission CheckResponse) PermissionListQueryBuilder(db *gorm.DB, field string, filterOutput FilterIDsOutput) *gorm.DB {
	query := db
	denyOwned := make([]string, 0)
	allowOwned := make([]string, 0)
	if permission.Allow.Owned && !permission.Deny.Owned {
		allowOwned = filterOutput.Owned
	} else if permission.Deny.Owned && !permission.Allow.Owned {
		denyOwned = filterOutput.Owned
	}
	if permission.Allow.All { //if permission has allow all
		if len(permission.Deny.Resources) > 0 || len(denyOwned) > 0 {
			query = query.Where(field+" NOT IN (?)", funk.UniqString(append(denyOwned, permission.Deny.Resources...))) //if there are denied resources, make a NOT IN statement
		}
		requireCondition := false
		for _, condition := range permission.Allow.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition || len(allowOwned) > 0 {
			query = query.Where(field+" IN (?)", funk.UniqString(append(append(filterOutput.Conditions, permission.Allow.Resources...), allowOwned...))) //if there are conditions, only return object matches with conditions & allowed resources
		}
	} else if permission.Deny.All { //if permission has deny all
		if len(permission.Allow.Resources) > 0 || len(allowOwned) > 0 {
			query = query.Where(field+" IN (?)", funk.UniqString(append(allowOwned, permission.Allow.Resources...))) //if there are allowed resources, make IN statement
		}

		requireCondition := false
		for _, condition := range permission.Deny.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition || len(denyOwned) > 0 {
			query = query.Where(field+" NOT IN (?)", funk.UniqString(append(append(filterOutput.Conditions, permission.Deny.Resources...), allowOwned...))) //if there are conditions, only deny object matches with the conditions and denied resources
		}
		if !requireCondition && len(permission.Allow.Resources) <= 0 {
			query = query.Where("user_id = ?", -1) //if no allowed resources and does not have any conditions, deny all of them
		}
	} else { //if permission has individual deny and allow
		if len(permission.Deny.Resources) > 0 || len(denyOwned) > 0 {
			query = query.Where(field+" NOT IN (?)", funk.UniqString(append(denyOwned, permission.Deny.Resources...))) //filter denied resources
		}
		if len(permission.Allow.Resources) > 0 || len(allowOwned) > 0 {
			query = query.Where(field+" IN (?)", funk.UniqString(append(allowOwned, permission.Allow.Resources...))) //only allow allowed resources
		}
		if len(permission.Deny.Resources) <= 0 && len(permission.Allow.Resources) <= 0 && len(denyOwned) <= 0 && len(allowOwned) <= 0 {
			query = query.Where(field+" = ?", -1) //deny all if permission is empty
		}
	}
	return query
}
