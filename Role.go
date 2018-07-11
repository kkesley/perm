package perm

//Role holds both allow and deny policies
type Role struct {
	AllowPolicy Policy
	DenyPolicy  Policy
}
