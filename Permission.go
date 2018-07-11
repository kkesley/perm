package perm

//Permission holds the permissions of a certain resources
type Permission struct {
	Actions    []string
	Resources  []*Resource
	AllActions bool `json:"-"`
}
