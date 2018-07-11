package perm

type ResourceBank struct {
	Allow ResourceXpression `json:"allow"`
	Deny  ResourceXpression `json:"deny"`
}
