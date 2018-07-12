package perm

//RawPolicy reflects the role document
type RawPolicy struct {
	Effect    string   `json:"Effect"`
	Resources []string `json:"Resources"`
	Actions   []string `json:"Actions"`
}
