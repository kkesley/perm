package checker

import (
	"strings"

	"github.com/jinzhu/gorm"
)

//PermissionFilterIDs return query for filtering group. Note that the table field must be the same
//field must have
//arn (primary key)
//resource_group (index)
//group_key (primary key)
//group_value
//created_at
func (p CheckResponse) PermissionFilterIDs(db *gorm.DB, arnGroup string, arnRequester string) *gorm.DB {
	conditions := make([]map[string]string, 0)
	if p.Allow.All {
		conditions = p.Allow.Conditions
	} else if p.Deny.All {
		conditions = p.Deny.Conditions
	}
	if p.Allow.Owned || p.Deny.Owned {
		conditions = append(conditions, map[string]string{"owner": arnRequester})
	}

	if len(conditions) <= 0 {
		return nil
	}

	groupQuery := db.Select([]string{"arn", "group_key"}).Where("active = ?", true).Where("resource_group = ?", arnGroup)
	conditionQuery := make([]string, 0)
	conditionObj := make([]interface{}, 0)

	havingQueryExt := make([]string, 0)
	isCondition := false
	for _, condition := range conditions {
		havingQuery := make([]string, 0)
		conditionHit := false
		for key, val := range condition {
			if len(key) <= 0 || len(val) <= 0 {
				continue
			}
			conditionHit = true
			conditionQuery = append(conditionQuery, "(group_key = ? AND group_value = ?)")
			conditionObj = append(conditionObj, []interface{}{key, val}...)

			havingQuery = append(havingQuery, "SUM(group_key = ? AND group_value = ?) > 0")
			isCondition = true
		}
		if conditionHit {
			havingQueryExt = append(havingQueryExt, "("+strings.Join(havingQuery, " AND ")+")")
		}
	}

	if !isCondition {
		return nil
	}

	groupQuery = groupQuery.Having(strings.Join(havingQueryExt, " OR "), conditionObj...)
	groupQuery = groupQuery.Where(strings.Join(conditionQuery, " OR "), conditionObj...)
	groupQuery = groupQuery.Group("arn")
	return groupQuery
}
