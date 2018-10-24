package perm

import (
	"github.com/jinzhu/gorm"
	funk "github.com/thoas/go-funk"
)

//PermissionBatchQueryBuilder build logic of permission query for batch operations (list obviously)
func (permission CheckResponse) PermissionBatchQueryBuilder(db *gorm.DB, field string, conditionIDs []string) *gorm.DB {
	query := db
	if permission.Allow.All { //if permission has allow all
		if len(permission.Deny.Resources) > 0 {
			query = query.Where(field+" NOT IN (?)", permission.Deny.Resources) //if there are denied resources, make a NOT IN statement
		}
		requireCondition := false
		for _, condition := range permission.Allow.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition {
			query = query.Where(field+" IN (?)", funk.Uniq(append(conditionIDs, permission.Allow.Resources...))) //if there are conditions, only return object matches with conditions & allowed resources
		}
	} else if permission.Deny.All { //if permission has deny all
		if len(permission.Allow.Resources) > 0 {
			query = query.Where(field+" IN (?)", permission.Allow.Resources) //if there are allowed resources, make IN statement
		}

		requireCondition := false
		for _, condition := range permission.Deny.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition {
			query = query.Where(field+" NOT IN (?)", funk.Uniq(append(conditionIDs, permission.Deny.Resources...))) //if there are conditions, only deny object matches with the conditions and denied resources
		}
		if !requireCondition && len(permission.Allow.Resources) <= 0 {
			query = query.Where("user_id = ?", -1) //if no allowed resources and does not have any conditions, deny all of them
		}
	} else { //if permission has individual deny and allow
		if len(permission.Deny.Resources) > 0 {
			query = query.Where(field+" NOT IN (?)", permission.Deny.Resources) //filter denied resources
		}
		if len(permission.Allow.Resources) > 0 {
			query = query.Where(field+" IN (?)", permission.Allow.Resources) //only allow allowed resources
		}
		if len(permission.Deny.Resources) <= 0 && len(permission.Allow.Resources) <= 0 {
			query = query.Where("user_id = ?", -1) //deny all if permission is empty
		}
	}
	return query
}
