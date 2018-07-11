package perm

//Resource holds the detail of the permission
type Resource struct {
	Target     string                 // id / *
	Conditions map[string]string      // for groupings
	Children   map[string]*Permission //children of the resource
	All        bool                   `json:"-"`
}
