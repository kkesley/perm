package perm

//Role holds both allow and deny policies
type Role struct {
	AllowPolicy Policy
	DenyPolicy  Policy
}

//Policy holds permissions grouped by resource group
type Policy struct {
	Groups map[string]map[string]map[string]map[string][]*Permission //client-id -> platform -> resource group -> resource name -> Resource
}

//Permission holds the permissions of a certain resources
type Permission struct {
	Actions    []string
	Resources  []*Resource
	AllActions bool `json:"-"`
}

//Resource holds the detail of the permission
type Resource struct {
	Target     string                 // id / *
	Conditions map[string]string      // for groupings
	Children   map[string]*Permission //children of the resource
	All        bool                   `json:"-"`
}
