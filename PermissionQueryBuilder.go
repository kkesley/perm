package perm

import (
	"github.com/jinzhu/gorm"
	funk "github.com/thoas/go-funk"
)

//PermissionQueryBuilder build logic of permission query
func (permission CheckResponse) PermissionQueryBuilder(db *gorm.DB, field string, conditionIDs []string) *gorm.DB {
	query := db
	if permission.Allow.All {
		if len(permission.Deny.Resources) > 0 {
			query = query.Where(field+" NOT IN (?)", permission.Deny.Resources)
		}
		requireCondition := false
		for _, condition := range permission.Allow.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition {
			query = query.Where(field+" IN (?)", funk.Uniq(append(conditionIDs, permission.Allow.Resources...)))
		}
	} else if permission.Deny.All {
		if len(permission.Allow.Resources) > 0 {
			query = query.Where(field+" IN (?)", permission.Allow.Resources)
		}

		requireCondition := false
		for _, condition := range permission.Deny.Conditions {
			if len(condition) > 0 {
				requireCondition = true
				break
			}
		}

		if requireCondition {
			query = query.Where(field+" NOT IN (?)", funk.Uniq(append(conditionIDs, permission.Deny.Resources...)))
		}
		if !requireCondition && len(permission.Allow.Resources) <= 0 {
			query = query.Where("user_id = ?", -1)
		}
	} else {
		if len(permission.Deny.Resources) > 0 {
			query = query.Where(field+" NOT IN (?)", permission.Deny.Resources)
		}
		if len(permission.Allow.Resources) > 0 {
			query = query.Where(field+" IN (?)", permission.Allow.Resources)
		}
	}
	return query
}
