package perm

import "github.com/jinzhu/gorm"

//PermissionQueryBuilder build logic of permission query
func (permission CheckResponse) PermissionQueryBuilder(db *gorm.DB, field string, conditionIDs []string) *gorm.DB {
	query := db
	if permission.Allow.All {
		if len(permission.Deny.Resources) > 0 {
			query = query.Where(field+" NOT IN (?)", permission.Deny.Resources)
		}
		if len(conditionIDs) > 0 && len(permission.Allow.Resources) > 0 {
			if len(permission.Allow.Resources) > 0 {
				query = query.Where(field+" IN (?) OR "+field+" IN (?)", conditionIDs, permission.Allow.Resources)
			} else {
				query = query.Where(field+" IN (?)", conditionIDs)
			}
		}
	} else if permission.Deny.All {
		if len(permission.Allow.Resources) > 0 {
			query = query.Where(field+" IN (?)", permission.Allow.Resources)
		}
		if len(conditionIDs) > 0 && len(permission.Deny.Resources) > 0 {
			if len(permission.Deny.Resources) > 0 {
				query = query.Where(field+" NOT IN (?) AND "+field+" NOT IN (?)", conditionIDs, permission.Deny.Resources)
			} else {
				query = query.Where(field+" NOT IN (?)", conditionIDs)
			}
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
