package checker

//FilterIDsOutput will return ids of conditions and owned resources
type FilterIDsOutput struct {
	Conditions []string `json:"conditions"`
	Owned      []string `json:"owned"`
}
